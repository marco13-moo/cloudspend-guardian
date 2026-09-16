# ADR-0015: Deduplicate provider and native recommendations

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Recommendation engine
- Related issues: MAR-13

## Context

AWS and CloudSpend Guardian may identify the same resource opportunity through different evidence, causing duplicated remediation and overstated savings.

## Proposed decision

Construct a canonical opportunity key from organization, resource, action class, target state, and overlapping effective window. Merge corroborating evidence while preserving every source recommendation.

## Consequences

Operators see one actionable opportunity and savings are not double counted. Ambiguous overlaps require conservative grouping.

## Guardrails

- Source recommendations remain independently traceable.
- Canonical savings use the most conservative compatible estimate.
- Incompatible target states remain separate alternatives.

## Validation

- Test exact duplicates, overlapping rightsizing, conflicting targets, and recurrence after remediation.

## Revisit triggers

- False merges or missed duplicates exceed an agreed error threshold.
