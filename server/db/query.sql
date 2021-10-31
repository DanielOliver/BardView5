-- name: GetAclBySubject :many
SELECT ra.user_id
     , rp.subject
     , rp.conditions
     , rp.action
     , rp.subject_id
     , r.name "role_name"
     , r.role_id
FROM "role_assignment" ra
         INNER JOIN "role" r on ra.role_id = r.role_id AND r.is_active = true
         INNER JOIN role_permission rp on r.role_id = rp.role_id AND rp.is_active = true
WHERE rp.subject = @subject
  AND (rp.subject_id IS NULL
    OR rp.subject_id = @subject_id)
  AND ra.user_id = @user_id
  AND ra.is_active = true;

-- name: UserFindByEmail :many
SELECT *
FROM "user" u
WHERE u.email = @email;

-- name: UserFindById :many
SELECT *
FROM "user" u
WHERE u.user_id = @user_id;

-- name: UserInsert :execrows
INSERT INTO "user" as u (user_id, uuid, "name", email, user_tags, system_tags, created_by, common_access)
VALUES (@user_id, @uuid, @name, @email, @user_tags, @system_tags, @created_by, @common_access)
ON CONFLICT (email) DO NOTHING;

-- name: UserUpdate :many
UPDATE "user" as u
SET name          = @name
  , user_tags     = @user_tags
  , system_tags   = @system_tags
  , common_access = @common_access
  , version       = version + 1
WHERE u.user_id = @user_id
  AND u.version = @version
RETURNING u.version, u.user_id;

