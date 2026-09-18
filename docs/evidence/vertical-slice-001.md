# Vertical slice 001: synthetic cost recommendation

- Status: Implemented and locally validated
- Date: 2026-09-18
- Governing ADRs: 0004, 0006, 0007, 0010, 0011, 0013, 0014, 0037
- Scope: Offline synthetic evidence only

## Objective

Prove the smallest safe product loop from a deterministic billing artifact to an explainable recommendation and a deliberate non-recommendation.

## Implemented path

```text
seeded synthetic records
  → FOCUS-compatible subset importer
  → canonical identities and source SHA-256
  → integer-minor-unit reconciliation
  → idle-compute/v1 deterministic rule
  → recommendation or explicit abstention
```

## Evidence

`make demo` produced:

- schema `focus-subset/v1`;
- source SHA-256 `618aafc12978658c14c030dcdfa40a31303a892110628c67e8c21ef88c0b3bf7`;
- 28 imported and reconciled records;
- total synthetic billed cost of EUR 105.00;
- one recommendation for `arn:synthetic:ec2:idle-1`;
- projected—not realized—monthly savings of EUR 37.50;
- transparent confidence `0.979375` with four named factors;
- two evidence statements, two assumptions, two exclusions, and rollback metadata;
- one abstention for `arn:synthetic:ec2:active-1` because it is not underutilized.

`make verify` passed formatting, vetting, race-enabled tests, compilation, and package coverage. Tests exercise deterministic replay, duplicate rejection, missing columns, lossy currency precision, insufficient duration, unreconciled input, successful recommendation, and reliability veto.

## Architectural assessment

- **ADR-0004:** provider SDK types do not enter the domain model.
- **ADR-0006:** partial evidence exists for canonical mapping and monetary reconciliation; correction semantics remain incomplete.
- **ADR-0007:** golden evidence covers positive, insufficient-evidence, and reliability-veto paths; broader false-positive calibration remains incomplete.
- **ADR-0011:** deterministic generation is demonstrated locally; clean-clone verification remains outstanding.
- **ADR-0013:** duration, utilization, memory, and error-budget evidence gate recommendations.
- **ADR-0014:** confidence factors are deterministic and visible, but empirical calibration remains outstanding.
- **ADR-0037:** the first end-to-end slice is reproducible locally without cloud credentials.

## Residual limitations

- This is a FOCUS-compatible subset, not full FOCUS conformance.
- Persistence, correction lineage, durable jobs, OPA, GitHub Apps, Terraform remediation, and realized-savings measurement are not implemented.
- The utilization fields are synthetic extensions to the billing record for this first slice; a later adapter must source them from observability systems.
- No proposed ADR is promoted to Accepted by this evidence alone.

## Rollback

Revert the slice commit. It introduces no external state, cloud identity, database migration, or persistent contract consumed by a released version.
