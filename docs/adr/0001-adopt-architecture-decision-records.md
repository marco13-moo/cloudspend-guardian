# ADR-0001: Adopt architecture decision records

- Status: Accepted
- Date: 2026-09-15
- Owners: Repository maintainers
- Decision scope: System
- Related issues: MAR-13
- Supersedes: None
- Superseded by: None

## Context

CloudSpend Guardian will reconcile financial, reliability, security, and developer-experience concerns. These forces will produce consequential choices whose rationale must remain reviewable after the original context disappears.

## Decision

Maintain numbered Markdown ADRs in `docs/adr`. Accepted decisions are immutable; changed decisions are recorded by explicit supersession. Each ADR includes alternatives, consequences, guardrails, validation, and measurable revisit triggers.

## Consequences

Decision-making becomes auditable and implementation can be tested against explicit invariants. Maintainers incur a modest documentation obligation for consequential changes.

## Guardrails

- Trust-boundary, persistent-contract, and major quality-attribute changes require an ADR.
- ADRs cannot be used to conceal unresolved disagreement or absent evidence.
- Implementation pull requests identify the ADRs they implement or supersede.

## Validation

- CI will eventually verify ledger links and required headings.

## Revisit triggers

- The format produces repetitive records without improving review quality.
