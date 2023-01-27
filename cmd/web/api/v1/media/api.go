package media

import "github.com/gin-gonic/gin"

type MediaAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
}

type mediaAPI struct {
}

// NewMediaAPI get media service instance
func NewMediaAPI() MediaAPI {
	return &mediaAPI{}
}
