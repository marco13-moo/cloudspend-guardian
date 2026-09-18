// Package measurement records savings claims without rewriting history.
package measurement

import (
	"errors"
	"sync"
	"time"

	"github.com/marco13-moo/cloudspend-guardian/internal/domain"
)

type Entry struct {
	ID               string
	OpportunityKey   string
	RecommendationID string
	Kind             string
	Amount           domain.Money
	Baseline         domain.Money
	WindowStart      time.Time
	WindowEnd        time.Time
	RecordedAt       time.Time
	Counterfactual   string
}

type Ledger struct {
	mu      sync.RWMutex
	entries []Entry
	ids     map[string]struct{}
}

func NewLedger() *Ledger { return &Ledger{ids: make(map[string]struct{})} }

func (l *Ledger) Append(entry Entry) error {
	if entry.ID == "" || entry.OpportunityKey == "" || entry.RecommendationID == "" || entry.Kind == "" || entry.Amount.Currency == "" || entry.Baseline.Currency == "" || entry.Counterfactual == "" || !entry.WindowStart.Before(entry.WindowEnd) || entry.RecordedAt.IsZero() {
		return errors.New("savings entry is incomplete")
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, exists := l.ids[entry.ID]; exists {
		return errors.New("savings entry already exists")
	}
	l.ids[entry.ID] = struct{}{}
	l.entries = append(l.entries, entry)
	return nil
}

func (l *Ledger) Entries() []Entry {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return append([]Entry(nil), l.entries...)
}
