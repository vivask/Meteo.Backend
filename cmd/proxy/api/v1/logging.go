package v1

import (
	"meteo/internal/errors"
	"meteo/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p proxyAPI) GetLogging(c *gin.Context) {
	body, err := log.GetLogging()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, body)
}

func (p proxyAPI) ClearLogging(c *gin.Context) {
	err := log.ClearLogging()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}
