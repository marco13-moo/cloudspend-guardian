# ADR-0019: Use counterfactual savings baselines

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Measurement
- Related issues: MAR-13

## Context

Simple before-and-after comparison attributes demand changes, seasonality, price changes, and unrelated deployments to the optimization.

## Proposed decision

Estimate the cost that would have occurred without remediation using documented, versioned baselines normalized by appropriate demand units.

## Consequences

Savings claims become more defensible but remain estimates with uncertainty and data requirements.

## Guardrails

- Publish baseline method, comparison cohort, excluded events, and uncertainty.
- Abstain when a defensible counterfactual cannot be established.
- Preserve original observations for later recalculation.

## Validation

- Evaluate baselines against withheld historical periods and synthetic interventions.

## Revisit triggers

- Forecast error exceeds the agreed threshold or a workload's economics require another model.
