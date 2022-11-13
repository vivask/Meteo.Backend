package main

import (
	api "meteo/cmd/sshclient/api/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAPIV1(router *gin.RouterGroup, db *gorm.DB) {
	sshclientAPI := api.NewSshClientAPI(db)
	sshclientAPI.RegisterAPIV1(router)
}
