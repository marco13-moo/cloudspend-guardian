# ADR-0003: Begin as a modular monolith

- Status: Accepted
- Date: 2026-09-15
- Owners: Repository maintainers
- Decision scope: System
- Related issues: MAR-13
- Supersedes: None
- Superseded by: None

## Context

The domain boundaries require discovery. Premature distribution would introduce deployment, consistency, and observability complexity before independent scaling or ownership needs exist.

## Decision

Implement an explicitly modular Go control plane deployed as one process. Modules communicate through narrow interfaces and retain ownership of their domain concepts.

## Consequences

Local development and transactional consistency remain simple. Module discipline must be enforced without network boundaries.

## Guardrails

- Transport types do not become domain models.
- Provider SDK objects terminate at adapter boundaries.
- Circular package dependencies are prohibited.

## Validation

- Go's import-cycle prohibition and future dependency fitness tests enforce boundaries.

## Revisit triggers

- A component requires independent scaling, failure isolation, deployment cadence, or ownership.
