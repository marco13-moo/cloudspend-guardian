# ADR-0008: Separate operational state from immutable source artifacts

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Data
- Related issues: MAR-13

## Context

Transactional workflow state and voluminous provider billing exports have dissimilar consistency, retention, and access patterns.

## Proposed decision

Use PostgreSQL for normalized operational state and object storage for immutable raw billing artifacts. Store content hashes and provenance in PostgreSQL; never treat object storage as the transactional authority.

## Consequences

The design gains relational integrity and inexpensive replayable evidence, at the cost of coordinating two persistence systems.

## Guardrails

- Raw artifacts are immutable and content-addressed.
- Database rows retain organization and source scope.
- Derived records always reference their source artifact and schema version.

## Validation

- Rebuild normalized state from retained artifacts and reconcile totals.

## Revisit triggers

- Operational volume exceeds measured PostgreSQL boundaries or object storage is unavailable in a supported deployment.
