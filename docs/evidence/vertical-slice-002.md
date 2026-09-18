# Vertical slice 002: replayable batch ingestion

- Status: Implemented and locally validated
- Date: 2026-09-18
- Governing ADRs: 0008, 0009, 0010, 0037
- Scope: In-memory operational boundary; no external credentials or cloud state

## Objective

Prove that a bounded billing import can retain its immutable source, be replayed
without duplicating state, and preserve correction lineage.

## Implemented path

```text
bounded source artifact
  -> validated FOCUS-compatible report
  -> content-addressed artifact retention
  -> scoped canonical record versions
  -> correction lineage or replay no-op
  -> completion cursor advanced after successful application
```

## Evidence

`go test ./internal/ingestion` verifies:

- identical source replay is a no-op;
- source bytes and schema version are retained by SHA-256;
- a changed record creates version 2 and retains its predecessor hash;
- invalid input does not advance the scope cursor;
- current state exposes one latest version per canonical source record.

## Architectural assessment

- **ADR-0008:** the store models the separation between immutable artifacts and
  operational state, but PostgreSQL and object-storage adapters remain outstanding.
- **ADR-0009:** bounded windows, explicit scope, completion time, and cursors are
  represented; scheduling and freshness telemetry remain outstanding.
- **ADR-0010:** deterministic identity, replay no-op, correction replacement,
  and predecessor lineage are implemented and tested. Crash recovery across
  external systems remains outstanding.

## Residual limitations

This is a local contract implementation. It is not a production persistence
deployment and does not yet provide durable asynchronous jobs, provider-specific
correction adapters, or transactional coordination across PostgreSQL and object
storage.

## Rollback

Revert the slice commit. It introduces no external state or persistent contract
consumed by a released version.