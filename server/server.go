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
			grpUsers.GET("/me", b.ApiRequireValidSession, b.WrapRequest(bv5.ApiGetUserThatIsMe))
			grpUsers.GET("/:userId", b.WrapRequest(bv5.ApiGetUsersById))
			grpUsers.POST("", b.ApiRequireValidSession, b.ApiPostUsersCreate)
			grpUsers.PATCH("/:userId", b.ApiRequireValidSession, b.ApiPatchUserById)
		}
		grpDnd5e := grpV1.Group("/dnd5e")
		{
			grpDnd5eSettings := grpDnd5e.Group("/settings")
			{
				grpDnd5eSettings.GET("/assigned", b.ApiRequireValidSession, b.WrapRequest(bv5.ApiGetMyDnd5eSettings))
				grpDnd5eSettings.GET("/:dnd5eSettingId/monsters", b.WrapRequest(bv5.ApiGetDnd5eMonstersBySettingId))
				grpDnd5eSettings.GET("/:dnd5eSettingId", b.WrapRequest(bv5.ApiGetDnd5eSettingById))
				grpDnd5eSettings.POST("/:dnd5eSettingId", b.WrapRequest(bv5.ApiPostDnd5eSettingsEdit))
				grpDnd5eSettings.GET("", b.ApiRequireValidSession, b.WrapRequest(bv5.ApiGetDnd5eSettings))
				grpDnd5eSettings.POST("", b.ApiRequireValidSession, b.WrapRequest(bv5.ApiPostDnd5eSettingsCreate))
			}
			grpDnd5eMonsters := grpDnd5e.Group("/monsters")
			{
				grpDnd5eMonsters.POST("", b.ApiRequireValidSession, b.WrapRequest(bv5.ApiPostDnd5eMonstersCreate))
			}
		}
		grpV1.GET("/session", b.ApiGetWhoAmI)
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
