# ADR-0020: Maintain an append-only savings ledger

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Audit and measurement
- Related issues: MAR-13

## Context

Savings evidence may be recalculated or disputed. Mutable aggregates cannot establish what was known, approved, or measured at a given time.

## Proposed decision

Record immutable ledger events for finding, estimate, approval, deployment, observation, verification, reversal, and correction. Derive current views from those events.

## Consequences

Every claim gains temporal lineage. Consumers must understand projections and event-derived views.

## Guardrails

- Corrections append; they never rewrite history.
- Every event carries actor, organization, source, algorithm version, and timestamp.
- Aggregate reports are reproducible from ledger events.

## Validation

- Reconstruct historical views and verify tamper-evident hashes across correction and reversal scenarios.

## Revisit triggers

- Legal or privacy deletion requirements cannot be reconciled with the event design; apply cryptographic erasure or scoped tombstones through a superseding ADR.
