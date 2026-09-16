# ADR-0023: Prefer workload identity and centralized secrets

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Identity and secrets
- Related issues: MAR-13

## Context

Static credentials increase disclosure risk, complicate rotation, and undermine workload attribution.

## Proposed decision

Use workload identity federation for cloud access. Store unavoidable GitHub App keys and database credentials in a supported secret manager and mount them ephemerally.

## Consequences

Most credentials become short-lived and attributable. Local development needs a documented, non-production authentication path.

## Guardrails

- Secrets never appear in Git, images, logs, traces, fixtures, or prompts.
- Rotation does not require an application rebuild.
- Authentication failure is fail-closed and observable.

## Validation

- Run secret scanning, rotation, expiration, revocation, and log-redaction tests.

## Revisit triggers

- The deployment environment cannot supply federated identity and an equivalently constrained alternative is demonstrated.
