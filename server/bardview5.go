package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"server/db"
)

type BardView5 struct {
	db      *sql.DB
	querier db.Querier
}

func (b *BardView5) DB() *sql.DB {
	return b.db
}

func (b *BardView5) Querier() db.Querier {
	return b.querier
}

func NewBardView5() (bv5 *BardView5, err error) {
	connectionString := viper.GetString("CONNECTION")
	if connectionString == "" {
		return nil, fmt.Errorf("expected bardview5 sql connection string")
	}
	pgConnection, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open bardview5 connection string")
	}

	return &BardView5{
		db:      pgConnection,
		querier: db.New(pgConnection),
	}, nil
}
