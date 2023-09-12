package v1

import "github.com/gin-gonic/gin"

func (p radiusAPI) RegisterAPIV1(router *gin.RouterGroup) {
	radius := router.Group("/radius")
	radius.GET("/health", p.GetHealth)
	radius.PUT("/stop", p.RadiusStart)
	radius.PUT("/start", p.RadiusStop)
	radius.GET("/logging", p.GetLogging)
	radius.PUT("/logging", p.ClearLogging)
	radius.GET("/logging/empty", p.IsEmptyLog)
	radius.PUT("/rotate", p.RotateRadiusLog)
}
