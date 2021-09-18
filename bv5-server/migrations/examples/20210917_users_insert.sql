INSERT INTO "user" (email, name, uid, tags) VALUES ('test@test2.com', 'Daniel', '', Array['test']);
INSERT INTO role_action (name) VALUES ('create');
INSERT INTO role_subject (name) VALUES ('user');
INSERT INTO role_type (name, multiple_assignments_allowed, system_managed) VALUES ('User Role', false, true);
INSERT INTO role (name, role_type_id, uid, tags) VALUES ('test@test.com', (SELECT r.id FROM role r limit 1), '', Array['test'])
INSERT INTO role_assignment(role_id, user_id, uid, tags) VALUES  ((SELECT r.id FROM role r limit 1),  (SELECT u.id FROM "user" u limit 1),'', Array['test']);
INSERT INTO role_permission(role_id, action, subject, conditions, fields, uid)  VALUES  ((SELECT r.id FROM role r limit 1),  'create', 'user', '{}', null, '');


SELECT rp.*
FROM "user" u
         INNER JOIN role_assignment ra on u.id = ra.user_id
         INNER JOIN role r on ra.role_id = r.id
         INNER JOIN role_permission rp on r.id = rp.role_id

