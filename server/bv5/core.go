package bv5

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"server/db"
)


type Generators struct {
	userNode *snowflake.Node
}

type BardView5 struct {
	db         *sql.DB
	querier    db.Querier
	generators *Generators
	dbMetrics  *db.WithDbMetrics
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
	metricsPg := db.NewDbMetrics(pgConnection, "bardview5")

	userNode, err := snowflake.NewNode(1)

	return &BardView5{
		db:      pgConnection,
		querier: db.New(metricsPg),
		generators: &Generators{
			userNode: userNode,
		},
		dbMetrics: metricsPg,
	}, nil
}

func (b *BardView5) Metrics() []prometheus.Collector {
	return b.dbMetrics.Collectors()
}

func (b *BardView5) DB() *sql.DB {
	return b.db
}

func (b *BardView5) Querier() db.Querier {
	return b.querier
}