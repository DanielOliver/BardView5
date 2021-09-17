INSERT INTO "user" (email, name) VALUES ('test@test.com', 'Daniel');
INSERT INTO role_action (name) VALUES ('create');
INSERT INTO role_subject (name) VALUES ('user');
INSERT INTO role_type (name, multiple_allowed) VALUES ('User Role', false);
INSERT INTO role (name, role_type_id) VALUES ('test@test.com', (SELECT r.id FROM role r limit 1))
INSERT INTO role_assignment(role_id, user_id) VALUES  ((SELECT r.id FROM role r limit 1),  (SELECT u.id FROM "user" u limit 1));
INSERT INTO role_permission(role_id, actions, subject, conditions, fields)  VALUES  ((SELECT r.id FROM role r limit 1),  ARRAY['create'], 'user', null, null);

SELECT rp.*
FROM "user" u
INNER JOIN role_assignment ra on u.id = ra.user_id
INNER JOIN role r on ra.role_id = r.id
INNER JOIN role_permission rp on r.id = rp.role_id
