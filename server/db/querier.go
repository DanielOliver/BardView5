// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	UserInsert(ctx context.Context, arg UserInsertParams) (int64, error)
	UsersFindByUid(ctx context.Context, arg UsersFindByUidParams) ([]User, error)
}

var _ Querier = (*Queries)(nil)