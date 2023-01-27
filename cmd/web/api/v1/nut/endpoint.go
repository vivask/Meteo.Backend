package nut

import "github.com/gin-gonic/gin"

func (p nutAPI) RegisterAPIV1(router *gin.RouterGroup) {
	nut := router.Group("/nut")
	nut.GET("/status", p.GetState)
}
