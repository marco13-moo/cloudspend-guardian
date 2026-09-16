# ADR-0025: Define measurable control-plane SLOs

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Reliability
- Related issues: MAR-13

## Context

Availability alone does not describe whether billing is fresh, recommendations are evidenced, or jobs complete on time.

## Proposed decision

Define SLOs for API availability, ingestion freshness, job completion, recommendation evidence completeness, and authorization correctness. Establish error budgets before production certification.

## Consequences

Reliability discussions gain measurable objectives. Instrumentation and operational response become product obligations.

## Guardrails

- Security invariants are absolute controls, not burnable error budgets.
- SLO windows and exclusions are versioned.
- User-facing freshness is derived from source completion, not process uptime.

## Validation

- Dashboards and alerts reproduce success, partial-source failure, backlog, and stale-data scenarios.

## Revisit triggers

- Observed user needs or operating cost demonstrate that targets are ineffective or disproportionate.
