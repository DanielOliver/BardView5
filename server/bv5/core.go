package bv5

import (
	"database/sql"
	"fmt"
	bigcache "github.com/allegro/bigcache/v3"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"server/bardlog"
	"server/db"
	"strconv"
	"time"
)

type Generators struct {
	userNode       *snowflake.Node
	dnd5eWorldNode *snowflake.Node
}

type BardView5Configuration struct {
	KratosBaseUrl string
}

type bardView5Sessions struct {
	sessionIdCache *bigcache.BigCache
}

type BardView5 struct {
	db         *sql.DB
	querier    db.Querier
	generators *Generators
	dbMetrics  *db.WithDbMetrics
	Conf       *BardView5Configuration
	sessions   *bardView5Sessions
}

type BardView5Http struct {
	BardView5 *BardView5
	Logger    zerolog.Logger
	Session   SessionContext
	Context   *gin.Context
}

type BardView5InitConfig struct {
	ConnectionString string
	KratosBaseUrl    string
}

func ConfigNewBardView5(config *BardView5InitConfig) (bv5 *BardView5, err error) {
	pgConnection, err := sql.Open("postgres", config.ConnectionString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open bardview5 connection string")
	}
	metricsPg := db.NewDbMetrics(pgConnection, "bardview5")

	userNode, _ := snowflake.NewNode(1)
	dnd5eWorldNode, _ := snowflake.NewNode(1)

	sessionIdCache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(1 * time.Minute))

	return &BardView5{
		db:      pgConnection,
		querier: db.New(metricsPg),
		generators: &Generators{
			userNode:       userNode,
			dnd5eWorldNode: dnd5eWorldNode,
		},
		dbMetrics: metricsPg,
		Conf: &BardView5Configuration{
			KratosBaseUrl: config.KratosBaseUrl,
		},
		sessions: &bardView5Sessions{
			sessionIdCache: sessionIdCache,
		},
	}, nil
}

func NewBardView5() (bv5 *BardView5, err error) {
	connectionString := viper.GetString("CONNECTION")
	if connectionString == "" {
		return nil, fmt.Errorf("expected bardview5 sql connection string")
	}
	return ConfigNewBardView5(&BardView5InitConfig{
		ConnectionString: connectionString,
		KratosBaseUrl:    "http://proxy.local",
	})
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

func (b *bardView5Sessions) SetSessionCache(sessionCookie string, userId int64) {
	b.sessionIdCache.Set(sessionCookie, []byte(strconv.FormatInt(userId, 10)))
}

func (b *bardView5Sessions) GetSessionCache(sessionCookie string) (int64, bool) {
	userIdStr, err := b.sessionIdCache.Get(sessionCookie)
	if err != nil {
		return 0, false
	}
	userId, err := strconv.ParseInt(string(userIdStr), 10, 64)
	if err != nil {
		return 0, false
	}
	return userId, true
}

func (b *BardView5) WrapRequest(pipe func(request *BardView5Http)) func(*gin.Context) {
	return func(c *gin.Context) {
		pipe(&BardView5Http{
			BardView5: b,
			Logger:    bardlog.GetLogger(c),
			Session:   *SessionCriteria(c),
			Context:   c,
		})
	}
}

func (b *BardView5Http) Querier() db.Querier {
	return b.BardView5.querier
}

func (b *BardView5Http) GenUser() *snowflake.Node {
	return b.BardView5.generators.userNode
}

func (b *BardView5Http) GenDnd5eWorld() *snowflake.Node {
	return b.BardView5.generators.dnd5eWorldNode
}
