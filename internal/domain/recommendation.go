// Package domain defines provider-neutral FinOps concepts and invariants.
package domain

import (
	"errors"
	"time"
)

// Recommendation describes a proposed optimization; it is evidence, not authority to mutate infrastructure.
type Recommendation struct {
	ID                      string
	ResourceID              string
	Summary                 string
	Evidence                []Evidence
	Confidence              float64
	ProjectedMonthlySavings Money
	RollbackPlan            string
	CreatedAt               time.Time
}

// Evidence records one observable fact supporting a recommendation.
type Evidence struct {
	Source      string
	Description string
	ObservedAt  time.Time
}

// Money uses integer minor units to avoid floating-point accounting errors.
type Money struct {
	Currency   string
	MinorUnits int64
}

// Validate enforces the minimum explainability contract for every recommendation.
func (r Recommendation) Validate() error {
	if r.ID == "" || r.ResourceID == "" || r.Summary == "" {
		return errors.New("identity, resource, and summary are required")
	}
	if len(r.Evidence) == 0 {
		return errors.New("at least one evidence item is required")
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
