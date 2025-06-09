package authz.introspection

import future.keywords.if
import future.keywords.in

default authorized := false

default authorized_project := ""

default authorized_pair := []

# Check if the input is authorized based on the policies and pairs provided.
authorized if {
	some input_sub in input.subjects
	some grant in data.policies[input_sub]

	some input_pair in input.pairs
	input_pair.resource == grant.pattern
	input_pair.action == grant.method
}

# Check if the input pair is authorized based on the policies and pairs provided.
authorized_pair := [pair] if {
	authorized

	some input_pair in input.pairs
	pair := {"resource": input_pair.resource, "action": input_pair.action}
}

# Check if the input is authorized for a specific project.
authorized_project := "api" if {
	authorized
}
