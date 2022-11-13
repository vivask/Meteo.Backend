package main

import (
	api "meteo/cmd/proxy/api/v1"
	"meteo/internal/config"
	"meteo/internal/log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAPIV1(router *gin.RouterGroup, db *gorm.DB) {
	proxyAPI := api.NewProxyAPI(db)
	proxyAPI.RegisterAPIV1(router)
	if config.Default.Proxy.Active {
		err := proxyAPI.Start()
		if err != nil {
			log.Fatalf("Can't start %s Server, error: %v", config.Default.Proxy.Title, err)
		}
	}
}
