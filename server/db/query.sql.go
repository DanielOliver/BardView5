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

const getAcl = `-- name: GetAcl :many
SELECT DISTINCT rp.action
              , rp.subject
              , rp.conditions
              , u.user_id
FROM "user" u
         INNER JOIN role_assignment ra ON u.user_id = ra.user_id AND ra.is_active = true
         INNER JOIN role r on ra.role_id = r.role_id AND r.is_active = true
         INNER JOIN role_permission rp on r.role_id = rp.role_id AND rp.is_active = true
WHERE rp.subject = $1
  AND (rp.subject_id = $2 OR rp.subject_id IS NULL)
  AND u.user_id = $3
`

type GetAclParams struct {
	Subject   string        `db:"subject"`
	SubjectID sql.NullInt64 `db:"subject_id"`
	UserID    int64         `db:"user_id"`
}

type GetAclRow struct {
	Action     string          `db:"action"`
	Subject    string          `db:"subject"`
	Conditions json.RawMessage `db:"conditions"`
	UserID     int64           `db:"user_id"`
}

func (q *Queries) GetAcl(ctx context.Context, arg GetAclParams) ([]GetAclRow, error) {
	rows, err := q.query(ctx, q.getAclStmt, getAcl, arg.Subject, arg.SubjectID, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAclRow{}
	for rows.Next() {
		var i GetAclRow
		if err := rows.Scan(
			&i.Action,
			&i.Subject,
			&i.Conditions,
			&i.UserID,
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
INSERT INTO "user" as u (user_id, uuid, "name", email, tags, created_by)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (email) DO NOTHING
`

type UserInsertParams struct {
	UserID    int64         `db:"user_id"`
	Uuid      uuid.UUID     `db:"uuid"`
	Name      string        `db:"name"`
	Email     string        `db:"email"`
	Tags      []string      `db:"tags"`
	CreatedBy sql.NullInt64 `db:"created_by"`
}

func (q *Queries) UserInsert(ctx context.Context, arg UserInsertParams) (int64, error) {
	result, err := q.exec(ctx, q.userInsertStmt, userInsert,
		arg.UserID,
		arg.Uuid,
		arg.Name,
		arg.Email,
		pq.Array(arg.Tags),
		arg.CreatedBy,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const usersFindByUid = `-- name: UsersFindByUid :many
SELECT DISTINCT u.user_id, u.uuid, u.created_by, u.created_at, u.effective_date, u.end_date, u.is_active, u.email, u.name, u.tags
FROM role_assignment ra
         INNER JOIN role r on ra.role_id = r.role_id
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
			&i.EffectiveDate,
			&i.EndDate,
			&i.IsActive,
			&i.Email,
			&i.Name,
			pq.Array(&i.Tags),
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
