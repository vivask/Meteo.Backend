package esp32

import (
	"meteo/internal/errors"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p esp32API) ResetAccessPoint(c *gin.Context) {
	err := p.repo.ResetAccessPoint()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) ResetStm32(c *gin.Context) {
	err := p.repo.ResetStm32()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) ResetAvr(c *gin.Context) {
	err := p.repo.ResetAvr(true)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) GetSensorsState(c *gin.Context) {
	data, err := p.repo.GetSensorsState()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) LockBmx280(c *gin.Context) {
	val := utils.StringToBool(c.Param("val"))
	err := p.repo.LockBmx280(val)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) LockDs18b20(c *gin.Context) {
	val := utils.StringToBool(c.Param("val"))
	err := p.repo.LockDs18b20(val)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) LockRadsens(c *gin.Context) {
	val := utils.StringToBool(c.Param("val"))
	err := p.repo.LockRadsens(val)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) LockMics6814(c *gin.Context) {
	val := utils.StringToBool(c.Param("val"))
	err := p.repo.LockMics6814(val)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) LockZe08(c *gin.Context) {
	val := utils.StringToBool(c.Param("val"))
	err := p.repo.LockZe08(val)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) LockAht25(c *gin.Context) {
	val := utils.StringToBool(c.Param("val"))
	err := p.repo.LockAht25(val)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}
