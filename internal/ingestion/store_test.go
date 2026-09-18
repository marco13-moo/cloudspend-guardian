package ingestion

import (
	"strings"
	"testing"
	"time"

	"github.com/marco13-moo/cloudspend-guardian/internal/demo"
)

func TestImportBatchIsIdempotentAndRetainsArtifact(t *testing.T) {
	store := NewStore()
	first, err := store.ImportBatch("org-1", day(1), day(15), strings.NewReader(demo.FixtureCSV()), day(16))
	if err != nil {
		t.Fatalf("first import: %v", err)
	}
	second, err := store.ImportBatch("org-1", day(1), day(15), strings.NewReader(demo.FixtureCSV()), day(17))
	if err != nil {
		t.Fatalf("replay import: %v", err)
	}
	if second.ArtifactHash != first.ArtifactHash || !second.Replay || second.RecordsApplied != 0 {
		t.Fatalf("expected no-op replay: %+v", second)
	}
	artifact, ok := store.Artifact(first.ArtifactHash)
	if !ok || artifact.SchemaVersion == "" || len(artifact.Content) == 0 {
		t.Fatal("expected immutable source artifact")
	}
	if len(store.CurrentRecords("org-1")) != 28 {
		t.Fatal("expected one current version per canonical billing record")
	}
}

func TestImportBatchPreservesCorrectionLineage(t *testing.T) {
	store := NewStore()
	original := demo.FixtureCSV()
	if _, err := store.ImportBatch("org-1", day(1), day(15), strings.NewReader(original), day(16)); err != nil {
		t.Fatalf("first import: %v", err)
	}
	corrected := strings.Replace(original, ",2.50,EUR,", ",3.50,EUR,", 1)
	result, err := store.ImportBatch("org-1", day(1), day(15), strings.NewReader(corrected), day(18))
	if err != nil {
		t.Fatalf("correction import: %v", err)
	}
	if result.CorrectionsApplied != 1 || result.Replay {
		t.Fatalf("expected one correction: %+v", result)
	}
	var correctedID string
	for _, current := range store.CurrentRecords("org-1") {
		if current.Record.ResourceID == "arn:synthetic:ec2:idle-1" && current.Record.BilledCost.MinorUnits == 350 {
			correctedID = current.Record.ID
			break
		}
	}
	versions := store.Versions("org-1", correctedID)
	if len(versions) != 2 || versions[1].Version != 2 || versions[1].Supersedes == "" {
		t.Fatalf("expected correction lineage: %+v", versions)
	}
}

func TestImportBatchDoesNotAdvanceCursorOnInvalidInput(t *testing.T) {
	store := NewStore()
	if _, err := store.ImportBatch("org-1", day(1), day(15), strings.NewReader("invalid"), day(16)); err == nil {
		t.Fatal("expected invalid import to fail")
	}
	if _, ok := store.Cursor("org-1"); ok {
		t.Fatal("invalid import advanced cursor")
	}
}

func day(dayOfMonth int) time.Time {
	return time.Date(2026, 9, dayOfMonth, 0, 0, 0, 0, time.UTC)
}
