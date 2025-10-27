-- 租户
TRUNCATE TABLE kratos_admin.public.sys_tenants RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_tenants(id, name, code, type, audit_status, status, create_time)
VALUES (1, '超级租户', 'super', 'PAID', 'APPROVED', 'ON', now()),
       (2, '测试租户', 'test', 'PAID', 'APPROVED',  'ON', now()),
       (3, '测试租户2', 'test2', 'PAID', 'APPROVED',  'ON', now())
;
SELECT setval('sys_tenants_id_seq', (SELECT MAX(id) FROM sys_tenants));

-- 组织
TRUNCATE TABLE kratos_admin.public.sys_organizations RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_organizations(id, parent_id, sort_id, manager_id, name, organization_type, is_legal_entity, business_scope, status, create_time)
VALUES (1, null, 1, 1,'虾米集团', 'GROUP', true, '综合型集团企业，涵盖多领域', 'ON', now()),
       (2, 1, 2, 1,'北京分公司', 'SUBSIDIARY', false, '负责华北区域业务运营', 'ON', now()),
       (3, 1, 3, 1,'上海子公司', 'FILIALE', true, '负责华东区域研发与生产', 'ON', now()),
       (4, 1, 4, 1,'新能源事业部', 'DIVISION', false, '新能源汽车技术研发与市场拓展', 'ON', now())
;
SELECT setval('sys_organizations_id_seq', (SELECT MAX(id) FROM sys_organizations));

-- 部门
TRUNCATE TABLE kratos_admin.public.sys_departments RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_departments(id, parent_id, sort_id, organization_id, manager_id, name, description, status, create_time)
VALUES (1, null, 1, 2, 1, '技术部', '负责北京分公司系统开发','ON', now()),
       (2, null, 2, 2, 1, '财务部', '负责北京分公司财务核算','ON', now()),
       (3, null, 3, 2, 1, '人力资源部', '负责北京分公司人员招聘','ON', now()),
       (4, null, 4, 3, 1, '研发一部', '上海子公司核心技术研发','ON', now()),
       (5, null, 5, 4, 1, '市场部', '新能源事业部市场推广','ON', now()),
       (6, 1, 1, 2, 1, '前端组', '技术部下属前端开发团队','ON', now())
;
SELECT setval('sys_departments_id_seq', (SELECT MAX(id) FROM sys_departments));

-- 职位
TRUNCATE TABLE kratos_admin.public.sys_positions RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_positions (id, name, code, parent_id, department_id, organization_id, quota, description, status, sort_id, create_time)
VALUES
-- 技术部(dept_id=1) 北京分公司(org_id=2)
(1, '技术总监', 'TECH-DIRECTOR-001', null, 1, 2, 1, '负责公司整体技术战略规划、团队管理及核心技术决策', 'ON', 1, now()),
(2, '技术部经理', 'TECH-MANAGER-001', 1, 1, 2, 1, '负责技术部日常管理、项目排期及团队协作', 'ON', 2, now()),
(3, '前端主管', 'TECH-FE-LEADER-001', 2, 1, 2, 1, '负责前端团队开发管理、技术方案评审及需求落地', 'ON', 3, now()),
(4, '后端主管', 'TECH-BE-LEADER-001', 2, 1, 2, 1, '负责后端服务架构设计、数据库优化及接口开发管理', 'ON', 4, now()),
(5, '前端开发专员', 'TECH-FE-DEV-001', 3, 1, 2, 5, '负责Web/移动端前端页面开发、交互实现及兼容性优化', 'ON', 5, now()),
(6, '后端开发专员', 'TECH-BE-DEV-001', 4, 1, 2, 5, '负责后端接口开发、业务逻辑实现及系统稳定性维护', 'ON', 6, now()),
(7, '测试工程师', 'TECH-TEST-001', 2, 1, 2, 3, '负责项目功能测试、性能测试及自动化测试脚本开发', 'ON', 7, now()),

