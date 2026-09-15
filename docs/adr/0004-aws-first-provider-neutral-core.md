# ADR-0004: Build AWS-first with a provider-neutral core

- Status: Accepted
- Date: 2026-09-15
- Owners: Repository maintainers
- Decision scope: System
- Related issues: MAR-13
- Supersedes: None
- Superseded by: None

## Context

Immediate multi-cloud implementation would multiply integrations before the product loop is validated. Embedding AWS concepts throughout the domain would make later extension unnecessarily expensive.

## Decision

Implement AWS as the first provider while keeping provider SDKs, identifiers, and billing terminology behind adapter boundaries. The domain expresses provider-neutral cost, resource, evidence, and recommendation concepts.

## Consequences

The first demonstrable outcome remains tractable, while portability is preserved structurally rather than claimed as an untested feature.

## Guardrails

- No AWS SDK types in `internal/domain`.
- Provider-specific data retains its provenance through normalization.
- Multi-cloud production readiness is not claimed until independently validated.

## Validation

- Static dependency tests will reject provider imports from domain packages.

## Revisit triggers

- A second provider is funded or necessary to validate the canonical model.
