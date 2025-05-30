package etcd

import (
	"context"
	"fmt"
	"path/filepath"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/tx7do/go-kratos-admin/tools/config-importer/internal/utils"
	"github.com/tx7do/go-kratos-admin/tools/config-importer/options"
)

type Importer struct {
	client  *clientv3.Client
	options *options.Options
}

func NewImporter(options *options.Options) *Importer {
	cli := &Importer{
		options: options,
	}

	cli.init()

	return cli
}

func (i *Importer) init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{i.options.Endpoint},
	})
	if err != nil {
		return
	}

	i.client = cli
}

// Import 导入所有的配置
func (i *Importer) Import() error {
	apps := utils.GetFolderNameList(i.options.ProjectRoot + "app/")
	for _, app := range apps {
		_ = i.ImportOneService(app)
	}

	return nil
}

// ImportOneService 导入单个配置
func (i *Importer) ImportOneService(app string) error {
	files := i.getConfigFileList(i.options.ProjectRoot, app)
	for _, file := range files {
		content := utils.ReadFile(i.options.ProjectRoot + "/" + file)
		key := i.getServiceConfigEtcdKey(i.options.ProjectName, app, filepath.Base(file))
		fmt.Println(key)
		if err := i.writeConfigToEtcd(key, content); err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}

// writeConfigToEtcd 写入配置到 Etcd
func (i *Importer) writeConfigToEtcd(key string, value []byte) error {
	_, err := i.client.Put(context.Background(), key, string(value))
	return err
}

// getServiceConfigFolder 获取某一个服务的配置文件夹路径
func (i *Importer) getServiceConfigFolder(root, app string) string {
	return root + "app/" + app + "/" + "service/configs/"
}

// getServiceConfigEtcdKey 获取配置的 Etcd Key
func (i *Importer) getServiceConfigEtcdKey(project, app, file string) string {
	return fmt.Sprintf("%s/%s/service/%s", project, app, file)
}

// getConfigFileList 获取配置文件列表
func (i *Importer) getConfigFileList(root, app string) []string {
	path := i.getServiceConfigFolder(root, app)
	return utils.GetFileList(path)
}
