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
insert into "dnd5e_world" (dnd5e_world_id, derived_from_world, common_access, created_by, is_active, system_tags,
                           user_tags, "name")
VALUES (@dnd5e_world_id, @derived_from_world, @common_access, @created_by, @is_active, @system_tags, @user_tags, @name);

-- name: Dnd5eWorldFindById :many
SELECT *
FROM "dnd5e_world" w
WHERE w.dnd5e_world_id = @dnd5e_world_id;

-- name: Dnd5eMonsterFindById :many
SELECT *
FROM "dnd5e_monster" m
WHERE m.dnd5e_monster_id = @dnd5e_monster_id;

-- name: Dnd5eInhabitantsFindByWorldAndMonster :many
SELECT wm.dnd5e_inhabitant_id,
       m.dnd5e_monster_id,
       wm.dnd5e_world_id,
       m.name,
       m.tags,
       m.monster_type,
       m.alignment,
       m.size_category,
       m.milli_challenge_rating,
       m.languages,
       m.description,
       wm.user_tags,
       wm.system_tags
FROM "dnd5e_monster" m
         INNER JOIN "dnd5e_inhabitant" wm ON wm.dnd5e_monster_id = m.dnd5e_monster_id
WHERE m.dnd5e_monster_id = @dnd5e_monster_id
  AND wm.dnd5e_monster_id = @dnd5e_monster_id
  AND wm.dnd5e_world_id = @dnd5e_world_id;

-- name: Dnd5eInhabitantsFindByWorld :many
SELECT wm.dnd5e_inhabitant_id,
       m.dnd5e_monster_id,
       wm.dnd5e_world_id,
       m.name,
       m.tags,
       m.monster_type,
       m.alignment,
       m.size_category,
       m.milli_challenge_rating,
       m.languages,
       m.description,
       wm.user_tags,
       wm.system_tags
FROM "dnd5e_monster" m
         INNER JOIN "dnd5e_inhabitant" wm ON wm.dnd5e_monster_id = m.dnd5e_monster_id
WHERE wm.dnd5e_world_id = @dnd5e_world_id
ORDER BY wm.dnd5e_world_id, wm.dnd5e_monster_id
OFFSET @row_offset LIMIT @row_limit;

-- name: Dnd5eSizeCategoryFindAll :many
SELECT *
FROM "dnd5e_size_category" s;

-- name: Dnd5eLanguageFindAll :many
SELECT *
FROM "dnd5e_language" l;
