# ADR-0027: Define retries and poison-message isolation

- Status: Proposed
- Date: 2026-09-15
- Decision scope: Reliability
- Related issues: MAR-13

## Context

Blind retries amplify outages, duplicate effects, and conceal permanently malformed input.

## Proposed decision

Classify failures as transient, rate-limited, permanent, conflict, or operator-actionable. Apply bounded exponential backoff with jitter and isolate terminal records with replay metadata.

## Consequences

Recovery is predictable and malformed data cannot indefinitely block a partition. Operators need an explicit replay procedure.

## Guardrails

- Retry count and elapsed time are bounded.
- Dead-letter records preserve original provenance and sanitized failure details.
- Replay requires corrected conditions and an attributable actor.

## Validation

- Inject timeouts, throttling, malformed records, dependency outages, worker death, and duplicate delivery.

## Revisit triggers

- Observed failure distributions justify different classification or backoff behaviour.
