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
INSERT INTO kratos_admin.public.organizations(id, parent_id, sort_id, manager_id, name, organization_type, is_legal_entity, business_scope, status, create_time)
VALUES (1, null, 1, 1,'虾米集团', 'GROUP', true, '综合型集团企业，涵盖多领域', 'ON', now()),
       (2, 1, 2, 1,'北京分公司', 'SUBSIDIARY', false, '负责华北区域业务运营', 'ON', now()),
       (3, 1, 3, 1,'上海子公司', 'FILIALE', true, '负责华东区域研发与生产', 'ON', now()),
       (4, 1, 4, 1,'新能源事业部', 'DIVISION', false, '新能源汽车技术研发与市场拓展', 'ON', now())
;
SELECT setval('organizations_id_seq', (SELECT MAX(id) FROM organizations));

-- 部门
TRUNCATE TABLE kratos_admin.public.departments RESTART IDENTITY;
INSERT INTO kratos_admin.public.departments(id, parent_id, sort_id, organization_id, manager_id, name, description, status, create_time)
VALUES (1, null, 1, 2, 1, '技术部', '负责北京分公司系统开发','ON', now()),
       (2, null, 2, 2, 1, '财务部', '负责北京分公司财务核算','ON', now()),
       (3, null, 3, 3, 1, '研发一部', '上海子公司核心技术研发','ON', now()),
       (4, null, 4, 4, 1, '市场部', '新能源事业部市场推广','ON', now()),
       (5, 1, 5, 2, 1, '前端组', '技术部下属前端开发团队','ON', now())
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
TRUNCATE TABLE kratos_admin.public.sys_tasks RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_tasks(type, type_name, task_payload, cron_spec, enable, create_time)
VALUES ('PERIODIC', 'backup', '{ "name": "test"}', '*/1 * * * ?', true, now())
;
SELECT setval('sys_tasks_id_seq', (SELECT MAX(id) FROM sys_tasks));

-- 后台登录限制
TRUNCATE TABLE kratos_admin.public.admin_login_restrictions RESTART IDENTITY;
INSERT INTO kratos_admin.public.admin_login_restrictions(id, target_id, type, method, value, reason, create_time)
VALUES (1, 1, 'BLACKLIST', 'IP', '127.0.0.1', '无理由', now()),
       (2, 1, 'WHITELIST', 'MAC', '00:1B:44:11:3A:B7 ', '无理由', now())
;
SELECT setval('admin_login_restrictions_id_seq', (SELECT MAX(id) FROM admin_login_restrictions));
