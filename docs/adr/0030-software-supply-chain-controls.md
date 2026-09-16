# ADR-0030: Enforce software supply-chain controls

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Supply chain
- Related issues: MAR-13

## Context

The platform processes sensitive cost data and proposes infrastructure changes; compromised build inputs or artifacts would undermine every higher-level control.

## Proposed decision

Pin CI dependencies, minimize permissions, generate SBOMs and provenance, scan dependencies and images, sign releases, and verify artifacts before deployment.

## Consequences

Artifacts become traceable and policy-enforceable. Release engineering becomes more elaborate and scanners may introduce triage work.

## Guardrails

- Release publication uses short-lived identity.
- Critical unresolved vulnerabilities block release unless a dated, attributable exception exists.
- Build provenance identifies the source commit and workflow.

## Validation

- Verify signatures, provenance, SBOM completeness, least-privilege workflow permissions, and tampered-artifact rejection.

## Revisit triggers

- Standards or platform capabilities materially improve the attestation model.
