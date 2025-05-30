# 远程配置导入工具

支持的远程配置系统：

- Consul
- Etcd

## 使用方法

```shell
go run -mod=mod github.com/tx7do/go-kratos-admin/tools/config-importer \
    --type=etcd \
    --addr=localhost:2379
    --proj=kratos_admin
```
