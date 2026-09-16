# ADR-0035: Minimize, retain, and redact sensitive data deliberately

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Privacy and data governance
- Related issues: MAR-13

## Context

Billing exports, resource metadata, repository details, and telemetry can expose commercially sensitive information even without conventional personal data.

## Proposed decision

Classify data by sensitivity and purpose, collect the minimum fields required, define retention per class, redact at ingestion and telemetry boundaries, and support scoped deletion obligations.

## Consequences

Disclosure impact and storage cost decrease. Some forensic and analytical capability expires intentionally.

## Guardrails

- Raw metadata is never logged by default.
- Retention jobs are observable, idempotent, and scope-aware.
- Public fixtures contain no derived private data.

## Validation

- Test redaction, retention expiry, scoped deletion, backup expiry, and log/trace leakage.

## Revisit triggers

- Legal, contractual, or product requirements change classification or retention periods.
