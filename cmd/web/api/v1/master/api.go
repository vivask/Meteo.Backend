package master

import "github.com/gin-gonic/gin"

type MainAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
}

type mainAPI struct {
}

// NewMainAPI get main server instance
func NewMainAPI() MainAPI {
	return &mainAPI{}
}
