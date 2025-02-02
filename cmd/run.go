package main

import (
	"log"
	"os"


	"github.com/gocql/gocql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"ssorc3/verkurzen/internal/controllers"
	"ssorc3/verkurzen/internal/data"
)

func main() {
    r := gin.Default()

    logger := log.New(os.Stdout, "[API]", 0)

    // Setup CORS
    corsConfig := cors.DefaultConfig()

    corsConfig.AllowOriginFunc = func(origin string) bool { return true }

    r.Use(cors.New(corsConfig))

    r.ForwardedByClientIP = true
    r.SetTrustedProxies([]string{"127.0.0.1"})

    // Setup database
    cluster := gocql.NewCluster("127.0.0.1")
    session, err := cluster.CreateSession()
    if err != nil {
        log.Fatal(err)
    }
    defer session.Close()
    
    shortenRepo := data.NewShortenRepo(session)
    shortenRepo.Migrate()

    // Create shorten controller
    shortenController := controllers.NewShortenController(shortenRepo, logger)
    shortenController.RegisterRoutes(r)

    r.Run("localhost:8081")
}
