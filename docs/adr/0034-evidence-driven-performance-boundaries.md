# ADR-0034: Establish evidence-driven performance boundaries

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Performance
- Related issues: MAR-13

## Context

Premature optimization obscures correctness, while absent capacity evidence makes production claims indefensible.

## Proposed decision

Define representative workloads and benchmark ingestion throughput, normalization cost, job latency, query latency, database growth, and recommendation duration before setting capacity limits.

## Consequences

Scaling decisions follow measurements rather than fashion. Benchmark fixtures and environments require versioning.

## Guardrails

- Publish hardware, dataset, concurrency, and percentile methodology.
- Correctness and reconciliation are checked during load tests.
- Benchmark results are not generalized beyond their tested envelope.

## Validation

- Run repeatable baselines, saturation tests, soak tests, and regression thresholds in controlled environments.

## Revisit triggers

- Workload shape, service objectives, or deployment architecture changes materially.
