package main

//go:generate genny -in=models/gen-api-models.go -out=models/api-models.go gen "ApiModel=RPG,RPGList"
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -o api/bardview5.go -package api -generate types,skip-prune bardview5.yaml
//tqwerqwer go:generate go-bindata -pkg main migrations

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
	"server/migrations"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/jteeuwen/go-bindata"
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
		if !debug {
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

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "The BardView5 database migration",
	Long:  `The BardView5 database migration`,
	Run: func(cmd *cobra.Command, args []string) {
		connectionString := viper.GetString("connection")
		if connectionString == "" {
			fmt.Println("Expected a connection string")
			os.Exit(1)
		}
		s := bindata.Resource(migrations.AssetNames(),
			func(name string) ([]byte, error) {
				return migrations.Asset(name)
			})
		d, err := bindata.WithInstance(s)
		if err != nil {
			panic(err)
		}
		m, err := migrate.NewWithSourceInstance("go-bindata", d, connectionString)
		if err != nil {
			panic(err)
		}
		err = m.Up()
		if err != nil {
			panic(err)
		}
	},
}

func configure() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.SetEnvPrefix("BARDVIEW5")
	viper.AutomaticEnv()
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

	pflag.BoolVarP(&debug, "debug", "D", false, "Debug mode")
	serveCmd.Flags().IntP("port", "p", 8080, "Port to run Application server on")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	migrateCmd.Flags().StringP("connection", "c", "", "Connection string to migrate postgresql database at")
	viper.BindPFlag("connection", migrateCmd.Flags().Lookup("connection"))
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

// Main function
func main() {
	configure()

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(migrateCmd)
	if debug {
		fmt.Println("Running in debug mode")
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
