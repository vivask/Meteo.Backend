package server

import "github.com/gin-gonic/gin"

type ServerAPI interface {
	RegisterMainAPIV1(router *gin.RouterGroup)
	RegisterBackupAPIV1(router *gin.RouterGroup)
}

type serverAPI struct {
}

// NewServerAPI get radius service instance
func NewServerAPI() ServerAPI {
	return &serverAPI{}
}
