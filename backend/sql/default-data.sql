-- 默认的超级管理员，默认账号：admin，密码：admin
TRUNCATE TABLE kratos_admin.public.users RESTART IDENTITY;
INSERT INTO kratos_admin.public.users (username, nickname, email, authority, roles)
VALUES ('admin', 'admin', 'admin@gmail.com', 'SYS_ADMIN', '["super"]');
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));

TRUNCATE TABLE user_credentials RESTART IDENTITY;
INSERT INTO user_credentials (user_id, identity_type, identifier, credential_type, credential, status, is_primary,
                              create_time)
VALUES (1, 'USERNAME', 'admin', 'PASSWORD_HASH', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a',
        'ENABLED', true, now()),
       (1, 'EMAIL', 'admin@gmail.com', 'PASSWORD_HASH', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a',
        'ENABLED', true, now())
;
SELECT setval('user_credentials_id_seq', (SELECT MAX(id) FROM user_credentials));

-- 默认的角色
TRUNCATE TABLE kratos_admin.public.sys_roles RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_roles(id, parent_id, create_by, sort_id, name, code, status, remark, menus, apis,
                                          create_time)
VALUES (1, null, 0, 1, '超级管理员', 'super', 'ON', '超级管理员拥有对系统的最高权限',
        '[1, 2, 10, 11, 12, 13, 14, 20, 21, 22, 15, 16, 17, 18, 23, 24, 25, 26, 27, 30, 31, 32]', '[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102]', now()),
       (2, null, 0, 2, '管理员', 'admin', 'ON', '系统管理员拥有对整个系统的管理权限',
        '[1, 2, 6, 7, 8, 9, 10, 11, 12, 13, 14]', '[]', now()),
       (3, null, 0, 3, '普通用户', 'user', 'ON', '普通用户没有管理权限，只有设备和APP的使用权限', '[]', '[]', now()),
       (4, null, 0, 4, '游客', 'guest', 'ON', '游客只有非常有限的数据读取权限', '[]', '[]', now());
SELECT setval('sys_roles_id_seq', (SELECT MAX(id) FROM sys_roles));

-- 后台目录
TRUNCATE TABLE kratos_admin.public.sys_menus RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_menus(id, parent_id, type, name, path, redirect, component, status, create_time,
                                          meta)
