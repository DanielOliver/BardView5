// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const dnd5eLanguageFindAll = `-- name: Dnd5eLanguageFindAll :many
SELECT dnd5e_language_id, created_by, created_at, version, name
FROM "dnd5e_language" l
`

func (q *Queries) Dnd5eLanguageFindAll(ctx context.Context) ([]Dnd5eLanguage, error) {
	rows, err := q.query(ctx, q.dnd5eLanguageFindAllStmt, dnd5eLanguageFindAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dnd5eLanguage{}
	for rows.Next() {
		var i Dnd5eLanguage
		if err := rows.Scan(
			&i.Dnd5eLanguageID,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const dnd5eMonsterFindById = `-- name: Dnd5eMonsterFindById :many
SELECT dnd5e_monster_id, created_by, created_at, version, dnd5e_setting_id, name, user_tags, system_tags, monster_type, alignment, size_category, milli_challenge_rating, languages, description
FROM "dnd5e_monster" m
WHERE m.dnd5e_monster_id = $1
`

func (q *Queries) Dnd5eMonsterFindById(ctx context.Context, dnd5eMonsterID int64) ([]Dnd5eMonster, error) {
	rows, err := q.query(ctx, q.dnd5eMonsterFindByIdStmt, dnd5eMonsterFindById, dnd5eMonsterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dnd5eMonster{}
	for rows.Next() {
		var i Dnd5eMonster
		if err := rows.Scan(
			&i.Dnd5eMonsterID,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.Dnd5eSettingID,
			&i.Name,
			pq.Array(&i.UserTags),
			pq.Array(&i.SystemTags),
			&i.MonsterType,
			&i.Alignment,
			&i.SizeCategory,
			&i.MilliChallengeRating,
			pq.Array(&i.Languages),
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const dnd5eMonstersFindBySetting = `-- name: Dnd5eMonstersFindBySetting :many
SELECT dnd5e_monster_id, created_by, created_at, version, dnd5e_setting_id, name, user_tags, system_tags, monster_type, alignment, size_category, milli_challenge_rating, languages, description
FROM "dnd5e_monster" m
WHERE m.dnd5e_setting_id = $1
ORDER BY m.dnd5e_setting_id, m.dnd5e_monster_id
OFFSET $2 LIMIT $3
`

type Dnd5eMonstersFindBySettingParams struct {
	Dnd5eSettingID sql.NullInt64 `db:"dnd5e_setting_id"`
	RowOffset      int32         `db:"row_offset"`
	RowLimit       int32         `db:"row_limit"`
}

func (q *Queries) Dnd5eMonstersFindBySetting(ctx context.Context, arg Dnd5eMonstersFindBySettingParams) ([]Dnd5eMonster, error) {
	rows, err := q.query(ctx, q.dnd5eMonstersFindBySettingStmt, dnd5eMonstersFindBySetting, arg.Dnd5eSettingID, arg.RowOffset, arg.RowLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dnd5eMonster{}
	for rows.Next() {
		var i Dnd5eMonster
		if err := rows.Scan(
			&i.Dnd5eMonsterID,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.Dnd5eSettingID,
			&i.Name,
			pq.Array(&i.UserTags),
			pq.Array(&i.SystemTags),
			&i.MonsterType,
			&i.Alignment,
			&i.SizeCategory,
			&i.MilliChallengeRating,
			pq.Array(&i.Languages),
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const dnd5eSettingFindAssignment = `-- name: Dnd5eSettingFindAssignment :many
SELECT wa.created_by, wa.created_at, wa.version, wa.user_id, wa.dnd5e_setting_id, wa.role_action
FROM "dnd5e_setting_assignment" wa
WHERE wa.user_id = $1
  AND wa.dnd5e_setting_id = $2
`

type Dnd5eSettingFindAssignmentParams struct {
	UserID         int64 `db:"user_id"`
	Dnd5eSettingID int64 `db:"dnd5e_setting_id"`
}

func (q *Queries) Dnd5eSettingFindAssignment(ctx context.Context, arg Dnd5eSettingFindAssignmentParams) ([]Dnd5eSettingAssignment, error) {
	rows, err := q.query(ctx, q.dnd5eSettingFindAssignmentStmt, dnd5eSettingFindAssignment, arg.UserID, arg.Dnd5eSettingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dnd5eSettingAssignment{}
	for rows.Next() {
		var i Dnd5eSettingAssignment
		if err := rows.Scan(
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.UserID,
			&i.Dnd5eSettingID,
			&i.RoleAction,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const dnd5eSettingFindByAssignment = `-- name: Dnd5eSettingFindByAssignment :many
