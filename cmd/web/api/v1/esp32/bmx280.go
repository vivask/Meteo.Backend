package esp32

import (
	"meteo/internal/dto"
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p esp32API) GetBmx280MinByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetBmx280MinByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetBmx280MaxByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetBmx280MaxByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetBmx280AvgByHours(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetBmx280AvgByHours(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetBmx280MinByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetBmx280MinByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetBmx280MaxByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetBmx280MaxByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetBmx280AvgByDays(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetBmx280AvgByDays(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetBmx280MinByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetBmx280MinByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetBmx280MaxByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetBmx280MaxByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p esp32API) GetBmx280AvgByMonths(c *gin.Context) {
	var period dto.Period
	if err := c.ShouldBind(&period); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	data, err := p.repo.GetBmx280AvgByMonths(period)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}
