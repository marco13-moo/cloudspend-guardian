package measurement

import (
	"testing"
	"time"

	"github.com/marco13-moo/cloudspend-guardian/internal/domain"
)

func TestLedgerIsAppendOnlyAndRejectsDuplicateEvents(t *testing.T) {
	entry := Entry{ID: "event-1", OpportunityKey: "org/resource", RecommendationID: "rec-1", Kind: "projected", Amount: domain.Money{Currency: "EUR", MinorUnits: 300}, Baseline: domain.Money{Currency: "EUR", MinorUnits: 600}, WindowStart: day(1), WindowEnd: day(15), RecordedAt: day(16), Counterfactual: "pre-change baseline"}
	ledger := NewLedger()
	if err := ledger.Append(entry); err != nil {
		t.Fatal(err)
	}
	if err := ledger.Append(entry); err == nil {
		t.Fatal("expected duplicate event to be rejected")
	}
	if len(ledger.Entries()) != 1 {
		t.Fatal("expected append-only ledger")
	}
}

func day(dayOfMonth int) time.Time { return time.Date(2026, 9, dayOfMonth, 0, 0, 0, 0, time.UTC) }
