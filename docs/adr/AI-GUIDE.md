# AI architecture execution guide

This document is the entry point for any AI agent modifying CloudSpend Guardian.

## Mandatory reading order

1. `README.md` and `docs/architecture/quality-attributes.md`.
2. `docs/architecture/system-context.md` and `docs/architecture/threat-model.md`.
3. This ledger and every ADR relevant to the requested change.
4. Current tests, policies, and deployment manifests touching the change boundary.

## Authority model

- **Accepted ADRs are architectural constraints.** Preserve their guardrails or create a proposed superseding ADR.
- **Proposed ADRs are falsifiable hypotheses.** Gather the validation evidence named in the ADR before changing the status to Accepted.
- **Rejected, Deprecated, and Superseded ADRs are historical evidence.** Do not revive them implicitly.
- An agent must never change an ADR status merely because implementation exists; validation and consequences must be reviewed explicitly.

## Delivery protocol

For every substantive change:

1. Identify the governing ADRs and phase gate.
2. State assumptions and unresolved decisions.
3. Implement the smallest end-to-end vertical slice consistent with accepted guardrails.
4. Add tests for success, failure, authorization, replay, and rollback where applicable.
5. Record evidence against each ADR's Validation section.
6. Update architecture documentation when boundaries or flows change.
7. Report unavailable checks as coverage gaps rather than successes.

## Phase gates

| Phase | Governing ADRs | Exit evidence |
|---|---|---|
| Foundation | 0001–0007 | Domain invariants, provider boundary, explainable recommendation fixture |
| Ingestion | 0008–0012 | Reconciled, replayable synthetic FOCUS import and durable job semantics |
| Recommendation | 0013–0016 | SLO-aware recommendation, confidence explanation, deduplication, policy decision |
| Remediation | 0005, 0017 | Repository-scoped pull-request proposal with rollback metadata |
| Measurement | 0018–0020 | Projected and realized savings with counterfactual lineage |
| Security | 0021–0024 | Threat review, read-only identity proof, secret isolation, scope tests |
| Operations | 0025–0028 | SLO telemetry, failure injection, restore and recovery evidence |
| Distribution | 0029–0032 | Reproducible deployment, signed artifacts, compatible upgrade and rollback |
| Graduation | 0033–0039 | Fitness, performance, privacy, exception, readiness, and demonstration evidence |

## Non-negotiable prohibitions

- No direct production-cloud mutation.
- No destructive remediation.
- No real credentials or customer billing data in fixtures, logs, prompts, or commits.
- No recommendation without evidence and rollback metadata.
- No savings claim without a measurement window and counterfactual baseline.
- No architecture status inflation: incomplete or unavailable evidence remains explicit.
