// Package domain defines provider-neutral FinOps concepts and invariants.
package domain

import (
	"errors"
	"time"
)

// Recommendation describes a proposed optimization; it is evidence, not authority to mutate infrastructure.
type Recommendation struct {
	ID                      string             `json:"id"`
	ResourceID              string             `json:"resourceId"`
	Summary                 string             `json:"summary"`
	RuleVersion             string             `json:"ruleVersion"`
	Evidence                []Evidence         `json:"evidence"`
	Assumptions             []string           `json:"assumptions"`
	Exclusions              []string           `json:"exclusions"`
	Confidence              float64            `json:"confidence"`
	ConfidenceFactors       map[string]float64 `json:"confidenceFactors"`
	ProjectedMonthlySavings Money              `json:"projectedMonthlySavings"`
	RollbackPlan            string             `json:"rollbackPlan"`
	CreatedAt               time.Time          `json:"createdAt"`
}

// Evidence records one observable fact supporting a recommendation.
type Evidence struct {
	Source      string    `json:"source"`
	Description string    `json:"description"`
	ObservedAt  time.Time `json:"observedAt"`
}

// Money uses integer minor units to avoid floating-point accounting errors.
type Money struct {
	Currency   string `json:"currency"`
	MinorUnits int64  `json:"minorUnits"`
}

// Validate enforces the minimum explainability contract for every recommendation.
func (r Recommendation) Validate() error {
	if r.ID == "" || r.ResourceID == "" || r.Summary == "" {
		return errors.New("identity, resource, and summary are required")
	}
	if len(r.Evidence) == 0 {
		return errors.New("at least one evidence item is required")
	}
	if r.RuleVersion == "" {
		return errors.New("rule version is required")
	}
	if len(r.Assumptions) == 0 {
		return errors.New("at least one assumption is required")
	}
	if len(r.ConfidenceFactors) == 0 {
		return errors.New("confidence factors are required")
	}
	if r.Confidence < 0 || r.Confidence > 1 {
		return errors.New("confidence must be between zero and one")
	}
	if r.ProjectedMonthlySavings.Currency == "" {
		return errors.New("savings currency is required")
	}
	if r.RollbackPlan == "" {
		return errors.New("rollback plan is required")
	}
	return nil
}
