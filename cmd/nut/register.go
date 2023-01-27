package main

import (
	api "meteo/cmd/nut/api/v1"

	"github.com/gin-gonic/gin"
)

func RegisterAPIV1(router *gin.RouterGroup) {
	nutAPI := api.NewNutAPI()
	nutAPI.RegisterAPIV1(router)
	nutAPI.StartService()
}
