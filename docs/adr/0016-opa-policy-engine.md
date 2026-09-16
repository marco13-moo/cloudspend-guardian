# ADR-0016: Use OPA for organizational policy

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Policy
- Related issues: MAR-13

## Context

Authorization and remediation governance must evolve independently from recommendation calculations and remain testable by operators.

## Proposed decision

Use Open Policy Agent and versioned Rego bundles for organizational decisions. Keep financial calculations and domain validation in Go; send bounded, schema-versioned inputs to policy.

## Consequences

Guardrails become declarative and portable. The project acquires another language, runtime, and bundle lifecycle.

## Guardrails

- Policy defaults deny protected or malformed operations.
- Bundles are signed, versioned, and attributable.
- Policy cannot manufacture missing evidence.

## Validation

- Test allow, deny, malformed input, protected environment, expired exception, and rollback scenarios.

## Revisit triggers

- Policy authors cannot safely operate Rego or another engine demonstrably improves correctness and maintainability.
