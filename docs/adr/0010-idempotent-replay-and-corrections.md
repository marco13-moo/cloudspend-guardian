# ADR-0010: Make ingestion replayable and correction-aware

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Ingestion
- Related issues: MAR-13

## Context

Retries, late records, provider corrections, and operator-initiated replays are normal billing behaviours, not exceptional cases.

## Proposed decision

Assign each source record a deterministic identity derived from provider scope, billing period, source key, and version. Apply imports transactionally and represent corrections through versioned replacement lineage.

## Consequences

Recovery becomes routine and duplicates become detectable. Identity and correction semantics require provider-specific tests.

## Guardrails

- Reprocessing identical input produces identical state.
- Corrections never destroy their predecessor's audit lineage.
- Partial imports do not advance completion cursors.

## Validation

- Replay the same fixture repeatedly and inject conflicting corrections, crashes, and reordered delivery.

## Revisit triggers

- Provider semantics cannot be represented without a different identity or temporal model.
