# GoWindAdmin (GoWind Admin System)

> GoWindAdmin — **Efficiently build enterprise-grade admin systems, making development as smooth as the wind.**

An out-of-the-box Golang full-stack admin system. The backend is based on the GO microservice framework [go-kratos](https://go-kratos.dev/), and the frontend is based on the Vue microservice framework [Vben Admin](https://doc.vben.pro/).

Although both use microservice frameworks, the frontend and backend can be developed and deployed as monoliths.

Easy to get started, feature-rich, suitable for rapid development of enterprise admin systems.

**English** | [中文](./README.md) | [日本語](./README.ja-JP.md)

## Demo

> Frontend: <http://124.221.26.30:8080/>
>
> Backend Swagger: <http://124.221.26.30:7788/docs/>
>
> Default account/password: `admin` / `admin`

## Tech Stack

- Backend: [Golang](https://go.dev/) + [go-kratos](https://go-kratos.dev/) + [wire](https://github.com/google/wire) + [ent](https://entgo.io/docs/getting-started/)
- Frontend: [Vue](https://vuejs.org/) + [TypeScript](https://www.typescriptlang.org/) + [Ant Design Vue](https://antdv.com/) + [Vben Admin](https://doc.vben.pro/)

## Quick Start

1. Install Docker and Go (see `backend/script/prepare_ubuntu.sh`, `backend/script/prepare_centos.sh`, `backend/script/prepare_rocky.sh`)
2. Go to the `backend` directory and run the following commands to compile the backend service `kratos-admin`, build Docker images and start services along with required dependent Docker services:
    ```bash
    make init
    make docker
    make compose-up
    ```
3. Install npm and pnpm (installation instructions can be requested from AI)
4. Go to the `frontend` directory and run the following commands to build and start the frontend (development mode):
    ```bash
    pnpm install
    pnpm dev
    ```
5. Access for testing

- Frontend: <http://localhost:5666>, login: `admin` / `admin`
- Backend: <http://localhost:7788/docs/openapi.yaml>

## Features

| Feature | Description |
|------|--------------------------------------------------------------------------|
| User Management | Manage and query users, support advanced search and department-linked users; enable/disable users, set/unset manager, reset password, configure multiple roles/departments/managers, one-click login as specified user. |
| Tenant Management | Manage tenants. Adding a tenant auto-initializes tenant departments, default roles, and admin. Support plan configuration, enable/disable, one-click login as tenant admin. |
| Role Management | Manage roles and role groups; support user selection by role, set menu and data permissions, batch add/remove employees. |
| Organization Management | Manage organizations with tree-view listing. |
| Department Management | Manage departments with tree-view listing. |
| Permission Management | Manage permission groups, menus, permission points with tree-view listing. |
| API Management | Manage APIs, support API synchronization (mainly for selecting interfaces when adding permission points), tree-view listing, configure operation log request parameters and responses. |
| Dictionary Management | Manage dictionary categories and entries, support category-linked entries, server-side multi-column sorting, data import/export. |
| Task Scheduler | Manage tasks and task run logs; support create, update, delete, start, pause, and run immediately. |
| File Management | Manage file uploads, search files, upload to OSS or local storage, download, copy file address, delete files, support image preview (large view). |
| Message Categories | Manage message categories (2-level custom categories) for message management category selection. |
| Message Management | Manage messages, send messages to specified users, view read status and read time. |
| Internal Mail | Manage internal messages, view details, delete, mark as read, mark all as read. |
| Personal Center | View and edit personal info, view last login info, change password, etc. |
| Cache Management | Query cache list and clear cache by key. |
| Login Logs | Query login logs for successful and failed logins; supports IP geolocation. |
| Operation Logs | Query operation logs for normal and abnormal operations; supports IP geolocation and viewing operation details. |

## Backend Screenshots

<table>
    <tr>
        <td><img src="./docs/images/admin_login_page.png" alt="Admin login page"/></td>
        <td><img src="./docs/images/admin_dashboard.png" alt="Admin dashboard"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_user_list.png" alt="User list"/></td>
        <td><img src="./docs/images/admin_user_create.png" alt="Create user"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_tenant_list.png" alt="Tenant list"/></td>
        <td><img src="./docs/images/admin_tenant_create.png" alt="Create tenant"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_organization_list.png" alt="Organization list"/></td>
        <td><img src="./docs/images/admin_organization_create.png" alt="Create organization"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_department_list.png" alt="Department list"/></td>
        <td><img src="./docs/images/admin_department_create.png" alt="Create department"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_position_list.png" alt="Position list"/></td>
        <td><img src="./docs/images/admin_position_create.png" alt="Create position"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_role_list.png" alt="Role list"/></td>
        <td><img src="./docs/images/admin_role_create.png" alt="Create role"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_menu_list.png" alt="Menu list"/></td>
        <td><img src="./docs/images/admin_menu_create.png" alt="Create menu"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_task_list.png" alt="Task list"/></td>
        <td><img src="./docs/images/admin_task_create.png" alt="Create task"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_dict_list.png" alt="Dictionary list"/></td>
        <td><img src="./docs/images/admin_dict_entry_create.png" alt="Create dictionary entry"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_internal_message_list.png" alt="Internal message list"/></td>
        <td><img src="./docs/images/admin_internal_message_publish.png" alt="Publish internal message"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_login_restriction_list.png" alt="Login restriction list"/></td>
        <td><img src="./docs/images/admin_api_resource_list.png" alt="API resource list"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/admin_operation_log_list.png" alt="Operation log list"/></td>
        <td><img src="./docs/images/admin_login_log_list.png" alt="Login log list"/></td>
    </tr>
    <tr>
        <td><img src="./docs/images/api_swagger_ui.png" alt="Backend Swagger UI"/></td>
    </tr>
</table>

## Contact

- WeChat: `yang_lin_bo` (note: `go-wind-admin`)
- Juejin column: [go-wind-admin](https://juejin.cn/column/7541283508041826367)

## Thanks to JetBrains for providing free GoLand & WebStorm

[![avatar](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)](https://jb.gg/OpenSource)
