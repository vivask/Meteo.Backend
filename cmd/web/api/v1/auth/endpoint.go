package auth

import "github.com/gin-gonic/gin"

func (p authAPI) RegisterAPIV1(router *gin.RouterGroup) {
	router.PUT("/signup", p.Signup)
}
