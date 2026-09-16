# ADR-0017: Use GitHub App authentication

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Repository integration
- Related issues: MAR-13

## Context

Long-lived personal tokens have excessive lifetime, weak installation scoping, and poor organizational attribution.

## Proposed decision

Use a GitHub App with narrowly scoped permissions and short-lived installation tokens. Separate pull-request proposal permissions from any deployment authority.

## Consequences

Credential exposure has a smaller blast radius and actions are attributable. App installation and key rotation add operational work.

## Guardrails

- No permission to merge protected remediation pull requests.
- Installation scope is explicit per repository.
- Private keys reside only in the configured secret manager.

## Validation

- Prove allowed proposal operations and denied merge, administration, and unrelated-repository operations.

## Revisit triggers

- GitHub offers a more constrained workload identity mechanism with equivalent integration support.
