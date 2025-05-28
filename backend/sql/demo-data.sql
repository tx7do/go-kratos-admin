-- 租户
TRUNCATE TABLE kratos_admin.public.tenants RESTART IDENTITY;
INSERT INTO kratos_admin.public.tenants(id, name, code, status, create_time)
VALUES (1, '超级租户', 'super', 'ON', now()),
       (2, '测试租户', 'test', 'ON', now()),
       (3, '测试租户2', 'test2', 'ON', now())
;
SELECT setval('tenants_id_seq', (SELECT MAX(id) FROM tenants));

-- 组织
TRUNCATE TABLE kratos_admin.public.organizations RESTART IDENTITY;
INSERT INTO kratos_admin.public.organizations(id, parent_id, sort_id, name, status, create_time)
VALUES (1, null, 1, '华东分部', 'ON', now()),
       (2, null, 1, '华南分部', 'ON', now()),
       (3, null, 2, '西北分部', 'ON', now())
;
SELECT setval('organizations_id_seq', (SELECT MAX(id) FROM organizations));

-- 部门
TRUNCATE TABLE kratos_admin.public.departments RESTART IDENTITY;
INSERT INTO kratos_admin.public.departments(id, parent_id, sort_id, name, status, create_time)
VALUES (1, null, 1, '华东分部', 'ON', now()),
       (10, 1, 1, '研发部', 'ON', now()),
       (11, 1, 2, '市场部', 'ON', now()),
       (12, 1, 3, '商务部', 'ON', now()),
       (13, 1, 4, '财务部', 'ON', now()),

       (2, null, 2, '华南分部', 'ON', now()),
       (20, 2, 1, '研发部', 'ON', now()),
       (21, 2, 2, '市场部', 'ON', now()),
       (22, 2, 3, '商务部', 'ON', now()),
       (23, 2, 4, '财务部', 'ON', now()),

       (3, null, 3, '西北分部', 'ON', now()),
       (30, 3, 1, '研发部', 'ON', now()),
       (31, 3, 2, '市场部', 'ON', now()),
       (32, 3, 3, '商务部', 'ON', now()),
       (33, 3, 4, '财务部', 'ON', now())
;
SELECT setval('departments_id_seq', (SELECT MAX(id) FROM departments));

-- 职位
TRUNCATE TABLE kratos_admin.public.positions RESTART IDENTITY;
INSERT INTO kratos_admin.public.positions (id, name, code, parent_id, status, sort_id, create_time)
VALUES (1, '开发工程师', 'dev_engineer', null, 'ON', 1, now()),
       (2, '测试工程师', 'test_engineer', null, 'ON', 2, now()),
       (3, '产品经理', 'product_manager', null, 'ON', 3, now()),
       (4, '项目经理', 'project_manager', null, 'ON', 4, now())
;
SELECT setval('positions_id_seq', (SELECT MAX(id) FROM positions));

-- 调度任务
TRUNCATE TABLE kratos_admin.public.tasks RESTART IDENTITY;
INSERT INTO kratos_admin.public.tasks(type, type_name, task_payload, cron_spec, enable, create_time)
VALUES ('Periodic', 'backup', '{ "name": "test"}', '*/1 * * * ?', true, now())
;
SELECT setval('tasks_id_seq', (SELECT MAX(id) FROM tasks));
