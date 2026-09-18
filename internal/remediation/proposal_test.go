package remediation

import (
	"testing"
	"time"

	"github.com/marco13-moo/cloudspend-guardian/internal/domain"
)

func TestEvaluateFailsClosedForDirectMutation(t *testing.T) {
	proposal := validProposal()
	proposal.DirectCloudMutation = true
	decision := Evaluate(proposal, "bundle/v1")
	if decision.Allowed || len(decision.Reasons) == 0 {
		t.Fatalf("expected denied decision: %+v", decision)
	}
}

func TestAuthorizeProposalRejectsMergeAndWrongRepository(t *testing.T) {
	scope := AppScope{Repository: "acme/infrastructure", CanPropose: true}
	if err := AuthorizeProposal(scope, "acme/other"); err == nil {
		t.Fatal("expected repository scope denial")
	}
	scope.CanMerge = true
	if err := AuthorizeProposal(scope, "acme/infrastructure"); err == nil {
		t.Fatal("expected merge authority denial")
	}
}

func validProposal() Proposal {
	return Proposal{ID: "proposal-1", Repository: "acme/infrastructure", Branch: "csg/rec-1", TerraformPatch: "- size = \\\"large\\\"\n+ size = \\\"medium\\\"", Recommendation: domain.Recommendation{ID: "rec-1", ResourceID: "resource-1", Summary: "Review rightsizing", RuleVersion: "rule/v1", Evidence: []domain.Evidence{{Source: "synthetic", Description: "low utilization", ObservedAt: time.Now()}}, Assumptions: []string{"representative demand"}, Confidence: 0.9, ConfidenceFactors: map[string]float64{"completeness": 1}, ProjectedMonthlySavings: domain.Money{Currency: "EUR", MinorUnits: 100}, RollbackPlan: "Restore the previous Terraform variable."}}
}
