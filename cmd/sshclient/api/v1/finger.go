package v1

import (
	"meteo/internal/entities"
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p sshclientAPI) GetHostFinger(c *gin.Context) {
	var touch entities.Touch
	if err := c.ShouldBind(&touch); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	finger, err := p.ssh.GetFinger(touch.User, touch.Host)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, finger)
}
