# ADR-0006: Adopt a FOCUS-compatible canonical cost model

- Status: Proposed
- Date: 2026-09-15
- Owners: Repository maintainers
- Decision scope: Data
- Related issues: MAR-13
- Supersedes: None
- Superseded by: None

## Context

Provider-native billing schemas are intricate, mutable, and mutually incompatible. The domain needs a stable vocabulary without discarding raw source evidence.

## Proposed decision

Store immutable provider records and map them into a versioned model compatible with the FinOps Open Cost and Usage Specification. Preserve source provenance and reconcile monetary totals at every ingestion boundary.

## Consequences

Analytics and policy can operate against a consistent model. Mapping complexity and specification-version migrations become explicit operational obligations.

## Guardrails

- Raw provider records remain immutable and retain provenance.
- Unknown or lossy mappings fail visibly rather than silently coercing values.
- Monetary reconciliation is mandatory before an import becomes authoritative.

## Validation

- Synthetic fixtures reconcile raw and normalized totals within a documented tolerance.
- Correction and late-arrival scenarios are replayable without duplication.

## Revisit triggers

- A required semantic cannot be represented without corrupting its provider meaning.
