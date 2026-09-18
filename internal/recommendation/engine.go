// Package recommendation evaluates deterministic, versioned FinOps rules.
package recommendation

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/marco13-moo/cloudspend-guardian/internal/billing"
	"github.com/marco13-moo/cloudspend-guardian/internal/domain"
)

const IdleComputeRuleVersion = "idle-compute/v1"

// Abstention explains why the engine deliberately declined to recommend action.
type Abstention struct {
	ResourceID  string `json:"resourceId"`
	RuleVersion string `json:"ruleVersion"`
	Reason      string `json:"reason"`
}

// Evaluation preserves actionable recommendations and deliberate non-actions.
type Evaluation struct {
	Recommendations []domain.Recommendation `json:"recommendations"`
	Abstentions     []Abstention            `json:"abstentions"`
}

// EvaluateIdleCompute evaluates a conservative synthetic rightsizing rule.
func EvaluateIdleCompute(report billing.ImportReport, evaluatedAt time.Time) (Evaluation, error) {
	if !report.ReconciliationOK {
		return Evaluation{}, fmt.Errorf("billing import must reconcile before evaluation")
	}
	grouped := make(map[string][]billing.Record)
	for _, record := range report.Records {
		grouped[record.ResourceID] = append(grouped[record.ResourceID], record)
	}

	resourceIDs := make([]string, 0, len(grouped))
	for resourceID := range grouped {
		resourceIDs = append(resourceIDs, resourceID)
	}
	sort.Strings(resourceIDs)

	result := Evaluation{}
	for _, resourceID := range resourceIDs {
		records := grouped[resourceID]
		sort.Slice(records, func(i, j int) bool { return records[i].ChargePeriodStart.Before(records[j].ChargePeriodStart) })
		days := distinctObservationDays(records)
		if days < 14 {
			result.Abstentions = append(result.Abstentions, Abstention{resourceID, IdleComputeRuleVersion, "fewer than 14 distinct observation days"})
			continue
		}

		var cost int64
		var cpu, memory, errorBudget float64
		currency := records[0].BilledCost.Currency
		for _, record := range records {
			if record.BilledCost.Currency != currency {
				return Evaluation{}, fmt.Errorf("resource %s contains mixed currencies", resourceID)
			}
			cost += record.BilledCost.MinorUnits
			cpu += record.CPUUtilizationP95
			memory += record.MemoryUtilizationP95
			errorBudget += record.SLOErrorBudgetRemaining
		}
		count := float64(len(records))
		avgCPU, avgMemory, avgErrorBudget := cpu/count, memory/count, errorBudget/count
		if avgErrorBudget < 0.5 {
			result.Abstentions = append(result.Abstentions, Abstention{resourceID, IdleComputeRuleVersion, "reliability veto: remaining SLO error budget is below 50%"})
			continue
		}
		if avgCPU > 10 || avgMemory > 20 {
			result.Abstentions = append(result.Abstentions, Abstention{resourceID, IdleComputeRuleVersion, "resource is not consistently underutilized"})
			continue
		}

		factors := map[string]float64{
			"dataCompleteness":    1,
			"observationDuration": min(float64(days)/14, 1),
			"telemetryQuality":    1,
			"workloadStability":   max(0, 1-(avgCPU/100)-(avgMemory/200)),
		}
		confidence := (factors["dataCompleteness"] + factors["observationDuration"] + factors["telemetryQuality"] + factors["workloadStability"]) / 4
		if cost > math.MaxInt64/30 {
			return Evaluation{}, fmt.Errorf("resource %s cost exceeds supported projection range", resourceID)
		}
		monthlyBaseline := cost * 30 / int64(days)
		projectedSavings := monthlyBaseline / 2
		idDigest := sha256.Sum256([]byte(IdleComputeRuleVersion + "\x00" + resourceID + "\x00" + report.SourceSHA256))

		recommendation := domain.Recommendation{
			ID:          hex.EncodeToString(idDigest[:]),
			ResourceID:  resourceID,
			Summary:     "Review a 50% rightsizing step for consistently underutilized compute",
			RuleVersion: IdleComputeRuleVersion,
			Evidence: []domain.Evidence{
				{Source: "synthetic-focus-fixture", Description: fmt.Sprintf("%d observation days; CPU p95 average %.2f%%; memory p95 average %.2f%%", days, avgCPU, avgMemory), ObservedAt: evaluatedAt.UTC()},
				{Source: "synthetic-slo-fixture", Description: fmt.Sprintf("remaining error budget average %.2f", avgErrorBudget), ObservedAt: evaluatedAt.UTC()},
			},
			Assumptions: []string{
				"the synthetic observation window represents future demand",
				"a 50% rightsizing step preserves workload compatibility",
			},
			Exclusions: []string{
				"commitment discounts and negotiated prices",
				"provider-specific migration and data-transfer charges",
			},
			Confidence:        confidence,
			ConfidenceFactors: factors,
			ProjectedMonthlySavings: domain.Money{
				Currency:   currency,
				MinorUnits: projectedSavings,
			},
			RollbackPlan: "Restore the previous infrastructure-as-code sizing value and redeploy through the ordinary reviewed pipeline.",
			CreatedAt:    evaluatedAt.UTC(),
		}
		if err := recommendation.Validate(); err != nil {
			return Evaluation{}, fmt.Errorf("validate recommendation for %s: %w", resourceID, err)
		}
		result.Recommendations = append(result.Recommendations, recommendation)
	}
	return result, nil
}

func distinctObservationDays(records []billing.Record) int {
	days := make(map[string]struct{}, len(records))
	for _, record := range records {
		days[record.ChargePeriodStart.UTC().Format(time.DateOnly)] = struct{}{}
	}
	return len(days)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
