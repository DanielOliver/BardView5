package main

import (
	"fmt"
	"github.com/dlmiddlecote/sqlstats"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"server/bardlog"
	"server/bardlogic"
	"server/bardmetric"
	"server/bv5"
)

func registerRoutes(router *gin.Engine, bardView5 *bv5.BardView5) {
	grpV1 := router.Group("/v1")
	{
		grpUsers := grpV1.Group("/users")
		{
			grpUsers.POST("", bardView5.PostUsersCreate)
			grpUsers.GET("/:userId", bardView5.GetUsersById)
		}
	}
}

func serve() {
	bardlogic.Init()

	bardView5, err := bv5.NewBardView5()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create bardview5")
	}
	collector := sqlstats.NewStatsCollector("bardview5", bardView5.DB())
	prometheus := bardmetric.NewPrometheus("bv5")
	prometheus.MustRegister(collector)
	for _, metric := range bardView5.Metrics() {
		prometheus.MustRegister(metric)
	}

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	prometheus.Use(router)
	router.Use(func(c *gin.Context) {
		//TODO: solve this atrocity.
		c.Set(bv5.SessionId, "1")
	})
	router.Use(bardlog.UseLoggingWithRequestId(log.Logger, []string{}, nil))

	registerRoutes(router, bardView5)

	log.Info().Int("port", viper.GetInt("port")).Msg("Running on port")
	if err := router.Run(fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
		log.Fatal().Err(err).Msg("Failed to create http")
	}
}
