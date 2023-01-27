package v1

import (
	"meteo/internal/errors"
	"meteo/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p nutAPI) GetLogging(c *gin.Context) {
	body, err := log.GetLogging()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, body)
}

func (p nutAPI) ClearLogging(c *gin.Context) {
	err := log.ClearLogging()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p nutAPI) LoggingMessage(c *gin.Context) {
	var message string
	if err := c.ShouldBind(&message); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	log.Info(message)

	c.Status(http.StatusOK)
}
