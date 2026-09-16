# ADR-0018: Measure realized savings separately

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Measurement
- Related issues: MAR-13

## Context

Projected savings are hypotheses. A merged remediation can underperform, be reverted, or coincide with workload changes.

## Proposed decision

Represent projected, approved, observed, and verified savings as distinct measurements with currencies, windows, assumptions, and provenance.

## Consequences

Portfolio claims become credible and optimization quality becomes measurable. Verification delays final outcomes and requires retained baselines.

## Guardrails

- Never overwrite projections with observations.
- A merged pull request is not evidence of realized savings.
- Reliability regressions invalidate successful-outcome classification.

## Validation

- Demonstrate accurate, underperforming, reverted, and inconclusive optimization outcomes.

## Revisit triggers

- Finance-reviewed accounting requirements demand another measurement taxonomy.
