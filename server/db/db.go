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
	if q.dnd5eLanguageFindAllStmt, err = db.PrepareContext(ctx, dnd5eLanguageFindAll); err != nil {
		return nil, fmt.Errorf("error preparing query Dnd5eLanguageFindAll: %w", err)
	}
	if q.dnd5eMonsterFindByIdStmt, err = db.PrepareContext(ctx, dnd5eMonsterFindById); err != nil {
		return nil, fmt.Errorf("error preparing query Dnd5eMonsterFindById: %w", err)
	}
	if q.dnd5eMonsterInsertStmt, err = db.PrepareContext(ctx, dnd5eMonsterInsert); err != nil {
		return nil, fmt.Errorf("error preparing query Dnd5eMonsterInsert: %w", err)
	}
	if q.dnd5eMonstersFindBySettingStmt, err = db.PrepareContext(ctx, dnd5eMonstersFindBySetting); err != nil {
		return nil, fmt.Errorf("error preparing query Dnd5eMonstersFindBySetting: %w", err)
	}
	if q.dnd5eSettingFindByAssignmentStmt, err = db.PrepareContext(ctx, dnd5eSettingFindByAssignment); err != nil {
		return nil, fmt.Errorf("error preparing query Dnd5eSettingFindByAssignment: %w", err)
	}
	if q.dnd5eSettingFindByIdStmt, err = db.PrepareContext(ctx, dnd5eSettingFindById); err != nil {
		return nil, fmt.Errorf("error preparing query Dnd5eSettingFindById: %w", err)
	}
	if q.dnd5eSettingFindByParamsStmt, err = db.PrepareContext(ctx, dnd5eSettingFindByParams); err != nil {
		return nil, fmt.Errorf("error preparing query Dnd5eSettingFindByParams: %w", err)
	}
	if q.dnd5eSettingInsertStmt, err = db.PrepareContext(ctx, dnd5eSettingInsert); err != nil {
		return nil, fmt.Errorf("error preparing query Dnd5eSettingInsert: %w", err)
	}
	if q.dnd5eSettingUpdateStmt, err = db.PrepareContext(ctx, dnd5eSettingUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query Dnd5eSettingUpdate: %w", err)
	}
	if q.dnd5eSizeCategoryFindAllStmt, err = db.PrepareContext(ctx, dnd5eSizeCategoryFindAll); err != nil {
		return nil, fmt.Errorf("error preparing query Dnd5eSizeCategoryFindAll: %w", err)
	}
	if q.roleAssignmentFindByScopeIdStmt, err = db.PrepareContext(ctx, roleAssignmentFindByScopeId); err != nil {
		return nil, fmt.Errorf("error preparing query RoleAssignmentFindByScopeId: %w", err)
	}
	if q.roleAssignmentUpsertDefaultAddStmt, err = db.PrepareContext(ctx, roleAssignmentUpsertDefaultAdd); err != nil {
		return nil, fmt.Errorf("error preparing query RoleAssignmentUpsertDefaultAdd: %w", err)
	}
	if q.roleAssignmentUpsertInitialStmt, err = db.PrepareContext(ctx, roleAssignmentUpsertInitial); err != nil {
		return nil, fmt.Errorf("error preparing query RoleAssignmentUpsertInitial: %w", err)
	}
	if q.userFindByEmailStmt, err = db.PrepareContext(ctx, userFindByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query UserFindByEmail: %w", err)
	}
	if q.userFindByIdStmt, err = db.PrepareContext(ctx, userFindById); err != nil {
		return nil, fmt.Errorf("error preparing query UserFindById: %w", err)
	}
	if q.userFindByUuidStmt, err = db.PrepareContext(ctx, userFindByUuid); err != nil {
		return nil, fmt.Errorf("error preparing query UserFindByUuid: %w", err)
	}
	if q.userInsertStmt, err = db.PrepareContext(ctx, userInsert); err != nil {
		return nil, fmt.Errorf("error preparing query UserInsert: %w", err)
	}
	if q.userUpdateStmt, err = db.PrepareContext(ctx, userUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query UserUpdate: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.dnd5eLanguageFindAllStmt != nil {
		if cerr := q.dnd5eLanguageFindAllStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing dnd5eLanguageFindAllStmt: %w", cerr)
		}
	}
	if q.dnd5eMonsterFindByIdStmt != nil {
		if cerr := q.dnd5eMonsterFindByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing dnd5eMonsterFindByIdStmt: %w", cerr)
		}
	}
	if q.dnd5eMonsterInsertStmt != nil {
		if cerr := q.dnd5eMonsterInsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing dnd5eMonsterInsertStmt: %w", cerr)
		}
	}
	if q.dnd5eMonstersFindBySettingStmt != nil {
		if cerr := q.dnd5eMonstersFindBySettingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing dnd5eMonstersFindBySettingStmt: %w", cerr)
		}
	}
	if q.dnd5eSettingFindByAssignmentStmt != nil {
		if cerr := q.dnd5eSettingFindByAssignmentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing dnd5eSettingFindByAssignmentStmt: %w", cerr)
		}
	}
	if q.dnd5eSettingFindByIdStmt != nil {
		if cerr := q.dnd5eSettingFindByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing dnd5eSettingFindByIdStmt: %w", cerr)
		}
	}
	if q.dnd5eSettingFindByParamsStmt != nil {
		if cerr := q.dnd5eSettingFindByParamsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing dnd5eSettingFindByParamsStmt: %w", cerr)
		}
	}
	if q.dnd5eSettingInsertStmt != nil {
		if cerr := q.dnd5eSettingInsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing dnd5eSettingInsertStmt: %w", cerr)
		}
	}
	if q.dnd5eSettingUpdateStmt != nil {
		if cerr := q.dnd5eSettingUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing dnd5eSettingUpdateStmt: %w", cerr)
		}
	}
	if q.dnd5eSizeCategoryFindAllStmt != nil {
		if cerr := q.dnd5eSizeCategoryFindAllStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing dnd5eSizeCategoryFindAllStmt: %w", cerr)
		}
	}
	if q.roleAssignmentFindByScopeIdStmt != nil {
		if cerr := q.roleAssignmentFindByScopeIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing roleAssignmentFindByScopeIdStmt: %w", cerr)
		}
	}
	if q.roleAssignmentUpsertDefaultAddStmt != nil {
		if cerr := q.roleAssignmentUpsertDefaultAddStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing roleAssignmentUpsertDefaultAddStmt: %w", cerr)
		}
	}
	if q.roleAssignmentUpsertInitialStmt != nil {
		if cerr := q.roleAssignmentUpsertInitialStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing roleAssignmentUpsertInitialStmt: %w", cerr)
		}
	}
	if q.userFindByEmailStmt != nil {
		if cerr := q.userFindByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFindByEmailStmt: %w", cerr)
		}
	}
	if q.userFindByIdStmt != nil {
		if cerr := q.userFindByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFindByIdStmt: %w", cerr)
		}
	}
	if q.userFindByUuidStmt != nil {
		if cerr := q.userFindByUuidStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFindByUuidStmt: %w", cerr)
		}
	}
	if q.userInsertStmt != nil {
		if cerr := q.userInsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userInsertStmt: %w", cerr)
		}
	}
	if q.userUpdateStmt != nil {
		if cerr := q.userUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userUpdateStmt: %w", cerr)
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
	db                                 DBTX
	tx                                 *sql.Tx
	dnd5eLanguageFindAllStmt           *sql.Stmt
	dnd5eMonsterFindByIdStmt           *sql.Stmt
	dnd5eMonsterInsertStmt             *sql.Stmt
	dnd5eMonstersFindBySettingStmt     *sql.Stmt
	dnd5eSettingFindByAssignmentStmt   *sql.Stmt
	dnd5eSettingFindByIdStmt           *sql.Stmt
	dnd5eSettingFindByParamsStmt       *sql.Stmt
	dnd5eSettingInsertStmt             *sql.Stmt
	dnd5eSettingUpdateStmt             *sql.Stmt
	dnd5eSizeCategoryFindAllStmt       *sql.Stmt
	roleAssignmentFindByScopeIdStmt    *sql.Stmt
	roleAssignmentUpsertDefaultAddStmt *sql.Stmt
	roleAssignmentUpsertInitialStmt    *sql.Stmt
	userFindByEmailStmt                *sql.Stmt
	userFindByIdStmt                   *sql.Stmt
	userFindByUuidStmt                 *sql.Stmt
	userInsertStmt                     *sql.Stmt
	userUpdateStmt                     *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                 tx,
		tx:                                 tx,
		dnd5eLanguageFindAllStmt:           q.dnd5eLanguageFindAllStmt,
		dnd5eMonsterFindByIdStmt:           q.dnd5eMonsterFindByIdStmt,
		dnd5eMonsterInsertStmt:             q.dnd5eMonsterInsertStmt,
		dnd5eMonstersFindBySettingStmt:     q.dnd5eMonstersFindBySettingStmt,
		dnd5eSettingFindByAssignmentStmt:   q.dnd5eSettingFindByAssignmentStmt,
		dnd5eSettingFindByIdStmt:           q.dnd5eSettingFindByIdStmt,
		dnd5eSettingFindByParamsStmt:       q.dnd5eSettingFindByParamsStmt,
		dnd5eSettingInsertStmt:             q.dnd5eSettingInsertStmt,
		dnd5eSettingUpdateStmt:             q.dnd5eSettingUpdateStmt,
		dnd5eSizeCategoryFindAllStmt:       q.dnd5eSizeCategoryFindAllStmt,
		roleAssignmentFindByScopeIdStmt:    q.roleAssignmentFindByScopeIdStmt,
		roleAssignmentUpsertDefaultAddStmt: q.roleAssignmentUpsertDefaultAddStmt,
		roleAssignmentUpsertInitialStmt:    q.roleAssignmentUpsertInitialStmt,
		userFindByEmailStmt:                q.userFindByEmailStmt,
		userFindByIdStmt:                   q.userFindByIdStmt,
		userFindByUuidStmt:                 q.userFindByUuidStmt,
		userInsertStmt:                     q.userInsertStmt,
		userUpdateStmt:                     q.userUpdateStmt,
	}
}
