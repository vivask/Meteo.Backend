package server

import (
	SSH "meteo/cmd/sshclient/api/v1/internal"

	"github.com/gin-gonic/gin"
)

type ServerAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
}

type serverAPI struct {
	ssh SSH.SshClient
}

// NewServerAPI get radius service instance
func NewServerAPI(c SSH.SshClient) ServerAPI {
	return &serverAPI{ssh: c}
}
