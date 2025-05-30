package options

type ImporterType string

const (
	Consul ImporterType = "consul"
	Etcd   ImporterType = "etcd"
)

// Options 参数
type Options struct {
	Service ImporterType

	Endpoint string // 配置服务器远程地址

	ProjectName string // 项目名
	ProjectRoot string // 项目根目录
}
