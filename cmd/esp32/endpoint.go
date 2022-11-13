package main

import (
	api "meteo/cmd/esp32/api/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAPIV1(router *gin.RouterGroup, db *gorm.DB) {
	esp32API := api.NewEsp32API(db)
	esp32API.RegisterAPIV1(router)
}
