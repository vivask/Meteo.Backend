package esp32

import (
	"meteo/internal/dto"
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p esp32API) GetRadsensMinByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetRadsensMinByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetRadsensMaxByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetRadsensMaxByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetRadsensAvgByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetRadsensAvgByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetRadsensMinByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetRadsensMinByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetRadsensMaxByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetRadsensMaxByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetRadsensAvgByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetRadsensAvgByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetRadsensMinByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetRadsensMinByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetRadsensMaxByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetRadsensMaxByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetRadsensAvgByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetRadsensAvgByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}
