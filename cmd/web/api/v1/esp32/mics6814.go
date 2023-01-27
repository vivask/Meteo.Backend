package esp32

import (
	"meteo/internal/dto"
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p esp32API) GetMics6814MinByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetMics6814MinByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetMics6814MaxByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetMics6814MaxByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetMics6814AvgByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetMics6814AvgByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetMics6814MinByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetMics6814MinByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetMics6814MaxByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetMics6814MaxByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetMics6814AvgByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetMics6814AvgByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetMics6814MinByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetMics6814MinByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetMics6814MaxByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetMics6814MaxByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetMics6814AvgByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetMics6814AvgByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}
