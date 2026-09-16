# ADR-0028: Prove backup, restore, and disaster recovery

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Resilience
- Related issues: MAR-13

## Context

Backups are unverified assumptions until restoration and reconciliation succeed under realistic failure conditions.

## Proposed decision

Define RPO and RTO for PostgreSQL, raw artifacts, policies, and audit events. Automate encrypted backups and perform scheduled restore exercises into isolated environments.

## Consequences

Recovery claims become evidence-based. Storage, rehearsal, and key-recovery procedures add cost.

## Guardrails

- Backup credentials are separate from runtime credentials.
- Restores verify referential integrity, artifact hashes, and monetary reconciliation.
- Recovery evidence records duration, data loss, and exceptions.

## Validation

- Exercise database loss, object loss, partial corruption, key unavailability, and regional outage scenarios.

## Revisit triggers

- Business impact analysis changes recovery objectives or architecture.
