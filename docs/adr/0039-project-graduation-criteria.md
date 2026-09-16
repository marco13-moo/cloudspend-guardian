# ADR-0039: Define explicit project graduation criteria

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Product completion
- Related issues: MAR-13

## Context

A portfolio project without falsifiable completion criteria can remain perpetually scaffolded or overstate maturity.

## Proposed decision

Graduate the initial project only when a clean clone can reproduce the complete synthetic cost-anomaly loop and its principal safety, recovery, and measurement evidence.

## Graduation evidence

- Reconciled FOCUS-compatible ingestion with correction and replay.
- SLO-aware, explainable recommendation with calibrated confidence.
- Deduplicated provider evidence and OPA policy decision.
- Repository-scoped Terraform pull-request proposal with rollback metadata.
- Projected and counterfactual realized-savings ledger entries.
- Read-only identity, threat-model, redaction, and cross-scope tests.
- Observable retry, restore, upgrade, rollback, and failure-injection exercises.
- Signed release, SBOM, provenance, Helm deployment, and published limitations.

## Consequences

The project has an honest finish line and an employer-reviewable demonstration. Optional multi-cloud, SaaS tenancy, and machine learning remain future work rather than implied capability.

## Guardrails

- No criterion is satisfied solely by documentation or mocked success output.
- Unavailable live infrastructure is reported as unverified.
- Graduation identifies residual risks and unsupported scenarios.

## Validation

- A reviewer follows the published runbook from a clean environment and reproduces evidence linked to the release commit.

## Revisit triggers

- The intended audience, product boundary, or initial release objective changes materially.
