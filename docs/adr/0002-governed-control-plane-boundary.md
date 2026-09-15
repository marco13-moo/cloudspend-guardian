# ADR-0002: Establish a governed control-plane boundary

- Status: Accepted
- Date: 2026-09-15
- Owners: Repository maintainers
- Decision scope: System
- Related issues: MAR-13
- Supersedes: None
- Superseded by: None

## Context

Autonomous deletion could maximize nominal savings while creating disproportionate operational and security risk. A passive dashboard, conversely, would not demonstrate an end-to-end engineering outcome.

## Decision

CloudSpend Guardian will be an advisory and governance control plane. It may ingest, analyze, prioritize, and propose remediations, but it will not mutate production cloud resources directly.

## Consequences

Changes remain auditable through ordinary repository controls and rollback procedures. Savings realization is slower and depends on an external deployment workflow.

## Guardrails

- Cloud access is read-only.
- Destructive remediation is denied.
- Production changes require policy, CI, and human approval.

## Validation

- Policy tests reject destructive and direct-mutation inputs.
- Deployment identity tests reject cloud write permissions.

## Revisit triggers

- None before a separately threat-modeled, narrowly scoped execution mechanism is proposed.
