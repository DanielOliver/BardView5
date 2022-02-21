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

-- name: Dnd5eSettingInsert :execrows
insert into "dnd5e_setting" (dnd5e_setting_id, common_access, created_by, is_active, system_tags,
                           user_tags, "name", module, description)
VALUES (@dnd5e_setting_id, @common_access, @created_by, @is_active,
        @system_tags, @user_tags, @name, @module, @description);

-- name: Dnd5eSettingUpdate :execrows
Update "dnd5e_setting" as s
SET common_access = @common_access
  ,is_active = @is_active
  ,system_tags = @system_tags
  ,user_tags = @user_tags
  ,"name" = @name
  ,module = @module
  ,description = @description
WHERE s.dnd5e_setting_id = @dnd5e_setting_id;

-- name: Dnd5eSettingFindById :many
SELECT *
FROM "dnd5e_setting" w
WHERE w.dnd5e_setting_id = @dnd5e_setting_id;

-- name: Dnd5eSettingUpsertAssignment :execrows
insert into "dnd5e_setting_assignment" (created_by, user_id, dnd5e_setting_id, role_action)
SELECT @created_by,
       @user_id,
       @dnd5e_setting_id,
       @role_action WHERE NOT EXISTS (
    SELECT 1 FROM dnd5e_setting_assignment
    WHERE user_id = @user_id AND dnd5e_setting_id = @dnd5e_setting_id AND role_action = @role_action
);

-- name: Dnd5eSettingFindByAssignment :many
SELECT DISTINCT w.*
FROM "dnd5e_setting" w
         INNER JOIN "dnd5e_setting_assignment" wa ON
    w.dnd5e_setting_id = wa.dnd5e_setting_id
WHERE wa.user_id = @user_id
ORDER BY w.dnd5e_setting_id desc;

-- name: Dnd5eSettingFindByParams :many
SELECT DISTINCT w.*
FROM "dnd5e_setting" w
         LEFT OUTER JOIN "dnd5e_setting_assignment" wa ON
        w.dnd5e_setting_id = wa.dnd5e_setting_id
    AND wa.user_id = @user_id
WHERE (wa.user_id IS NOT NULL
    OR w.common_access IN ('anyuser', 'public')
  )
    AND w.name LIKE @name
ORDER BY w.dnd5e_setting_id desc;

-- name: Dnd5eSettingFindAssignment :many
SELECT wa.*
FROM "dnd5e_setting_assignment" wa
WHERE wa.user_id = @user_id
  AND wa.dnd5e_setting_id = @dnd5e_setting_id;

-- name: Dnd5eMonsterFindById :many
SELECT *
FROM "dnd5e_monster" m
WHERE m.dnd5e_monster_id = @dnd5e_monster_id;

-- name: Dnd5eMonstersFindBySetting :many
SELECT *
FROM "dnd5e_monster" m
WHERE m.dnd5e_setting_id = @dnd5e_setting_id
ORDER BY m.dnd5e_setting_id, m.dnd5e_monster_id
OFFSET @row_offset LIMIT @row_limit;

-- name: Dnd5eMonsterInsert :execrows
INSERT INTO dnd5e_monster (dnd5e_monster_id, created_by, dnd5e_setting_id, name, sources,
                           user_tags, languages, environments, is_legendary, is_unique, monster_type, alignment,
                           size_category, milli_challenge_rating, armor_class, hit_points, description)
VALUES (
           @dnd5e_monster_id, @created_by, @dnd5e_setting_id, @name, @sources,
           @user_tags, @languages, @environments, @is_legendary, @is_unique, @monster_type, @alignment,
           @size_category, @milli_challenge_rating, @armor_class, @hit_points, @description
       );

-- name: Dnd5eSizeCategoryFindAll :many
SELECT *
FROM "dnd5e_size_category" s;

-- name: Dnd5eLanguageFindAll :many
SELECT *
FROM "dnd5e_language" l;
