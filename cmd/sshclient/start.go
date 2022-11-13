package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	configPath string
	startCmd   = &cobra.Command{
		Use:   "start",
		Short: "start sshclient",
		Long:  `start sshclient, default rest port is 13000`,
		Run:   startMessanger,
	}
	enablePprof       bool
	enableAutomigrate bool
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file (default is $PWD/config/default.yaml)")
	startCmd.PersistentFlags().Int("port", 13000, "Port to run Application server on")
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

func startMessanger(cmd *cobra.Command, agrs []string) {

	log.SetLogger(config.Default.SshClient.Title, config.Default.SshClient.LogLevel)

	log.Info("Starting sshclient...")
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

	entities.AutoMigrate(db, enableAutomigrate)

	kit.InitClient()

	router := gin.New()
	router.Use(errors.GinError())
	router.Use(gin.Recovery())
	if enablePprof {
		pprof.Register(router, "monitor/pprof")
	}
	apiV1Router := router.Group("/api/v1")
	RegisterAPIV1(apiV1Router, db)

	// run the rest server
	var address = config.Default.SshClient.Bind
	var port = config.Default.SshClient.Port
	if port > 0 {
		address = fmt.Sprintf("%s:%d", address, port)
	}

	var srv *http.Server
	// run the server with gracefull shutdown
	go func() {
		tlsMinVersion := tls.VersionTLS12

		caCert, err := os.ReadFile(config.Default.SshClient.Ca)
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
		log.Infof("starting %s server on: https://%s", config.Default.SshClient.Title, address)
		if err := srv.ListenAndServeTLS(config.Default.SshClient.Crt, config.Default.SshClient.Key); err != nil && err != http.ErrServerClosed {
			log.Error(fmt.Sprintf("listen: %s", err))
		}
	}()

	var healt *http.Server
	go func() {
		address := fmt.Sprintf("127.0.0.1:%d", config.Default.App.HealthPort)
		healt = &http.Server{
			Addr:    address,
			Handler: router,
		}
		log.Infof("starting %s health on: http://%s", config.Default.SshClient.Title, address)
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
	log.Infof("Shutdown %s Server ...", config.Default.SshClient.Title)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("%s Server Shutdown error: %s", config.Default.SshClient.Title, err)
	}
	if err := healt.Shutdown(ctx); err != nil {
		log.Errorf("%s Health Shutdown error: %s", config.Default.SshClient.Title, err)
	} // catching ctx.Done(). timeout of 1 seconds.
	<-ctx.Done()
	log.Infof("%s Server exiting", config.Default.SshClient.Title)
}