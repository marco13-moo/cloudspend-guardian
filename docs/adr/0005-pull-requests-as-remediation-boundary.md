# ADR-0005: Use pull requests as the remediation boundary

- Status: Accepted
- Date: 2026-09-15
- Owners: Repository maintainers
- Decision scope: System
- Related issues: MAR-13
- Supersedes: None
- Superseded by: None

## Context

Recommendations become operationally useful only when translated into deployable changes. Direct cloud mutation would bypass code review, drift management, policy, and ordinary rollback controls.

## Decision

Produce proposed infrastructure-as-code patches through pull requests. Repository CI, policy evaluation, ownership rules, and human approval remain authoritative.

## Consequences

Remediations become reviewable and reproducible. Non-IaC resources cannot be remediated automatically, and time-to-savings includes repository workflow latency.

## Guardrails

- Every proposal includes evidence, projected savings, assumptions, and rollback instructions.
- The platform cannot merge its own production remediation.
- Protected-environment policy fails closed.

## Validation

- End-to-end tests will demonstrate proposal, policy rejection, approval, deployment, and rollback scenarios.

## Revisit triggers

- A resource class has no IaC representation and a safer independently approved mechanism is designed.
