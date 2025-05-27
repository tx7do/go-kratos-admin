-- 默认的超级管理员，默认账号：admin，密码：admin
TRUNCATE TABLE kratos_admin.public.users;
INSERT INTO kratos_admin.public.users (id, username, nick_name, email, authority, role_id)
VALUES (1, 'admin', 'admin', 'admin@gmail.com', 'SYS_ADMIN', 1);

TRUNCATE TABLE user_credentials;
INSERT INTO user_credentials (id, create_time, user_id, identity_type, identifier, credential_type, credential, status, is_primary)
VALUES (1, now(), 1, 'PASSWORD', 'admin', 'PASSWORD_HASH', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a', 'ENABLED', true);

-- 默认的角色
TRUNCATE TABLE kratos_admin.public.roles;
INSERT INTO kratos_admin.public.roles(id, parent_id, create_by, sort_id, name, code, status, remark, menus, create_time)
VALUES (1, null, 0, 1, '超级管理员', 'super', 'ON', '超级管理员拥有对系统的最高权限', '[1, 2, 6, 7, 8, 9, 10, 11, 12, 13, 14, 20, 21, 22]', now()),
       (2, null, 0, 2, '管理员', 'admin', 'ON', '系统管理员拥有对整个系统的管理权限', '[1, 2, 6, 7, 8, 9, 10, 11, 12, 13, 14]', now()),
       (3, null, 0, 3, '普通用户', 'user', 'ON', '普通用户没有管理权限，只有设备和APP的使用权限', '[]', now()),
       (4, null, 0, 4, '游客', 'guest', 'ON', '游客只有非常有限的数据读取权限', '[]', now());

-- 后台目录
TRUNCATE TABLE kratos_admin.public.menus;
INSERT INTO kratos_admin.public.menus(id, parent_id, type, name, path, redirect, component, status, create_time, meta)
VALUES (1, null, 'FOLDER', 'Dashboard', '/', null, 'BasicLayout', 'ON', now(), '{"order":-1, "title":"page.dashboard.title", "icon":"lucide:layout-dashboard", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (2, 1, 'MENU', 'Analytics', '/analytics', null, 'dashboard/analytics/index.vue', 'ON', now(), '{"order":-1, "title":"page.dashboard.analytics", "icon":"lucide:area-chart", "affixTab": true, "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (10, null, 'FOLDER', 'Auth', '/auth', null, 'BasicLayout', 'ON', now(), '{"order":2000, "title":"menu.auth.moduleName", "icon":"lucide:shield-check", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (11, 10, 'MENU', 'UserManagement', 'users', null, 'app/auth/users/index.vue', 'ON', now(), '{"order":1, "title":"menu.auth.user", "icon":"lucide:users", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (12, 10, 'MENU', 'UserDetail', 'users/detail/:id', null, 'app/auth/users/detail/index.vue', 'ON', now(), '{"order":2, "title":"menu.auth.userDetail", "icon":"", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":true, "hideInTab":false}'),
       (13, 10, 'MENU', 'TenantManagement', 'tenants', null, 'app/auth/tenants/index.vue', 'ON', now(), '{"order":3, "title":"menu.auth.tenant", "icon":"lucide:book-user", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (14, 10, 'MENU', 'RoleManagement', 'roles', null, 'app/auth/role/index.vue', 'ON', now(), '{"order":4, "title":"menu.auth.role", "icon":"lucide:chef-hat", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (15, 10, 'MENU', 'OrganizationManagement', 'organizations', null, 'app/auth/org/index.vue', 'ON', now(), '{"order":5, "title":"menu.auth.org", "icon":"lucide:building-2", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (16, 10, 'MENU', 'DepartmentManagement', 'departments', null, 'app/auth/dept/index.vue', 'ON', now(), '{"order":6, "title":"menu.auth.dept", "icon":"lucide:network", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (17, 10, 'MENU', 'PositionManagement', 'positions', null, 'app/auth/position/index.vue', 'ON', now(), '{"order":7, "title":"menu.auth.position", "icon":"lucide:id-card", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (18, 10, 'MENU', 'MenuManagement', 'menus', null, 'app/auth/menu/index.vue', 'ON', now(), '{"order":8, "title":"menu.auth.menu", "icon":"lucide:square-menu", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (20, null, 'FOLDER', 'System', '/system', null, 'BasicLayout', 'ON', now(), '{"order":2001, "title":"menu.system.moduleName", "icon":"lucide:settings", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (21, 20, 'MENU', 'DictManagement', 'dict', null, 'app/system/dict/index.vue', 'ON', now(), '{"order":1, "title":"menu.system.dict", "icon":"lucide:library-big", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (22, 20, 'MENU', 'FileManagement', 'files', null, 'app/system/files/index.vue', 'ON', now(), '{"order":2, "title":"menu.system.file", "icon":"lucide:file", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (23, 20, 'MENU', 'TaskManagement', 'tasks', null, 'app/system/task/index.vue', 'ON', now(), '{"order":3, "title":"menu.system.task", "icon":"lucide:calendar-clock", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (24, 20, 'MENU', 'NotificationMessageManagement', 'notifications', null, 'app/system/notification_message/index.vue', 'ON', now(), '{"order":4, "title":"menu.system.notificationMessage", "icon":"lucide:bell", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (25, 20, 'MENU', 'NotificationMessageCategoryManagement', 'notification_categories', null, 'app/system/notification_message_category/index.vue', 'ON', now(), '{"order":5, "title":"menu.system.notificationMessageCategory", "icon":"lucide:bell-dot", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (26, 20, 'MENU', 'PrivateMessageManagement', 'private_messages', null, 'app/system/private_message/index.vue', 'ON', now(), '{"order":6, "title":"menu.system.privateMessage", "icon":"lucide:message-circle-more", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (30, null, 'FOLDER', 'Log', '/log', null, 'BasicLayout', 'ON', now(), '{"order":2002, "title":"menu.log.moduleName", "icon":"lucide:logs", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (31, 30, 'MENU', 'AdminLoginLog', 'login', null, 'app/log/admin_login_log/index.vue', 'ON', now(), '{"order":1, "title":"menu.log.adminLoginLog", "icon":"lucide:log-in", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (32, 30, 'MENU', 'AdminOperationLog', 'operation', null, 'app/log/admin_operation_log/index.vue', 'ON', now(), '{"order":2, "title":"menu.log.adminOperationLog", "icon":"lucide:arrow-up-down", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}');