SELECT DISTINCT w.dnd5e_setting_id, w.created_by, w.created_at, w.version, w.is_active, w.common_access, w.user_tags, w.system_tags, w.name, w.module, w.description
FROM "dnd5e_setting" w
         INNER JOIN "dnd5e_setting_assignment" wa ON
    w.dnd5e_setting_id = wa.dnd5e_setting_id
WHERE wa.user_id = $1
ORDER BY w.dnd5e_setting_id desc
`

func (q *Queries) Dnd5eSettingFindByAssignment(ctx context.Context, userID int64) ([]Dnd5eSetting, error) {
	rows, err := q.query(ctx, q.dnd5eSettingFindByAssignmentStmt, dnd5eSettingFindByAssignment, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dnd5eSetting{}
	for rows.Next() {
		var i Dnd5eSetting
		if err := rows.Scan(
			&i.Dnd5eSettingID,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.IsActive,
			&i.CommonAccess,
			pq.Array(&i.UserTags),
			pq.Array(&i.SystemTags),
			&i.Name,
			&i.Module,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const dnd5eSettingFindById = `-- name: Dnd5eSettingFindById :many
SELECT dnd5e_setting_id, created_by, created_at, version, is_active, common_access, user_tags, system_tags, name, module, description
FROM "dnd5e_setting" w
WHERE w.dnd5e_setting_id = $1
`

func (q *Queries) Dnd5eSettingFindById(ctx context.Context, dnd5eSettingID int64) ([]Dnd5eSetting, error) {
	rows, err := q.query(ctx, q.dnd5eSettingFindByIdStmt, dnd5eSettingFindById, dnd5eSettingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dnd5eSetting{}
	for rows.Next() {
		var i Dnd5eSetting
		if err := rows.Scan(
			&i.Dnd5eSettingID,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.IsActive,
			&i.CommonAccess,
			pq.Array(&i.UserTags),
			pq.Array(&i.SystemTags),
			&i.Name,
			&i.Module,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const dnd5eSettingFindByParams = `-- name: Dnd5eSettingFindByParams :many
SELECT DISTINCT w.dnd5e_setting_id, w.created_by, w.created_at, w.version, w.is_active, w.common_access, w.user_tags, w.system_tags, w.name, w.module, w.description
FROM "dnd5e_setting" w
         LEFT OUTER JOIN "dnd5e_setting_assignment" wa ON
        w.dnd5e_setting_id = wa.dnd5e_setting_id
    AND wa.user_id = $1
WHERE (wa.user_id IS NOT NULL
    OR w.common_access IN ('anyuser', 'public')
  )
    AND w.name LIKE $2
ORDER BY w.dnd5e_setting_id desc
`

type Dnd5eSettingFindByParamsParams struct {
	UserID int64  `db:"user_id"`
	Name   string `db:"name"`
}

func (q *Queries) Dnd5eSettingFindByParams(ctx context.Context, arg Dnd5eSettingFindByParamsParams) ([]Dnd5eSetting, error) {
	rows, err := q.query(ctx, q.dnd5eSettingFindByParamsStmt, dnd5eSettingFindByParams, arg.UserID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dnd5eSetting{}
	for rows.Next() {
		var i Dnd5eSetting
		if err := rows.Scan(
			&i.Dnd5eSettingID,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.IsActive,
			&i.CommonAccess,
			pq.Array(&i.UserTags),
			pq.Array(&i.SystemTags),
			&i.Name,
			&i.Module,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const dnd5eSettingInsert = `-- name: Dnd5eSettingInsert :execrows
insert into "dnd5e_setting" (dnd5e_setting_id, common_access, created_by, is_active, system_tags,
                           user_tags, "name", module, description)
VALUES ($1, $2, $3, $4,
        $5, $6, $7, $8, $9)
