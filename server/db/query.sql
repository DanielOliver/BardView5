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

-- name: WorldInsert :execrows
insert into "world" (world_id, derived_from_world, common_access, created_by, is_active, system_tags, user_tags,
                     "name")
VALUES (@world_id, @derived_from_world, @common_access, @created_by, @is_active, @system_tags, @user_tags, @name);

-- name: WorldFindById :many
SELECT *
FROM "world" w
WHERE w.world_id = @world_id;

-- name: MonsterFindById :many
SELECT *
FROM "monster" m
WHERE m.monster_id = @monster_id;

-- name: InhabitantsFindByWorldAndMonster :many
SELECT wm.inhabitant_id,
       m.monster_id,
       wm.world_id,
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
FROM "monster" m
         INNER JOIN "inhabitant" wm ON wm.monster_id = m.monster_id
WHERE m.monster_id = @monster_id
  AND wm.monster_id = @monster_id
  AND wm.world_id = @world_id;

-- name: InhabitantsFindByWorld :many
SELECT wm.inhabitant_id,
       m.monster_id,
       wm.world_id,
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
FROM "monster" m
         INNER JOIN "inhabitant" wm ON wm.monster_id = m.monster_id
WHERE wm.world_id = @world_id
ORDER BY wm.world_id, wm.monster_id
OFFSET @row_offset LIMIT @row_limit;

-- name: SizeCategoryFindAll :many
SELECT *
FROM "size_category" s;

-- name: LanguageFindAll :many
SELECT *
FROM "language" l;
