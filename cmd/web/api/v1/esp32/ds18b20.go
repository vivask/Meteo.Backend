package esp32

import (
	"meteo/internal/dto"
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p esp32API) GetDs18b20MinByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetDs18b20MinByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetDs18b20MaxByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetDs18b20MaxByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetDs18b20AvgByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetDs18b20AvgByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetDs18b20MinByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetDs18b20MinByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetDs18b20MaxByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetDs18b20MaxByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetDs18b20AvgByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetDs18b20AvgByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetDs18b20MinByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetDs18b20MinByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetDs18b20MaxByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetDs18b20MaxByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetDs18b20AvgByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetDs18b20AvgByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}
