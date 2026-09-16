# ADR-0031: Version every persistent contract

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Compatibility
- Related issues: MAR-13

## Context

Billing schemas, job payloads, policy inputs, APIs, ledger events, algorithms, and fixtures outlive individual deployments.

## Proposed decision

Assign explicit versions to every persisted or externally consumed contract and publish compatibility and migration rules.

## Consequences

Historical evidence remains interpretable and upgrades become deliberate. Maintainers must support or migrate declared versions.

## Guardrails

- Unknown versions fail visibly rather than being guessed.
- Algorithm version accompanies every recommendation and savings calculation.
- Breaking API changes require a new major contract version.

## Validation

- Maintain compatibility fixtures for the oldest supported version through current.

## Revisit triggers

- A contract is proven ephemeral and has no persistence, replay, or external consumers.
