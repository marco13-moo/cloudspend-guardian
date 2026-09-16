# ADR-0022: Use read-only cloud identities

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Cloud security
- Related issues: MAR-13

## Context

The control plane needs billing, inventory, recommendation, and telemetry visibility but does not need cloud mutation authority.

## Proposed decision

Issue dedicated least-privilege collector identities per organization and environment. Express permissions as reviewed infrastructure as code and deny mutating actions.

## Consequences

A compromised collector cannot directly alter resources. Some provider recommendations may require additional read scopes that must be reviewed explicitly.

## Guardrails

- No wildcard write actions.
- Separate demonstration, development, and production identities.
- Continuously analyze effective permissions, not only declared policies.

## Validation

- Positive tests enumerate required reads; negative tests attempt representative mutations and unrelated-account access.

## Revisit triggers

- A future feature proposes mutation; it requires a new trust boundary and superseding ADR.
