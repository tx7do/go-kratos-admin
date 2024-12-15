-- 默认的超级管理员，默认账号：admin，密码：admin
TRUNCATE TABLE kratos_admin.public.users;
INSERT INTO kratos_admin.public.users (id, username, nick_name, email, password, authority)
VALUES (1, 'admin', 'admin', 'admin@gmail.com', '$2a$10$yajZDX20Y40FkG0Bu4N19eXNqRizez/S9fK63.JxGkfLq.RoNKR/a', 'SYS_ADMIN');
