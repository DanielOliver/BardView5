package main

import (
	"context"
	"fmt"
	"github.com/dlmiddlecote/sqlstats"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"net/http"
	"server/bardlog"
	"server/bardlogic"
	"server/bardmetric"
	"server/db"
)

func serve() {
	bardlogic.Init()

	bardView5, err := NewBardView5()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create bardview5")
	}
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

		users, err := bardView5.Querier().UsersFindByUid(context.Background(), db.UsersFindByUidParams{SessionID: 322, Uuid: uuid.New()})
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
	router.GET("/users/:user_id/acls/:subject", bardView5.GetUserAcl)
	router.POST("/users", bardView5.CreateNewUser)

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	log.Info().Int("port", viper.GetInt("port")).Msg("Running on port")
	if err := router.Run(fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
		log.Fatal().Err(err).Msg("Failed to create http")
	}
}
