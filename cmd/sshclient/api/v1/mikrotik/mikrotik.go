package mikrotik

import (
	SSH "meteo/cmd/sshclient/api/v1/internal"
	repo "meteo/internal/repo/proxy"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MikrotikAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
}

type mikrotikAPI struct {
	ssh  SSH.SshClient
	repo repo.ProxyService
}

// NewMikrotikAPI get radius service instance
func NewMikrotikAPI(c SSH.SshClient, db *gorm.DB) MikrotikAPI {
	return &mikrotikAPI{ssh: c, repo: repo.NewProxyService(db)}
}
