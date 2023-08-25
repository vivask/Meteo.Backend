package v1

import (
	"github.com/gin-gonic/gin"
)

func (p esp32API) RegisterAPIV1(router *gin.RouterGroup) *gin.RouterGroup {
	esp32 := router.Group("/esp32")
	esp32.GET("/settings", p.GetSettings)
	esp32.GET("/orders", p.GetOrders)
	esp32.POST("/meashure", p.SetMeashure)
	esp32.PUT("/reset/ap", p.ResetAccessPoint)
	esp32.PUT("/reset/stm32", p.ResetStm32)
	esp32.PUT("/reset/radsens", p.ResetRadsens)
	esp32.PUT("/reset/journal", p.ResetJournal)
	esp32.PUT("/reset/avr", p.ResetAvr)
	esp32.POST("/logging", p.SetLoggingEsp32)
	esp32.POST("/upload", p.UploadFirmware)
	esp32.GET("/health", p.GetHealth)
	esp32.GET("/logging", p.GetLogging)
	esp32.PUT("/logging", p.ClearLogging)
	esp32.GET("/logging/empty", p.IsEmptyLog)
	return esp32
}
