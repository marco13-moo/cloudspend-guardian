# ADR-0029: Adopt a progressive deployment model

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Distribution
- Related issues: MAR-13

## Context

Evaluators need a low-friction local path while production-like validation requires Kubernetes and reproducible infrastructure.

## Proposed decision

Support three explicit tiers: local process/Compose, Kubernetes through Helm, and a Terraform-managed AWS reference deployment. Do not promise behavioural parity without conformance tests.

## Consequences

The project becomes approachable and operationally demonstrable. Packaging and compatibility matrices require maintenance.

## Guardrails

- One configuration vocabulary spans tiers.
- Secure defaults do not weaken for convenience.
- Each tier states supported and unsupported capabilities.

## Validation

- Execute the same synthetic acceptance suite in every supported tier.

## Revisit triggers

- A tier has no users or its maintenance cost exceeds demonstrated value.
