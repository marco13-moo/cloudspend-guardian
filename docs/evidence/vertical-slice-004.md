# Vertical slice 004: deduplicated opportunities and savings lineage

- Status: Implemented and locally validated
- Date: 2026-09-18
- Governing ADRs: 0015, 0018, 0019, 0020, 0037
- Scope: In-memory recommendation and measurement contracts

## Evidence

Focused tests verify that:

- provider and native recommendations with the same opportunity identity merge;
- every source recommendation remains traceable;
- the merged opportunity uses the most conservative projected savings;
- savings events require a measurement window and counterfactual explanation;
- duplicate event IDs cannot append a second ledger entry.

## Architectural assessment

- **ADR-0015:** canonical grouping and conservative savings are implemented;
  recurrence and incompatible target-state cases need broader fixtures.
- **ADRs-0018 to 0020:** projected and realized events can be recorded as
  append-only entries with baseline lineage. Production persistence and actual
  post-change usage measurement remain outstanding.

## Rollback

Revert the slice commit. No external state or cloud mutation is introduced.