# Vertical slice 003: durable job semantics

- Status: Implemented and locally validated
- Date: 2026-09-18
- Governing ADRs: 0012, 0027, 0037
- Scope: In-memory durable-state contract; no external queue or database

## Objective

Prove that long-running workflows have explicit, auditable state and that
transient failure recovery cannot turn malformed input into an infinite retry.

## Evidence

`go test ./internal/jobs` verifies:

- idempotency keys prevent duplicate jobs;
- leased work is reclaimable after worker expiry;
- transient failures return to pending with a next-attempt time;
- permanent failures enter dead-letter state with sanitized reason metadata;
- completion requires the current lease owner;
- every state transition is retained in job history.

## Architectural assessment

- **ADR-0012:** job identity, at-least-once delivery, leases, attempts, and
  monotonic audited transitions are represented. PostgreSQL persistence and
  worker crash-injection evidence remain outstanding.
- **ADR-0027:** failure classes, bounded attempts, backoff scheduling, and
  dead-letter isolation are represented. Jitter policy and operator replay API
  remain outstanding.

## Rollback

Revert the slice commit. No worker or external state is introduced yet.