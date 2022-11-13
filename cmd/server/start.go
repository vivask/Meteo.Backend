package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"meteo/docs"
	"meteo/internal/config"
	entities "meteo/internal/entities/migration"
	"meteo/internal/errors"
	"meteo/internal/kit"
	"meteo/internal/log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

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
		Long:  `start server, default rest port is 12000`,
		Run:   startServer,
	}
	enablePprof       bool
	enableAutomigrate bool
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file (default is $PWD/config/default.yaml)")
	startCmd.PersistentFlags().Int("port", 12000, "Port to run Application server on")
	startCmd.PersistentFlags().BoolVarP(&enablePprof, "pprof", "p", false, "enable pprof mode (default: false)")
	startCmd.PersistentFlags().BoolVarP(&enableAutomigrate, "migrate", "m", false, "enable auto migrate (default: false)")
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

	log.SetLogger(config.Default.Server.Title, config.Default.Server.LogLevel)

	log.Info("Starting Server...")
	db_url := config.Default.Database.URL
	if config.Default.App.Server == "main" {
		db_url = config.Default.Database.UrlLocal
	}

	db, err := gorm.Open(postgres.Open(db_url))
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

	entities.AutoMigrate(db, enableAutomigrate)

	setupDoc()

	kit.InitClient()

	router := gin.New()
	router.Use(errors.GinError())
	router.Use(gin.Recovery())
	if enablePprof {
		pprof.Register(router, "monitor/pprof")
	}
	apiV1Router := router.Group("/api/v1")
	RegisterAPIV1(apiV1Router, db)
	// use swagger middleware to serve the API docs
	router.GET("/doc/*any", swagger.WrapHandler(swaggerfiles.Handler))

	// run the rest server
	address := config.Default.Server.Bind
	port := config.Default.Server.Port
	if port > 0 {
		address = fmt.Sprintf("%s:%d", address, port)
	}

	var srv *http.Server
	// run the server with gracefull shutdown
	go func() {
		tlsMinVersion := tls.VersionTLS12

		caCert, err := os.ReadFile(config.Default.Server.Ca)
		if err != nil {
			log.Fatalf("error read CA: %w", err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		cfg := &tls.Config{
			ClientAuth:       tls.RequireAndVerifyClientCert,
			ClientCAs:        caCertPool,
			MinVersion:       uint16(tlsMinVersion),
			CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		}

		srv = &http.Server{
			Addr:      address,
			TLSConfig: cfg,
			Handler:   router,
		}
		log.Infof("starting %s server on: https://%s", config.Default.Server.Title, address)
		if err := srv.ListenAndServeTLS(config.Default.Server.Crt, config.Default.Server.Key); err != nil && err != http.ErrServerClosed {
			log.Error(fmt.Sprintf("listen: %s", err))
		}
	}()

	var healt *http.Server
	go func() {
		address := fmt.Sprintf("127.0.0.1:%d", 25000)
		healt = &http.Server{
			Addr:    address,
			Handler: router,
		}
		log.Infof("starting %s health on: http://%s", config.Default.Server.Title, address)
		if err := healt.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(fmt.Sprintf("listen: %s", err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infof("Shutdown %s Server ...", config.Default.Server.Title)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("%s Server Shutdown error: %s", config.Default.Server.Title, err)
	}
	if err := healt.Shutdown(ctx); err != nil {
		log.Errorf("%s Health Shutdown error: %s", config.Default.Server.Title, err)
	}
	// catching ctx.Done(). timeout of 1 seconds.
	<-ctx.Done()
	log.Infof("%s Server exiting", config.Default.Server.Title)
}

func setupDoc() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Meteo API"
	docs.SwaggerInfo.Description = "This is meteo golang server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", config.Default.Web.Listen, config.Default.Web.Port)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}