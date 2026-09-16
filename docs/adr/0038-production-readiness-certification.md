# ADR-0038: Require evidence-based production-readiness certification

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Operational governance
- Related issues: MAR-13

## Context

Passing unit tests does not establish deployability, recoverability, least privilege, or truthful savings measurement.

## Proposed decision

Gate production-ready claims on a versioned certification suite covering identity, policy, deployment, upgrade, rollback, backup, restore, failure injection, SLO telemetry, reconciliation, and audit lineage.

## Consequences

Readiness becomes reviewable evidence rather than rhetoric. Certification adds environments, fixtures, execution time, and residual-risk reporting.

## Guardrails

- Missing infrastructure or external dependencies are coverage gaps, not passing results.
- Evidence identifies commit, artifact, environment, time, and tool versions.
- Waivers follow ADR-0036 and appear in the certification summary.

## Validation

- Independently reproduce certification from a signed release and compare evidence hashes.

## Revisit triggers

- Deployment scope or threat model changes the required evidence matrix.
