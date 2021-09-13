package main

//go:generate genny -in=gen-api-models.go -out=api-models.go gen "ApiModel=string,int"

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
	"net/http"
)

type RPG struct {
	Id   string
	Name string
}

var rpgs = []RPG{
	{
		"dnd5e",
		"Dungeons & Dragons 5e",
	},
}

// ListRPGs lists the available RPGs.
func ListRPGs(c *gin.Context) ([]RPG, error) {
	return rpgs, nil
}

// Main function
func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fizzHandler := fizz.NewFromEngine(router)

	// Initialize the information of
	// the API that will be served with
	// the specification.
	infos := &openapi.Info{
		Title:       "Bard View 5",
		Description: `RPG universe management server. Focus on Dungeons & Dragons 5e compliance.`,
		Version:     "0.0.1",
	}
	// Create a new route that serve the OpenAPI spec.
	fizzHandler.GET("/openapi.json", nil, fizzHandler.OpenAPI(infos, "json"))

	// Setup routes.
	rpgsGroup := fizzHandler.Group("/rpgs", "rpg", "The different RPGs supported")
	// List all available RPGs.
	rpgsGroup.GET("", []fizz.OperationOption{
		fizz.Summary("List the rpgs available"),
		fizz.Response("400", "Bad request", nil, nil, nil),
	}, tonic.Handler(ListRPGs, 200))

	srv := &http.Server{
		Addr:    ":4242",
		Handler: fizzHandler,
	}
	srv.ListenAndServe()
}
