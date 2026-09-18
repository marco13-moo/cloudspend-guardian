package recommendation

import (
	"testing"
	"time"

	"github.com/marco13-moo/cloudspend-guardian/internal/domain"
)

func TestDeduplicateMergesCorroborationAndUsesConservativeSavings(t *testing.T) {
	base := domain.Recommendation{ID: "provider-1", ProjectedMonthlySavings: domain.Money{Currency: "EUR", MinorUnits: 500}, Evidence: []domain.Evidence{{Source: "provider"}}}
	native := base
	native.ID = "native-1"
	native.Evidence = append([]domain.Evidence(nil), base.Evidence...)
	native.ProjectedMonthlySavings.MinorUnits = 300
	candidates := []Candidate{
		{Organization: "org", ResourceID: "resource", ActionClass: "rightsizing", TargetState: "small", WindowStart: day(1), WindowEnd: day(15), Source: "provider", Recommendation: base},
		{Organization: "org", ResourceID: "resource", ActionClass: "rightsizing", TargetState: "small", WindowStart: day(1), WindowEnd: day(15), Source: "native", Recommendation: native},
	}
	opportunities, err := Deduplicate(candidates)
	if err != nil || len(opportunities) != 1 {
		t.Fatalf("expected one opportunity: %+v %v", opportunities, err)
	}
	if opportunities[0].Recommendation.ProjectedMonthlySavings.MinorUnits != 300 || len(opportunities[0].Recommendation.Evidence) != 2 {
		t.Fatalf("expected conservative merged opportunity: %+v", opportunities[0])
	}
}

func day(dayOfMonth int) time.Time { return time.Date(2026, 9, dayOfMonth, 0, 0, 0, 0, time.UTC) }
