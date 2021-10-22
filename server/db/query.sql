-- name: UsersFindByUid :many
SELECT DISTINCT u.*
FROM role_assignment ra
         INNER JOIN role r on ra.role_id = r.role_id
         INNER JOIN role_permission rp on r.role_id = rp.role_id
         INNER JOIN "user" u on evaluate_access_user(rp.conditions, @session_id::bigint, u.id)
WHERE ra.user_id = @session_id::bigint
  AND rp.subject = 'user'
  AND rp.is_active = true
  AND ra.is_active = true
  AND r.is_active = true
  AND u.uuid = @uuid
ORDER BY u.user_id;

-- name: GetAcl :many
SELECT DISTINCT rp.action
              , rp.subject
              , rp.conditions
              , u.user_id
FROM "user" u
         INNER JOIN role_assignment ra ON u.user_id = ra.user_id AND ra.is_active = true
         INNER JOIN role r on ra.role_id = r.role_id AND r.is_active = true
         INNER JOIN role_permission rp on r.role_id = rp.role_id AND rp.is_active = true
WHERE rp.subject = @subject
  AND (rp.subject_id = @subject_id OR rp.subject_id IS NULL)
  AND u.user_id = @user_id;

-- name: UserInsert :execrows
INSERT INTO "user" as u (user_id, uuid, "name", email, tags, created_by)
VALUES (@user_id, @uuid, @name, @email, @tags, @created_by)
ON CONFLICT (email) DO NOTHING;


