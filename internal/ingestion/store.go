// Package ingestion coordinates replayable, correction-aware billing imports.
package ingestion

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sync"
	"time"

	"github.com/marco13-moo/cloudspend-guardian/internal/billing"
)

// Artifact is an immutable, content-addressed source export.
type Artifact struct {
	Hash          string
	SchemaVersion string
	Content       []byte
}

// RecordVersion preserves the complete lineage of a canonical source record.
type RecordVersion struct {
	Scope      string
	Record     billing.Record
	SourceHash string
	Version    int
	Supersedes string
	ImportedAt time.Time
}

// Cursor records the last successfully applied billing window for a scope.
type Cursor struct {
	Scope       string
	WindowStart time.Time
	WindowEnd   time.Time
	SourceHash  string
	CompletedAt time.Time
}

// ImportResult describes whether an import created state, replayed state, or corrected it.
type ImportResult struct {
	ArtifactHash       string
	RecordsApplied     int
	CorrectionsApplied int
	Replay             bool
	Cursor             Cursor
}

// Store is an in-memory implementation of the operational persistence contract.
// Production adapters can retain the same transactional semantics in PostgreSQL
// and object storage without changing the ingestion coordinator.
type Store struct {
	mu        sync.RWMutex
	artifacts map[string]Artifact
	records   map[string][]RecordVersion
	cursors   map[string]Cursor
}

// NewStore creates an empty operational store.
func NewStore() *Store {
	return &Store{
		artifacts: make(map[string]Artifact),
		records:   make(map[string][]RecordVersion),
		cursors:   make(map[string]Cursor),
	}
}

// ImportBatch validates and atomically applies one bounded billing window.
func (s *Store) ImportBatch(scope string, windowStart, windowEnd time.Time, source io.Reader, importedAt time.Time) (ImportResult, error) {
	if scope == "" {
		return ImportResult{}, errors.New("import scope is required")
	}
	if !windowStart.Before(windowEnd) {
		return ImportResult{}, errors.New("import window start must precede window end")
	}
	raw, err := io.ReadAll(source)
	if err != nil {
		return ImportResult{}, fmt.Errorf("read source artifact: %w", err)
	}
	report, err := billing.ImportCSV(bytes.NewReader(raw))
	if err != nil {
		return ImportResult{}, fmt.Errorf("validate source artifact: %w", err)
	}
	if importedAt.IsZero() {
		return ImportResult{}, errors.New("import timestamp is required")
	}
	importedAt = importedAt.UTC()

	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.artifacts[report.SourceSHA256]; exists {
		cursor := s.cursors[scope]
		return ImportResult{ArtifactHash: report.SourceSHA256, Replay: true, Cursor: cursor}, nil
	}

	// Stage all mutations first. Any validation failure leaves the store and cursor unchanged.
	staged := make(map[string][]RecordVersion)
	corrections := 0
	for _, record := range report.Records {
		key := scope + "\x00" + record.ID
		versions := append([]RecordVersion(nil), s.records[key]...)
		if len(versions) > 0 {
			current := versions[len(versions)-1]
			if reflect.DeepEqual(current.Record, record) {
				continue
			}
			corrections++
			versions = append(versions, RecordVersion{
				Scope: scope, Record: record, SourceHash: report.SourceSHA256,
				Version: len(versions) + 1, Supersedes: current.SourceHash, ImportedAt: importedAt,
			})
		} else {
			versions = append(versions, RecordVersion{
				Scope: scope, Record: record, SourceHash: report.SourceSHA256,
				Version: 1, ImportedAt: importedAt,
			})
		}
		staged[key] = versions
	}

	for key, versions := range staged {
		s.records[key] = versions
	}
	s.artifacts[report.SourceSHA256] = Artifact{
		Hash: report.SourceSHA256, SchemaVersion: report.SchemaVersion, Content: append([]byte(nil), raw...),
	}
	cursor := Cursor{Scope: scope, WindowStart: windowStart.UTC(), WindowEnd: windowEnd.UTC(), SourceHash: report.SourceSHA256, CompletedAt: importedAt}
	s.cursors[scope] = cursor
	return ImportResult{ArtifactHash: report.SourceSHA256, RecordsApplied: len(report.Records), CorrectionsApplied: corrections, Cursor: cursor}, nil
}

// Artifact returns an immutable copy of a retained source artifact.
func (s *Store) Artifact(hash string) (Artifact, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	artifact, ok := s.artifacts[hash]
	if !ok {
		return Artifact{}, false
	}
	artifact.Content = append([]byte(nil), artifact.Content...)
	return artifact, true
}

// CurrentRecords returns the latest version of every record in a scope.
func (s *Store) CurrentRecords(scope string) []RecordVersion {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]RecordVersion, 0)
	for key, versions := range s.records {
		if len(versions) > 0 && len(key) >= len(scope)+1 && key[:len(scope)] == scope && key[len(scope)] == '\x00' {
			result = append(result, versions[len(versions)-1])
		}
	}
	return result
}

// Versions returns the complete correction lineage for one canonical record.
func (s *Store) Versions(scope, recordID string) []RecordVersion {
	s.mu.RLock()
	defer s.mu.RUnlock()
	versions := s.records[scope+"\x00"+recordID]
	return append([]RecordVersion(nil), versions...)
}

// Cursor returns the last completed import for a scope.
func (s *Store) Cursor(scope string) (Cursor, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	cursor, ok := s.cursors[scope]
	return cursor, ok
}

// HashSource is exposed for callers that need to compare an artifact without storing it.
func HashSource(source []byte) string {
	digest := sha256.Sum256(source)
	return hex.EncodeToString(digest[:])
}
