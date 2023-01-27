package v1

import "github.com/gin-gonic/gin"

// MediaAPI api interface
type MediaAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
	StartService()
}

type mediaAPI struct {
}

// NewMediaAPI get mesanger service instance
func NewMediaAPI() MediaAPI {
	return &mediaAPI{}
}
