# ADR-0026: Standardize observability on OpenTelemetry

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Observability
- Related issues: MAR-13

## Context

Billing ingestion, recommendation, policy, repository proposal, and later verification form one distributed logical workflow even in a modular monolith.

## Proposed decision

Instrument traces, metrics, and structured logs through OpenTelemetry semantic conventions extended by a small versioned FinOps vocabulary.

## Consequences

Backends remain replaceable and workflows become correlatable. Cardinality and sensitive-field discipline require enforcement.

## Guardrails

- No credentials, raw billing records, or unrestricted resource tags in telemetry.
- Correlation identifiers are opaque and organization-scoped.
- Metrics labels have explicit cardinality budgets.

## Validation

- Trace the complete synthetic workflow and verify redaction, sampling, and failure linkage.

## Revisit triggers

- Required signals cannot be represented or OpenTelemetry creates measured operational harm.
