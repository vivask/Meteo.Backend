package esp32

import (
	"meteo/internal/dto"
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p esp32API) GetZe08MinByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetZe08MinByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetZe08MaxByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetZe08MaxByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetZe08AvgByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetZe08AvgByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetZe08MinByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetZe08MinByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetZe08MaxByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetZe08MaxByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetZe08AvgByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetZe08AvgByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetZe08MinByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetZe08MinByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetZe08MaxByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetZe08MaxByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetZe08AvgByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetZe08AvgByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}
