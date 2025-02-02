package main

import (
	"log"
	"os"

	"github.com/gocql/gocql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"ssorc3/verkurzen/internal/config"
	"ssorc3/verkurzen/internal/controllers"
	"ssorc3/verkurzen/internal/data"
)

func main() {
    logger := log.New(os.Stdout, "[API]", log.Default().Flags())
    r := gin.Default()

    configFile, err := os.Open("config.yaml")
    if err != nil {
        logger.Fatal("Failed to read config")
    }
    defer configFile.Close()
    config, err := config.LoadConfig(configFile)
    if err != nil {
        logger.Fatal("Failed to read config")
    }

    // Setup CORS
    corsConfig := cors.DefaultConfig()

    corsConfig.AllowOrigins = config.AllowedOrigins

    r.Use(cors.New(corsConfig))

    // Setup database
    cluster := gocql.NewCluster(config.Database.ConnectionString)
    session, err := cluster.CreateSession()
    if err != nil {
        log.Fatal(err)
    }
    defer session.Close()
    
    shortenRepo := data.NewShortenRepo(session)
    shortenRepo.Migrate()

    // Create shorten controller
    shortenController := controllers.NewShortenController(config, shortenRepo, logger)
    shortenController.RegisterRoutes(r)

    r.Run(config.BaseUrl)
}
