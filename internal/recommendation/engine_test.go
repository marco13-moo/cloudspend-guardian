package recommendation

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/marco13-moo/cloudspend-guardian/internal/billing"
)

const testHeader = "BillingPeriodStart,BillingPeriodEnd,ChargePeriodStart,ChargePeriodEnd,ProviderName,BillingAccountId,ResourceId,ServiceName,BilledCost,Currency,UsageQuantity,UsageUnit,CPUUtilizationP95,MemoryUtilizationP95,SLOErrorBudgetRemaining\n"

func TestEvaluateIdleComputeProducesExplainableRecommendation(t *testing.T) {
	report := importFixture(t, 14, 4, 8, 0.9)
	evaluation, err := EvaluateIdleCompute(report, time.Date(2026, 9, 15, 12, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("unexpected evaluation error: %v", err)
	}
	if len(evaluation.Recommendations) != 1 || len(evaluation.Abstentions) != 0 {
		t.Fatalf("unexpected evaluation: %+v", evaluation)
	}
	recommendation := evaluation.Recommendations[0]
	if recommendation.RuleVersion != IdleComputeRuleVersion || len(recommendation.Evidence) != 2 || recommendation.Confidence < 0.9 {
		t.Fatalf("recommendation lacks expected explanation: %+v", recommendation)
	}
	if recommendation.ProjectedMonthlySavings.MinorUnits != 3750 {
		t.Fatalf("unexpected projected savings: %+v", recommendation.ProjectedMonthlySavings)
	}
}

func TestEvaluateIdleComputeAbstainsOnInsufficientEvidence(t *testing.T) {
	evaluation, err := EvaluateIdleCompute(importFixture(t, 13, 4, 8, 0.9), time.Now())
	if err != nil {
		t.Fatalf("unexpected evaluation error: %v", err)
	}
	if len(evaluation.Recommendations) != 0 || !strings.Contains(evaluation.Abstentions[0].Reason, "14") {
		t.Fatalf("expected duration abstention: %+v", evaluation)
	}
}

func TestEvaluateIdleComputeHonorsReliabilityVeto(t *testing.T) {
	evaluation, err := EvaluateIdleCompute(importFixture(t, 14, 4, 8, 0.2), time.Now())
	if err != nil {
		t.Fatalf("unexpected evaluation error: %v", err)
	}
	if len(evaluation.Recommendations) != 0 || !strings.Contains(evaluation.Abstentions[0].Reason, "reliability veto") {
		t.Fatalf("expected reliability abstention: %+v", evaluation)
	}
}

func TestEvaluateIdleComputeRequiresReconciledInput(t *testing.T) {
	if _, err := EvaluateIdleCompute(billing.ImportReport{}, time.Now()); err == nil {
		t.Fatal("expected unreconciled input to be rejected")
	}
}

func importFixture(t *testing.T, days int, cpu, memory, errorBudget float64) billing.ImportReport {
	t.Helper()
	var fixture strings.Builder
	fixture.WriteString(testHeader)
	start := time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC)
	for day := 0; day < days; day++ {
		chargeStart := start.AddDate(0, 0, day)
		chargeEnd := chargeStart.AddDate(0, 0, 1)
		fmt.Fprintf(&fixture, "2026-09-01T00:00:00Z,2026-10-01T00:00:00Z,%s,%s,AWS,synthetic-account,arn:synthetic:ec2:idle-1,Amazon EC2,2.50,EUR,24,hours,%.2f,%.2f,%.2f\n", chargeStart.Format(time.RFC3339), chargeEnd.Format(time.RFC3339), cpu, memory, errorBudget)
	}
	report, err := billing.ImportCSV(strings.NewReader(fixture.String()))
	if err != nil {
		t.Fatalf("import fixture: %v", err)
	}
	return report
}
