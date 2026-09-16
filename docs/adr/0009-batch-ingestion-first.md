# ADR-0009: Begin with scheduled batch ingestion

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Ingestion
- Related issues: MAR-13

## Context

Cloud billing records commonly arrive with latency and corrections. Streaming would not make source data instantaneous but would add coordination complexity.

## Proposed decision

Run bounded hourly or daily imports with explicit billing windows, cursors, reconciliation, and freshness telemetry.

## Consequences

Ingestion remains comprehensible and replayable. Detection latency is bounded by provider availability and the configured schedule.

## Guardrails

- Freshness is exposed as data, never implied.
- Overlapping import windows are safe.
- Scheduling is independent from import execution.

## Validation

- Demonstrate initial, overlapping, late-arriving, and corrected imports.

## Revisit triggers

- A provider exposes reliable event delivery and a validated use case requires materially lower latency.
