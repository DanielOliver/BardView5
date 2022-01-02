// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	GetAclBySubject(ctx context.Context, arg GetAclBySubjectParams) ([]GetAclBySubjectRow, error)
	MonsterFindById(ctx context.Context, monsterID int64) ([]Monster, error)
	SizeFindAll(ctx context.Context) ([]SizeCategory, error)
	UserFindByEmail(ctx context.Context, email string) ([]User, error)
	UserFindById(ctx context.Context, userID int64) ([]User, error)
	UserFindByUuid(ctx context.Context, uuid uuid.UUID) ([]User, error)
	UserInsert(ctx context.Context, arg UserInsertParams) (int64, error)
	UserUpdate(ctx context.Context, arg UserUpdateParams) ([]User, error)
	WorldFindById(ctx context.Context, worldID int64) ([]World, error)
	WorldInsert(ctx context.Context, arg WorldInsertParams) (int64, error)
}

var _ Querier = (*Queries)(nil)
