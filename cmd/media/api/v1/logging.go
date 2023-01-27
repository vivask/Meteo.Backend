package v1

import (
	"meteo/internal/errors"
	"meteo/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p mediaAPI) GetLogging(c *gin.Context) {
	body, err := log.GetLogging()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, body)
}

func (p mediaAPI) ClearLogging(c *gin.Context) {
	err := log.ClearLogging()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mediaAPI) RotateMediaLogs(c *gin.Context) {
	err := p.SambaLogRotate()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	err = p.TransmissionLogRotate()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}
