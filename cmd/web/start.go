package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"meteo/docs"
	"meteo/internal/config"
	entities "meteo/internal/entities/migration"
	"meteo/internal/kit"
	"meteo/internal/log"
	"meteo/internal/wait"
	"strings"

	"github.com/spf13/cobra"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	configPath string
	startCmd   = &cobra.Command{
		Use:   "start",
		Short: "start web",
		Long:  `start web, default port is 443`,
		Run:   startWebServer,
	}
	enablePprof       bool
	enableAutomigrate bool
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file (default is $PWD/config/default.yaml)")
	startCmd.PersistentFlags().BoolVarP(&enablePprof, "pprof", "p", false, "enable pprof mode (default: false)")
	startCmd.PersistentFlags().BoolVarP(&enableAutomigrate, "migrate", "m", false, "enable auto migrate (default: false)")
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

func startWebServer(cmd *cobra.Command, agrs []string) {

	log.SetLogger(config.Default.Web.Title, config.Default.Web.LogLevel)

	WaitForStartCluster()

	log.Info("Starting http-server...")

	db, err := gorm.Open(postgres.Open(getDbUrl(config.Default.Web.DbLink)))
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

	const path = "./firmware"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	router := SetupRouter(db, config.Default.Web.Ui)

	// run the server
	var address = config.Default.Web.Listen
	var port = config.Default.Web.Port
	if port > 0 {
		address = fmt.Sprintf("%s:%d", address, port)
	}

	var srv *http.Server

	// run the server with gracefull shutdown
	go func() {
		if !config.Default.Web.Ssl {
			srv = &http.Server{
				Addr:    address,
				Handler: router,
			}

			log.Infof("starting %s server on: http://%s", config.Default.Web.Title, address)
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Error(fmt.Sprintf("listen: %s", err))
			}
		} else {
			tlsMinVersion := tls.VersionTLS12
			switch config.Default.Web.TlsMin {
			case "TLS13":
				tlsMinVersion = tls.VersionTLS13
			case "TLS12":
				tlsMinVersion = tls.VersionTLS12
			case "TLS11":
				tlsMinVersion = tls.VersionTLS11
			case "TLS10":
				tlsMinVersion = tls.VersionTLS10
			}

			caCert, err := os.ReadFile(config.Default.Web.Ca)
			if err != nil {
				log.Fatalf("error read CA: %w", err)
			}
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(caCert)

			cfg := &tls.Config{
				ClientAuth:       tls.NoClientCert,
				ClientCAs:        caCertPool,
				MinVersion:       uint16(tlsMinVersion),
				CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			}

			srv = &http.Server{
				Addr:      address,
				TLSConfig: cfg,
				Handler:   router,
			}
			log.Infof("starting %s server on: https://%s", config.Default.Web.Title, address)
			if err := srv.ListenAndServeTLS(config.Default.Web.Crt, config.Default.Web.Key); err != nil && err != http.ErrServerClosed {
				log.Error(fmt.Sprintf("listen: %s", err))
			}
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
	log.Infof("Shutdown %s Server ...", config.Default.Web.Title)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("%s Server Shutdown error: %v", config.Default.Web.Title, err)
	}
	// catching ctx.Done(). timeout of 2 seconds.
	<-ctx.Done()
	log.Infof("%s Server exiting", config.Default.Web.Title)
}

func setupDoc() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Meteo API"
	docs.SwaggerInfo.Description = "This is meteo golang server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", config.Default.Web.Listen, config.Default.Web.Port)
	docs.SwaggerInfo.BasePath = config.Default.App.Api
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
