-- name: UserFindByEmail :many
SELECT *
FROM "user" u
WHERE u.email = @email;

-- name: UserFindById :many
SELECT *
FROM "user" u
WHERE u.user_id = @user_id;

-- name: UserFindByUuid :many
SELECT *
FROM "user" u
WHERE u.uuid = @uuid;

-- name: UserInsert :execrows
INSERT INTO "user" as u (user_id, uuid, "name", email, user_tags, system_tags, created_by, common_access, is_active)
VALUES (@user_id, @uuid, @name, @email, @user_tags, @system_tags, @created_by, @common_access, @is_active)
ON CONFLICT (email) DO NOTHING;

-- name: UserUpdate :many
UPDATE "user" as u
SET name          = @name
  , user_tags     = @user_tags
  , system_tags   = @system_tags
  , common_access = @common_access
  , version       = version + 1
  , is_active     = @is_active
WHERE u.user_id = @user_id
  AND u.version = @version RETURNING *;

-- name: Dnd5eWorldInsert :execrows
insert into "dnd5e_world" (dnd5e_world_id, common_access, created_by, is_active, system_tags,
                           user_tags, "name", module, description)
VALUES (@dnd5e_world_id, @common_access, @created_by, @is_active,
        @system_tags, @user_tags, @name, @module, @description);

-- name: Dnd5eWorldFindById :many
SELECT *
FROM "dnd5e_world" w
WHERE w.dnd5e_world_id = @dnd5e_world_id;

-- name: Dnd5eWorldUpsertAssignment :execrows
insert into "dnd5e_world_assignment" (created_by, user_id, dnd5e_world_id, role_action)
SELECT @created_by,
       @user_id,
       @dnd5e_world_id,
       @role_action WHERE NOT EXISTS (
    SELECT 1 FROM dnd5e_world_assignment
    WHERE user_id = @user_id AND dnd5e_world_id = @dnd5e_world_id AND role_action = @role_action
);

-- name: Dnd5eWorldFindByAssignment :many
SELECT DISTINCT w.*
FROM "dnd5e_world" w
         INNER JOIN "dnd5e_world_assignment" wa ON
    w.dnd5e_world_id = wa.dnd5e_world_id
WHERE wa.user_id = @user_id
ORDER BY w.dnd5e_world_id desc;

-- name: Dnd5eWorldFindAssignment :many
SELECT wa.*
FROM "dnd5e_world_assignment" wa
WHERE wa.user_id = @user_id
  AND wa.dnd5e_world_id = @dnd5e_world_id;

-- name: Dnd5eMonsterFindById :many
SELECT *
FROM "dnd5e_monster" m
WHERE m.dnd5e_monster_id = @dnd5e_monster_id;

-- name: Dnd5eMonstersFindByWorld :many
SELECT *
FROM "dnd5e_monster" m
WHERE m.dnd5e_world_id = @dnd5e_world_id
ORDER BY m.dnd5e_world_id, m.dnd5e_monster_id
OFFSET @row_offset LIMIT @row_limit;

-- name: Dnd5eSizeCategoryFindAll :many
SELECT *
FROM "dnd5e_size_category" s;

-- name: Dnd5eLanguageFindAll :many
SELECT *
FROM "dnd5e_language" l;
