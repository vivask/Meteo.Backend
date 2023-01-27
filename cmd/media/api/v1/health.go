package v1

import (
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p mediaAPI) GetHealth(c *gin.Context) {
	err := p.HealthStorage()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	err = p.HealthSamba()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	err = p.HealthTransmission()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, "healthy")
}

func (p mediaAPI) GetHealthSamba(c *gin.Context) {
	err := p.HealthSamba()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, "healthy")
}

func (p mediaAPI) GetHealthStore(c *gin.Context) {
	err := p.HealthStorage()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, "healthy")
}

func (p mediaAPI) GetHealthTransmission(c *gin.Context) {
	err := p.HealthTransmission()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, "healthy")
}
