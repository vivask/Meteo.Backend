package sshclient

import "github.com/gin-gonic/gin"

func (p sshclientAPI) RegisterAPIV1(router *gin.RouterGroup) {
	sshclient := router.Group("/sshclient")
	sshclient.GET("/sshkeys/get", p.GetAllSshKeys)
	sshclient.POST("/sshkeys/add", p.AddSshKey)
	sshclient.DELETE("/sshkeys/:id", p.DelSshKey)
	sshclient.GET("/gitusers/get", p.GetAllGitUsers)
	sshclient.POST("/gitusers/add", p.AddGitUser)
	sshclient.POST("/gitusers/edit", p.EditGitUser)
	sshclient.DELETE("/gitusers/:id", p.DelGitUser)
	sshclient.GET("/gitservices/get", p.GetAllGitServices)
	sshclient.GET("/sshhosts/get", p.GetAllSshHosts)
	sshclient.POST("/sshhosts/add", p.AddSshHost)
	sshclient.POST("/sshhosts/edit", p.EditSshHost)
	sshclient.DELETE("/sshhosts/:id", p.DelSshHost)
}
