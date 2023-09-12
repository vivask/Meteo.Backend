package radius

import "github.com/gin-gonic/gin"

func (p radiusAPI) RegisterAPIV1(router *gin.RouterGroup) {
	radius := router.Group("/radius")
	radius.GET("/auth", p.GetAllUsers)
	radius.PUT("/auth", p.AddUser)
	radius.POST("/auth", p.EditUser)
	radius.DELETE("/auth/:id", p.DeleteUser)
	radius.POST("/accounting/all", p.GetAllAccounting)
	radius.POST("/accounting/verified", p.GetVerifiedAccounting)
	radius.POST("/accounting/unverified", p.GetNotVerifiedAccounting)
	radius.POST("/accounting/suspect", p.GetAlarmAccounting)
	radius.PUT("/accounting/verified/:id", p.Verified)
	radius.PUT("/accounting/clear/:id", p.ClearAccounting)
	radius.GET("/verified", p.GetAllVerified)
	radius.DELETE("/verified/:id", p.ExcludeUser)
}
