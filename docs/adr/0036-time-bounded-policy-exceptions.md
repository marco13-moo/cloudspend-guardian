# ADR-0036: Make policy exceptions attributable and time-bounded

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Governance
- Related issues: MAR-13

## Context

Permanent undocumented exceptions gradually nullify safety and cost controls.

## Proposed decision

Represent exceptions as reviewed policy data containing scope, justification, owner, approver, creation time, expiry, and compensating controls.

## Consequences

Exceptional delivery remains possible without normalizing unmanaged bypasses. Owners incur renewal and remediation obligations.

## Guardrails

- Exceptions fail closed after expiry.
- Wildcard organization or production scope requires a separately reviewed ADR.
- Security invariants explicitly marked non-exemptible cannot be bypassed.

## Validation

- Test active, expired, malformed, over-broad, unauthorized, and revoked exceptions.

## Revisit triggers

- Exception frequency reveals an invalid base policy requiring redesign.
