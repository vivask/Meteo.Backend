package esp32

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p esp32API) SetRebootEsp32(c *gin.Context) {
	err := p.repo.Esp32Reboot()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
