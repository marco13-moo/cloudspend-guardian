# Quality attributes

The project optimizes for the following attributes in descending order.

1. **Safety:** a recommendation cannot directly mutate or delete production infrastructure.
2. **Explainability:** every recommendation exposes evidence, assumptions, confidence, and calculation.
3. **Auditability:** decisions, approvals, deployments, and measured outcomes remain attributable.
4. **Reliability:** optimization must not violate workload SLOs or recovery objectives.
5. **Portability:** provider-specific integrations terminate at explicit adapter boundaries.
6. **Operability:** failure modes are observable, replayable, and recoverable.
7. **Performance:** throughput is optimized only after the preceding attributes are preserved.

## Initial fitness functions

- Reject a recommendation without evidence or rollback metadata.
- Reject direct-cloud-mutation and destructive-remediation requests.
- Reconcile normalized billing totals with their raw source totals.
- Require idempotency for ingestion and asynchronous jobs.
- Prevent cloud write permissions in deployable identity policies.
- Record projected and realized savings as separate measurements.
