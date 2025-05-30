package internal

import (
	"flag"

	"github.com/tx7do/go-kratos-admin/tools/config-importer/internal/consul"
	"github.com/tx7do/go-kratos-admin/tools/config-importer/internal/etcd"
	"github.com/tx7do/go-kratos-admin/tools/config-importer/options"
)

// Importer 远程配置导入器
type Importer interface {
	// Import 导入所有的配置
	Import() error

	// ImportOneService 导入单个配置
	ImportOneService(app string) error
}

func NewImporter() Importer {
	var opts options.Options

	flag.StringVar((*string)(&opts.Service), "type", "consul", "remote config service name, eg: -type consul")
	flag.StringVar(&opts.Endpoint, "addr", "127.0.0.1:8500", "consul address, eg: -addr 127.0.0.1:8500")
	flag.StringVar(&opts.ProjectName, "proj", "kratos_admin", "project name, eg: -proj kratos_admin")
	flag.StringVar(&opts.ProjectRoot, "root", "./", "project root dir, eg: -root ./")

	flag.Parse()

	switch opts.Service {
	default:
		fallthrough
	case options.Consul:
		return consul.NewImporter(&opts)

	case options.Etcd:
		return etcd.NewImporter(&opts)
	}
}
