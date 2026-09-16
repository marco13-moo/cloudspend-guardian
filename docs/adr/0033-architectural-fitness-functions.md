# ADR-0033: Automate architectural fitness functions

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Architecture governance
- Related issues: MAR-13

## Context

Written constraints decay when violations are cheap and detection depends on reviewer memory.

## Proposed decision

Encode critical ADR guardrails as executable CI checks covering dependencies, permissions, schemas, recommendation completeness, policy, and savings lineage.

## Consequences

Architecture becomes continuously verifiable. Poorly designed checks can fossilize implementation rather than protect intent.

## Guardrails

- Test invariants, not arbitrary file layouts.
- Every fitness function cites its governing ADR.
- Exceptions follow ADR-0036 and cannot silently disable controls.

## Validation

- Mutation tests deliberately violate each invariant and prove CI rejection.

## Revisit triggers

- A check creates recurrent false positives or no longer represents its ADR's intent.
