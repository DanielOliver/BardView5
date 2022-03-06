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
  , is_active     = @is_active
  , system_tags   = @system_tags
  , user_tags     = @user_tags
  , "name"        = @name
  , module        = @module
  , description   = @description
WHERE s.dnd5e_setting_id = @dnd5e_setting_id;

-- name: Dnd5eSettingFindById :many
SELECT *
FROM "dnd5e_setting" w
WHERE w.dnd5e_setting_id = @dnd5e_setting_id;

-- name: RoleAssignmentUpsertInitial :execrows
insert into "role_assignment" (created_by, user_id, role_id, scope_id)
SELECT @created_by,
       @user_id,
       (SELECT MIN(role_id)
        FROM "role" r
        WHERE r.scope_id IS NULL
          AND r.role_subject = @role_subject
          AND r.assign_on_create = true),
       @scope_id WHERE NOT EXISTS (
    SELECT 1 FROM role_assignment ra
    INNER JOIN "role" r ON r.role_id = ra.role_id
    WHERE ra.user_id = @user_id
  AND ra.scope_id = @scope_id
  AND r.scope_id IS NULL
  AND r.assign_on_create = true
  AND r.role_subject = @role_subject
    );

-- name: RoleAssignmentUpsertDefaultAdd :execrows
insert into "role_assignment" (created_by, user_id, role_id, scope_id)
SELECT @created_by,
       @user_id,
       (SELECT MIN(role_id)
        FROM "role" r
        WHERE r.scope_id IS NULL
          AND r.role_subject = @role_subject
          AND r.assign_on_add = true),
       @scope_id WHERE NOT EXISTS (
    SELECT 1 FROM role_assignment ra
    INNER JOIN "role" r ON r.role_id = ra.role_id
    WHERE ra.user_id = @user_id
  AND ra.scope_id = @scope_id
  AND r.scope_id IS NULL
  AND r.assign_on_add = true
  AND r.role_subject = @role_subject
    );

-- name: RoleAssignmentFindByScopeId :many
SELECT wa.scope_id
    ,r.role_id
    ,r.name
    ,r.role_type
    ,r.role_subject
    ,r.capabilities
FROM "role_assignment" wa
         INNER JOIN "role" r ON
        r.role_id = wa.role_id
WHERE wa.user_id = @user_id
  AND wa.scope_id = @scope_id
  AND r.role_subject = @role_subject;

-- name: Dnd5eSettingFindByAssignment :many
SELECT DISTINCT w.*
FROM "dnd5e_setting" w
         INNER JOIN "role_assignment" wa ON
    w.dnd5e_setting_id = wa.scope_id
         INNER JOIN "role" r ON
            r.role_id = wa.role_id
        AND r.role_subject = 'dnd5esetting'
WHERE wa.user_id = @user_id
ORDER BY w.dnd5e_setting_id desc;

-- name: Dnd5eSettingFindByParams :many
SELECT DISTINCT w.*
FROM "dnd5e_setting" w
         LEFT OUTER JOIN "role_assignment" wa ON
            w.dnd5e_setting_id = wa.scope_id
        AND wa.user_id = @user_id
        AND EXISTS(
                    SELECT 1
                    FROM "role" r
                    WHERE r.role_id = wa.role_id
                      AND r.role_subject = 'dnd5esetting')
WHERE (wa.user_id IS NOT NULL
    OR w.common_access IN ('anyuser', 'public')
    )
  AND w.name LIKE @name
ORDER BY w.dnd5e_setting_id desc;

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
                           size_category, milli_challenge_rating, armor_class, hit_points, description,
                           str_score, int_score, wis_score, dex_score, con_score, cha_score)
VALUES (@dnd5e_monster_id, @created_by, @dnd5e_setting_id, @name, @sources,
        @user_tags, @languages, @environments, @is_legendary, @is_unique, @monster_type, @alignment,
        @size_category, @milli_challenge_rating, @armor_class, @hit_points, @description,
        @str_score, @int_score, @wis_score, @dex_score, @con_score, @cha_score);

-- name: Dnd5eSizeCategoryFindAll :many
SELECT *
FROM "dnd5e_size_category" s;

-- name: Dnd5eLanguageFindAll :many
SELECT *
FROM "dnd5e_language" l;
