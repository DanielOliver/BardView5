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
  AND u.uid = @uid
ORDER BY u.user_id;

-- name: UserInsert :execrows
INSERT INTO "user" as u (user_id, uid, "name", email, tags, created_by)
VALUES (@user_id, @uid, @name, @email, @tags, @created_by)
ON CONFLICT (email) DO NOTHING;


