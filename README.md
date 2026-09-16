# CloudSpend Guardian

CloudSpend Guardian is an evidence-driven FinOps control plane that converts cloud-cost signals into safe, reviewable infrastructure changes.

Instead of deleting resources or mutating production directly, the platform assembles utilization evidence, evaluates policy, and proposes remediations through pull requests. Every recommendation carries its assumptions, confidence, projected savings, rollback plan, and—after deployment—measured savings.

> **Status:** architectural scaffold. The initial vertical slice exposes health/readiness endpoints and codifies the safety boundary for future recommendation and remediation workflows.

## Why this project exists

Most cost tools stop at visibility. CloudSpend Guardian is intended to demonstrate the harder platform-engineering problem: reconciling financial efficiency with reliability, security, developer autonomy, and auditable governance.

The target workflow is:

```text
FOCUS-compatible billing + utilization + Terraform
                         │
                         ▼
               normalized cost model
                         │
                         ▼
             recommendation and policy
                         │
                         ▼
             reviewable Terraform pull request
                         │
                         ▼
             deployment and savings evidence
```

## Architectural invariants

- Cloud credentials are read-only.
- The control plane never deletes or mutates production resources directly.
- Remediation crosses a pull-request and human-approval boundary.
- No recommendation is emitted without evidence, assumptions, confidence, and rollback metadata.
- Projected savings and realized savings are recorded separately.
- Reliability guardrails can veto an otherwise economical recommendation.

## Repository map

| Path | Responsibility |
|---|---|
| `cmd/server` | Control-plane process entry point |
| `internal/api` | HTTP transport and operational endpoints |
| `internal/config` | Environment-derived runtime configuration |
| `internal/domain` | Provider-neutral recommendation model |
| `policy` | Policy-as-code guardrails |
| `docs/architecture` | System boundaries, quality attributes, and threat model |
| `docs/adr` | Immutable architecture decision ledger |
| `deploy/helm` | Kubernetes packaging |

## Quick start

Prerequisites: Go 1.27+ and Docker.

```bash
make verify
make run
curl http://localhost:8080/healthz
curl http://localhost:8080/readyz
```

Or run the containerized service:

```bash
docker compose up --build
```

## Roadmap

1. Ingest deterministic synthetic FOCUS billing fixtures.
2. Reconcile raw and normalized billing totals.
3. Detect an intentionally injected cost anomaly.
4. Evaluate the recommendation through OPA policy.
5. Generate a reviewable Terraform remediation patch.
6. Record projected and realized savings without double counting.

The complete 39-decision sequence is maintained in [the ADR index](docs/adr/README.md). AI-assisted implementation begins with the [AI architecture execution guide](docs/adr/AI-GUIDE.md), which defines authority, phase gates, evidence requirements, and non-negotiable prohibitions.

## Security posture

This repository must contain no cloud credentials or customer billing data. Development fixtures will be synthetic and deterministic. See [SECURITY.md](SECURITY.md) and the [threat model](docs/architecture/threat-model.md).

## License

Apache-2.0. See [LICENSE](LICENSE).
