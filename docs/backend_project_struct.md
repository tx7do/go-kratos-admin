# 后端项目结构说明

```text
├─.docker
│  └─compose
├─api
│  ├─gen
│  │  └─go
│  │      ├─admin
│  │      │  └─service
│  │      │      └─v1
│  │      ├─file
│  │      │  └─service
│  │      │      └─v1
│  │      ├─system
│  │      │  └─service
│  │      │      └─v1
│  │      └─user
│  │          └─service
│  │              └─v1
│  └─protos
│      ├─admin
│      │  └─service
│      │      └─v1
│      ├─file
│      │  └─service
│      │      └─v1
│      ├─system
│      │  └─service
│      │      └─v1
│      └─user
│          └─service
│              └─v1
├─app
│  └─admin
│      └─service
│          ├─cmd
│          │  └─server
│          │      └─assets
│          ├─configs
│          └─internal
│              ├─data
│              │  └─ent
│              ├─middleware
│              │  ├─auth
│              │  └─logging
│              ├─server
│              └─service
├─pkg
├─script
├─sql
└─tools
    └─config-importer
```

1. `.docker`：存放 Docker 相关配置文件
2. `api`：存放 API 相关代码，API使用Protobuf定义，使用Buf进行编译管理。
    - `gen`：存放 API 服务生成的代码，目前只有 Go 语言的代码；
    - `protos`：存放 API 服务的 Protobuf的 proto 文件，它的目录结构是`{服务名}/service/{版本号}`。
    - `buf.gen.yaml`：buf 配置文件，用于生成 API 服务的 Go 代码。
    - `buf.admin.openapi.gen.yaml` buf 配置文件，用于生成 Admin 服务的 OpenAPI 文档。
    - `buf.admin.typescript.gen.yaml` buf 配置文件，用于生成 Admin 服务的 TypeScript 代码。
    - `buf.yaml` buf 配置文件。
3. `app`：存放应用服务相关代码，它的目录结构是`{服务名/service}`，目前只有 Admin 服务。
    - `admin/service`：存放 Admin 服务相关代码
        - `Makefile`：Makefile 文件，调用项目根目录下的`app.mk`，用于构建、运行、测试 Admin 服务。
        - `cmd`：存放 Admin 服务的命令行代码
            - `server`：存放 Admin 服务的入口代码
                - `assets`：存放 Admin 服务的静态资源文件，现在只存放了OpenAPI的静态资源文件。
        - `configs`：存放 Admin 服务的配置文件
        - `internal`：存放 Admin 服务的内部代码，使用internal目录是为了避免被外部代码引用。
            - `data`：存放 Admin 服务的数据访问代码
                - `ent`：存放 Admin 服务的 Ent 数据库 ORM 代码
            - `middleware`：存放 Admin 服务的中间件代码
                - `auth`：存放 Admin 服务的认证中间件代码
                - `logging`：存放 Admin 服务的日志中间件代码
            - `server`：存放 Admin 服务的服务端代码
            - `service`：存放 Admin 服务的服务代码
4. `pkg`：存放通用代码
5. `script`：存放Shell脚本代码，用于项目的构建、部署等，如果系统新装，需要执行`prepare_{OS}.sh`脚本，安装依赖软件；以及
   `install_docker_dependences.sh`安装第三方Docker中间件的容器；平时就使用`build_install.sh`安装项目的服务，使用PM2进行进程管理。
6. `sql`：存放 SQL 文件，里面存放了数据库的初始化 SQL 文件，以及一些测试数据的 SQL 文件。
7. `tools`：存放工具代码
    - `config-importer`：存放配置导入工具代码
8. `app.mk`：存放应用服务使用的 Makefile 文件，它由`app/{服务名}/service`下的`Makefile`调用，用于构建、运行、测试应用服务。
9. `Makefile`：项目根目录下的 Makefile 文件，可以用来安装cli，生成api代码等。
10. `README.md`：项目后端的说明文档。
