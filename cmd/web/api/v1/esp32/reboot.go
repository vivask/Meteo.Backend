package esp32

import (
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p esp32API) SetRebootEsp32(c *gin.Context) {
	err := p.repo.Esp32Reboot()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}
