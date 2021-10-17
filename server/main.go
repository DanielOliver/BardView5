package main

//go:generate genny -in=models/gen-api-models.go -out=models/api-models.go gen "ApiModel=RPG,RPGList"
//go:generate oapi-codegen -o api/bardview5.go -package api -generate types,skip-prune bardview5.yaml
//go:generate struct2ts -o userget.ts api.UserGet api.User

//docker-compose -f docker-compose-local.yml exec db "pg_dump -U postgres -s bardview5 > /sql_dump/snapshot.sql"

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
	"net/http"
	"server/models"
)

// ListRPGs lists the available RPGs.
func ListRPGs(c *gin.Context) (models.RPGListResponseType, error) {
	return models.RPGListResponseType{
		Data: models.AllRPGs,
	}, nil
}

type RPGQueryInput struct {
	Id string `path:"id" validate:"required" description:"RPG Id"`
}

// GetRPG gets a single RPG.
func GetRPG(c *gin.Context, in *RPGQueryInput) (models.RPGResponseType, error) {
	for _, rpg := range models.AllRPGs {
		if rpg.Id == in.Id {
			return models.RPGResponseType{
				Data: models.AllRPGs[0],
			}, nil
		}
	}
	c.Status(http.StatusNotFound)
	return models.RPGResponseType{
		Errors: map[string]interface{}{
			"id": fmt.Sprintf("Id %v not found", in.Id),
		},
	}, nil
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
	rpgsGroup.GET(":id", []fizz.OperationOption{
		fizz.Summary("Gets a single RPG"),
		fizz.Response("400", "Bad request", nil, nil, nil),
	}, tonic.Handler(GetRPG, 200))

	srv := &http.Server{
		Addr:    ":4242",
		Handler: fizzHandler,
	}
	srv.ListenAndServe()
}
