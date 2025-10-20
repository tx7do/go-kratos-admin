-- 插入4个权限的用户
TRUNCATE TABLE kratos_admin.public.users RESTART IDENTITY;
INSERT INTO kratos_admin.public.users (username, nickname, realname, email, authority, role_ids, gender)
VALUES ('admin', '鹳狸猿', '喵个咪', 'admin@gmail.com', 'SYS_ADMIN', '[1]', 'MALE'),
       -- 2. 租户管理员（TENANT_ADMIN）
       ('tenant_admin', '租户管理', '张管理员', 'tenant@company.com', 'TENANT_ADMIN', '[2]', 'MALE'),
       -- 3. 普通用户（CUSTOMER_USER）
       ('normal_user', '普通用户', '李用户', 'user@company.com', 'CUSTOMER_USER', '[3]', 'FEMALE'),
       -- 4. 访客（GUEST）
       ('guest_user', '临时访客', '王访客', 'guest@company.com', 'GUEST', '[4]', 'SECRET')
;
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));

-- 插入4个用户的凭证（密码统一为admin，哈希值与原admin一致，方便测试）
TRUNCATE TABLE user_credentials RESTART IDENTITY;
INSERT INTO user_credentials (user_id, identity_type, identifier, credential_type, credential, status, is_primary, create_time)
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
SELECT setval('user_credentials_id_seq', (SELECT MAX(id) FROM user_credentials));

