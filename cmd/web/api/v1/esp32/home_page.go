package esp32

import (
	"meteo/internal/errors"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p esp32API) GetHomePageData(c *gin.Context) {
	hp, err := p.repo.GetHomePageData()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": hp})
}

func (p esp32API) Mics6814CoChk(c *gin.Context) {
	err := p.repo.Mics6814CoChk()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) Mics6814No2Chk(c *gin.Context) {
	err := p.repo.Mics6814No2Chk()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) Mics6814Nh3Chk(c *gin.Context) {
	err := p.repo.Mics6814Nh3Chk()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) Bme280TemperatureChk(c *gin.Context) {
	err := p.repo.Bme280TemperatureChk()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) RadsensStaticChk(c *gin.Context) {
	err := p.repo.RadsensStaticChk()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) RadsensDynamicChk(c *gin.Context) {
	err := p.repo.RadsensDynamicChk()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) RadsensHVSet(c *gin.Context) {
	err := p.repo.RadsensHVSet()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) RadsensSetSens(c *gin.Context) {
	val, err := utils.StringToUint(c.Param("val"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	err = p.repo.RadsensSetSens(val)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) Ds18b20TemperatureChk(c *gin.Context) {
	err := p.repo.Ds18b20TemperatureChk()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) Ze08ch2oChk(c *gin.Context) {
	err := p.repo.Ze08ch2oChk()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}
