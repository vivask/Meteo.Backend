package v1

import "github.com/gin-gonic/gin"

// MessangerAPI api interface
type MessangerAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
	SendTelegram(*gin.Context)
}

type messangerAPI struct {
}

// NewMessangerAPI get mesanger service instance
func NewMessangerAPI() MessangerAPI {
	return &messangerAPI{}
}
