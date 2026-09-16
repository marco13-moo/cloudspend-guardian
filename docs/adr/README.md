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
| [0008](0008-persistence-boundaries.md) | Separate operational state from immutable source artifacts | Proposed |
| [0009](0009-batch-ingestion-first.md) | Begin with scheduled batch ingestion | Proposed |
| [0010](0010-idempotent-replay-and-corrections.md) | Make ingestion replayable and correction-aware | Proposed |
| [0011](0011-synthetic-public-demonstration.md) | Use deterministic synthetic data for the public demonstration | Proposed |
| [0012](0012-durable-asynchronous-jobs.md) | Use durable asynchronous jobs for long-running workflows | Proposed |
| [0013](0013-require-utilization-and-slo-evidence.md) | Require utilization and SLO evidence | Proposed |
| [0014](0014-transparent-confidence-scoring.md) | Use transparent confidence scoring | Proposed |
| [0015](0015-deduplicate-recommendations.md) | Deduplicate provider and native recommendations | Proposed |
| [0016](0016-opa-policy-engine.md) | Use OPA for organizational policy | Proposed |
| [0017](0017-github-app-authentication.md) | Use GitHub App authentication | Proposed |
| [0018](0018-measure-realized-savings.md) | Measure realized savings separately | Proposed |
| [0019](0019-counterfactual-savings-baselines.md) | Use counterfactual savings baselines | Proposed |
| [0020](0020-append-only-savings-ledger.md) | Maintain an append-only savings ledger | Proposed |
| [0021](0021-continuous-threat-modeling.md) | Maintain a continuously reviewed threat model | Proposed |
| [0022](0022-read-only-cloud-identities.md) | Use read-only cloud identities | Proposed |
| [0023](0023-workload-identity-and-secrets.md) | Prefer workload identity and centralized secrets | Proposed |
| [0024](0024-single-organization-tenancy.md) | Begin with an explicit single-organization tenancy boundary | Proposed |
| [0025](0025-control-plane-slos.md) | Define measurable control-plane SLOs | Proposed |
| [0026](0026-opentelemetry-observability.md) | Standardize observability on OpenTelemetry | Proposed |
| [0027](0027-retry-and-dead-letter-semantics.md) | Define retries and poison-message isolation | Proposed |
| [0028](0028-backup-restore-and-disaster-recovery.md) | Prove backup, restore, and disaster recovery | Proposed |
| [0029](0029-progressive-deployment-model.md) | Adopt a progressive deployment model | Proposed |
| [0030](0030-software-supply-chain-controls.md) | Enforce software supply-chain controls | Proposed |
| [0031](0031-version-persistent-contracts.md) | Version every persistent contract | Proposed |
| [0032](0032-compatible-upgrades-and-rollback.md) | Require compatible upgrades and rehearsed rollback | Proposed |
| [0033](0033-architectural-fitness-functions.md) | Automate architectural fitness functions | Proposed |
| [0034](0034-evidence-driven-performance-boundaries.md) | Establish evidence-driven performance boundaries | Proposed |
| [0035](0035-data-retention-and-redaction.md) | Minimize, retain, and redact sensitive data deliberately | Proposed |
| [0036](0036-time-bounded-policy-exceptions.md) | Make policy exceptions attributable and time-bounded | Proposed |
| [0037](0037-vertical-slice-delivery.md) | Deliver through demonstrable vertical slices | Proposed |
| [0038](0038-production-readiness-certification.md) | Require evidence-based production-readiness certification | Proposed |
| [0039](0039-project-graduation-criteria.md) | Define explicit project graduation criteria | Proposed |

## AI-assisted delivery

AI agents must begin with [AI-GUIDE.md](AI-GUIDE.md), inspect the status of every relevant ADR, and preserve accepted guardrails. Proposed ADRs are hypotheses to validate—not authority to silently implement.

An ADR is required when a decision is costly to reverse, crosses a trust or ownership boundary, changes a persistent contract, or materially affects a quality attribute. Routine implementation minutiae do not warrant ADRs.
