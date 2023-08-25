package v1

import "github.com/gin-gonic/gin"

func (p webAPI) RegisterPublicAPIV1(router *gin.RouterGroup) {
	p.esp32.RegisterPublicAPIV1(router)
	web := router.Group("/web")
	web.GET("/health", p.GetHealth)
	web.GET("/logging", p.GetLogging)
	web.PUT("/logging", p.ClearLogging)
	web.GET("/logging/empty", p.IsEmptyLog)
}