`

type Dnd5eSettingInsertParams struct {
	Dnd5eSettingID int64          `db:"dnd5e_setting_id"`
	CommonAccess   string         `db:"common_access"`
	CreatedBy      sql.NullInt64  `db:"created_by"`
	IsActive       bool           `db:"is_active"`
	SystemTags     []string       `db:"system_tags"`
	UserTags       []string       `db:"user_tags"`
	Name           string         `db:"name"`
	Module         sql.NullString `db:"module"`
	Description    string         `db:"description"`
}

func (q *Queries) Dnd5eSettingInsert(ctx context.Context, arg Dnd5eSettingInsertParams) (int64, error) {
	result, err := q.exec(ctx, q.dnd5eSettingInsertStmt, dnd5eSettingInsert,
		arg.Dnd5eSettingID,
		arg.CommonAccess,
		arg.CreatedBy,
		arg.IsActive,
		pq.Array(arg.SystemTags),
		pq.Array(arg.UserTags),
		arg.Name,
		arg.Module,
		arg.Description,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const dnd5eSettingUpdate = `-- name: Dnd5eSettingUpdate :execrows
Update "dnd5e_setting" as s
SET common_access = $1
  ,is_active = $2
  ,system_tags = $3
  ,user_tags = $4
  ,"name" = $5
  ,module = $6
  ,description = $7
WHERE s.dnd5e_setting_id = $8
`

type Dnd5eSettingUpdateParams struct {
	CommonAccess   string         `db:"common_access"`
	IsActive       bool           `db:"is_active"`
	SystemTags     []string       `db:"system_tags"`
	UserTags       []string       `db:"user_tags"`
	Name           string         `db:"name"`
	Module         sql.NullString `db:"module"`
	Description    string         `db:"description"`
	Dnd5eSettingID int64          `db:"dnd5e_setting_id"`
}

func (q *Queries) Dnd5eSettingUpdate(ctx context.Context, arg Dnd5eSettingUpdateParams) (int64, error) {
	result, err := q.exec(ctx, q.dnd5eSettingUpdateStmt, dnd5eSettingUpdate,
		arg.CommonAccess,
		arg.IsActive,
		pq.Array(arg.SystemTags),
		pq.Array(arg.UserTags),
		arg.Name,
		arg.Module,
		arg.Description,
		arg.Dnd5eSettingID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const dnd5eSettingUpsertAssignment = `-- name: Dnd5eSettingUpsertAssignment :execrows
insert into "dnd5e_setting_assignment" (created_by, user_id, dnd5e_setting_id, role_action)
SELECT $1,
       $2,
       $3,
       $4 WHERE NOT EXISTS (
    SELECT 1 FROM dnd5e_setting_assignment
    WHERE user_id = $2 AND dnd5e_setting_id = $3 AND role_action = $4
)
`

type Dnd5eSettingUpsertAssignmentParams struct {
	CreatedBy      sql.NullInt64 `db:"created_by"`
	UserID         int64         `db:"user_id"`
	Dnd5eSettingID int64         `db:"dnd5e_setting_id"`
	RoleAction     string        `db:"role_action"`
}

func (q *Queries) Dnd5eSettingUpsertAssignment(ctx context.Context, arg Dnd5eSettingUpsertAssignmentParams) (int64, error) {
	result, err := q.exec(ctx, q.dnd5eSettingUpsertAssignmentStmt, dnd5eSettingUpsertAssignment,
		arg.CreatedBy,
		arg.UserID,
		arg.Dnd5eSettingID,
		arg.RoleAction,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const dnd5eSizeCategoryFindAll = `-- name: Dnd5eSizeCategoryFindAll :many
SELECT created_by, created_at, version, name, space
FROM "dnd5e_size_category" s
`

func (q *Queries) Dnd5eSizeCategoryFindAll(ctx context.Context) ([]Dnd5eSizeCategory, error) {
	rows, err := q.query(ctx, q.dnd5eSizeCategoryFindAllStmt, dnd5eSizeCategoryFindAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dnd5eSizeCategory{}
	for rows.Next() {
		var i Dnd5eSizeCategory
		if err := rows.Scan(
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.Name,
			&i.Space,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const userFindByEmail = `-- name: UserFindByEmail :many
