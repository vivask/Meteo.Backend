package v1

import (
	"github.com/gin-gonic/gin"
)

func (p esp32API) RegisterAPIV1(router *gin.RouterGroup) *gin.RouterGroup {
	esp32 := router.Group("/esp32")
	esp32.POST("", p.Handler)
	esp32.POST("/upload", p.UploadFirmware)
	esp32.GET("/health", p.GetHealth)
	esp32.GET("/logging", p.GetLogging)
	esp32.PUT("/logging", p.ClearLogging)
	return esp32
}
