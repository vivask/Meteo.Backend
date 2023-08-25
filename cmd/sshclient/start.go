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
	"meteo/internal/wait"
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
		Run:   startSshClient,
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

func getDbUrl(link string) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.Default.Database.User,
		config.Default.Database.Password,
		link,
		config.Default.Database.Port,
		config.Default.Database.Name)
}

func WaitForStartCluster() {
	const timeout = 10
	var link = fmt.Sprintf("%s:%d", config.Default.Client.Local, config.Default.Cluster.Api.Port)
	wait.List.Set(link)
	ok := wait.List.Wait(timeout)
	if !ok {
		log.Fatalf("timeout occured after waiting for %d seconds", timeout)
	}
}

func startSshClient(cmd *cobra.Command, agrs []string) {

	log.SetLogger(config.Default.SshClient.Title, config.Default.SshClient.LogLevel)

	WaitForStartCluster()

	log.Info("Starting sshclient...")

	db, err := gorm.Open(postgres.Open(getDbUrl(config.Default.SshClient.DbLink)))
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
	apiV1Router := router.Group(config.Default.App.Api)
	RegisterAPIV1(apiV1Router, db)

	// run the rest server
	var address = config.Default.SshClient.Api.Bind
	var port = config.Default.SshClient.Api.Port
	if port > 0 {
		address = fmt.Sprintf("%s:%d", address, port)
	}

	var srv *http.Server
	// run the server with gracefull shutdown
	go func() {
		tlsMinVersion := tls.VersionTLS12

		caCert, err := os.ReadFile(config.Default.SshClient.Api.Ca)
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
		if err := srv.ListenAndServeTLS(config.Default.SshClient.Api.Crt, config.Default.SshClient.Api.Key); err != nil && err != http.ErrServerClosed {
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
	// catching ctx.Done(). timeout of 2 seconds.
	<-ctx.Done()
	log.Infof("%s Server exiting", config.Default.SshClient.Title)
}
