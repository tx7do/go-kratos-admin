-- 插入4个权限的用户
TRUNCATE TABLE kratos_admin.public.sys_users RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_users (username, nickname, realname, email, authority, role_ids, gender, create_time)
VALUES ('admin', '鹳狸猿', '喵个咪', 'admin@gmail.com', 'SYS_ADMIN', '[1]', 'MALE', now()),
       -- 2. 租户管理员（TENANT_ADMIN）
       ('tenant_admin', '租户管理', '张管理员', 'tenant@company.com', 'TENANT_ADMIN', '[2]', 'MALE', now()),
       -- 3. 普通用户（CUSTOMER_USER）
       ('normal_user', '普通用户', '李用户', 'user@company.com', 'CUSTOMER_USER', '[3]', 'FEMALE', now()),
       -- 4. 访客（GUEST）
       ('guest_user', '临时访客', '王访客', 'guest@company.com', 'GUEST', '[4]', 'SECRET', now())
;
SELECT setval('sys_users_id_seq', (SELECT MAX(id) FROM sys_users));

-- 插入4个用户的凭证（密码统一为admin，哈希值与原admin一致，方便测试）
TRUNCATE TABLE sys_user_credentials RESTART IDENTITY;
INSERT INTO sys_user_credentials (user_id, identity_type, identifier, credential_type, credential, status, is_primary, create_time)
VALUES (1, 'USERNAME', 'admin', 'PASSWORD_HASH', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a', 'ENABLED', true, now()),
       (1, 'EMAIL', 'admin@gmail.com', 'PASSWORD_HASH', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a', 'ENABLED', false, now()),
       -- 租户管理员（对应users表id=2）
       (2, 'USERNAME', 'tenant_admin', 'PASSWORD_HASH', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a', 'ENABLED', true, now()),
       (2, 'EMAIL', 'tenant@company.com', 'PASSWORD_HASH', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a', 'ENABLED', false, now()),

       -- 普通用户（对应users表id=3）
       (3, 'USERNAME', 'normal_user', 'PASSWORD_HASH', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a', 'ENABLED', true, now()),
       (3, 'EMAIL', 'user@company.com', 'PASSWORD_HASH', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a', 'ENABLED', false, now()),

       -- 访客（对应users表id=4）
       (4, 'USERNAME', 'guest_user', 'PASSWORD_HASH', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a', 'ENABLED', true, now()),
       (4, 'EMAIL', 'guest@company.com', 'PASSWORD_HASH', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a', 'ENABLED', false, now())
;
SELECT setval('sys_user_credentials_id_seq', (SELECT MAX(id) FROM sys_user_credentials));

-- 默认的角色
TRUNCATE TABLE kratos_admin.public.sys_roles RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_roles(id, parent_id, create_by, sort_id, name, code, status, remark, menus, apis, create_time)
VALUES (1, null, 0, 1, '超级管理员', 'super', 'ON', '拥有系统所有功能的操作权限，可管理租户、用户、角色及所有资源',
        '[1, 2, 10, 11, 20, 21, 22, 23, 24, 25, 30, 31, 32, 40, 41, 42, 43, 50, 51, 52, 60, 61, 62, 63, 64, 65]', '[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108]', now()),
       (2, null, 0, 2, '租户管理员', 'tenant_admin', 'ON', '管理当前租户下的用户、角色及资源，无跨租户操作权限', '[1, 2, 20, 21, 22, 23, 24, 25, 50, 51, 52]', '[105, 104, 35, 34, 16, 106, 93, 14, 1, 92, 91, 85, 79, 46, 24, 23, 78, 56, 55, 8, 7, 52, 51, 6, 5, 4, 31, 30, 20, 19, 53, 15]', now()),
       (3, null, 0, 3, '普通用户', 'user', 'ON', '可访问和使用租户内授权的资源，无管理权限', '[]', '[]', now()),
       (4, null, 0, 4, '访客用户', 'guest', 'ON', '仅可访问公开资源，无修改和管理权限，会话过期后自动失效', '[]', '[]', now()),
       (5, null, 0, 5, '审计员', 'auditor', 'ON', '仅可查看系统操作日志和数据记录，无修改权限', '[]', '[]', now())
;
SELECT setval('sys_roles_id_seq', (SELECT MAX(id) FROM sys_roles));

-- 后台目录
TRUNCATE TABLE kratos_admin.public.sys_menus RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_menus(id, parent_id, type, name, path, redirect, component, status, create_time, meta)
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
INSERT INTO public.sys_api_resources (
    id, create_time, update_time, delete_time,
    create_by, update_by, description, module,
    module_description, operation, path, method, scope
) VALUES
      (1, now(), null, null, null, null, '登录', 'AuthenticationService', '用户后台登录认证服务', 'AuthenticationService_Login', '/admin/v1/login', 'POST', 'ADMIN'),
      (2, now(), null, null, null, null, '删除角色', 'RoleService', '角色管理服务', 'RoleService_Delete', '/admin/v1/roles/{id}', 'DELETE', 'ADMIN'),
      (3, now(), null, null, null, null, '查询角色详情', 'RoleService', '角色管理服务', 'RoleService_Get', '/admin/v1/roles/{id}', 'GET', 'ADMIN'),
      (4, now(), null, null, null, null, '删除部门', 'DepartmentService', '部门管理服务', 'DepartmentService_Delete', '/admin/v1/departments/{id}', 'DELETE', 'ADMIN'),
      (5, now(), null, null, null, null, '查询部门详情', 'DepartmentService', '部门管理服务', 'DepartmentService_Get', '/admin/v1/departments/{id}', 'GET', 'ADMIN'),
      (6, now(), null, null, null, null, '更新部门', 'DepartmentService', '部门管理服务', 'DepartmentService_Update', '/admin/v1/departments/{data.id}', 'PUT', 'ADMIN'),
      (7, now(), null, null, null, null, '查询职位列表', 'PositionService', '职位管理服务', 'PositionService_List', '/admin/v1/positions', 'GET', 'ADMIN'),
      (8, now(), null, null, null, null, '创建职位', 'PositionService', '职位管理服务', 'PositionService_Create', '/admin/v1/positions', 'POST', 'ADMIN'),
      (9, now(), null, null, null, null, '更新API资源', 'ApiResourceService', 'API资源管理服务', 'ApiResourceService_Update', '/admin/v1/api-resources/{data.id}', 'PUT', 'ADMIN'),
      (10, now(), null, null, null, null, '删除私信消息', 'PrivateMessageService', '私信消息管理服务', 'PrivateMessageService_Delete', '/admin/v1/private_messages/{id}', 'DELETE', 'ADMIN'),
      (11, now(), null, null, null, null, '查询私信消息详情', 'PrivateMessageService', '私信消息管理服务', 'PrivateMessageService_Get', '/admin/v1/private_messages/{id}', 'GET', 'ADMIN'),
      (12, now(), null, null, null, null, '更新通知消息分类', 'NotificationMessageCategoryService', '通知消息分类管理服务', 'NotificationMessageCategoryService_Update', '/admin/v1/notifications:categories/{data.id}', 'PUT', 'ADMIN'),
      (13, now(), null, null, null, null, '更新文件', 'FileService', '文件管理服务', 'FileService_Update', '/admin/v1/files/{data.id}', 'PUT', 'ADMIN'),
      (14, now(), null, null, null, null, '登出', 'AuthenticationService', '用户后台登录认证服务', 'AuthenticationService_Logout', '/admin/v1/logout', 'POST', 'ADMIN'),
      (15, now(), null, null, null, null, '查询后台登录日志列表', 'AdminLoginLogService', '后台登录日志管理服务', 'AdminLoginLogService_List', '/admin/v1/admin_login_logs', 'GET', 'ADMIN'),
      (16, now(), null, null, null, null, '更新组织', 'OrganizationService', '组织管理服务', 'OrganizationService_Update', '/admin/v1/organizations/{data.id}', 'PUT', 'ADMIN'),
      (17, now(), null, null, null, null, '查询通知消息详情', 'NotificationMessageService', '通知消息管理服务', 'NotificationMessageService_Get', '/admin/v1/notifications/{id}', 'GET', 'ADMIN'),
      (18, now(), null, null, null, null, '删除通知消息', 'NotificationMessageService', '通知消息管理服务', 'NotificationMessageService_Delete', '/admin/v1/notifications/{id}', 'DELETE', 'ADMIN'),
      (19, now(), null, null, null, null, 'UEditor API', 'UEditorService', 'UEditor后端服务', 'UEditorService_UEditorAPI', '/admin/v1/ueditor', 'GET', 'ADMIN'),
      (20, now(), null, null, null, null, '上传文件', 'UEditorService', 'UEditor后端服务', 'UEditorService_UploadFile', '/admin/v1/ueditor', 'POST', 'ADMIN'),
      (21, now(), null, null, null, null, '获取租户列表', 'TenantService', '租户管理服务', 'TenantService_List', '/admin/v1/tenants', 'GET', 'ADMIN'),
      (22, now(), null, null, null, null, '创建租户', 'TenantService', '租户管理服务', 'TenantService_Create', '/admin/v1/tenants', 'POST', 'ADMIN'),
      (23, now(), null, null, null, null, '删除用户', 'UserService', '用户管理服务', 'UserService_Delete', '/admin/v1/users/{id}', 'DELETE', 'ADMIN'),
      (24, now(), null, null, null, null, '获取用户数据', 'UserService', '用户管理服务', 'UserService_Get', '/admin/v1/users/{id}', 'GET', 'ADMIN'),
      (25, now(), null, null, null, null, '查询权限码列表', 'RouterService', '网站后台动态路由服务', 'RouterService_ListPermissionCode', '/admin/v1/perm-codes', 'GET', 'ADMIN'),
      (26, now(), null, null, null, null, '查询私信消息列表', 'PrivateMessageService', '私信消息管理服务', 'PrivateMessageService_List', '/admin/v1/private_messages', 'GET', 'ADMIN'),
      (27, now(), null, null, null, null, '创建私信消息', 'PrivateMessageService', '私信消息管理服务', 'PrivateMessageService_Create', '/admin/v1/private_messages', 'POST', 'ADMIN'),
      (28, now(), null, null, null, null, '更新私信消息', 'PrivateMessageService', '私信消息管理服务', 'PrivateMessageService_Update', '/admin/v1/private_messages/{data.id}', 'PUT', 'ADMIN'),
      (29, now(), null, null, null, null, '停止所有的调度任务', 'TaskService', '调度任务管理服务', 'TaskService_StopAllTask', '/admin/v1/tasks:stop', 'POST', 'ADMIN'),
      (30, now(), null, null, null, null, '获取用户资料', 'UserProfileService', '用户个人资料服务', 'UserProfileService_GetUser', '/admin/v1/me', 'GET', 'ADMIN'),
      (31, now(), null, null, null, null, '更新用户资料', 'UserProfileService', '用户个人资料服务', 'UserProfileService_UpdateUser', '/admin/v1/me', 'PUT', 'ADMIN'),
      (32, now(), null, null, null, null, '查询字典详情', 'DictService', '字典管理服务', 'DictService_Get', '/admin/v1/dict/{id}', 'GET', 'ADMIN'),
      (33, now(), null, null, null, null, '删除字典', 'DictService', '字典管理服务', 'DictService_Delete', '/admin/v1/dict/{id}', 'DELETE', 'ADMIN'),
      (34, now(), null, null, null, null, '查询组织列表', 'OrganizationService', '组织管理服务', 'OrganizationService_List', '/admin/v1/organizations', 'GET', 'ADMIN'),
      (35, now(), null, null, null, null, '创建组织', 'OrganizationService', '组织管理服务', 'OrganizationService_Create', '/admin/v1/organizations', 'POST', 'ADMIN'),
      (36, now(), null, null, null, null, '启动所有的调度任务', 'TaskService', '调度任务管理服务', 'TaskService_StartAllTask', '/admin/v1/tasks:start', 'POST', 'ADMIN'),
      (37, now(), null, null, null, null, '查询调度任务列表', 'TaskService', '调度任务管理服务', 'TaskService_List', '/admin/v1/tasks', 'GET', 'ADMIN'),
      (38, now(), null, null, null, null, '创建调度任务', 'TaskService', '调度任务管理服务', 'TaskService_Create', '/admin/v1/tasks', 'POST', 'ADMIN'),
      (39, now(), null, null, null, null, '查询路由数据', 'ApiResourceService', 'API资源管理服务', 'ApiResourceService_GetWalkRouteData', '/admin/v1/api-resources/walk-route', 'GET', 'ADMIN'),
      (40, now(), null, null, null, null, '查询通知消息接收者列表', 'NotificationMessageRecipientService', '通知消息接收者管理服务', 'NotificationMessageRecipientService_List', '/admin/v1/notifications:recipients', 'GET', 'ADMIN'),
      (41, now(), null, null, null, null, '创建通知消息接收者', 'NotificationMessageRecipientService', '通知消息接收者管理服务', 'NotificationMessageRecipientService_Create', '/admin/v1/notifications:recipients', 'POST', 'ADMIN'),
      (42, now(), null, null, null, null, '重启所有的调度任务', 'TaskService', '调度任务管理服务', 'TaskService_RestartAllTask', '/admin/v1/tasks:restart', 'POST', 'ADMIN'),
      (43, now(), null, null, null, null, '删除API资源', 'ApiResourceService', 'API资源管理服务', 'ApiResourceService_Delete', '/admin/v1/api-resources/{id}', 'DELETE', 'ADMIN'),
      (44, now(), null, null, null, null, '查询API资源详情', 'ApiResourceService', 'API资源管理服务', 'ApiResourceService_Get', '/admin/v1/api-resources/{id}', 'GET', 'ADMIN'),
      (45, now(), null, null, null, null, '更新调度任务', 'TaskService', '调度任务管理服务', 'TaskService_Update', '/admin/v1/tasks/{data.id}', 'PUT', 'ADMIN'),
      (46, now(), null, null, null, null, '修改用户密码', 'UserService', '用户管理服务', 'UserService_EditUserPassword', '/admin/v1/users/{userId}/password', 'POST', 'ADMIN'),
      (47, now(), null, null, null, null, '删除菜单', 'MenuService', '后台菜单管理服务', 'MenuService_Delete', '/admin/v1/menus/{id}', 'DELETE', 'ADMIN'),
      (48, now(), null, null, null, null, '查询菜单详情', 'MenuService', '后台菜单管理服务', 'MenuService_Get', '/admin/v1/menus/{id}', 'GET', 'ADMIN'),
      (49, now(), null, null, null, null, '更新租户', 'TenantService', '租户管理服务', 'TenantService_Update', '/admin/v1/tenants/{data.id}', 'PUT', 'ADMIN'),
      (50, now(), null, null, null, null, '查询后台操作日志列表', 'AdminOperationLogService', '后台操作日志管理服务', 'AdminOperationLogService_List', '/admin/v1/admin_operation_logs', 'GET', 'ADMIN'),
      (51, now(), null, null, null, null, '查询部门列表', 'DepartmentService', '部门管理服务', 'DepartmentService_List', '/admin/v1/departments', 'GET', 'ADMIN'),
      (52, now(), null, null, null, null, '创建部门', 'DepartmentService', '部门管理服务', 'DepartmentService_Create', '/admin/v1/departments', 'POST', 'ADMIN'),
      (53, now(), null, null, null, null, '查询后台登录日志详情', 'AdminLoginLogService', '后台登录日志管理服务', 'AdminLoginLogService_Get', '/admin/v1/admin_login_logs/{id}', 'GET', 'ADMIN'),
      (54, now(), null, null, null, null, '创建租户及管理员用户', 'TenantService', '租户管理服务', 'TenantService_CreateTenantWithAdminUser', '/admin/v1/tenants_with_admin', 'POST', 'ADMIN'),
      (55, now(), null, null, null, null, '删除职位', 'PositionService', '职位管理服务', 'PositionService_Delete', '/admin/v1/positions/{id}', 'DELETE', 'ADMIN'),
      (56, now(), null, null, null, null, '查询职位详情', 'PositionService', '职位管理服务', 'PositionService_Get', '/admin/v1/positions/{id}', 'GET', 'ADMIN'),
      (57, now(), null, null, null, null, '控制调度任务', 'TaskService', '调度任务管理服务', 'TaskService_ControlTask', '/admin/v1/tasks:control', 'POST', 'ADMIN'),
      (58, now(), null, null, null, null, '删除通知消息接收者', 'NotificationMessageRecipientService', '通知消息接收者管理服务', 'NotificationMessageRecipientService_Delete', '/admin/v1/notifications:recipients/{id}', 'DELETE', 'ADMIN'),
      (59, now(), null, null, null, null, '查询通知消息接收者详情', 'NotificationMessageRecipientService', '通知消息接收者管理服务', 'NotificationMessageRecipientService_Get', '/admin/v1/notifications:recipients/{id}', 'GET', 'ADMIN'),
      (60, now(), null, null, null, null, '更新角色', 'RoleService', '角色管理服务', 'RoleService_Update', '/admin/v1/roles/{data.id}', 'PUT', 'ADMIN'),
      (61, now(), null, null, null, null, '同步API资源', 'ApiResourceService', 'API资源管理服务', 'ApiResourceService_SyncApiResources', '/admin/v1/api-resources/sync', 'POST', 'ADMIN'),
      (62, now(), null, null, null, null, '查询后台登录限制列表', 'AdminLoginRestrictionService', '后台登录限制管理服务', 'AdminLoginRestrictionService_List', '/admin/v1/login-restrictions', 'GET', 'ADMIN'),
      (63, now(), null, null, null, null, '创建后台登录限制', 'AdminLoginRestrictionService', '后台登录限制管理服务', 'AdminLoginRestrictionService_Create', '/admin/v1/login-restrictions', 'POST', 'ADMIN'),
      (64, now(), null, null, null, null, '更新通知消息接收者', 'NotificationMessageRecipientService', '通知消息接收者管理服务', 'NotificationMessageRecipientService_Update', '/admin/v1/notifications:recipients/{data.id}', 'PUT', 'ADMIN'),
      (65, now(), null, null, null, null, '租户是否存在', 'TenantService', '租户管理服务', 'TenantService_TenantExists', '/admin/v1/tenants_exists', 'GET', 'ADMIN'),
      (66, now(), null, null, null, null, '任务类型名称列表', 'TaskService', '调度任务管理服务', 'TaskService_ListTaskTypeName', '/admin/v1/tasks:type-names', 'GET', 'ADMIN'),
      (67, now(), null, null, null, null, '更新后台登录限制', 'AdminLoginRestrictionService', '后台登录限制管理服务', 'AdminLoginRestrictionService_Update', '/admin/v1/login-restrictions/{data.id}', 'PUT', 'ADMIN'),
      (68, now(), null, null, null, null, '删除租户', 'TenantService', '租户管理服务', 'TenantService_Delete', '/admin/v1/tenants/{id}', 'DELETE', 'ADMIN'),
      (69, now(), null, null, null, null, '获取租户数据', 'TenantService', '租户管理服务', 'TenantService_Get', '/admin/v1/tenants/{id}', 'GET', 'ADMIN'),
      (70, now(), null, null, null, null, '查询角色列表', 'RoleService', '角色管理服务', 'RoleService_List', '/admin/v1/roles', 'GET', 'ADMIN'),
      (71, now(), null, null, null, null, '创建角色', 'RoleService', '角色管理服务', 'RoleService_Create', '/admin/v1/roles', 'POST', 'ADMIN'),
      (72, now(), null, null, null, null, '查询路由列表', 'RouterService', '网站后台动态路由服务', 'RouterService_ListRoute', '/admin/v1/routes', 'GET', 'ADMIN'),
      (73, now(), null, null, null, null, '查询通知消息列表', 'NotificationMessageService', '通知消息管理服务', 'NotificationMessageService_List', '/admin/v1/notifications', 'GET', 'ADMIN'),
      (74, now(), null, null, null, null, '创建通知消息', 'NotificationMessageService', '通知消息管理服务', 'NotificationMessageService_Create', '/admin/v1/notifications', 'POST', 'ADMIN'),
      (75, now(), null, null, null, null, '获取对象存储（OSS）上传用的预签名链接', 'OssService', 'OSS服务', 'OssService_OssUploadUrl', '/admin/v1/file:upload-url', 'POST', 'ADMIN'),
      (76, now(), null, null, null, null, '删除文件', 'FileService', '文件管理服务', 'FileService_Delete', '/admin/v1/files/{id}', 'DELETE', 'ADMIN'),
      (77, now(), null, null, null, null, '查询文件详情', 'FileService', '文件管理服务', 'FileService_Get', '/admin/v1/files/{id}', 'GET', 'ADMIN'),
      (78, now(), null, null, null, null, '更新职位', 'PositionService', '职位管理服务', 'PositionService_Update', '/admin/v1/positions/{data.id}', 'PUT', 'ADMIN'),
      (79, now(), null, null, null, null, '更新用户', 'UserService', '用户管理服务', 'UserService_Update', '/admin/v1/users/{data.id}', 'PUT', 'ADMIN'),
      (80, now(), null, null, null, null, 'POST方法上传文件', 'OssService', 'OSS服务', 'OssService_PostUploadFile', '/admin/v1/file:upload', 'POST', 'ADMIN'),
      (81, now(), null, null, null, null, 'PUT方法上传文件', 'OssService', 'OSS服务', 'OssService_PutUploadFile', '/admin/v1/file:upload', 'PUT', 'ADMIN'),
      (82, now(), null, null, null, null, '更新通知消息', 'NotificationMessageService', '通知消息管理服务', 'NotificationMessageService_Update', '/admin/v1/notifications/{data.id}', 'PUT', 'ADMIN'),
      (83, now(), null, null, null, null, '删除通知消息分类', 'NotificationMessageCategoryService', '通知消息分类管理服务', 'NotificationMessageCategoryService_Delete', '/admin/v1/notifications:categories/{id}', 'DELETE', 'ADMIN'),
      (84, now(), null, null, null, null, '查询通知消息分类详情', 'NotificationMessageCategoryService', '通知消息分类管理服务', 'NotificationMessageCategoryService_Get', '/admin/v1/notifications:categories/{id}', 'GET', 'ADMIN'),
      (85, now(), null, null, null, null, '用户是否存在', 'UserService', '用户管理服务', 'UserService_UserExists', '/admin/v1/users_exists', 'GET', 'ADMIN'),
      (86, now(), null, null, null, null, '查询字典列表', 'DictService', '字典管理服务', 'DictService_List', '/admin/v1/dict', 'GET', 'ADMIN'),
      (87, now(), null, null, null, null, '创建字典', 'DictService', '字典管理服务', 'DictService_Create', '/admin/v1/dict', 'POST', 'ADMIN'),
      (88, now(), null, null, null, null, '查询文件列表', 'FileService', '文件管理服务', 'FileService_List', '/admin/v1/files', 'GET', 'ADMIN'),
      (89, now(), null, null, null, null, '创建文件', 'FileService', '文件管理服务', 'FileService_Create', '/admin/v1/files', 'POST', 'ADMIN'),
      (90, now(), null, null, null, null, '查询后台操作日志详情', 'AdminOperationLogService', '后台操作日志管理服务', 'AdminOperationLogService_Get', '/admin/v1/admin_operation_logs/{id}', 'GET', 'ADMIN'),
      (91, now(), null, null, null, null, '获取用户列表', 'UserService', '用户管理服务', 'UserService_List', '/admin/v1/users', 'GET', 'ADMIN'),
      (92, now(), null, null, null, null, '创建用户', 'UserService', '用户管理服务', 'UserService_Create', '/admin/v1/users', 'POST', 'ADMIN'),
      (93, now(), null, null, null, null, '刷新认证令牌', 'AuthenticationService', '用户后台登录认证服务', 'AuthenticationService_RefreshToken', '/admin/v1/refresh_token', 'POST', 'ADMIN'),
      (94, now(), null, null, null, null, '更新字典', 'DictService', '字典管理服务', 'DictService_Update', '/admin/v1/dict/{data.id}', 'PUT', 'ADMIN'),
      (95, now(), null, null, null, null, '查询API资源列表', 'ApiResourceService', 'API资源管理服务', 'ApiResourceService_List', '/admin/v1/api-resources', 'GET', 'ADMIN'),
      (96, now(), null, null, null, null, '创建API资源', 'ApiResourceService', 'API资源管理服务', 'ApiResourceService_Create', '/admin/v1/api-resources', 'POST', 'ADMIN'),
      (97, now(), null, null, null, null, '查询通知消息分类列表', 'NotificationMessageCategoryService', '通知消息分类管理服务', 'NotificationMessageCategoryService_List', '/admin/v1/notifications:categories', 'GET', 'ADMIN'),
      (98, now(), null, null, null, null, '创建通知消息分类', 'NotificationMessageCategoryService', '通知消息分类管理服务', 'NotificationMessageCategoryService_Create', '/admin/v1/notifications:categories', 'POST', 'ADMIN'),
      (99, now(), null, null, null, null, '更新菜单', 'MenuService', '后台菜单管理服务', 'MenuService_Update', '/admin/v1/menus/{data.id}', 'PUT', 'ADMIN'),
      (100, now(), null, null, null, null, '查询菜单列表', 'MenuService', '后台菜单管理服务', 'MenuService_List', '/admin/v1/menus', 'GET', 'ADMIN'),
      (101, now(), null, null, null, null, '创建菜单', 'MenuService', '后台菜单管理服务', 'MenuService_Create', '/admin/v1/menus', 'POST', 'ADMIN'),
      (102, now(), null, null, null, null, '删除调度任务', 'TaskService', '调度任务管理服务', 'TaskService_Delete', '/admin/v1/tasks/{id}', 'DELETE', 'ADMIN'),
      (103, now(), null, null, null, null, '查询调度任务详情', 'TaskService', '调度任务管理服务', 'TaskService_Get', '/admin/v1/tasks/{id}', 'GET', 'ADMIN'),
      (104, now(), null, null, null, null, '删除组织', 'OrganizationService', '组织管理服务', 'OrganizationService_Delete', '/admin/v1/organizations/{id}', 'DELETE', 'ADMIN'),
      (105, now(), null, null, null, null, '查询组织详情', 'OrganizationService', '组织管理服务', 'OrganizationService_Get', '/admin/v1/organizations/{id}', 'GET', 'ADMIN'),
      (106, now(), null, null, null, null, '修改用户密码', 'AuthenticationService', '用户后台登录认证服务', 'AuthenticationService_ChangePassword', '/admin/v1/change_password', 'POST', 'ADMIN'),
      (107, now(), null, null, null, null, '删除后台登录限制', 'AdminLoginRestrictionService', '后台登录限制管理服务', 'AdminLoginRestrictionService_Delete', '/admin/v1/login-restrictions/{id}', 'DELETE', 'ADMIN'),
      (108, now(), null, null, null, null, '查询后台登录限制详情', 'AdminLoginRestrictionService', '后台登录限制管理服务', 'AdminLoginRestrictionService_Get', '/admin/v1/login-restrictions/{id}', 'GET', 'ADMIN')
;
SELECT setval('sys_api_resources_id_seq', (SELECT MAX(id) FROM sys_api_resources));
