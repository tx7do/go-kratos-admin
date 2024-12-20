package consul

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/hashicorp/consul/api"

	"config-importer/internal/utils"
	"config-importer/options"
)

type Importer struct {
	client  *api.Client
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
	client, err := api.NewClient(&api.Config{
		Address: i.options.Address,
	})
	if err != nil {
		panic(err)
	}

	i.client = client
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
		key := i.getServiceConfigConsulKey(i.options.ProjectName, app, filepath.Base(file))
		fmt.Println(key)
		if err := i.writeConfigToConsul(key, content); err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}

// writeConfigToConsul 写入配置到Consul
func (i *Importer) writeConfigToConsul(key string, value []byte) error {
	if _, err := i.client.KV().Put(&api.KVPair{Key: key, Value: value}, nil); err != nil {
		return err
	}
	return nil
}

// getServiceConfigFolder 获取某一个服务的配置文件夹路径
func (i *Importer) getServiceConfigFolder(root, app string) string {
	return root + "app/" + app + "/" + "service/configs/"
}

func (i *Importer) getServiceConfigConsulKey(project, app, file string) string {
	key := project + "/" + app + "/" + "service/" + file
	key = strings.Replace(key, "\\", "/", -1)
	return key
}

// getConfigFileList 获取配置文件列表
func (i *Importer) getConfigFileList(root, app string) []string {
	path := i.getServiceConfigFolder(root, app)
	return utils.GetFileList(path)
}
