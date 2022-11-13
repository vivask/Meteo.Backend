package v1

import "github.com/gin-gonic/gin"

func (p authAPI) RegisterAPIV1(router *gin.RouterGroup) {
	router.POST("/signup", p.Signup)
}
