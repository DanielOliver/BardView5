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
SELECT dnd5e_monster_id, created_by, created_at, version, dnd5e_world_id, name, user_tags, system_tags, monster_type, alignment, size_category, milli_challenge_rating, languages, description
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
			&i.Dnd5eWorldID,
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

const dnd5eMonstersFindByWorld = `-- name: Dnd5eMonstersFindByWorld :many
SELECT dnd5e_monster_id, created_by, created_at, version, dnd5e_world_id, name, user_tags, system_tags, monster_type, alignment, size_category, milli_challenge_rating, languages, description
FROM "dnd5e_monster" m
WHERE wm.dnd5e_world_id = $1
ORDER BY wm.dnd5e_world_id, wm.dnd5e_monster_id
OFFSET $2 LIMIT $3
`

type Dnd5eMonstersFindByWorldParams struct {
	Dnd5eWorldID sql.NullInt64 `db:"dnd5e_world_id"`
	RowOffset    int32         `db:"row_offset"`
	RowLimit     int32         `db:"row_limit"`
}

func (q *Queries) Dnd5eMonstersFindByWorld(ctx context.Context, arg Dnd5eMonstersFindByWorldParams) ([]Dnd5eMonster, error) {
	rows, err := q.query(ctx, q.dnd5eMonstersFindByWorldStmt, dnd5eMonstersFindByWorld, arg.Dnd5eWorldID, arg.RowOffset, arg.RowLimit)
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
			&i.Dnd5eWorldID,
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

const dnd5eWorldFindAssignment = `-- name: Dnd5eWorldFindAssignment :many
SELECT wa.created_by, wa.created_at, wa.version, wa.user_id, wa.dnd5e_world_id, wa.role_action
FROM "dnd5e_world_assignment" wa
WHERE wa.user_id = $1
  AND wa.dnd5e_world_id = $2
`

type Dnd5eWorldFindAssignmentParams struct {
	UserID       int64 `db:"user_id"`
	Dnd5eWorldID int64 `db:"dnd5e_world_id"`
}

func (q *Queries) Dnd5eWorldFindAssignment(ctx context.Context, arg Dnd5eWorldFindAssignmentParams) ([]Dnd5eWorldAssignment, error) {
	rows, err := q.query(ctx, q.dnd5eWorldFindAssignmentStmt, dnd5eWorldFindAssignment, arg.UserID, arg.Dnd5eWorldID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dnd5eWorldAssignment{}
	for rows.Next() {
		var i Dnd5eWorldAssignment
		if err := rows.Scan(
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Version,
			&i.UserID,
			&i.Dnd5eWorldID,
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

const dnd5eWorldFindByAssignment = `-- name: Dnd5eWorldFindByAssignment :many
SELECT DISTINCT w.dnd5e_world_id, w.created_by, w.created_at, w.version, w.is_active, w.common_access, w.user_tags, w.system_tags, w.name, w.module, w.description, w.external_source_id, w.external_source_key
FROM "dnd5e_world" w
         INNER JOIN "dnd5e_world_assignment" wa ON
    w.dnd5e_world_id = wa.dnd5e_world_id
WHERE wa.user_id = $1
ORDER BY w.dnd5e_world_id desc
`

func (q *Queries) Dnd5eWorldFindByAssignment(ctx context.Context, userID int64) ([]Dnd5eWorld, error) {
	rows, err := q.query(ctx, q.dnd5eWorldFindByAssignmentStmt, dnd5eWorldFindByAssignment, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dnd5eWorld{}
	for rows.Next() {
		var i Dnd5eWorld
		if err := rows.Scan(
			&i.Dnd5eWorldID,
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
			&i.ExternalSourceID,
			&i.ExternalSourceKey,
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

const dnd5eWorldFindById = `-- name: Dnd5eWorldFindById :many
SELECT dnd5e_world_id, created_by, created_at, version, is_active, common_access, user_tags, system_tags, name, module, description, external_source_id, external_source_key
FROM "dnd5e_world" w
WHERE w.dnd5e_world_id = $1
`

func (q *Queries) Dnd5eWorldFindById(ctx context.Context, dnd5eWorldID int64) ([]Dnd5eWorld, error) {
	rows, err := q.query(ctx, q.dnd5eWorldFindByIdStmt, dnd5eWorldFindById, dnd5eWorldID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dnd5eWorld{}
	for rows.Next() {
		var i Dnd5eWorld
		if err := rows.Scan(
			&i.Dnd5eWorldID,
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
			&i.ExternalSourceID,
			&i.ExternalSourceKey,
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

const dnd5eWorldInsert = `-- name: Dnd5eWorldInsert :execrows
insert into "dnd5e_world" (dnd5e_world_id, common_access, created_by, is_active, system_tags,
                           user_tags, "name", module, description)
VALUES ($1, $2, $3, $4,
        $5, $6, $7, $8, $9)
`

type Dnd5eWorldInsertParams struct {
	Dnd5eWorldID int64          `db:"dnd5e_world_id"`
	CommonAccess string         `db:"common_access"`
	CreatedBy    sql.NullInt64  `db:"created_by"`
	IsActive     bool           `db:"is_active"`
	SystemTags   []string       `db:"system_tags"`
	UserTags     []string       `db:"user_tags"`
	Name         string         `db:"name"`
	Module       sql.NullString `db:"module"`
	Description  string         `db:"description"`
}

func (q *Queries) Dnd5eWorldInsert(ctx context.Context, arg Dnd5eWorldInsertParams) (int64, error) {
	result, err := q.exec(ctx, q.dnd5eWorldInsertStmt, dnd5eWorldInsert,
		arg.Dnd5eWorldID,
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

const dnd5eWorldUpsertAssignment = `-- name: Dnd5eWorldUpsertAssignment :execrows
insert into "dnd5e_world_assignment" (created_by, user_id, dnd5e_world_id, role_action)
SELECT $1,
       $2,
       $3,
       $4 WHERE NOT EXISTS (
    SELECT 1 FROM dnd5e_world_assignment
    WHERE user_id = $2 AND dnd5e_world_id = $3 AND role_action = $4
)
`

type Dnd5eWorldUpsertAssignmentParams struct {
	CreatedBy    sql.NullInt64 `db:"created_by"`
	UserID       int64         `db:"user_id"`
	Dnd5eWorldID int64         `db:"dnd5e_world_id"`
	RoleAction   string        `db:"role_action"`
}

func (q *Queries) Dnd5eWorldUpsertAssignment(ctx context.Context, arg Dnd5eWorldUpsertAssignmentParams) (int64, error) {
	result, err := q.exec(ctx, q.dnd5eWorldUpsertAssignmentStmt, dnd5eWorldUpsertAssignment,
		arg.CreatedBy,
		arg.UserID,
		arg.Dnd5eWorldID,
		arg.RoleAction,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
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
