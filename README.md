# Kratos Admin

开箱即用的Golang全栈Admin。

其后端是基于的GO微服务框架[go-kratos](https://go-kratos.dev/)
，前端也是基于Vue微服务框架的[Vben Admin](https://doc.vben.pro/)。

虽然都是使用的微服务的框架，但是前后端都是可以使用单体架构的方式进行开发和部署的。

上手容易，功能丰富，适合快速开发企业级管理系统。

## 演示地址

> 前端地址：<http://124.221.26.30:8080/>
>
> 后端Swagger地址：<http://124.221.26.30:7788/docs/>
>
> 默认账号密码: `admin` / `admin`

## 技术栈

* 后端基于 [Golang](https://go.dev/) + [go-kratos](https://go-kratos.dev/) + [wire](https://github.com/google/wire) + [ent](https://entgo.io/docs/getting-started/)
* 前端基于 [Vue](https://vuejs.org/) + [TypeScript](https://www.typescriptlang.org/) + [Ant Design Vue](https://antdv.com/) + [Vben Admin](https://doc.vben.pro/)

## 快速上手指南
1. 安装完成docker和golang（参考backend/script/prepare_ubuntu.sh、backend/script/prepare_centos.sh、backend/script/prepare_rocky.sh）
2. 进入backend目录，执行以下命令，完成后端服务kratos-admin编译和docker镜像构建并启动，同时启动依赖的docker镜像服务
    ```bash
    make init
    make docker
    make compose-up
    ```
3. 安装npm和pnpm(安装方法可询问AI)
4. 进入frontend目录，执行以下命令，完成前端的编译并启动(开发模式)
    ```bash
    pnpm install
    pnpm dev
    ```
5. 访问测试
- 前端：http://localhost:5666， 登录账号：admin，密码：admin
- 后端：http://localhost:7788/docs/openapi.yaml

## 功能列表

| 功能   | 说明                                                                       |
|------|--------------------------------------------------------------------------|
| 用户管理 | 管理和查询用户，支持高级查询和按部门联动用户，用户可禁用/启用、设置/取消主管、重置密码、配置多角色、多部门和上级主管、一键登录指定用户等功能。 |
| 租户管理 | 管理租户，新增租户后自动初始化租户部门、默认角色和管理员。支持配置套餐、禁用/启用、一键登录租户管理员功能。                   |
| 角色管理 | 管理角色和角色分组，支持按角色联动用户，设置菜单和数据权限，批量添加和移除员工。                                 |
| 组织管理 | 管理组织，支持树形列表展示。                                                           |
| 部门管理 | 管理部门，支持树形列表展示。                                                           |
| 权限管理 | 管理权限分组、菜单、权限点，支持树形列表展示。                                                  |
| 接口管理 | 管理接口，支持接口同步功能，主要用于新增权限点时选择接口，支持树形列表展示、操作日志请求参数和响应结果配置。                   |
| 字典管理 | 管理数据字典大类及其小类，支持按字典大类联动字典小类、服务端多列排序、数据导入和导出。                              |
| 任务调度 | 管理和查看任务及其任务运行日志，支持任务新增、修改、删除、启动、暂停、立即执行。                                 |
| 文件管理 | 管理文件上传，支持文件查询、上传到OSS或本地、下载、复制文件地址、删除文件、图片支持查看大图功能。                       |
| 消息分类 | 管理消息分类，支持2级自定义消息分类，用于消息管理消息分类选择。                                         |
| 消息管理 | 管理消息，支持发送指定用户消息，可查看用户是否已读和已读时间。                                          |
| 站内信  | 站内消息管理，支持消息详细查看、删除、标为已读、全部已读功能。                                          |
| 个人中心 | 个人信息展示和修改，查看最后登录信息，密码修改等功能。                                              |
| 缓存管理 | 缓存列表查询，支持根据缓存键清除缓存。                                                      |
| 登录日志 | 登录日志列表查询，记录用户登录成功和失败日志，支持IP归属地记录。                                        |
| 操作日志 | 操作日志列表查询，记录用户操作正常和异常日志，支持IP归属地记录，查看操作日志详情。                               |

## 后台截图

- 后台用户登录界面

![后台用户登录界面](./docs/images/admin_login_page.png)

- 后台分析界面

![后台分析界面](./docs/images/admin_dashboard.png)

- 用户管理界面

![后台用户管理界面](./docs/images/admin_user_management.png)

- 后台创建用户界面

![后台创建用户界面](./docs/images/admin_create_user.png)

- 后台目录管理界面

![后台目录管理界面](./docs/images/admin_menu_management.png)

- 组织管理界面

![组织管理界面](./docs/images/admin_organization_management.png)

- 部门管理界面

![部门管理界面](./docs/images/admin_department_management.png)

- 用户角色管理界面

![用户角色管理界面](./docs/images/admin_role_management.png)

- 用户职位管理界面

![用户职位管理界面](./docs/images/admin_position_management.png)

- 租户管理界面

![租户管理界面](./docs/images/admin_tenant_management.png)

- 调度任务管理界面

![调度任务管理界面](./docs/images/admin_task_management.png)

- 字典管理界面

![字典管理界面](./docs/images/admin_dict_management.png)

- 后台登录日志界面

![后台登录日志界面](./docs/images/admin_operation_log.png)

- 后台操作日志界面

![后台操作日志界面](./docs/images/admin_login_log.png)

- 后端内置Swagger UI界面

![后端内置Swagger UI界面](./docs/images/api_swagger_ui.png)
