# ADR-0024: Begin with an explicit single-organization tenancy boundary

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Tenancy
- Related issues: MAR-13

## Context

Production multi-tenancy requires pervasive isolation and adversarial validation. A personal project should not claim guarantees it has not proven.

## Proposed decision

Support one organization per deployment initially while including organization scope in persistent identities, jobs, metrics, and authorization interfaces.

## Consequences

The initial system is simpler and honest. Shared SaaS economics and cross-tenant administration are deferred.

## Guardrails

- No implicit global resource lookup.
- Organization identifiers are immutable scope, not presentation metadata.
- Documentation states the single-organization limitation.

## Validation

- Scope tests reject mismatched organization identifiers at every repository and storage boundary.

## Revisit triggers

- A real multi-organization use case funds isolation design, migration, and penetration testing.
