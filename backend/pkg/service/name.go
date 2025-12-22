package service

// 服务名称命名规则
//
// Consul：字母、数字、破折号；
// Etcd:
// Nacos:

const (
	Project = "gowind"

	AdminService = "admin-gateway" // 后台BFF
)

// NewDiscoveryName 构建服务发现名称
func NewDiscoveryName(serviceName string) string {
	return Project + "/" + serviceName
}