-- 人力资源部(dept_id=3) 北京分公司(org_id=2)
(8, '人力总监', 'HR-DIRECTOR-001', null, 3, 2, 1, '负责人力资源战略规划、组织架构设计及人才梯队建设', 'ON', 1, now()),
(9, '招聘主管', 'HR-RECRUIT-LEADER-001', 8, 3, 2, 2, '负责公司各部门招聘需求对接、简历筛选及面试安排', 'ON', 2, now()),
(10, '薪酬绩效专员', 'HR-C&P-001', 8, 3, 2, 2, '负责员工薪酬核算、绩效考核制度落地及社保公积金管理', 'ON', 3, now()),
(11, 'HRBP', 'HR-BP-001', 8, 3, 2, 3, '对接业务部门，提供人力资源支持（入离职、员工关系等）', 'ON', 4, now()),

-- 财务部(dept_id=2) 北京分公司(org_id=2)
(12, '财务总监', 'FIN-DIRECTOR-001', null, 2, 2, 1, '负责公司财务战略、预算管理及财务风险控制', 'ON', 1, now()),
(13, '会计主管', 'FIN-ACCOUNT-LEADER-001', 12, 2, 2, 1, '负责账务处理、财务报表编制及税务申报管理', 'ON', 2, now()),
(14, '出纳专员', 'FIN-CASHIER-001', 13, 2, 2, 2, '负责日常资金收付、银行对账及票据管理', 'ON', 3, now()),
(15, '成本会计', 'FIN-COST-001', 13, 2, 2, 1, '负责成本核算、成本分析及成本控制方案制定', 'ON', 4, now()),

-- 市场部(dept_id=5) 新能源事业部(org_id=4)
(16, '市场总监', 'MKT-DIRECTOR-001', null, 5, 4, 1, '负责市场战略规划、品牌建设及营销活动策划', 'ON', 1, now()),
(17, '新媒体运营主管', 'MKT-NEWS-LEADER-001', 16, 5, 4, 1, '负责新媒体平台（微信、抖音等）内容运营及用户增长', 'ON', 2, now()),
(18, '活动策划专员', 'MKT-EVENT-001', 16, 5, 4, 3, '负责线下活动策划、执行及效果复盘', 'ON', 3, now()),
(19, '市场调研专员', 'MKT-RESEARCH-001', 16, 5, 4, 2, '负责行业动态调研、竞品分析及市场趋势报告撰写', 'ON', 4, now()),

-- 禁用职位（示例：已废弃的“行政助理”）
(20, '行政助理', 'ADMIN-ASSIST-001', 8, 3, 2, 0, '负责办公用品采购、会议安排等行政工作（已合并至HRBP）', 'OFF', 5, now())
;
SELECT setval('sys_positions_id_seq', (SELECT MAX(id) FROM sys_positions));

-- 调度任务
TRUNCATE TABLE kratos_admin.public.sys_tasks RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_tasks(type, type_name, task_payload, cron_spec, enable, create_time)
VALUES ('PERIODIC', 'backup', '{ "name": "test"}', '*/1 * * * ?', true, now())
;
SELECT setval('sys_tasks_id_seq', (SELECT MAX(id) FROM sys_tasks));

-- 后台登录限制
TRUNCATE TABLE kratos_admin.public.sys_admin_login_restrictions RESTART IDENTITY;
INSERT INTO kratos_admin.public.sys_admin_login_restrictions(id, target_id, type, method, value, reason, create_time)
VALUES (1, 1, 'BLACKLIST', 'IP', '127.0.0.1', '无理由', now()),
       (2, 1, 'WHITELIST', 'MAC', '00:1B:44:11:3A:B7 ', '无理由', now())
;
SELECT setval('sys_admin_login_restrictions_id_seq', (SELECT MAX(id) FROM sys_admin_login_restrictions));
