# Vertical slice 005: governed remediation proposal

- Status: Implemented and locally validated
- Date: 2026-09-18
- Governing ADRs: 0005, 0016, 0017, 0037
- Scope: Proposal and policy contract; no GitHub or cloud credentials

## Evidence

Focused tests verify that proposals require a repository-scoped Terraform patch
and validated recommendation evidence, direct cloud mutation fails closed, and
GitHub App scopes cannot merge, administer, or access another repository.

The policy input is bounded and versioned at the proposal boundary. The existing
Rego policy remains the deployable policy authority; this Go validator prevents
malformed proposals from reaching it.

## Residual limitations

OPA bundle execution, signed bundle distribution, GitHub App token exchange, and
end-to-end pull-request creation remain deployment work. No proposal path grants
merge or cloud mutation authority.