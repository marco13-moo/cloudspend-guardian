package domain

import (
	"testing"
	"time"
)

func TestRecommendationRequiresEvidence(t *testing.T) {
	recommendation := Recommendation{
		ID:                      "rec-1",
		ResourceID:              "arn:example",
		Summary:                 "Downsize an underutilized resource",
		RuleVersion:             "idle-compute/v1",
		Assumptions:             []string{"synthetic fixture"},
		ConfidenceFactors:       map[string]float64{"completeness": 1},
		Confidence:              0.9,
		ProjectedMonthlySavings: Money{Currency: "EUR", MinorUnits: 7100},
		RollbackPlan:            "Restore the previous Terraform variable.",
		CreatedAt:               time.Now(),
	}
	if err := recommendation.Validate(); err == nil {
		t.Fatal("expected recommendation without evidence to be rejected")
	}
}

func TestRecommendationAcceptsCompleteEvidence(t *testing.T) {
	recommendation := Recommendation{
		ID:                      "rec-1",
		ResourceID:              "arn:example",
		Summary:                 "Downsize an underutilized resource",
		RuleVersion:             "idle-compute/v1",
		Evidence:                []Evidence{{Source: "prometheus", Description: "CPU p95 below 5%", ObservedAt: time.Now()}},
		Assumptions:             []string{"the observation window is representative"},
		Exclusions:              []string{"reserved-instance pricing"},
		ConfidenceFactors:       map[string]float64{"dataCompleteness": 1},
		Confidence:              0.9,
		ProjectedMonthlySavings: Money{Currency: "EUR", MinorUnits: 7100},
		RollbackPlan:            "Restore the previous Terraform variable.",
		CreatedAt:               time.Now(),
	}
	if err := recommendation.Validate(); err != nil {
		t.Fatalf("expected recommendation to be valid: %v", err)
	}
}
