package main

import (
	"fmt"
	//"meteo/docs"
	"meteo/docs"
	v1 "meteo/internal/api/v1"
	"meteo/internal/config"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/log"
	"strings"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	swaggerfiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		log.Fatalf("Load config from file [%s]: %v", config.Viper().ConfigFileUsed(), err)
	}
	config.Parse()
}

func startServer(cmd *cobra.Command, agrs []string) {
	log.Info("Start http-server")
	db, err := gorm.Open(postgres.Open(config.Default.Database.URL))
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Can't connect database")
	}
	sqlDB.SetMaxOpenConns(int(config.Default.Database.Pool.Max))
	defer func() {
		sqlDB.Close()
		log.Info("Closed db connection")
	}()
	go entities.AutoMigrate(db)
	setupDoc()
	router := gin.New()
	router.Use(errors.GinError())
	router.Use(gin.Recovery())
	if enablePprof {
		pprof.Register(router, "monitor/pprof")
	}
	apiV1Router := router.Group("/api/v1")
	v1.RegisterRouterAPIV1(apiV1Router, db)
	// use swagger middleware to serve the API docs
	router.GET("/doc/*any", swagger.WrapHandler(swaggerfiles.Handler))

	router.Run(fmt.Sprintf("%s:%d", config.Default.Server.Host, config.Default.Server.Port))
}

func setupDoc() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Go Example API"
	docs.SwaggerInfo.Description = "This is example golang server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", config.Default.Server.Host, config.Default.Server.Port)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
