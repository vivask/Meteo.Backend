package sshclient

import "github.com/gin-gonic/gin"

func (p sshclientAPI) RegisterAPIV1(router *gin.RouterGroup) {
	sshclient := router.Group("/sshclient")
	sshclient.GET("/sshkeys", p.GetAllSshKeys)
	sshclient.PUT("/sshkeys", p.AddSshKey)
	sshclient.POST("/sshkeys", p.EditSshKey)
	sshclient.DELETE("/sshkeys/:id", p.DelSshKey)
	sshclient.GET("/gitusers", p.GetAllGitUsers)
	sshclient.PUT("/gitusers", p.AddGitUser)
	sshclient.POST("/gitusers", p.EditGitUser)
	sshclient.DELETE("/gitusers/:id", p.DelGitUser)
	sshclient.GET("/gitservices", p.GetAllGitServices)
	sshclient.GET("/sshhosts", p.GetAllSshHosts)
	sshclient.PUT("/sshhosts", p.AddSshHost)
	sshclient.POST("/sshhosts", p.EditSshHost)
	sshclient.DELETE("/sshhosts/:id", p.DelSshHost)
}
