package main

import (
	"fmt"
	"ssorc3/verkurzen/docs"
	"ssorc3/verkurzen/internal/config"
	"ssorc3/verkurzen/internal/controllers"
	"ssorc3/verkurzen/internal/data"
	"ssorc3/verkurzen/internal/log"
	"strings"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/spf13/cobra"

    swagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

var (
	configPath string
	startCmd   = &cobra.Command{
		Use:   "start",
		Short: "start server",
		Long:  `start server, default port is 5000`,
		Run:   startServer,
	}
	enablePprof bool
)

func init() {
    log.InitLogger()
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file (default is $PWD/config/default.yaml)")
	startCmd.PersistentFlags().Int("port", 5000, "Port to run Application server on")
	startCmd.PersistentFlags().BoolVarP(&enablePprof, "pprof", "p", false, "enable pprof mode (default: false)")
	config.Viper().BindPFlag("port", startCmd.PersistentFlags().Lookup("port"))
}

func initConfig() {
	if len(configPath) != 0 {
		config.Viper().SetConfigFile(configPath)
	} else {
		config.Viper().AddConfigPath("./config")
		config.Viper().SetConfigName("default")
	}
	config.Viper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.Viper().AutomaticEnv()
	if err := config.Viper().ReadInConfig(); err != nil {
		log.Logger().Fatalf("Load config from file [%s]: %v", config.Viper().ConfigFileUsed(), err)
	}
	config.Parse()
}

func startServer(cmd *cobra.Command, args []string) {
	cluster := gocql.NewCluster(config.Default.Database.URL)
	session, err := cluster.CreateSession()
	if err != nil {
		log.Logger().Fatal("Failed to connect to database", err)
	}
	defer session.Close()

    router := gin.New()

    router.Use(gin.Recovery())
    if enablePprof {
        pprof.Register(router, "monitor/pprof")
    }
    rootRouter := router.Group("/")
    rootRouter.GET("/doc/*any", swagger.WrapHandler(swaggerFiles.Handler))

    shortenRepo := data.NewShortenRepo(session)
    shortenController := controllers.NewShortenController(shortenRepo)
    shortenController.RegisterRoutes(rootRouter)

    router.Run(fmt.Sprintf("%s:%d", config.Default.Server.Host, config.Default.Server.Port))
}

func setupDoc() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Verkürzen"
	docs.SwaggerInfo.Description = "Basic URL shortener"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", config.Default.Server.Host, config.Default.Server.Port)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
