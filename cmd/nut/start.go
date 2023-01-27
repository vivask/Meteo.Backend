package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"meteo/internal/config"
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
)

var (
	configPath string
	startCmd   = &cobra.Command{
		Use:   "start",
		Short: "start nut",
		Long:  `start nut, default rest port is 16000`,
		Run:   startNut,
	}
	enablePprof bool
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file (default is $PWD/config/default.yaml)")
	startCmd.PersistentFlags().Int("port", 16000, "Port to run Application server on")
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

func startNut(cmd *cobra.Command, agrs []string) {

	log.SetLogger(config.Default.Nut.Title, config.Default.Nut.LogLevel)

	log.Info("Starting nut...")

	kit.InitClient()

	router := gin.New()
	router.Use(errors.GinError())
	router.Use(gin.Recovery())
	if enablePprof {
		pprof.Register(router, "monitor/pprof")
	}
	apiV1Router := router.Group(config.Default.App.Api)
	RegisterAPIV1(apiV1Router)

	// run the rest server
	var address = config.Default.Nut.Api.Bind
	var port = config.Default.Nut.Api.Port
	if port > 0 {
		address = fmt.Sprintf("%s:%d", address, port)
	}

	var srv *http.Server
	// run the server with gracefull shutdown
	go func() {
		tlsMinVersion := tls.VersionTLS12

		caCert, err := os.ReadFile(config.Default.Nut.Api.Ca)
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
		log.Infof("starting %s server on: https://%s", config.Default.Nut.Title, address)
		if err := srv.ListenAndServeTLS(config.Default.Nut.Api.Crt, config.Default.Nut.Api.Key); err != nil && err != http.ErrServerClosed {
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
	log.Infof("Shutdown %s Server ...", config.Default.Nut.Title)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("%s Server Shutdown error: %s", config.Default.Nut.Title, err)
	}
	// catching ctx.Done(). timeout of 1 seconds.
	<-ctx.Done()
	log.Infof("%s Server exiting", config.Default.Nut.Title)
}
