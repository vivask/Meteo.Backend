package v1

import (
	"github.com/gin-gonic/gin"
)

func (p webAPI) RegisterProtectedAPIV1(router *gin.RouterGroup) {
	p.proxy.RegisterAPIV1(router)
	p.schedule.RegisterAPIV1(router)
	p.sshclient.RegisterAPIV1(router)
	p.database.RegisterAPIV1(router)
	p.esp32.RegisterProtectedAPIV1(router)
	p.radius.RegisterAPIV1(router)
	p.server.RegisterMainAPIV1(router)
	p.server.RegisterBackupAPIV1(router)
}

func (p webAPI) RegisterPublicAPIV1(router *gin.RouterGroup) {
	p.esp32.RegisterPublicAPIV1(router)
	web := router.Group("/web")
	web.GET("/health", p.GetHealth)
}
