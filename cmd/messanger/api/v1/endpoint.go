package v1

import "github.com/gin-gonic/gin"

func (p messangerAPI) RegisterAPIV1(router *gin.RouterGroup) {
	messanger := router.Group("/messanger")
	messanger.POST("/telegram", p.SendTelegram)
	messanger.PUT("/telegram/schedule", p.ScheduleSendTelegram)
	messanger.GET("/health", p.GetHealth)
	messanger.GET("/logging", p.GetLogging)
	messanger.PUT("/logging", p.ClearLogging)
	messanger.GET("/logging/empty", p.IsEmptyLog)
}
