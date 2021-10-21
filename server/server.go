package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"net/http"
	"server/bardlog"
	"server/bardlogic"
	"server/bardmetric"
)


func serve() {
	bardlogic.Init()

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	prometheus := bardmetric.NewPrometheus("gin")
	prometheus.Use(router)
	router.Use(bardlog.UseLoggingWithRequestId(log.Logger, []string{}, nil))
	router.GET("/ping", func(c *gin.Context) {

		logger := bardlog.GetLogger(c)
		logger.Info().Msg("ping pong!")
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
