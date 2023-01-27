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
	p.media.RegisterAPIV1(router)
	p.main.RegisterAPIV1(router)
	p.backup.RegisterAPIV1(router)
	p.nut.RegisterAPIV1(router)
}
