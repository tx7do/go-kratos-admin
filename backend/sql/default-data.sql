-- 默认的超级管理员，默认账号：admin，密码：admin
TRUNCATE TABLE kratos_admin.public.users;
INSERT INTO kratos_admin.public.users (id, username, nick_name, email, password, authority, role_id)
VALUES (1, 'admin', 'admin', 'admin@gmail.com', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a', 'SYS_ADMIN', 1);

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

       (6, null, 'FOLDER', 'System', '/system', null, 'BasicLayout', 'ON', now(), '{"order":2000, "title":"menu.system.moduleName", "icon":"lucide:settings", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (7, 6, 'MENU', 'DictManagement', 'dict', null, 'app/system/dict/index.vue', 'ON', now(), '{"order":1, "title":"menu.system.dict", "icon":"lucide:library-big", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (8, 6, 'MENU', 'MenuManagement', 'menus', null, 'app/system/menu/index.vue', 'ON', now(), '{"order":2, "title":"menu.system.menu", "icon":"lucide:square-menu", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (9, 6, 'MENU', 'UserManagement', 'users', null, 'app/system/users/index.vue', 'ON', now(), '{"order":3, "title":"menu.system.user", "icon":"lucide:users", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (10, 6, 'MENU', 'UserDetail', 'users/detail/:id', null, 'app/system/users/detail/index.vue', 'ON', now(), '{"order":4, "title":"menu.system.user_detail", "icon":"", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":true, "hideInTab":false}'),
       (11, 6, 'MENU', 'RoleManagement', 'roles', null, 'app/system/role/index.vue', 'ON', now(), '{"order":5, "title":"menu.system.role", "icon":"lucide:shirt", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (12, 6, 'MENU', 'OrganizationManagement', 'organizations', null, 'app/system/org/index.vue', 'ON', now(), '{"order":6, "title":"menu.system.org", "icon":"lucide:building-2", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (13, 6, 'MENU', 'DepartmentManagement', 'departments', null, 'app/system/dept/index.vue', 'ON', now(), '{"order":7, "title":"menu.system.dept", "icon":"lucide:network", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (14, 6, 'MENU', 'PositionManagement', 'positions', null, 'app/system/position/index.vue', 'ON', now(), '{"order":8, "title":"menu.system.position", "icon":"lucide:id-card", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (20, null, 'FOLDER', 'Log', '/log', null, 'BasicLayout', 'ON', now(), '{"order":2001, "title":"menu.log.moduleName", "icon":"lucide:logs", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (21, 20, 'MENU', 'AdminLoginLog', 'login', null, 'app/log/admin_login_log/index.vue', 'ON', now(), '{"order":1, "title":"menu.log.admin_login_log", "icon":"lucide:log-in", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (22, 20, 'MENU', 'AdminOperationLog', 'operation', null, 'app/log/admin_operation_log/index.vue', 'ON', now(), '{"order":2, "title":"menu.log.admin_operation_log", "icon":"lucide:arrow-up-down", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}');
