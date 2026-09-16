# ADR-0014: Use transparent confidence scoring

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Recommendation engine
- Related issues: MAR-13

## Context

A confidence percentage without a reproducible derivation creates false precision and weakens trust.

## Proposed decision

Calculate confidence from versioned factors: data completeness, observation duration, workload stability, telemetry quality, and corroborating recommendations. Persist both factors and final score.

## Consequences

Operators can challenge decisions and thresholds. Calibration becomes an ongoing evidence obligation.

## Guardrails

- Scores are deterministic for identical inputs and rule versions.
- Display every factor, weight, and missing signal.
- Confidence is not probability unless empirically calibrated as such.

## Validation

- Back-test labelled fixtures and publish calibration, precision, and abstention results.

## Revisit triggers

- Observed outcomes demonstrate systematic miscalibration or factor redundancy.
