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
VALUES (1, null, 'FOLDER', 'Dashboard', '/', null, 'BasicLayout', 'ON', now(),
        '{"order":-1, "title":"page.dashboard.title", "icon":"lucide:layout-dashboard", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (2, 1, 'MENU', 'Analytics', '/analytics', null, 'dashboard/analytics/index.vue', 'ON', now(),
        '{"order":-1, "title":"page.dashboard.analytics", "icon":"lucide:area-chart", "affixTab": true, "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (10, null, 'FOLDER', 'TenantManagement', '/tenant', null, 'BasicLayout', 'ON', now(),
        '{"order":2001, "title":"menu.tenant.moduleName", "icon":"lucide:earth", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (11, 10, 'MENU', 'TenantMemberManagement', 'members', null, 'app/tenant/tenant/index.vue', 'ON', now(),
        '{"order":1, "title":"menu.tenant.member", "icon":"lucide:book-user", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (20, null, 'FOLDER', 'OrganizationalPersonnelManagement', '/opm', null, 'BasicLayout', 'ON', now(),
        '{"order":2002, "title":"menu.opm.moduleName", "icon":"lucide:shield-check", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (21, 20, 'MENU', 'UserManagement', 'users', null, 'app/opm/users/index.vue', 'ON', now(),
        '{"order":1, "title":"menu.opm.user", "icon":"lucide:users", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (22, 20, 'MENU', 'UserDetail', 'users/detail/:id', null, 'app/opm/users/detail/index.vue', 'ON', now(),
        '{"order":2, "title":"menu.opm.userDetail", "icon":"", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":true, "hideInTab":false}'),
       (25, 20, 'MENU', 'OrganizationManagement', 'organizations', null, 'app/opm/org/index.vue', 'ON', now(),
        '{"order":3, "title":"menu.opm.org", "icon":"lucide:building-2", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (26, 20, 'MENU', 'DepartmentManagement', 'departments', null, 'app/opm/dept/index.vue', 'ON', now(),
        '{"order":4, "title":"menu.opm.dept", "icon":"lucide:network", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (27, 20, 'MENU', 'PositionManagement', 'positions', null, 'app/opm/position/index.vue', 'ON', now(),
        '{"order":5, "title":"menu.opm.position", "icon":"lucide:id-card", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (30, null, 'FOLDER', 'PermissionManagement', '/permission', null, 'BasicLayout', 'ON', now(),
        '{"order":2003, "title":"menu.permission.moduleName", "icon":"lucide:shield-check", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (31, 30, 'MENU', 'RoleManagement', 'roles', null, 'app/permission/role/index.vue', 'ON', now(),
        '{"order":1, "title":"menu.permission.role", "icon":"lucide:user-round-cog", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (32, 30, 'MENU', 'MenuManagement', 'menus', null, 'app/permission/menu/index.vue', 'ON', now(),
        '{"order":2, "title":"menu.permission.menu", "icon":"lucide:layout-list", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (40, null, 'FOLDER', 'LogAuditManagement', '/log', null, 'BasicLayout', 'ON', now(),
        '{"order":2002, "title":"menu.log.moduleName", "icon":"lucide:logs", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (41, 40, 'MENU', 'AdminLoginLog', 'login', null, 'app/log/admin_login_log/index.vue', 'ON', now(),
        '{"order":1, "title":"menu.log.adminLoginLog", "icon":"lucide:file-symlink", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (42, 40, 'MENU', 'AdminOperationLog', 'operation', null, 'app/log/admin_operation_log/index.vue', 'ON', now(),
        '{"order":2, "title":"menu.log.adminOperationLog", "icon":"lucide:file-sliders", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (50, null, 'FOLDER', 'System', '/system', null, 'BasicLayout', 'ON', now(),
        '{"order":2004, "title":"menu.system.moduleName", "icon":"lucide:settings", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (51, 50, 'MENU', 'DictManagement', 'dict', null, 'app/system/dict/index.vue', 'ON', now(),
        '{"order":1, "title":"menu.system.dict", "icon":"lucide:library-big", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (52, 50, 'MENU', 'FileManagement', 'files', null, 'app/system/files/index.vue', 'ON', now(),
        '{"order":2, "title":"menu.system.file", "icon":"lucide:file", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (53, 50, 'MENU', 'TaskManagement', 'tasks', null, 'app/system/task/index.vue', 'ON', now(),
        '{"order":3, "title":"menu.system.task", "icon":"lucide:calendar-clock", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (54, 50, 'MENU', 'APIResourceManagement', 'apis', null, 'app/system/api_resource/index.vue', 'ON', now(),
        '{"order":4, "title":"menu.system.apiResource", "icon":"lucide:route", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (55, 50, 'MENU', 'NotificationMessageManagement', 'notifications', null,
        'app/system/notification_message/index.vue', 'ON', now(),
        '{"order":5, "title":"menu.system.notificationMessage", "icon":"lucide:bell", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (56, 50, 'MENU', 'NotificationMessageCategoryManagement', 'notification_categories', null,
        'app/system/notification_message_category/index.vue', 'ON', now(),
        '{"order":6, "title":"menu.system.notificationMessageCategory", "icon":"lucide:bell-dot", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (57, 50, 'MENU', 'PrivateMessageManagement', 'private_messages', null, 'app/system/private_message/index.vue',
        'ON', now(),
        '{"order":7, "title":"menu.system.privateMessage", "icon":"lucide:message-circle-more", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}');
SELECT setval('sys_menus_id_seq', (SELECT MAX(id) FROM sys_menus));
