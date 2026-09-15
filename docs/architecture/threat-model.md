# Threat model

## Assets

- Cloud billing and utilization data
- Provider and GitHub credentials
- Recommendation and approval integrity
- Infrastructure-as-code repositories
- Savings and audit records

## Principal threats

| Threat | Consequence | Initial control |
|---|---|---|
| Compromised collector | Cloud reconnaissance or mutation | Read-only, least-privilege identity |
| Forged billing record | Misleading recommendation | Provenance, schema validation, reconciliation |
| Malicious resource metadata | Instruction or content injection | Treat metadata as inert, untrusted data |
| Recommendation tampering | Unsafe infrastructure patch | Signed audit trail, policy and CI validation |
| GitHub credential theft | Repository compromise | Short-lived, repository-scoped GitHub App token |
| Cross-account data leakage | Confidentiality breach | Explicit account scope in every storage and query key |
| Double-counted savings | Fraudulent portfolio evidence | Deduplication and immutable measurement lineage |
| Availability attack | Delayed cost detection | Bounded retries, dead-letter isolation, replayable ingestion |

## Explicit non-goals for the scaffold

- Production-grade multi-tenancy
- Autonomous infrastructure mutation
- Storage of real customer billing information
- Machine-learning-generated remediation

The threat model must be revisited before adding real cloud credentials, external contributors, or autonomous execution.