VALUES (1, null, 'FOLDER', 'Dashboard', '/', null, 'BasicLayout', 'ON', now(), '{"order":-1, "title":"page.dashboard.title", "icon":"lucide:layout-dashboard", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (2, 1, 'MENU', 'Analytics', '/analytics', null, 'dashboard/analytics/index.vue', 'ON', now(), '{"order":-1, "title":"page.dashboard.analytics", "icon":"lucide:area-chart", "affixTab": true, "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (10, null, 'FOLDER', 'TenantManagement', '/tenant', null, 'BasicLayout', 'ON', now(), '{"order":2000, "title":"menu.tenant.moduleName", "icon":"lucide:building-2", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (11, 10, 'MENU', 'TenantMemberManagement', 'members', null, 'app/tenant/tenant/index.vue', 'ON', now(), '{"order":1, "title":"menu.tenant.member", "icon":"lucide:building-2", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (20, null, 'FOLDER', 'OrganizationalPersonnelManagement', '/opm', null, 'BasicLayout', 'ON', now(), '{"order":2001, "title":"menu.opm.moduleName", "icon":"lucide:users", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (21, 20, 'MENU', 'OrganizationManagement', 'organizations', null, 'app/opm/org/index.vue', 'ON', now(), '{"order":1, "title":"menu.opm.org", "icon":"lucide:building-2", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (22, 20, 'MENU', 'DepartmentManagement', 'departments', null, 'app/opm/dept/index.vue', 'ON', now(), '{"order":2, "title":"menu.opm.dept", "icon":"lucide:folder-tree", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (23, 20, 'MENU', 'PositionManagement', 'positions', null, 'app/opm/position/index.vue', 'ON', now(), '{"order":3, "title":"menu.opm.position", "icon":"lucide:briefcase", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (24, 20, 'MENU', 'UserManagement', 'users', null, 'app/opm/users/index.vue', 'ON', now(), '{"order":4, "title":"menu.opm.user", "icon":"lucide:users", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (25, 20, 'MENU', 'UserDetail', 'users/detail/:id', null, 'app/opm/users/detail/index.vue', 'ON', now(), '{"order":1, "title":"menu.opm.userDetail", "icon":"", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":true, "hideInTab":false}'),

       (30, null, 'FOLDER', 'PermissionManagement', '/permission', null, 'BasicLayout', 'ON', now(), '{"order":2002, "title":"menu.permission.moduleName", "icon":"lucide:shield-check", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (31, 30, 'MENU', 'RoleManagement', 'roles', null, 'app/permission/role/index.vue', 'ON', now(), '{"order":1, "title":"menu.permission.role", "icon":"lucide:shield-user", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (32, 30, 'MENU', 'MenuManagement', 'menus', null, 'app/permission/menu/index.vue', 'ON', now(), '{"order":2, "title":"menu.permission.menu", "icon":"lucide:square-menu", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (40, null, 'FOLDER', 'InternalMessageManagement', '/internal_message', null, 'BasicLayout', 'ON', now(), '{"order":2003, "title":"menu.internalMessage.moduleName", "icon":"lucide:mail", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (41, 40, 'MENU', 'NotificationMessageManagement', 'notifications', null, 'app/internal_message/notification_message/index.vue', 'ON', now(), '{"order": 1, "title":"menu.internalMessage.notificationMessage", "icon":"lucide:bell", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (42, 40, 'MENU', 'NotificationMessageCategoryManagement', 'notification_categories', null, 'app/internal_message/notification_message_category/index.vue', 'ON', now(), '{"order":2, "title":"menu.internalMessage.notificationMessageCategory", "icon":"lucide:calendar-check", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (43, 40, 'MENU', 'PrivateMessageManagement', 'private_messages', null, 'app/internal_message/private_message/index.vue', 'ON', now(), '{"order":3, "title":"menu.internalMessage.privateMessage", "icon":"lucide:message-circle-more", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (50, null, 'FOLDER', 'LogAuditManagement', '/log', null, 'BasicLayout', 'ON', now(), '{"order":2004, "title":"menu.log.moduleName", "icon":"lucide:activity", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (51, 50, 'MENU', 'AdminLoginLog', 'login', null, 'app/log/admin_login_log/index.vue', 'ON', now(), '{"order":1, "title":"menu.log.adminLoginLog", "icon":"lucide:user-lock", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (52, 50, 'MENU', 'AdminOperationLog', 'operation', null, 'app/log/admin_operation_log/index.vue', 'ON', now(), '{"order":2, "title":"menu.log.adminOperationLog", "icon":"lucide:file-clock", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (60, null, 'FOLDER', 'System', '/system', null, 'BasicLayout', 'ON', now(), '{"order":2005, "title":"menu.system.moduleName", "icon":"lucide:settings", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (61, 60, 'MENU', 'DictManagement', 'dict', null, 'app/system/dict/index.vue', 'ON', now(), '{"order":1, "title":"menu.system.dict", "icon":"lucide:library-big", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (62, 60, 'MENU', 'FileManagement', 'files', null, 'app/system/files/index.vue', 'ON', now(), '{"order":2, "title":"menu.system.file", "icon":"lucide:file-search", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (63, 60, 'MENU', 'TaskManagement', 'tasks', null, 'app/system/task/index.vue', 'ON', now(), '{"order":3, "title":"menu.system.task", "icon":"lucide:list-todo", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (64, 60, 'MENU', 'APIResourceManagement', 'apis', null, 'app/system/api_resource/index.vue', 'ON', now(), '{"order":4, "title":"menu.system.apiResource", "icon":"lucide:route", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (65, 60, 'MENU', 'AdminLoginRestrictionManagement', 'admin_login_restriction', null, 'app/system/admin_login_restriction/index.vue', 'ON', now(), '{"order":5, "title":"menu.system.adminLoginRestriction", "icon":"lucide:shield-x", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}')
;
SELECT setval('sys_menus_id_seq', (SELECT MAX(id) FROM sys_menus));

-- API资源表数据
TRUNCATE TABLE kratos_admin.public.sys_api_resources RESTART IDENTITY;
INSERT INTO public.sys_api_resources (id, create_time, update_time, delete_time, create_by, update_by, operation, description, module, path, method, module_description)
VALUES (1, '2025-09-11 01:49:30.203766 +00:00', null, null, null, null, 'PositionService_List', ' 查询职位列表 ', 'PositionService', '/admin/v1/positions', 'GET', ' 职位管理服务 '),
       (2, '2025-09-11 01:49:30.210298 +00:00', null, null, null, null, 'PositionService_Create', ' 创建职位 ', 'PositionService', '/admin/v1/positions', 'POST', ' 职位管理服务 '),
       (3, '2025-09-11 01:49:30.216269 +00:00', null, null, null, null, 'UserService_Update', ' 更新用户 ', 'UserService', '/admin/v1/users/{data.id}', 'PUT', ' 用户管理服务 '),
       (4, '2025-09-11 01:49:30.222746 +00:00', null, null, null, null, 'AdminLoginLogService_List', ' 查询后台登录日志列表 ', 'AdminLoginLogService', '/admin/v1/admin_login_logs', 'GET', ' 后台登录日志管理服务 '),
       (5, '2025-09-11 01:49:30.224638 +00:00', null, null, null, null, 'PrivateMessageService_Create', ' 创建私信消息 ', 'PrivateMessageService', '/admin/v1/private_messages', 'POST', ' 私信消息管理服务 '),
       (6, '2025-09-11 01:49:30.230644 +00:00', null, null, null, null, 'PrivateMessageService_List', ' 查询私信消息列表 ', 'PrivateMessageService', '/admin/v1/private_messages', 'GET', ' 私信消息管理服务 '),
       (7, '2025-09-11 01:49:30.234092 +00:00', null, null, null, null, 'RoleService_Get', ' 查询角色详情 ', 'RoleService', '/admin/v1/roles/{id}', 'GET', ' 角色管理服务 '),
       (8, '2025-09-11 01:49:30.238413 +00:00', null, null, null, null, 'RoleService_Delete', ' 删除角色 ', 'RoleService', '/admin/v1/roles/{id}', 'DELETE', ' 角色管理服务 '),
       (9, '2025-09-11 01:49:30.241427 +00:00', null, null, null, null, 'FileService_Update', ' 更新文件 ', 'FileService', '/admin/v1/files/{data.id}', 'PUT', ' 文件管理服务 '),
       (10, '2025-09-11 01:49:30.243677 +00:00', null, null, null, null, 'AdminLoginRestrictionService_Update', ' 更新后台登录限制 ', 'AdminLoginRestrictionService', '/admin/v1/login-restrictions/{data.id}', 'PUT', ' 后台登录限制管理服务 '),
       (11, '2025-09-11 01:49:30.247730 +00:00', null, null, null, null, 'DepartmentService_Update', ' 更新部门 ', 'DepartmentService', '/admin/v1/departments/{data.id}', 'PUT', ' 部门管理服务 '),
       (12, '2025-09-11 01:49:30.251382 +00:00', null, null, null, null, 'MenuService_Delete', ' 删除菜单 ', 'MenuService', '/admin/v1/menus/{id}', 'DELETE', ' 后台菜单管理服务 '),
       (13, '2025-09-11 01:49:30.255007 +00:00', null, null, null, null, 'MenuService_Get', ' 查询菜单详情 ', 'MenuService', '/admin/v1/menus/{id}', 'GET', ' 后台菜单管理服务 '),
       (14, '2025-09-11 01:49:30.257650 +00:00', null, null, null, null, 'TaskService_RestartAllTask', ' 重启所有的调度任务 ', 'TaskService', '/admin/v1/tasks:restart', 'POST', ' 调度任务管理服务 '),
       (15, '2025-09-11 01:49:30.260613 +00:00', null, null, null, null, 'NotificationMessageService_Delete', ' 删除通知消息 ', 'NotificationMessageService', '/admin/v1/notifications/{id}', 'DELETE', ' 通知消息管理服务 '),
       (16, '2025-09-11 01:49:30.263794 +00:00', null, null, null, null, 'NotificationMessageService_Get', ' 查询通知消息详情 ', 'NotificationMessageService', '/admin/v1/notifications/{id}', 'GET', ' 通知消息管理服务 '),
       (17, '2025-09-11 01:49:30.267011 +00:00', null, null, null, null, 'AdminOperationLogService_List', ' 查询后台操作日志列表 ', 'AdminOperationLogService', '/admin/v1/admin_operation_logs', 'GET', ' 后台操作日志管理服务 '),
       (18, '2025-09-11 01:49:30.270947 +00:00', null, null, null, null, 'ApiResourceService_SyncApiResources', ' 同步 API 资源 ', 'ApiResourceService', '/admin/v1/api-resources/sync', 'POST', 'API 资源管理服务 '),
       (19, '2025-09-11 01:49:30.273581 +00:00', null, null, null, null, 'AuthenticationService_Login', ' 登录 ', 'AuthenticationService', '/admin/v1/login', 'POST', ' 用户后台登录认证服务 '),
       (20, '2025-09-11 01:49:30.275791 +00:00', null, null, null, null, 'PositionService_Update', ' 更新职位 ', 'PositionService', '/admin/v1/positions/{data.id}', 'PUT', ' 职位管理服务 '),
       (21, '2025-09-11 01:49:30.278158 +00:00', null, null, null, null, 'RoleService_Update', ' 更新角色 ', 'RoleService', '/admin/v1/roles/{data.id}', 'PUT', ' 角色管理服务 '),
       (22, '2025-09-11 01:49:30.279881 +00:00', null, null, null, null, 'TenantService_Update', ' 更新租户 ', 'TenantService', '/admin/v1/tenants/{data.id}', 'PUT', ' 租户管理服务 '),
       (23, '2025-09-11 01:49:30.284905 +00:00', null, null, null, null, 'ApiResourceService_Delete', ' 删除 API 资源 ', 'ApiResourceService', '/admin/v1/api-resources/{id}', 'DELETE', 'API 资源管理服务 '),
       (24, '2025-09-11 01:49:30.286989 +00:00', null, null, null, null, 'ApiResourceService_Get', ' 查询 API 资源详情 ', 'ApiResourceService', '/admin/v1/api-resources/{id}', 'GET', 'API 资源管理服务 '),
       (25, '2025-09-11 01:49:30.289782 +00:00', null, null, null, null, 'DictService_Delete', ' 删除字典 ', 'DictService', '/admin/v1/dict/{id}', 'DELETE', ' 字典管理服务 '),
       (26, '2025-09-11 01:49:30.292776 +00:00', null, null, null, null, 'DictService_Get', ' 查询字典详情 ', 'DictService', '/admin/v1/dict/{id}', 'GET', ' 字典管理服务 '),
       (27, '2025-09-11 01:49:30.294947 +00:00', null, null, null, null, 'AuthenticationService_RefreshToken', ' 刷新认证令牌 ', 'AuthenticationService', '/admin/v1/refresh_token', 'POST', ' 用户后台登录认证服务 '),
       (28, '2025-09-11 01:49:30.297965 +00:00', null, null, null, null, 'TaskService_ControlTask', ' 控制调度任务 ', 'TaskService', '/admin/v1/tasks:control', 'POST', ' 调度任务管理服务 '),
       (29, '2025-09-11 01:49:30.300988 +00:00', null, null, null, null, 'TenantService_Delete', ' 删除租户 ', 'TenantService', '/admin/v1/tenants/{id}', 'DELETE', ' 租户管理服务 '),
       (30, '2025-09-11 01:49:30.302299 +00:00', null, null, null, null, 'TenantService_Get', ' 获取租户数据 ', 'TenantService', '/admin/v1/tenants/{id}', 'GET', ' 租户管理服务 '),
       (31, '2025-09-11 01:49:30.305189 +00:00', null, null, null, null, 'NotificationMessageService_List', ' 查询通知消息列表 ', 'NotificationMessageService', '/admin/v1/notifications', 'GET', ' 通知消息管理服务 '),
       (32, '2025-09-11 01:49:30.308531 +00:00', null, null, null, null, 'NotificationMessageService_Create', ' 创建通知消息 ', 'NotificationMessageService', '/admin/v1/notifications', 'POST', ' 通知消息管理服务 '),
       (33, '2025-09-11 01:49:30.314753 +00:00', null, null, null, null, 'TaskService_Update', ' 更新调度任务 ', 'TaskService', '/admin/v1/tasks/{data.id}', 'PUT', ' 调度任务管理服务 '),
       (34, '2025-09-11 01:49:30.318192 +00:00', null, null, null, null, 'ApiResourceService_List', ' 查询 API 资源列表 ', 'ApiResourceService', '/admin/v1/api-resources', 'GET', 'API 资源管理服务 '),
       (35, '2025-09-11 01:49:30.320366 +00:00', null, null, null, null, 'ApiResourceService_Create', ' 创建 API 资源 ', 'ApiResourceService', '/admin/v1/api-resources', 'POST', 'API 资源管理服务 '),
       (36, '2025-09-11 01:49:30.322789 +00:00', null, null, null, null, 'TenantService_List', ' 获取租户列表 ', 'TenantService', '/admin/v1/tenants', 'GET', ' 租户管理服务 '),
       (37, '2025-09-11 01:49:30.325738 +00:00', null, null, null, null, 'TenantService_Create', ' 创建租户 ', 'TenantService', '/admin/v1/tenants', 'POST', ' 租户管理服务 '),
       (38, '2025-09-11 01:49:30.328649 +00:00', null, null, null, null, 'FileService_List', ' 查询文件列表 ', 'FileService', '/admin/v1/files', 'GET', ' 文件管理服务 '),
       (39, '2025-09-11 01:49:30.331671 +00:00', null, null, null, null, 'FileService_Create', ' 创建文件 ', 'FileService', '/admin/v1/files', 'POST', ' 文件管理服务 '),
       (40, '2025-09-11 01:49:30.335467 +00:00', null, null, null, null, 'OssService_OssUploadUrl', ' 获取对象存储（OSS）上传用的预签名链接 ', 'OssService', '/admin/v1/file:upload-url', 'POST', 'OSS 服务 '),
       (41, '2025-09-11 01:49:30.337915 +00:00', null, null, null, null, 'NotificationMessageService_Update', ' 更新通知消息 ', 'NotificationMessageService', '/admin/v1/notifications/{data.id}', 'PUT', ' 通知消息管理服务 '),
       (42, '2025-09-11 01:49:30.340448 +00:00', null, null, null, null, 'ApiResourceService_Update', ' 更新 API 资源 ', 'ApiResourceService', '/admin/v1/api-resources/{data.id}', 'PUT', 'API 资源管理服务 '),
       (43, '2025-09-11 01:49:30.345413 +00:00', null, null, null, null, 'ApiResourceService_GetWalkRouteData', ' 查询路由数据 ', 'ApiResourceService', '/admin/v1/api-resources/walk-route', 'GET', 'API 资源管理服务 '),
       (44, '2025-09-11 01:49:30.345413 +00:00', null, null, null, null, 'AdminLoginRestrictionService_Create', ' 创建后台登录限制 ', 'AdminLoginRestrictionService', '/admin/v1/login-restrictions', 'POST', ' 后台登录限制管理服务 '),
       (45, '2025-09-11 01:49:30.351530 +00:00', null, null, null, null, 'AdminLoginRestrictionService_List', ' 查询后台登录限制列表 ', 'AdminLoginRestrictionService', '/admin/v1/login-restrictions', 'GET', ' 后台登录限制管理服务 '),
       (46, '2025-09-11 01:49:30.355066 +00:00', null, null, null, null, 'OssService_PutUploadFile', 'PUT 方法上传文件 ', 'OssService', '/admin/v1/file:upload', 'PUT', 'OSS 服务 '),
       (47, '2025-09-11 01:49:30.357292 +00:00', null, null, null, null, 'OssService_PostUploadFile', 'POST 方法上传文件 ', 'OssService', '/admin/v1/file:upload', 'POST', 'OSS 服务 '),
       (48, '2025-09-11 01:49:30.360233 +00:00', null, null, null, null, 'TaskService_Delete', ' 删除调度任务 ', 'TaskService', '/admin/v1/tasks/{id}', 'DELETE', ' 调度任务管理服务 '),
       (49, '2025-09-11 01:49:30.363099 +00:00', null, null, null, null, 'TaskService_Get', ' 查询调度任务详情 ', 'TaskService', '/admin/v1/tasks/{id}', 'GET', ' 调度任务管理服务 '),
       (50, '2025-09-11 01:49:30.366148 +00:00', null, null, null, null, 'DictService_List', ' 查询字典列表 ', 'DictService', '/admin/v1/dict', 'GET', ' 字典管理服务 '),
       (51, '2025-09-11 01:49:30.368329 +00:00', null, null, null, null, 'DictService_Create', ' 创建字典 ', 'DictService', '/admin/v1/dict', 'POST', ' 字典管理服务 '),
       (52, '2025-09-11 01:49:30.370768 +00:00', null, null, null, null, 'MenuService_Update', ' 更新菜单 ', 'MenuService', '/admin/v1/menus/{data.id}', 'PUT', ' 后台菜单管理服务 '),
       (53, '2025-09-11 01:49:30.372538 +00:00', null, null, null, null, 'NotificationMessageRecipientService_Delete', ' 删除通知消息接收者 ', 'NotificationMessageRecipientService', '/admin/v1/notifications:recipients/{id}', 'DELETE', ' 通知消息接收者管理服务 '),
       (54, '2025-09-11 01:49:30.374109 +00:00', null, null, null, null, 'NotificationMessageRecipientService_Get', ' 查询通知消息接收者详情 ', 'NotificationMessageRecipientService', '/admin/v1/notifications:recipients/{id}', 'GET', ' 通知消息接收者管理服务 '),
       (55, '2025-09-11 01:49:30.379477 +00:00', null, null, null, null, 'PositionService_Delete', ' 删除职位 ', 'PositionService', '/admin/v1/positions/{id}', 'DELETE', ' 职位管理服务 '),
       (56, '2025-09-11 01:49:30.382696 +00:00', null, null, null, null, 'PositionService_Get', ' 查询职位详情 ', 'PositionService', '/admin/v1/positions/{id}', 'GET', ' 职位管理服务 '),
       (57, '2025-09-11 01:49:30.385397 +00:00', null, null, null, null, 'FileService_Delete', ' 删除文件 ', 'FileService', '/admin/v1/files/{id}', 'DELETE', ' 文件管理服务 '),
       (58, '2025-09-11 01:49:30.387649 +00:00', null, null, null, null, 'FileService_Get', ' 查询文件详情 ', 'FileService', '/admin/v1/files/{id}', 'GET', ' 文件管理服务 '),
       (59, '2025-09-11 01:49:30.390135 +00:00', null, null, null, null, 'OrganizationService_Get', ' 查询组织详情 ', 'OrganizationService', '/admin/v1/organizations/{id}', 'GET', ' 组织管理服务 '),
       (60, '2025-09-11 01:49:30.392734 +00:00', null, null, null, null, 'OrganizationService_Delete', ' 删除组织 ', 'OrganizationService', '/admin/v1/organizations/{id}', 'DELETE', ' 组织管理服务 '),
       (61, '2025-09-11 01:49:30.395842 +00:00', null, null, null, null, 'UserService_Delete', ' 删除用户 ', 'UserService', '/admin/v1/users/{id}', 'DELETE', ' 用户管理服务 '),
       (62, '2025-09-11 01:49:30.397978 +00:00', null, null, null, null, 'UserService_Get', ' 获取用户数据 ', 'UserService', '/admin/v1/users/{id}', 'GET', ' 用户管理服务 '),
       (63, '2025-09-11 01:49:30.400504 +00:00', null, null, null, null, 'NotificationMessageCategoryService_Delete', ' 删除通知消息分类 ', 'NotificationMessageCategoryService', '/admin/v1/notifications:categories/{id}', 'DELETE', ' 通知消息分类管理服务 '),
       (64, '2025-09-11 01:49:30.403234 +00:00', null, null, null, null, 'NotificationMessageCategoryService_Get', ' 查询通知消息分类详情 ', 'NotificationMessageCategoryService', '/admin/v1/notifications:categories/{id}', 'GET', ' 通知消息分类管理服务 '),
       (65, '2025-09-11 01:49:30.405244 +00:00', null, null, null, null, 'OrganizationService_Update', ' 更新组织 ', 'OrganizationService', '/admin/v1/organizations/{data.id}', 'PUT', ' 组织管理服务 '),
       (66, '2025-09-11 01:49:30.406295 +00:00', null, null, null, null, 'AdminLoginLogService_Get', ' 查询后台登录日志详情 ', 'AdminLoginLogService', '/admin/v1/admin_login_logs/{id}', 'GET', ' 后台登录日志管理服务 '),
       (67, '2025-09-11 01:49:30.410871 +00:00', null, null, null, null, 'AdminOperationLogService_Get', ' 查询后台操作日志详情 ', 'AdminOperationLogService', '/admin/v1/admin_operation_logs/{id}', 'GET', ' 后台操作日志管理服务 '),
       (68, '2025-09-11 01:49:30.413220 +00:00', null, null, null, null, 'AuthenticationService_ChangePassword', ' 修改用户密码 ', 'AuthenticationService', '/admin/v1/change_password', 'POST', ' 用户后台登录认证服务 '),
       (69, '2025-09-11 01:49:30.415529 +00:00', null, null, null, null, 'UserProfileService_GetUser', ' 获取用户资料 ', 'UserProfileService', '/admin/v1/me', 'GET', ' 用户个人资料服务 '),
       (70, '2025-09-11 01:49:30.419186 +00:00', null, null, null, null, 'UserProfileService_UpdateUser', ' 更新用户资料 ', 'UserProfileService', '/admin/v1/me', 'PUT', ' 用户个人资料服务 '),
       (71, '2025-09-11 01:49:30.421231 +00:00', null, null, null, null, 'PrivateMessageService_Update', ' 更新私信消息 ', 'PrivateMessageService', '/admin/v1/private_messages/{data.id}', 'PUT', ' 私信消息管理服务 '),
       (72, '2025-09-11 01:49:30.423688 +00:00', null, null, null, null, 'AuthenticationService_Logout', ' 登出 ', 'AuthenticationService', '/admin/v1/logout', 'POST', ' 用户后台登录认证服务 '),
       (73, '2025-09-11 01:49:30.426459 +00:00', null, null, null, null, 'RoleService_List', ' 查询角色列表 ', 'RoleService', '/admin/v1/roles', 'GET', ' 角色管理服务 '),
       (74, '2025-09-11 01:49:30.429653 +00:00', null, null, null, null, 'RoleService_Create', ' 创建角色 ', 'RoleService', '/admin/v1/roles', 'POST', ' 角色管理服务 '),
       (75, '2025-09-11 01:49:30.431935 +00:00', null, null, null, null, 'DictService_Update', ' 更新字典 ', 'DictService', '/admin/v1/dict/{data.id}', 'PUT', ' 字典管理服务 '),
       (76, '2025-09-11 01:49:30.434618 +00:00', null, null, null, null, 'RouterService_ListPermissionCode', ' 查询权限码列表 ', 'RouterService', '/admin/v1/perm-codes', 'GET', ' 网站后台动态路由服务 '),
       (77, '2025-09-11 01:49:30.437105 +00:00', null, null, null, null, 'NotificationMessageRecipientService_Create', ' 创建通知消息接收者 ', 'NotificationMessageRecipientService', '/admin/v1/notifications:recipients', 'POST', ' 通知消息接收者管理服务 '),
       (78, '2025-09-11 01:49:30.439112 +00:00', null, null, null, null, 'NotificationMessageRecipientService_List', ' 查询通知消息接收者列表 ', 'NotificationMessageRecipientService', '/admin/v1/notifications:recipients', 'GET', ' 通知消息接收者管理服务 '),
       (79, '2025-09-11 01:49:30.441546 +00:00', null, null, null, null, 'TaskService_StopAllTask', ' 停止所有的调度任务 ', 'TaskService', '/admin/v1/tasks:stop', 'POST', ' 调度任务管理服务 '),
       (80, '2025-09-11 01:49:30.443729 +00:00', null, null, null, null, 'MenuService_List', ' 查询菜单列表 ', 'MenuService', '/admin/v1/menus', 'GET', ' 后台菜单管理服务 '),
       (81, '2025-09-11 01:49:30.446878 +00:00', null, null, null, null, 'MenuService_Create', ' 创建菜单 ', 'MenuService', '/admin/v1/menus', 'POST', ' 后台菜单管理服务 '),
       (82, '2025-09-11 01:49:30.449328 +00:00', null, null, null, null, 'RouterService_ListRoute', ' 查询路由列表 ', 'RouterService', '/admin/v1/routes', 'GET', ' 网站后台动态路由服务 '),
       (83, '2025-09-11 01:49:30.452158 +00:00', null, null, null, null, 'AdminLoginRestrictionService_Delete', ' 删除后台登录限制 ', 'AdminLoginRestrictionService', '/admin/v1/login-restrictions/{id}', 'DELETE', ' 后台登录限制管理服务 '),
       (84, '2025-09-11 01:49:30.454799 +00:00', null, null, null, null, 'AdminLoginRestrictionService_Get', ' 查询后台登录限制详情 ', 'AdminLoginRestrictionService', '/admin/v1/login-restrictions/{id}', 'GET', ' 后台登录限制管理服务 '),
       (85, '2025-09-11 01:49:30.457433 +00:00', null, null, null, null, 'TaskService_List', ' 查询调度任务列表 ', 'TaskService', '/admin/v1/tasks', 'GET', ' 调度任务管理服务 '),
       (86, '2025-09-11 01:49:30.459400 +00:00', null, null, null, null, 'TaskService_Create', ' 创建调度任务 ', 'TaskService', '/admin/v1/tasks', 'POST', ' 调度任务管理服务 '),
       (87, '2025-09-11 01:49:30.469103 +00:00', null, null, null, null, 'OrganizationService_List', ' 查询组织列表 ', 'OrganizationService', '/admin/v1/organizations', 'GET', ' 组织管理服务 '),
       (88, '2025-09-11 01:49:30.477770 +00:00', null, null, null, null, 'OrganizationService_Create', ' 创建组织 ', 'OrganizationService', '/admin/v1/organizations', 'POST', ' 组织管理服务 '),
       (89, '2025-09-11 01:49:30.495221 +00:00', null, null, null, null, 'NotificationMessageCategoryService_Update', ' 更新通知消息分类 ', 'NotificationMessageCategoryService', '/admin/v1/notifications:categories/{data.id}', 'PUT', ' 通知消息分类管理服务 '),
       (90, '2025-09-11 01:49:30.514043 +00:00', null, null, null, null, 'NotificationMessageCategoryService_List', ' 查询通知消息分类列表 ', 'NotificationMessageCategoryService', '/admin/v1/notifications:categories', 'GET', ' 通知消息分类管理服务 '),
       (91, '2025-09-11 01:49:30.523874 +00:00', null, null, null, null, 'NotificationMessageCategoryService_Create', ' 创建通知消息分类 ', 'NotificationMessageCategoryService', '/admin/v1/notifications:categories', 'POST', ' 通知消息分类管理服务 '),
       (92, '2025-09-11 01:49:30.541257 +00:00', null, null, null, null, 'NotificationMessageRecipientService_Update', ' 更新通知消息接收者 ', 'NotificationMessageRecipientService', '/admin/v1/notifications:recipients/{data.id}', 'PUT', ' 通知消息接收者管理服务 '),
       (93, '2025-09-11 01:49:30.547578 +00:00', null, null, null, null, 'DepartmentService_Delete', ' 删除部门 ', 'DepartmentService', '/admin/v1/departments/{id}', 'DELETE', ' 部门管理服务 '),
       (94, '2025-09-11 01:49:30.553878 +00:00', null, null, null, null, 'DepartmentService_Get', ' 查询部门详情 ', 'DepartmentService', '/admin/v1/departments/{id}', 'GET', ' 部门管理服务 '),
       (95, '2025-09-11 01:49:30.558164 +00:00', null, null, null, null, 'PrivateMessageService_Delete', ' 删除私信消息 ', 'PrivateMessageService', '/admin/v1/private_messages/{id}', 'DELETE', ' 私信消息管理服务 '),
       (96, '2025-09-11 01:49:30.560634 +00:00', null, null, null, null, 'PrivateMessageService_Get', ' 查询私信消息详情 ', 'PrivateMessageService', '/admin/v1/private_messages/{id}', 'GET', ' 私信消息管理服务 '),
       (97, '2025-09-11 01:49:30.563137 +00:00', null, null, null, null, 'DepartmentService_List', ' 查询部门列表 ', 'DepartmentService', '/admin/v1/departments', 'GET', ' 部门管理服务 '),
       (98, '2025-09-11 01:49:30.565688 +00:00', null, null, null, null, 'DepartmentService_Create', ' 创建部门 ', 'DepartmentService', '/admin/v1/departments', 'POST', ' 部门管理服务 '),
       (99, '2025-09-11 01:49:30.568053 +00:00', null, null, null, null, 'UEditorService_UploadFile', ' 上传文件 ', 'UEditorService', '/admin/v1/ueditor', 'POST', 'UEditor 后端服务 '),
       (100, '2025-09-11 01:49:30.569955 +00:00', null, null, null, null, 'UEditorService_UEditorAPI', 'UEditor API', 'UEditorService', '/admin/v1/ueditor', 'GET', 'UEditor 后端服务 '),
       (101, '2025-09-11 01:49:30.572235 +00:00', null, null, null, null, 'UserService_List', ' 获取用户列表 ', 'UserService', '/admin/v1/users', 'GET', ' 用户管理服务 '),
       (102, '2025-09-11 01:49:30.575284 +00:00', null, null, null, null, 'UserService_Create', ' 创建用户 ', 'UserService', '/admin/v1/users', 'POST', ' 用户管理服务 ');
