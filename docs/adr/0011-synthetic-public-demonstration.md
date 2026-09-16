# ADR-0011: Use deterministic synthetic data for the public demonstration

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Demonstration and privacy
- Related issues: MAR-13

## Context

A public portfolio must be reproducible without exposing credentials, customer data, or an expensive cloud estate.

## Proposed decision

Ship versioned generators and fixtures representing normal spend, runaway scaling, idle resources, duplicate recommendations, corrections, and SLO vetoes.

## Consequences

Anyone can reproduce the product narrative. Synthetic performance and savings cannot be presented as production evidence.

## Guardrails

- Fixtures contain no transformed production data.
- Random generation uses published seeds.
- Synthetic outcomes are labelled conspicuously.

## Validation

- Execute the complete demo offline from a clean clone and verify deterministic hashes.

## Revisit triggers

- A legally distributable, anonymized benchmark dataset is independently reviewed and adds material value.
