package v1

import "github.com/gin-gonic/gin"

func (p sshclientAPI) RegisterAPIV1(router *gin.RouterGroup) {
	sshclient := router.Group("/sshclient")
	sshclient.PUT("/finger/get", p.GetHostFinger)
	sshclient.GET("/health", p.GetHealth)
	sshclient.GET("/logging", p.GetLogging)
	sshclient.PUT("/logging", p.ClearLogging)

	p.server.RegisterAPIV1(sshclient)
	p.mikrotik.RegisterAPIV1(sshclient)
}
