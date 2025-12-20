package authz.introspection

import rego.v1

# 安全的 policies 访问：当 data.policies 缺失时返回空对象
policies := data.policies if data.policies
else := {}

# defaults
default authorized := false
default authorized_project := ""
default authorized_pair := []

# 输入安全取值：如果 input.subjects 或 input.pairs 缺失，则使用空数组
subjects := input.subjects if input.subjects
else := []

pairs := input.pairs if input.pairs
else := []

# 判断是否有任一 subject 对任一 pair 被授权
authorized if {
	some s in subjects
	some grant in policies[s]
	some p in pairs
	p.resource == grant.pattern
	p.action == grant.method
}

# 返回所有被授权的 (resource, action) 对
authorized_pair := [p |
	some s in subjects
	some grant in policies[s]
	some p in pairs
	p.resource == grant.pattern
	p.action == grant.method
]

# 项目字段目前写死为 "api"
authorized_project := "api" if authorized
