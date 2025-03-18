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
make docker-compose
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
