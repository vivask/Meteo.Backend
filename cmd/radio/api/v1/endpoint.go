package v1

import "github.com/gin-gonic/gin"

func (p radioAPI) RegisterAPIV1(router *gin.RouterGroup) {
	radio := router.Group("/radio")
	radio.GET("/health", p.GetHealth)
	// radio.GET("/logging", p.GetLogging)
	// radio.PUT("/logging", p.ClearLogging)
	// radio.GET("/logging/empty", p.IsEmptyLog)
}
