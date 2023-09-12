package main

import (
	api "meteo/cmd/radius/api/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAPIV1(router *gin.RouterGroup, db *gorm.DB) {
	radiusAPI := api.NewRadiusAPI(db)
	radiusAPI.RegisterAPIV1(router)
	radiusAPI.StartService()
}
