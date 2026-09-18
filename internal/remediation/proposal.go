// Package remediation contains reviewable, non-mutating remediation proposals.
package remediation

import (
	"errors"
	"strings"

	"github.com/marco13-moo/cloudspend-guardian/internal/domain"
)

type Proposal struct {
	ID                  string
	Repository          string
	Branch              string
	TerraformPatch      string
	Recommendation      domain.Recommendation
	DirectCloudMutation bool
	Destructive         bool
}

type Decision struct {
	Allowed       bool
	BundleVersion string
	Reasons       []string
}

func (p Proposal) PolicyInput() map[string]any {
	return map[string]any{
		"delivery": "pull_request", "rollback_plan": p.Recommendation.RollbackPlan,
		"evidence": p.Recommendation.Evidence, "confidence": p.Recommendation.Confidence,
		"direct_cloud_mutation": p.DirectCloudMutation, "destructive": p.Destructive,
	}
}

func Evaluate(p Proposal, bundleVersion string) Decision {
	decision := Decision{BundleVersion: bundleVersion, Allowed: true}
	if bundleVersion == "" {
		decision.Allowed = false
		decision.Reasons = append(decision.Reasons, "policy bundle version is required")
	}
	if err := p.Validate(); err != nil {
		decision.Allowed = false
		decision.Reasons = append(decision.Reasons, err.Error())
	}
	input := p.PolicyInput()
	if input["delivery"] != "pull_request" || p.DirectCloudMutation || p.Destructive || p.Recommendation.Confidence < 0.8 {
		decision.Allowed = false
		decision.Reasons = append(decision.Reasons, "policy denies non-reviewable or unsafe remediation")
	}
	return decision
}

func (p Proposal) Validate() error {
	if p.ID == "" || p.Repository == "" || p.Branch == "" || !strings.Contains(p.Repository, "/") || p.TerraformPatch == "" {
		return errors.New("proposal identity, repository, branch, and Terraform patch are required")
	}
	if err := p.Recommendation.Validate(); err != nil {
		return err
	}
	return nil
}

type AppScope struct {
	Repository    string
	CanPropose    bool
	CanMerge      bool
	CanAdminister bool
}

func AuthorizeProposal(scope AppScope, repository string) error {
	if scope.Repository == "" || scope.Repository != repository || !scope.CanPropose || scope.CanMerge || scope.CanAdminister {
		return errors.New("GitHub App scope cannot propose this repository safely")
	}
	return nil
}
