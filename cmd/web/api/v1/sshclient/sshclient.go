package sshclient

import (
	repo "meteo/internal/repo/sshclient"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SshClientAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
	GetAllSshKeys(c *gin.Context)
	AddSshKey(c *gin.Context)
	DelSshKey(c *gin.Context)
	GetAllGitUsers(c *gin.Context)
	AddGitUser(c *gin.Context)
	EditGitUser(c *gin.Context)
	DelGitUser(c *gin.Context)
}

type sshclientAPI struct {
	repo repo.SshClientService
}

// NewSshClientAPI get sshclient service instance
func NewSshClientAPI(db *gorm.DB) SshClientAPI {
	return &sshclientAPI{repo: repo.NewSshClientService(db)}
}
