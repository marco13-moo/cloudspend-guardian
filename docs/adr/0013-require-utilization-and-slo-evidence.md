# ADR-0013: Require utilization and SLO evidence

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Recommendation engine
- Related issues: MAR-13

## Context

Low expenditure or low average CPU alone cannot establish safe rightsizing; workload peaks, memory, latency, and reliability headroom may dominate.

## Proposed decision

Require a rule-specific evidence contract covering observation duration, utilization percentiles, workload stability, and applicable SLO/error-budget signals. Insufficient evidence produces no actionable remediation.

## Consequences

Recommendations become defensible but require richer telemetry and may remain advisory when evidence is incomplete.

## Guardrails

- Reliability vetoes financial optimization.
- Missing telemetry lowers confidence or blocks action; it never becomes zero utilization.
- Evidence windows and exclusions are displayed.

## Validation

- Golden tests cover safe rightsizing, bursty workloads, missing memory, and exhausted error budgets.

## Revisit triggers

- Empirical review demonstrates that a required signal adds no decision value.
