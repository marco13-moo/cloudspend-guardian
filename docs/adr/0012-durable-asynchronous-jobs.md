# ADR-0012: Use durable asynchronous jobs for long-running workflows

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Execution
- Related issues: MAR-13

## Context

Ingestion, analysis, pull-request generation, and post-change measurement outlive request deadlines and must survive process interruption.

## Proposed decision

Represent workflows as durable jobs with explicit states, leases, idempotency keys, attempt history, deadlines, and cancellation semantics. Begin with a PostgreSQL-backed queue before introducing a broker.

## Consequences

The system gains recoverability without premature infrastructure. Database contention and polling require measurement.

## Guardrails

- Delivery is at least once; handlers are idempotent.
- Leases expire and are reclaimable.
- State transitions are monotonic and audited.

## Validation

- Kill workers at every transition and prove eventual completion without duplicate effects.

## Revisit triggers

- Measured throughput, isolation, or fan-out requirements exceed the database-backed design.
