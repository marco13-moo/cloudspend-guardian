# Contributing

1. Open an issue describing the problem and intended outcome.
2. Create or amend an ADR when the change crosses a trust boundary, changes a persistent contract, or materially affects a quality attribute.
3. Add tests that demonstrate the intended behaviour and important failure modes.
4. Run `make verify` before opening a pull request.
5. Explain which architectural invariants the change preserves.

Commit messages should use an imperative Conventional Commit subject, for example `feat: ingest synthetic FOCUS records`.
