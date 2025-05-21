-- 租户
TRUNCATE TABLE kratos_admin.public.tenants;
INSERT INTO kratos_admin.public.tenants(id, name, code, status, create_time)
VALUES (1, '超级租户', 'super', 'ON', now()),
       (2, '测试租户', 'test', 'ON', now()),
       (3, '测试租户2', 'test2', 'ON', now())
;

-- 组织
TRUNCATE TABLE kratos_admin.public.organizations;
INSERT INTO kratos_admin.public.organizations(id, parent_id, sort_id, name, status, create_time)
VALUES (1, null, 1, '华东分部', 'ON', now()),
       (2, null, 1, '华南分部', 'ON', now()),
       (3, null, 2, '西北分部', 'ON', now())
;

-- 部门
TRUNCATE TABLE kratos_admin.public.departments;
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

-- 调度任务
TRUNCATE TABLE kratos_admin.public.tasks;
INSERT INTO kratos_admin.public.tasks(id, type, type_name, task_payload, task_id, cron_spec, enable, create_time)
VALUES (1, 'Periodic', 'backup', '{ "name": "test"}', 'backup', '*/1 * * * ?', true, now())
;
