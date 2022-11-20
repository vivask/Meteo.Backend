package controller

import (
	"github.com/gin-gonic/gin"
)

func (p esp32API) RegisterAPIV1(router *gin.RouterGroup) *gin.RouterGroup {
	esp32 := router.Group("/esp32")
	esp32.POST("", p.Handler)
	esp32.POST("/upload", p.UploadFirmware)
	esp32.GET("/health", p.GetHealth)
	return esp32
}
