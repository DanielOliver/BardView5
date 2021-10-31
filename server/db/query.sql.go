// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getAclBySubject = `-- name: GetAclBySubject :many
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
WHERE rp.subject = $1
  AND (rp.subject_id IS NULL
    OR rp.subject_id = $2)
  AND ra.user_id = $3
  AND ra.is_active = true
`

type GetAclBySubjectParams struct {
	Subject   string        `db:"subject"`
	SubjectID sql.NullInt64 `db:"subject_id"`
	UserID    int64         `db:"user_id"`
}

type GetAclBySubjectRow struct {
	UserID     int64           `db:"user_id"`
	Subject    string          `db:"subject"`
	Conditions json.RawMessage `db:"conditions"`
	Action     string          `db:"action"`
	SubjectID  sql.NullInt64   `db:"subject_id"`
	RoleName   string          `db:"role_name"`
	RoleID     int64           `db:"role_id"`
}

func (q *Queries) GetAclBySubject(ctx context.Context, arg GetAclBySubjectParams) ([]GetAclBySubjectRow, error) {
	rows, err := q.query(ctx, q.getAclBySubjectStmt, getAclBySubject, arg.Subject, arg.SubjectID, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAclBySubjectRow{}
	for rows.Next() {
		var i GetAclBySubjectRow
		if err := rows.Scan(
			&i.UserID,
			&i.Subject,
			&i.Conditions,
			&i.Action,
			&i.SubjectID,
			&i.RoleName,
			&i.RoleID,
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

const userFindByIdOrEmailOrUuid = `-- name: UserFindByIdOrEmailOrUuid :many
SELECT user_id, uuid, created_by, created_at, version, effective_date, end_date, is_active, common_access, email, name, user_tags, system_tags
FROM "user" u
WHERE u.user_id = $1
   OR u.email = $2
   or u.uuid = $3
ORDER BY CASE WHEN u.user_id = $1 THEN 0 ELSE 1 END
`

type UserFindByIdOrEmailOrUuidParams struct {
	UserID int64     `db:"user_id"`
	Email  string    `db:"email"`
	Uuid   uuid.UUID `db:"uuid"`
}

func (q *Queries) UserFindByIdOrEmailOrUuid(ctx context.Context, arg UserFindByIdOrEmailOrUuidParams) ([]User, error) {
	rows, err := q.query(ctx, q.userFindByIdOrEmailOrUuidStmt, userFindByIdOrEmailOrUuid, arg.UserID, arg.Email, arg.Uuid)
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
INSERT INTO "user" as u (user_id, uuid, "name", email, user_tags, system_tags, created_by, common_access)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
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
WHERE u.user_id = $5
  AND u.version = $6
RETURNING u.version, u.user_id
`

type UserUpdateParams struct {
	Name         string   `db:"name"`
	UserTags     []string `db:"user_tags"`
	SystemTags   []string `db:"system_tags"`
	CommonAccess string   `db:"common_access"`
	UserID       int64    `db:"user_id"`
	Version      int64    `db:"version"`
}

type UserUpdateRow struct {
	Version int64 `db:"version"`
	UserID  int64 `db:"user_id"`
}

func (q *Queries) UserUpdate(ctx context.Context, arg UserUpdateParams) ([]UserUpdateRow, error) {
	rows, err := q.query(ctx, q.userUpdateStmt, userUpdate,
		arg.Name,
		pq.Array(arg.UserTags),
		pq.Array(arg.SystemTags),
		arg.CommonAccess,
		arg.UserID,
		arg.Version,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserUpdateRow{}
	for rows.Next() {
		var i UserUpdateRow
		if err := rows.Scan(&i.Version, &i.UserID); err != nil {
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

const usersFindByUid = `-- name: UsersFindByUid :many
SELECT DISTINCT u.user_id, u.uuid, u.created_by, u.created_at, u.version, u.effective_date, u.end_date, u.is_active, u.common_access, u.email, u.name, u.user_tags, u.system_tags
FROM role_assignment ra
         INNER JOIN "role" r on ra.role_id = r.role_id
         INNER JOIN role_permission rp on r.role_id = rp.role_id
         INNER JOIN "user" u on evaluate_access_user(rp.conditions, $1::bigint, u.user_id)
WHERE ra.user_id = $1::bigint
  AND rp.subject = 'user'
  AND rp.is_active = true
  AND ra.is_active = true
  AND r.is_active = true
  AND u.uuid = $2
ORDER BY u.user_id
`

type UsersFindByUidParams struct {
	SessionID int64     `db:"session_id"`
	Uuid      uuid.UUID `db:"uuid"`
}

func (q *Queries) UsersFindByUid(ctx context.Context, arg UsersFindByUidParams) ([]User, error) {
	rows, err := q.query(ctx, q.usersFindByUidStmt, usersFindByUid, arg.SessionID, arg.Uuid)
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
