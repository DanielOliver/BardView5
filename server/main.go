package main

//go:generate genny -in=models/gen-api-models.go -out=models/api-models.go gen "ApiModel=RPG,RPGList"
//go:generate oapi-codegen -o api/bardview5.go -package api -generate types,skip-prune bardview5.yaml

//struct2ts -o userget.ts api.UserGet api.User
//docker-compose -f docker-compose-local.yml exec db "pg_dump -U postgres -s bardview5 > /sql_dump/snapshot.sql"
//PowerShell: docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

var debug = false

var rootCmd = &cobra.Command{
	Use:   "bardview5",
	Short: "BardView5 TTRPG manager",
	Long:  `BardView5 TTRPG manager`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "The BardView5 website",
	Long:  `The BardView5 website`,
	Run: func(cmd *cobra.Command, args []string) {
		if debug {
			fmt.Println("Running server in debug mode")
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
		router := gin.Default()
		router.Use(cors.Default())
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		fmt.Println("Port:", viper.GetInt("port"))
		if err := router.Run(fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
			panic(err)
		}
	},
}

// Main function
func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.SetEnvPrefix("BARDVIEW5_")
	viper.AddConfigPath("/etc/bardview5/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.bardview5") // call multiple times to add many search paths
	viper.AddConfigPath(".")                // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	}
	viper.AutomaticEnv()
	pflag.BoolVarP(&debug, "debug", "D", false, "Debug mode")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	serveCmd.Flags().Int("port", 8080, "Port to run Application server on")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	rootCmd.AddCommand(serveCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
