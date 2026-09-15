# Architecture decision ledger

ADRs are immutable after acceptance except for status metadata and typographical corrections. A changed decision receives a new ADR that explicitly supersedes its predecessor.

## Status vocabulary

- **Proposed:** awaiting evidence or review.
- **Accepted:** authoritative for implementation.
- **Rejected:** evaluated and deliberately declined.
- **Deprecated:** retained but no longer preferred.
- **Superseded:** replaced by a named later ADR.

## Current records

| ADR | Decision | Status |
|---|---|---|
| [0001](0001-adopt-architecture-decision-records.md) | Adopt architecture decision records | Accepted |
| [0002](0002-governed-control-plane-boundary.md) | Establish a governed control-plane boundary | Accepted |
| [0003](0003-modular-monolith-first.md) | Begin as a modular monolith | Accepted |
| [0004](0004-aws-first-provider-neutral-core.md) | Build AWS-first with a provider-neutral core | Accepted |
| [0005](0005-pull-requests-as-remediation-boundary.md) | Use pull requests as the remediation boundary | Accepted |
| [0006](0006-focus-compatible-cost-model.md) | Adopt a FOCUS-compatible canonical cost model | Proposed |
| [0007](0007-deterministic-recommendations-first.md) | Prefer deterministic recommendations initially | Proposed |

## Planned decision sequence

The project will create additional records only when implementation evidence makes the decision timely:

1. operational state and immutable raw-artifact storage;
2. batch versus streaming ingestion;
3. replay and correction semantics;
4. synthetic-fixture isolation;
5. observability and SLO evidence requirements;
6. recommendation confidence scoring;
7. provider-recommendation deduplication;
8. policy-engine selection;
9. GitHub App authentication;
10. counterfactual savings baselines;
11. append-only savings ledger;
12. cloud identity and secret management;
13. tenancy boundary;
14. OpenTelemetry and operational SLOs;
15. retries, dead letters, backup and restore;
16. deployment, supply-chain, versioning and rollback strategy;
17. architectural fitness functions and graduation criteria.

An ADR is required when a decision is costly to reverse, crosses a trust or ownership boundary, changes a persistent contract, or materially affects a quality attribute. Routine implementation minutiae do not warrant ADRs.
