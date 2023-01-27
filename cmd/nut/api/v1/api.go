package v1

import "github.com/gin-gonic/gin"

// NutAPI api interface
type NutAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
	StartService()
}

type nutAPI struct {
}

// NewNutAPI get mesanger service instance
func NewNutAPI() NutAPI {
	return &nutAPI{}
}
