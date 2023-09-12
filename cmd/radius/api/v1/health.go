package v1

import (
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p radiusAPI) GetHealth(c *gin.Context) {
	err := p.HealthRadius()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, "healthy")
}
