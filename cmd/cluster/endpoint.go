package main

import (
	api "meteo/cmd/cluster/api/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAPIV1(router *gin.RouterGroup, db *gorm.DB) {
	clusterAPI := api.NewClusterAPI(db)
	clusterAPI.RegisterAPIV1(router)
}
