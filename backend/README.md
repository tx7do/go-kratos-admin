# GO后端

## API文档

### Swagger UI

- [Admin Swagger UI](http://localhost:7788/docs/)

### openapi.yaml

- [Admin openapi.yaml](http://localhost:7788/docs/openapi.yaml)

## Buf.Build使用

使用[buf.build](https://buf.build/)进行Protobuf API的工程化构建。

相关命令行工具和插件的具体安装方法请参见：[Kratos微服务框架API工程化指南](https://juejin.cn/post/7191095845096259641)

在`backend/api`目录下执行命令：

### 更新buf.lock

```bash
buf dep update
```

### 生成GO代码

```bash
buf generate
```

### 生成OpenAPI v3文档

```bash
buf generate --template buf.admin.openapi.gen.yaml
```

### 生成Dart代码

首先安装插件：

```bash
flutter pub global activate protoc_plugin
```

或者

```bash
dart pub global activate protoc_plugin
```

使用方法：

```bash
buf generate --template buf.front.dart.gen.yaml
```

### 生成Typescript代码

首先安装插件：

```bash
npm install -g ts-proto
```

使用方法：

```bash
buf generate --template buf.admin.typescript.gen.yaml
```

## Make构建

### 在后端项目根目录下执行：

#### 初始化开发环境

安装命令行工具和插件：

```bash
make init
```

#### 安装更新命令行工具

```bash
make cli
```

#### 安装更新插件

```bash
make plugin
```

#### 使用docker-compose部署整套服务（包含三方中间件）

```bash
make compose-up
```

#### 停止docker-compose部署的整套服务（包含三方中间件）

```bash
make compose-down
```

### 在`app/{服务名}/service`下执行：

#### 生成API的go代码

```bash
make api
```

#### 生成API的OpenAPI v3 文档

```bash
make openapi
```

#### 生成API的dart代码

```bash
make dart
```

#### 生成API的typescript代码

```bash
make ts
```

#### 生成ent代码

```bash
make ent
```

#### 生成wire代码

```bash
make wire
```

#### 构建二进制文件

```bash
make build
```

#### 调试运行

```bash
make run
```

#### 构建Docker镜像

```bash
make docker
```

## 部署

- 所有的Docker配置文件都在`backend`目录下。
- 所有的部署脚本都在`backend/script`目录下。

Shell脚本需要赋予执行权限：

```bash
chmod +x ./script/*.sh
```

### 初始化操作系统环境

在我们拿到服务器后，首先要做的就是初始化操作系统环境。我们需要安装一些必要的工具和软件包。

#### Centos

```bash
./script/prepare_centos.sh
```

#### Rocky

```bash
./script/prepare_rocky.sh
```

#### Ubuntu

```bash
./script/prepare_ubuntu.sh
```

### Docker安装三方依赖中间件

后端需要依赖一些三方中间件，比如：postgresql、redis、consul……，我们通过Docker来安装，这样会比较简单一些。

```bash
./script/docker_compose_install_depends.sh
```

### 一键部署整个项目

部署项目有两种方法：

1. 三方中间件和微服务都运行在Docker之下；
2. 三方中间件运行在Docker下，微服务通过PM2管理运行在OS下。

#### 1. 三方中间件和微服务都运行在Docker之下

```bash
./script/docker_compose_install.sh
```

#### 2. 三方中间件运行在Docker下，微服务运行在OS下

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

## 联系我们

- 微信个人号：`yang_lin_bo`（备注：go-kratos-admin）
- 掘金专栏：[go-kratos-admin](https://juejin.cn/column/7541283508041826367)

## [感谢JetBrains提供的免费GoLand](https://jb.gg/OpenSource)

[![avatar](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)](https://jb.gg/OpenSource)
