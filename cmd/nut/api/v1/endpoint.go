package v1

import "github.com/gin-gonic/gin"

func (p nutAPI) RegisterAPIV1(router *gin.RouterGroup) {
	nut := router.Group("/nut")
	nut.GET("/health", p.GetHealth)
	nut.GET("/logging", p.GetLogging)
	nut.PUT("/logging", p.ClearLogging)
	nut.GET("/logging/empty", p.IsEmptyLog)
	nut.POST("/message", p.LoggingMessage)
}
