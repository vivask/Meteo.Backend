package v1

import (
	"meteo/cmd/proxy/api/v1/tools"

	"github.com/gin-gonic/gin"
)

func (p proxyAPI) ReloadUnlocker(c *gin.Context) {
	p.dns.SetUnlocker(tools.LoadUnlocker(p.repo))
}
