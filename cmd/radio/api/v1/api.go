package v1

import "github.com/gin-gonic/gin"

// RadioAPI api interface
type RadioAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
	StartService()
}

type radioAPI struct {
}

// NewRadioAPI get mesanger service instance
func NewRadioAPI() RadioAPI {
	return &radioAPI{}
}
