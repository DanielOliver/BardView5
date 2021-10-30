// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.getAclBySubjectStmt, err = db.PrepareContext(ctx, getAclBySubject); err != nil {
		return nil, fmt.Errorf("error preparing query GetAclBySubject: %w", err)
	}
	if q.userFindByIdStmt, err = db.PrepareContext(ctx, userFindById); err != nil {
		return nil, fmt.Errorf("error preparing query UserFindById: %w", err)
	}
	if q.userInsertStmt, err = db.PrepareContext(ctx, userInsert); err != nil {
		return nil, fmt.Errorf("error preparing query UserInsert: %w", err)
	}
	if q.usersFindByUidStmt, err = db.PrepareContext(ctx, usersFindByUid); err != nil {
		return nil, fmt.Errorf("error preparing query UsersFindByUid: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.getAclBySubjectStmt != nil {
		if cerr := q.getAclBySubjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAclBySubjectStmt: %w", cerr)
		}
	}
	if q.userFindByIdStmt != nil {
		if cerr := q.userFindByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFindByIdStmt: %w", cerr)
		}
	}
	if q.userInsertStmt != nil {
		if cerr := q.userInsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userInsertStmt: %w", cerr)
		}
	}
	if q.usersFindByUidStmt != nil {
		if cerr := q.usersFindByUidStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing usersFindByUidStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                  DBTX
	tx                  *sql.Tx
	getAclBySubjectStmt *sql.Stmt
	userFindByIdStmt    *sql.Stmt
	userInsertStmt      *sql.Stmt
	usersFindByUidStmt  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                  tx,
		tx:                  tx,
		getAclBySubjectStmt: q.getAclBySubjectStmt,
		userFindByIdStmt:    q.userFindByIdStmt,
		userInsertStmt:      q.userInsertStmt,
		usersFindByUidStmt:  q.usersFindByUidStmt,
	}
}