-- 默认的角色
TRUNCATE TABLE kratos_admin.public.sys_roles RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_roles(id, parent_id, create_by, sort_id, name, code, status, remark, menus, apis, create_time)
VALUES (1, null, 0, 1, '超级管理员', 'super', 'ON', '拥有系统所有功能的操作权限，可管理租户、用户、角色及所有资源',
        '[1, 2, 10, 11, 20, 21, 22, 23, 24, 25, 30, 31, 32, 40, 41, 42, 43, 50, 51, 52, 60, 61, 62, 63, 64, 65]', '[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102]', now()),
       (2, null, 0, 2, '租户管理员', 'tenant_admin', 'ON', '管理当前租户下的用户、角色及资源，无跨租户操作权限', '[1, 2, 6, 7, 8, 9, 10, 11, 12, 13, 14]', '[]', now()),
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
INSERT INTO public.sys_api_resources
(id, create_time, update_time, delete_time, create_by, update_by, operation, description, module, path, method, module_description)
VALUES
    (1, now(), null, null, null, null, 'TenantService_Delete', '删除租户', 'TenantService', '/admin/v1/tenants/{id}', 'DELETE', '租户管理服务'),
    (2, now(), null, null, null, null, 'TenantService_Get', '获取租户数据', 'TenantService', '/admin/v1/tenants/{id}', 'GET', '租户管理服务'),
    (3, now(), null, null, null, null, 'TaskService_Delete', '删除调度任务', 'TaskService', '/admin/v1/tasks/{id}', 'DELETE', '调度任务管理服务'),
    (4, now(), null, null, null, null, 'TaskService_Get', '查询调度任务详情', 'TaskService', '/admin/v1/tasks/{id}', 'GET', '调度任务管理服务'),
    (5, now(), null, null, null, null, 'TaskService_Update', '更新调度任务', 'TaskService', '/admin/v1/tasks/{data.id}', 'PUT', '调度任务管理服务'),
    (6, now(), null, null, null, null, 'NotificationMessageRecipientService_Delete', '删除通知消息接收者', 'NotificationMessageRecipientService', '/admin/v1/notifications:recipients/{id}', 'DELETE', '通知消息接收者管理服务'),
    (7, now(), null, null, null, null, 'NotificationMessageRecipientService_Get', '查询通知消息接收者详情', 'NotificationMessageRecipientService', '/admin/v1/notifications:recipients/{id}', 'GET', '通知消息接收者管理服务'),
    (8, now(), null, null, null, null, 'UserService_Delete', '删除用户', 'UserService', '/admin/v1/users/{id}', 'DELETE', '用户管理服务'),
    (9, now(), null, null, null, null, 'UserService_Get', '获取用户数据', 'UserService', '/admin/v1/users/{id}', 'GET', '用户管理服务'),
    (10, now(), null, null, null, null, 'PositionService_Delete', '删除职位', 'PositionService', '/admin/v1/positions/{id}', 'DELETE', '职位管理服务'),
    (11, now(), null, null, null, null, 'PositionService_Get', '查询职位详情', 'PositionService', '/admin/v1/positions/{id}', 'GET', '职位管理服务'),
    (12, now(), null, null, null, null, 'UserService_Update', '更新用户', 'UserService', '/admin/v1/users/{data.id}', 'PUT', '用户管理服务'),
    (13, now(), null, null, null, null, 'PositionService_Update', '更新职位', 'PositionService', '/admin/v1/positions/{data.id}', 'PUT', '职位管理服务'),
    (14, now(), null, null, null, null, 'ApiResourceService_GetWalkRouteData', '查询路由数据', 'ApiResourceService', '/admin/v1/api-resources/walk-route', 'GET', 'API资源管理服务'),
    (15, now(), null, null, null, null, 'OssService_OssUploadUrl', '获取对象存储（OSS）上传用的预签名链接', 'OssService', '/admin/v1/file:upload-url', 'POST', 'OSS服务'),
    (16, now(), null, null, null, null, 'UEditorService_UEditorAPI', 'UEditor API', 'UEditorService', '/admin/v1/ueditor', 'GET', 'UEditor后端服务'),
    (17, now(), null, null, null, null, 'UEditorService_UploadFile', '上传文件', 'UEditorService', '/admin/v1/ueditor', 'POST', 'UEditor后端服务'),
    (18, now(), null, null, null, null, 'NotificationMessageService_List', '查询通知消息列表', 'NotificationMessageService', '/admin/v1/notifications', 'GET', '通知消息管理服务'),
    (19, now(), null, null, null, null, 'NotificationMessageService_Create', '创建通知消息', 'NotificationMessageService', '/admin/v1/notifications', 'POST', '通知消息管理服务'),
    (20, now(), null, null, null, null, 'TaskService_RestartAllTask', '重启所有的调度任务', 'TaskService', '/admin/v1/tasks:restart', 'POST', '调度任务管理服务'),
    (21, now(), null, null, null, null, 'AdminOperationLogService_List', '查询后台操作日志列表', 'AdminOperationLogService', '/admin/v1/admin_operation_logs', 'GET', '后台操作日志管理服务'),
    (22, now(), null, null, null, null, 'UserProfileService_GetUser', '获取用户资料', 'UserProfileService', '/admin/v1/me', 'GET', '用户个人资料服务'),
    (23, now(), null, null, null, null, 'UserProfileService_UpdateUser', '更新用户资料', 'UserProfileService', '/admin/v1/me', 'PUT', '用户个人资料服务'),
    (24, now(), null, null, null, null, 'DictService_Create', '创建字典', 'DictService', '/admin/v1/dict', 'POST', '字典管理服务'),
    (25, now(), null, null, null, null, 'DictService_List', '查询字典列表', 'DictService', '/admin/v1/dict', 'GET', '字典管理服务'),
    (26, now(), null, null, null, null, 'AdminLoginRestrictionService_List', '查询后台登录限制列表', 'AdminLoginRestrictionService', '/admin/v1/login-restrictions', 'GET', '后台登录限制管理服务'),
    (27, now(), null, null, null, null, 'AdminLoginRestrictionService_Create', '创建后台登录限制', 'AdminLoginRestrictionService', '/admin/v1/login-restrictions', 'POST', '后台登录限制管理服务'),
    (28, now(), null, null, null, null, 'RoleService_Delete', '删除角色', 'RoleService', '/admin/v1/roles/{id}', 'DELETE', '角色管理服务'),
    (29, now(), null, null, null, null, 'RoleService_Get', '查询角色详情', 'RoleService', '/admin/v1/roles/{id}', 'GET', '角色管理服务'),
    (30, now(), null, null, null, null, 'AuthenticationService_ChangePassword', '修改用户密码', 'AuthenticationService', '/admin/v1/change_password', 'POST', '用户后台登录认证服务'),
    (31, now(), null, null, null, null, 'NotificationMessageService_Update', '更新通知消息', 'NotificationMessageService', '/admin/v1/notifications/{data.id}', 'PUT', '通知消息管理服务'),
    (32, now(), null, null, null, null, 'TaskService_List', '查询调度任务列表', 'TaskService', '/admin/v1/tasks', 'GET', '调度任务管理服务'),
    (33, now(), null, null, null, null, 'TaskService_Create', '创建调度任务', 'TaskService', '/admin/v1/tasks', 'POST', '调度任务管理服务'),
    (34, now(), null, null, null, null, 'FileService_Update', '更新文件', 'FileService', '/admin/v1/files/{data.id}', 'PUT', '文件管理服务'),
    (35, now(), null, null, null, null, 'AdminOperationLogService_Get', '查询后台操作日志详情', 'AdminOperationLogService', '/admin/v1/admin_operation_logs/{id}', 'GET', '后台操作日志管理服务'),
    (36, now(), null, null, null, null, 'MenuService_Delete', '删除菜单', 'MenuService', '/admin/v1/menus/{id}', 'DELETE', '后台菜单管理服务'),
    (37, now(), null, null, null, null, 'MenuService_Get', '查询菜单详情', 'MenuService', '/admin/v1/menus/{id}', 'GET', '后台菜单管理服务'),
    (38, now(), null, null, null, null, 'TaskService_StopAllTask', '停止所有的调度任务', 'TaskService', '/admin/v1/tasks:stop', 'POST', '调度任务管理服务'),
    (39, now(), null, null, null, null, 'PrivateMessageService_Update', '更新私信消息', 'PrivateMessageService', '/admin/v1/private_messages/{data.id}', 'PUT', '私信消息管理服务'),
    (40, now(), null, null, null, null, 'PrivateMessageService_Delete', '删除私信消息', 'PrivateMessageService', '/admin/v1/private_messages/{id}', 'DELETE', '私信消息管理服务'),
    (41, now(), null, null, null, null, 'PrivateMessageService_Get', '查询私信消息详情', 'PrivateMessageService', '/admin/v1/private_messages/{id}', 'GET', '私信消息管理服务'),
    (42, now(), null, null, null, null, 'DictService_Delete', '删除字典', 'DictService', '/admin/v1/dict/{id}', 'DELETE', '字典管理服务'),
    (43, now(), null, null, null, null, 'DictService_Get', '查询字典详情', 'DictService', '/admin/v1/dict/{id}', 'GET', '字典管理服务'),
    (44, now(), null, null, null, null, 'FileService_List', '查询文件列表', 'FileService', '/admin/v1/files', 'GET', '文件管理服务'),
    (45, now(), null, null, null, null, 'FileService_Create', '创建文件', 'FileService', '/admin/v1/files', 'POST', '文件管理服务'),
    (46, now(), null, null, null, null, 'FileService_Delete', '删除文件', 'FileService', '/admin/v1/files/{id}', 'DELETE', '文件管理服务'),
    (47, now(), null, null, null, null, 'FileService_Get', '查询文件详情', 'FileService', '/admin/v1/files/{id}', 'GET', '文件管理服务'),
    (48, now(), null, null, null, null, 'RoleService_List', '查询角色列表', 'RoleService', '/admin/v1/roles', 'GET', '角色管理服务'),
    (49, now(), null, null, null, null, 'RoleService_Create', '创建角色', 'RoleService', '/admin/v1/roles', 'POST', '角色管理服务'),
    (50, now(), null, null, null, null, 'AdminLoginRestrictionService_Update', '更新后台登录限制', 'AdminLoginRestrictionService', '/admin/v1/login-restrictions/{data.id}', 'PUT', '后台登录限制管理服务'),
    (51, now(), null, null, null, null, 'NotificationMessageCategoryService_List', '查询通知消息分类列表', 'NotificationMessageCategoryService', '/admin/v1/notifications:categories', 'GET', '通知消息分类管理服务'),
    (52, now(), null, null, null, null, 'NotificationMessageCategoryService_Create', '创建通知消息分类', 'NotificationMessageCategoryService', '/admin/v1/notifications:categories', 'POST', '通知消息分类管理服务'),
    (53, now(), null, null, null, null, 'DepartmentService_Update', '更新部门', 'DepartmentService', '/admin/v1/departments/{data.id}', 'PUT', '部门管理服务'),
    (54, now(), null, null, null, null, 'UserService_List', '获取用户列表', 'UserService', '/admin/v1/users', 'GET', '用户管理服务'),
    (55, now(), null, null, null, null, 'UserService_Create', '创建用户', 'UserService', '/admin/v1/users', 'POST', '用户管理服务'),
    (56, now(), null, null, null, null, 'RouterService_ListRoute', '查询路由列表', 'RouterService', '/admin/v1/routes', 'GET', '网站后台动态路由服务'),
    (57, now(), null, null, null, null, 'AdminLoginLogService_List', '查询后台登录日志列表', 'AdminLoginLogService', '/admin/v1/admin_login_logs', 'GET', '后台登录日志管理服务'),
    (58, now(), null, null, null, null, 'AdminLoginLogService_Get', '查询后台登录日志详情', 'AdminLoginLogService', '/admin/v1/admin_login_logs/{id}', 'GET', '后台登录日志管理服务'),
    (59, now(), null, null, null, null, 'DepartmentService_List', '查询部门列表', 'DepartmentService', '/admin/v1/departments', 'GET', '部门管理服务'),
    (60, now(), null, null, null, null, 'DepartmentService_Create', '创建部门', 'DepartmentService', '/admin/v1/departments', 'POST', '部门管理服务'),
    (61, now(), null, null, null, null, 'PositionService_List', '查询职位列表', 'PositionService', '/admin/v1/positions', 'GET', '职位管理服务'),
    (62, now(), null, null, null, null, 'PositionService_Create', '创建职位', 'PositionService', '/admin/v1/positions', 'POST', '职位管理服务'),
    (63, now(), null, null, null, null, 'ApiResourceService_List', '查询API资源列表', 'ApiResourceService', '/admin/v1/api-resources', 'GET', 'API资源管理服务'),
    (64, now(), null, null, null, null, 'ApiResourceService_Create', '创建API资源', 'ApiResourceService', '/admin/v1/api-resources', 'POST', 'API资源管理服务'),
    (65, now(), null, null, null, null, 'PrivateMessageService_List', '查询私信消息列表', 'PrivateMessageService', '/admin/v1/private_messages', 'GET', '私信消息管理服务'),
    (66, now(), null, null, null, null, 'PrivateMessageService_Create', '创建私信消息', 'PrivateMessageService', '/admin/v1/private_messages', 'POST', '私信消息管理服务'),
    (67, now(), null, null, null, null, 'OssService_PostUploadFile', 'POST方法上传文件', 'OssService', '/admin/v1/file:upload', 'POST', 'OSS服务'),
    (68, now(), null, null, null, null, 'OssService_PutUploadFile', 'PUT方法上传文件', 'OssService', '/admin/v1/file:upload', 'PUT', 'OSS服务'),
    (69, now(), null, null, null, null, 'NotificationMessageCategoryService_Update', '更新通知消息分类', 'NotificationMessageCategoryService', '/admin/v1/notifications:categories/{data.id}', 'PUT', '通知消息分类管理服务'),
    (70, now(), null, null, null, null, 'NotificationMessageService_Delete', '删除通知消息', 'NotificationMessageService', '/admin/v1/notifications/{id}', 'DELETE', '通知消息管理服务'),
    (71, now(), null, null, null, null, 'NotificationMessageService_Get', '查询通知消息详情', 'NotificationMessageService', '/admin/v1/notifications/{id}', 'GET', '通知消息管理服务'),
    (72, now(), null, null, null, null, 'DictService_Update', '更新字典', 'DictService', '/admin/v1/dict/{data.id}', 'PUT', '字典管理服务'),
    (73, now(), null, null, null, null, 'OrganizationService_Update', '更新组织', 'OrganizationService', '/admin/v1/organizations/{data.id}', 'PUT', '组织管理服务'),
    (74, now(), null, null, null, null, 'TaskService_ControlTask', '控制调度任务', 'TaskService', '/admin/v1/tasks:control', 'POST', '调度任务管理服务'),
    (75, now(), null, null, null, null, 'AdminLoginRestrictionService_Delete', '删除后台登录限制', 'AdminLoginRestrictionService', '/admin/v1/login-restrictions/{id}', 'DELETE', '后台登录限制管理服务'),
    (76, now(), null, null, null, null, 'AdminLoginRestrictionService_Get', '查询后台登录限制详情', 'AdminLoginRestrictionService', '/admin/v1/login-restrictions/{id}', 'GET', '后台登录限制管理服务'),
    (77, now(), null, null, null, null, 'ApiResourceService_SyncApiResources', '同步API资源', 'ApiResourceService', '/admin/v1/api-resources/sync', 'POST', 'API资源管理服务'),
    (78, now(), null, null, null, null, 'RouterService_ListPermissionCode', '查询权限码列表', 'RouterService', '/admin/v1/perm-codes', 'GET', '网站后台动态路由服务'),
    (79, now(), null, null, null, null, 'RoleService_Update', '更新角色', 'RoleService', '/admin/v1/roles/{data.id}', 'PUT', '角色管理服务'),
    (80, now(), null, null, null, null, 'NotificationMessageRecipientService_Update', '更新通知消息接收者', 'NotificationMessageRecipientService', '/admin/v1/notifications:recipients/{data.id}', 'PUT', '通知消息接收者管理服务'),
    (81, now(), null, null, null, null, 'MenuService_List', '查询菜单列表', 'MenuService', '/admin/v1/menus', 'GET', '后台菜单管理服务'),
    (82, now(), null, null, null, null, 'MenuService_Create', '创建菜单', 'MenuService', '/admin/v1/menus', 'POST', '后台菜单管理服务'),
    (83, now(), null, null, null, null, 'NotificationMessageCategoryService_Delete', '删除通知消息分类', 'NotificationMessageCategoryService', '/admin/v1/notifications:categories/{id}', 'DELETE', '通知消息分类管理服务'),
    (84, now(), null, null, null, null, 'NotificationMessageCategoryService_Get', '查询通知消息分类详情', 'NotificationMessageCategoryService', '/admin/v1/notifications:categories/{id}', 'GET', '通知消息分类管理服务'),
    (85, now(), null, null, null, null, 'TenantService_List', '获取租户列表', 'TenantService', '/admin/v1/tenants', 'GET', '租户管理服务'),
    (86, now(), null, null, null, null, 'TenantService_Create', '创建租户', 'TenantService', '/admin/v1/tenants', 'POST', '租户管理服务'),
    (87, now(), null, null, null, null, 'NotificationMessageRecipientService_List', '查询通知消息接收者列表', 'NotificationMessageRecipientService', '/admin/v1/notifications:recipients', 'GET', '通知消息接收者管理服务'),
    (88, now(), null, null, null, null, 'NotificationMessageRecipientService_Create', '创建通知消息接收者', 'NotificationMessageRecipientService', '/admin/v1/notifications:recipients', 'POST', '通知消息接收者管理服务'),
    (89, now(), null, null, null, null, 'AuthenticationService_Logout', '登出', 'AuthenticationService', '/admin/v1/logout', 'POST', '用户后台登录认证服务'),
    (90, now(), null, null, null, null, 'TenantService_Update', '更新租户', 'TenantService', '/admin/v1/tenants/{data.id}', 'PUT', '租户管理服务'),
    (91, now(), null, null, null, null, 'ApiResourceService_Delete', '删除API资源', 'ApiResourceService', '/admin/v1/api-resources/{id}', 'DELETE', 'API资源管理服务'),
    (92, now(), null, null, null, null, 'ApiResourceService_Get', '查询API资源详情', 'ApiResourceService', '/admin/v1/api-resources/{id}', 'GET', 'API资源管理服务'),
    (93, now(), null, null, null, null, 'DepartmentService_Delete', '删除部门', 'DepartmentService', '/admin/v1/departments/{id}', 'DELETE', '部门管理服务'),
    (94, now(), null, null, null, null, 'DepartmentService_Get', '查询部门详情', 'DepartmentService', '/admin/v1/departments/{id}', 'GET', '部门管理服务'),
    (95, now(), null, null, null, null, 'ApiResourceService_Update', '更新API资源', 'ApiResourceService', '/admin/v1/api-resources/{data.id}', 'PUT', 'API资源管理服务'),
    (96, now(), null, null, null, null, 'AuthenticationService_Login', '登录', 'AuthenticationService', '/admin/v1/login', 'POST', '用户后台登录认证服务'),
    (97, now(), null, null, null, null, 'OrganizationService_List', '查询组织列表', 'OrganizationService', '/admin/v1/organizations', 'GET', '组织管理服务'),
    (98, now(), null, null, null, null, 'OrganizationService_Create', '创建组织', 'OrganizationService', '/admin/v1/organizations', 'POST', '组织管理服务'),
    (99, now(), null, null, null, null, 'AuthenticationService_RefreshToken', '刷新认证令牌', 'AuthenticationService', '/admin/v1/refresh_token', 'POST', '用户后台登录认证服务'),
    (100, now(), null, null, null, null, 'OrganizationService_Delete', '删除组织', 'OrganizationService', '/admin/v1/organizations/{id}', 'DELETE', '组织管理服务'),
    (101, now(), null, null, null, null, 'OrganizationService_Get', '查询组织详情', 'OrganizationService', '/admin/v1/organizations/{id}', 'GET', '组织管理服务'),
    (102, now(), null, null, null, null, 'MenuService_Update', '更新菜单', 'MenuService', '/admin/v1/menus/{data.id}', 'PUT', '后台菜单管理服务');
SELECT setval('sys_api_resources_id_seq', (SELECT MAX(id) FROM sys_api_resources));
