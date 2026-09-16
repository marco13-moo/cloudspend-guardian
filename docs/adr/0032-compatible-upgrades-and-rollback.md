# ADR-0032: Require compatible upgrades and rehearsed rollback

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Release engineering
- Related issues: MAR-13

## Context

Database migrations, concurrent workers, policies, and queued jobs create mixed-version periods where naive rollback can corrupt state.

## Proposed decision

Use expand-migrate-contract database changes, backward-compatible job readers, versioned policy bundles, and explicitly rehearsed application and data rollback procedures.

## Consequences

Deployments become safer but some changes require multiple releases and delayed cleanup.

## Guardrails

- A release cannot depend on a destructive schema contraction until every old reader is retired.
- Rollback compatibility is tested before promotion.
- Irreversible migrations require backup and restoration evidence.

## Validation

- Run old-new, new-old, interrupted migration, queued-old-message, and policy-rollback scenarios.

## Revisit triggers

- Deployment topology changes the concurrency or compatibility model.
