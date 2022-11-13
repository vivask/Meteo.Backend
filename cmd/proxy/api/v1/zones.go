package v1

import (
	"meteo/cmd/proxy/api/v1/tools"

	"github.com/gin-gonic/gin"
)

func (p proxyAPI) ReloadZones(c *gin.Context) {
	p.dns.SetZones(tools.LoadZones(p.repo))
}
