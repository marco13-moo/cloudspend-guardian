package cloudspendguardian.remediation

import rego.v1

# Production remediations must remain reviewable and reversible.
default allow := false

allow if {
	input.delivery == "pull_request"
	input.rollback_plan != ""
	count(input.evidence) > 0
	input.confidence >= 0.8
	not input.direct_cloud_mutation
	not input.destructive
}
