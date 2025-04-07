# 后端项目部署

- 所有的Docker配置文件都在`backend`目录下。
- 所有的部署脚本都在`backend/script`目录下。

Shell脚本需要赋予执行权限：

```bash
chmod +x ./script/*.sh
```

## 初始化操作系统环境

在我们拿到服务器后，首先要做的就是初始化操作系统环境。我们需要安装一些必要的工具和软件包。

### Centos

```bash
./script/prepare_centos.sh
```

### Rocky

```bash
./script/prepare_rocky.sh
```

### Ubuntu

```bash
./script/prepare_ubuntu.sh
```

## Docker安装三方依赖中间件

后端需要依赖一些三方中间件，比如：postgresql、redis、consul……，我们通过Docker来安装，这样会比较简单一些。

```bash
./script/docker_compose_install_depends.sh
```

## 一键部署整个项目

部署项目有两种方法：

1. 三方中间件和微服务都运行在Docker之下；
2. 三方中间件运行在Docker下，微服务通过PM2管理运行在OS下。

### 1. 三方中间件和微服务都运行在Docker之下

```bash
./script/docker_compose_install.sh
```

### 2. 三方中间件运行在Docker下，微服务运行在OS下

```bash
./script/build_install.sh
```

我们需要修改`hosts`文件，修改需要管理员权限，其配置文件路径在：

- Linux：`/etc/hosts`
- MacOS: `/private/etc/hosts`
- Windows: `C:\Windows\System32\drivers\etc\hosts`

增加以下内容：

```ini
127.0.0.1 postgres
127.0.0.1 redis
127.0.0.1 minio
127.0.0.1 consul
127.0.0.1 jaeger
```

> 注意：如果注册中心使用Consul，consul的地址填写为`consul`会返回`502`，使用`localhost`或者`127.0.0.1`都可以。
> ```yaml
> registry:
> type: "consul"
>
> consul:
> address: "localhost:8500"
> ```
