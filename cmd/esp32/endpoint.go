package main

import (
	esp32 "meteo/cmd/esp32/api/v1/controller"
	database "meteo/cmd/esp32/api/v1/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAPIV1(router *gin.RouterGroup, db *gorm.DB) {
	esp32API := esp32.NewEsp32API(db)
	esp32_r := esp32API.RegisterAPIV1(router)
	database := database.NewDatabaseAPI(db)
	database.RegisterAPIV1(esp32_r)
}
