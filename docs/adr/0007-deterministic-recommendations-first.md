# ADR-0007: Prefer deterministic recommendations initially

- Status: Proposed
- Date: 2026-09-15
- Owners: Repository maintainers
- Decision scope: Recommendation engine
- Related issues: MAR-13
- Supersedes: None
- Superseded by: None

## Context

Opaque recommendations undermine trust and make safety validation difficult. The initial project has neither representative training data nor evidence that machine learning improves decision quality.

## Proposed decision

Begin with deterministic, versioned recommendation rules. Each rule exposes its evidence window, assumptions, exclusions, confidence derivation, and rollback requirements.

## Consequences

Behaviour is explainable and testable, although sophisticated workload patterns may initially evade detection.

## Validation

- Golden fixtures demonstrate true-positive, false-positive, insufficient-evidence, and reliability-veto scenarios.

## Revisit triggers

- A representative dataset demonstrates that a statistical model materially improves precision or recall while preserving explainability and safety.
