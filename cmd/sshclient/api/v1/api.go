package v1

import (
	SSH "meteo/cmd/sshclient/api/v1/internal"
	"meteo/cmd/sshclient/api/v1/mikrotik"
	"meteo/cmd/sshclient/api/v1/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SshClientAPI api interface
type SshClientAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
}

type sshclientAPI struct {
	ssh      SSH.SshClient
	server   server.ServerAPI
	mikrotik mikrotik.MikrotikAPI
}

// NewSshClientAPI get sshclient service instance
func NewSshClientAPI(db *gorm.DB) SshClientAPI {
	ssh := SSH.NewSshClient(db)
	return &sshclientAPI{
		ssh:      ssh,
		server:   server.NewServerAPI(ssh),
		mikrotik: mikrotik.NewMikrotikAPI(ssh),
	}
}
