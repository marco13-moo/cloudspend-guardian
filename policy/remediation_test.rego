package cloudspendguardian.remediation_test

import data.cloudspendguardian.remediation.allow
import rego.v1

test_reviewable_remediation_is_allowed if {
	allow with input as {
		"delivery": "pull_request",
		"rollback_plan": "Restore the previous Terraform variable.",
		"evidence": [{"source": "prometheus"}],
		"confidence": 0.91,
		"direct_cloud_mutation": false,
		"destructive": false,
	}
}

test_direct_mutation_is_denied if {
	not allow with input as {
		"delivery": "pull_request",
		"rollback_plan": "Restore the previous Terraform variable.",
		"evidence": [{"source": "prometheus"}],
		"confidence": 0.91,
		"direct_cloud_mutation": true,
		"destructive": false,
	}
}