SELECT user_id, uuid, created_by, created_at, version, effective_date, end_date, is_active, common_access, email, name, user_tags, system_tags
FROM "user" u
WHERE u.email = $1
`

func (q *Queries) UserFindByEmail(ctx context.Context, email string) ([]User, error) {
	rows, err := q.query(ctx, q.userFindByEmailStmt, userFindByEmail, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Uuid,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.EffectiveDate,
			&i.EndDate,
			&i.IsActive,
			&i.CommonAccess,
			&i.Email,
			&i.Name,
			pq.Array(&i.UserTags),
			pq.Array(&i.SystemTags),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const userFindById = `-- name: UserFindById :many
SELECT user_id, uuid, created_by, created_at, version, effective_date, end_date, is_active, common_access, email, name, user_tags, system_tags
FROM "user" u
WHERE u.user_id = $1
`

func (q *Queries) UserFindById(ctx context.Context, userID int64) ([]User, error) {
	rows, err := q.query(ctx, q.userFindByIdStmt, userFindById, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Uuid,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.EffectiveDate,
			&i.EndDate,
			&i.IsActive,
			&i.CommonAccess,
			&i.Email,
			&i.Name,
			pq.Array(&i.UserTags),
			pq.Array(&i.SystemTags),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const userFindByUuid = `-- name: UserFindByUuid :many
SELECT user_id, uuid, created_by, created_at, version, effective_date, end_date, is_active, common_access, email, name, user_tags, system_tags
FROM "user" u
WHERE u.uuid = $1
`

func (q *Queries) UserFindByUuid(ctx context.Context, uuid uuid.UUID) ([]User, error) {
	rows, err := q.query(ctx, q.userFindByUuidStmt, userFindByUuid, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Uuid,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.EffectiveDate,
			&i.EndDate,
			&i.IsActive,
			&i.CommonAccess,
			&i.Email,
			&i.Name,
			pq.Array(&i.UserTags),
			pq.Array(&i.SystemTags),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const userInsert = `-- name: UserInsert :execrows
INSERT INTO "user" as u (user_id, uuid, "name", email, user_tags, system_tags, created_by, common_access, is_active)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT (email) DO NOTHING
`

type UserInsertParams struct {
	UserID       int64         `db:"user_id"`
	Uuid         uuid.UUID     `db:"uuid"`
	Name         string        `db:"name"`
	Email        string        `db:"email"`
	UserTags     []string      `db:"user_tags"`
	SystemTags   []string      `db:"system_tags"`
	CreatedBy    sql.NullInt64 `db:"created_by"`
	CommonAccess string        `db:"common_access"`
	IsActive     bool          `db:"is_active"`
}

func (q *Queries) UserInsert(ctx context.Context, arg UserInsertParams) (int64, error) {
	result, err := q.exec(ctx, q.userInsertStmt, userInsert,
		arg.UserID,
		arg.Uuid,
		arg.Name,
		arg.Email,
		pq.Array(arg.UserTags),
		pq.Array(arg.SystemTags),
		arg.CreatedBy,
		arg.CommonAccess,
		arg.IsActive,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const userUpdate = `-- name: UserUpdate :many
UPDATE "user" as u
SET name          = $1
  , user_tags     = $2
  , system_tags   = $3
  , common_access = $4
  , version       = version + 1
  , is_active     = $5
WHERE u.user_id = $6
  AND u.version = $7 RETURNING user_id, uuid, created_by, created_at, version, effective_date, end_date, is_active, common_access, email, name, user_tags, system_tags
`

type UserUpdateParams struct {
	Name         string   `db:"name"`
	UserTags     []string `db:"user_tags"`
	SystemTags   []string `db:"system_tags"`
	CommonAccess string   `db:"common_access"`
	IsActive     bool     `db:"is_active"`
	UserID       int64    `db:"user_id"`
	Version      int64    `db:"version"`
}

func (q *Queries) UserUpdate(ctx context.Context, arg UserUpdateParams) ([]User, error) {
	rows, err := q.query(ctx, q.userUpdateStmt, userUpdate,
		arg.Name,
		pq.Array(arg.UserTags),
		pq.Array(arg.SystemTags),
		arg.CommonAccess,
		arg.IsActive,
		arg.UserID,
		arg.Version,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Uuid,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.EffectiveDate,
			&i.EndDate,
			&i.IsActive,
			&i.CommonAccess,
			&i.Email,
			&i.Name,
			pq.Array(&i.UserTags),
			pq.Array(&i.SystemTags),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
