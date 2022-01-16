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

func registerRoutes(router *gin.Engine, b *bv5.BardView5) {
	grpV1 := router.Group("/api/v1")
	{
		grpUsers := grpV1.Group("/users")
		{
			grpUsers.GET("/me", b.WrapRequest(bv5.GetUserThatIsMe))
			grpUsers.GET("/:userId", b.WrapRequest(bv5.GetUsersById))
			grpUsers.POST("", b.RequireValidSession, b.PostUsersCreate)
			grpUsers.PATCH("/:userId", b.RequireValidSession, b.PatchUserById)
		}
		grpDnd5e := grpV1.Group("/dnd5e")
		{
			grpDnd5eWorlds := grpDnd5e.Group("/worlds")
			{
				grpDnd5eWorlds.GET("/assigned", b.RequireValidSession, b.WrapRequest(bv5.GetMyDnd5eWorlds))
				grpDnd5eWorlds.GET("/:dnd5eWorldId", b.WrapRequest(bv5.GetDnd5eWorldById))
				grpDnd5eWorlds.POST("", b.RequireValidSession, b.WrapRequest(bv5.PostDnd5eWorldsCreate))
			}
		}
		grpV1.GET("/session", b.GetWhoAmI)
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
	router.Use(bardView5.AddSessionToContext)
	router.Use(bardlog.UseLoggingWithRequestId(log.Logger, []string{}, nil))

	registerRoutes(router, bardView5)

	log.Info().Int("port", viper.GetInt("port")).Msg("Running on port")
	if err := router.Run(fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
		log.Fatal().Err(err).Msg("Failed to create http")
	}
}
