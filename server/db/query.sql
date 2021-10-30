-- name: GetAclBySubject :many
SELECT
    ra.user_id
     ,rp.subject
     ,rp.conditions
     ,rp.action
     ,rp.subject_id
     ,r.name "role_name"
     ,r.role_id
FROM "role_assignment" ra
         INNER JOIN role r on ra.role_id = r.role_id AND r.is_active = true
         INNER JOIN role_permission rp on r.role_id = rp.role_id AND rp.is_active = true
WHERE rp.subject = @subject
    AND (  rp.subject_id IS NULL
        OR rp.subject_id = @subject_id)
    AND ra.user_id = @user_id
    AND ra.is_active = true;

-- name: UserFindById :one
SELECT * FROM "user" u WHERE u.user_id = @user_id;

-- name: UsersFindByUid :many
SELECT DISTINCT u.*
FROM role_assignment ra
         INNER JOIN role r on ra.role_id = r.role_id
         INNER JOIN role_permission rp on r.role_id = rp.role_id
         INNER JOIN "user" u on evaluate_access_user(rp.conditions, @session_id::bigint, u.user_id)
WHERE ra.user_id = @session_id::bigint
  AND rp.subject = 'user'
  AND rp.is_active = true
  AND ra.is_active = true
  AND r.is_active = true
  AND u.uuid = @uuid
ORDER BY u.user_id;

-- name: UserInsert :execrows
INSERT INTO "user" as u (user_id, uuid, "name", email, tags, created_by)
VALUES (@user_id, @uuid, @name, @email, @tags, @created_by)
ON CONFLICT (email) DO NOTHING;


