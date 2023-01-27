package main

import (
	api "meteo/cmd/messanger/api/v1"

	"github.com/gin-gonic/gin"
)

func RegisterAPIV1(router *gin.RouterGroup) {
	messangerAPI := api.NewMessangerAPI()
	messangerAPI.RegisterAPIV1(router)
}
