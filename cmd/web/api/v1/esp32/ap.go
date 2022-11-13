package esp32

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p esp32API) SetApMode(c *gin.Context) {
	err := p.repo.SetAccesPointMode()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "ESP32ERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
