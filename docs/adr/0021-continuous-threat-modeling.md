# ADR-0021: Maintain a continuously reviewed threat model

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Security
- Related issues: MAR-13

## Context

New providers, credentials, AI assistance, repository writes, and multi-organization scope continually change the attack surface.

## Proposed decision

Treat `docs/architecture/threat-model.md` as a version-controlled security artifact reviewed at every new trust boundary and release gate.

## Consequences

Security assumptions remain visible and testable. Feature work includes explicit threat-analysis effort.

## Guardrails

- Resource names, tags, issue text, and model output are untrusted data.
- Threat controls link to tests or operational evidence.
- Residual risks have owners and review dates.

## Validation

- Run abuse-case reviews for credential theft, forged billing, recommendation tampering, cross-scope access, and prompt injection.

## Revisit triggers

- A formal compliance regime mandates a different threat-model methodology.
