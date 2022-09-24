package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	v1 "meteo/internal/api/v1"
	"meteo/internal/errors"
	"meteo/internal/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func noRoute(c *gin.Context) {
	path := strings.Split(c.Request.URL.Path, "/")
	if (path[1] != "") && (path[1] == "api") {
		c.JSON(http.StatusNotFound, gin.H{"msg": "no route", "body": nil})
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
}

// Logger provides logrus logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		// after request
		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path
		message := fmt.Sprintf("[server] | %3d | %12v |%s | %-7s %s %s",
			statusCode,
			latency,
			clientIP,
			method,
			path,
			c.Errors.String(),
		)

		switch {
		case statusCode >= 400 && statusCode <= 499:
			log.Warning(message)
		case statusCode >= 500:
			log.Error(message)
		default:
			log.Info(message)
		}
	}
}

func SetupRouter(db *gorm.DB, ui, gin_mode string) *gin.Engine {

	gin.DefaultWriter = colorable.NewColorableStdout()
	if viper.GetString("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(Logger())
	router.Use(errors.GinError())
	router.Use(gin.Recovery())
	if enablePprof {
		pprof.Register(router, "monitor/pprof")
		log.Infof("Monitor pprof enabled")
	}
	apiV1Router := router.Group("/api/v1")
	v1.RegisterRouterAPIV1(apiV1Router, db)
	// use swagger middleware to serve the API docs
	router.GET("/doc/*any", swagger.WrapHandler(swaggerfiles.Handler))

	// set CORS
	// router.Use(cors.Default())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	// no route, bad url
	router.NoRoute(noRoute)

	router.Use(static.Serve("/", static.LocalFile(ui, true)))

	// handle 404 and due to Vue history mode return home page
	router.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(".", "client", "dist", "index.html"))
	})

	return router
}