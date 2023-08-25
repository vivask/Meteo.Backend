package v1

import (
	"meteo/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p proxyAPI) GetHealth(c *gin.Context) {
	msg := p.dns.Lookup(config.Default.Proxy.HealthHost)
	if len(msg.Answer) > 0 && p.dns.Status() {
		c.JSON(http.StatusOK, "healthy")
	} else {
		c.JSON(http.StatusOK, "unhealthy")
	}
}
