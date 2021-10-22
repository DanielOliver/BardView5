package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dlmiddlecote/sqlstats"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"net/http"
	"server/bardlog"
	"server/bardlogic"
	"server/bardmetric"
	"server/db"
)

type BardView5 struct {
	db *sql.DB
}

func (b *BardView5) DB() *sql.DB {
	return b.db
}

func NewBardView5() (*BardView5, error) {
	result := &BardView5{}

	db, err := sql.Open("postgres", viper.GetString("CONNECTION"))
	if err != nil {
		return nil, err
	}
	result.db = db
	return result, nil
}

func serve() {
	bardlogic.Init()

	bardView5, err := NewBardView5()
	if err != nil {
		panic(err)
	}
	querier := db.New(bardView5.DB())
	collector := sqlstats.NewStatsCollector("bardview5", bardView5.DB())
	prometheus := bardmetric.NewPrometheus("bv5")
	prometheus.MustRegister(collector)

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	prometheus.Use(router)
	router.Use(bardlog.UseLoggingWithRequestId(log.Logger, []string{}, nil))
	router.GET("/ping", func(c *gin.Context) {
		logger := bardlog.GetLogger(c)
		logger.Info().Msg("ping pong!")

		users, err := querier.UsersFindByUid(context.Background(), db.UsersFindByUidParams{SessionID: 322, Uid: "2"})
		if err != nil {
			logger.Err(err).Msg("can't find users")
		} else {
			logger.Info().Fields(users).Msg("results?")
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/ping2", func(c *gin.Context) {
		c.JSON(500, gin.H{
			"message": "oh no!",
		})
	})
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	log.Info().Int("port", viper.GetInt("port")).Msg("Running on port")
	if err := router.Run(fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
		panic(err)
	}
}
