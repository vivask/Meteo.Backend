package mikrotik

import (
	SSH "meteo/cmd/sshclient/api/v1/internal"

	"github.com/gin-gonic/gin"
)

type MikrotikAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
}

type mikrotikAPI struct {
	ssh SSH.SshClient
}

// NewMikrotikAPI get radius service instance
func NewMikrotikAPI(c SSH.SshClient) MikrotikAPI {
	return &mikrotikAPI{ssh: c}
}
