package main

import (
	api "meteo/cmd/radio/api/v1"

	"github.com/gin-gonic/gin"
)

func RegisterAPIV1(router *gin.RouterGroup) {
	radioAPI := api.NewRadioAPI()
	radioAPI.RegisterAPIV1(router)
	radioAPI.StartService()
}
