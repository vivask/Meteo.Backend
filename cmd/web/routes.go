package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	auth "meteo/internal/api/v1/auth"
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
			log.Debug(message)
		}
	}
}

func SetupRouter(db *gorm.DB, ui string) *gin.Engine {

	gin.DefaultWriter = colorable.NewColorableStdout()
	if viper.GetString("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	public := gin.New()
	public.MaxMultipartMemory = 8 << 20
	public.Use(Logger())
	public.Use(errors.GinError())
	public.Use(gin.Recovery())
	if enablePprof {
		pprof.Register(public, "monitor/pprof")
		log.Infof("Monitor pprof enabled")
	}
	apiV1Router := public.Group("/api/v1")
	RegisterPublicAPIV1(apiV1Router, db)
	// use swagger middleware to serve the API docs
	public.GET("/doc/*any", swagger.WrapHandler(swaggerfiles.Handler))

	// set CORS
	// router.Use(cors.Default())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	public.Use(cors.New(config))

	// no route, bad url
	public.NoRoute(noRoute)

	// set Auth
	authAPI := auth.NewAuthAPI(public, db)
	authAPI.LinkJWT()
	protected := authAPI.Protected()
	RegisterProtectedAPIV1(protected, db)

	public.Use(static.Serve("/", static.LocalFile(ui, true)))

	// handle 404 and due to Vue history mode return home page
	public.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(".", "client", "dist", "index.html"))
	})

	return public
}
