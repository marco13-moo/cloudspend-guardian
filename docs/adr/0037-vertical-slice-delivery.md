# ADR-0037: Deliver through demonstrable vertical slices

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Delivery strategy
- Related issues: MAR-13

## Context

Building horizontal layers for months can produce impressive infrastructure without proving that the product closes the cost-optimization loop.

## Proposed decision

Deliver incremental end-to-end demonstrations: synthetic import, reconciled finding, policy decision, Terraform proposal, and verified outcome. Each slice preserves production-intent boundaries while minimizing scope.

## Consequences

Value and architectural assumptions are tested early. Temporary adapters may be needed to avoid prematurely completing adjacent systems.

## Guardrails

- No slice bypasses accepted safety controls.
- Demonstrations distinguish simulated from live evidence.
- Each slice has a documented rollback and observable completion criterion.

## Validation

- A clean-clone script reproduces every declared slice and its failure scenarios.

## Revisit triggers

- A foundational migration must precede an end-to-end slice and is explicitly justified.
