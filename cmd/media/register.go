package main

import (
	api "meteo/cmd/media/api/v1"

	"github.com/gin-gonic/gin"
)

func RegisterAPIV1(router *gin.RouterGroup) {
	mediaAPI := api.NewMediaAPI()
	mediaAPI.RegisterAPIV1(router)
	mediaAPI.StartService()
}
