package database

import (
	"meteo/internal/entities"
	"meteo/internal/errors"
	"net/http"

	lock "meteo/internal/repo/esp32"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) GetNotSyncDs18b20(c *gin.Context) {
	bmx280, err := p.repo.GetNotSyncDs18b20()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, bmx280)
}

func (p databaseAPI) GetAllDs18b20(c *gin.Context) {
	bmx280, err := p.repo.GetAllDs18b20()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, bmx280)
}

func (p databaseAPI) AddSyncDs18b20(c *gin.Context) {

	var bmx280 []entities.Ds18b20

	if err := c.ShouldBind(&bmx280); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.AddSyncDs18b20(bmx280)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) LockDs18b20(c *gin.Context) {
	err := lock.LockDs18b20(false)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) UnlockDs18b20(c *gin.Context) {
	err := lock.UnlockDs18b20(false)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) SyncDs18b20(c *gin.Context) {
	err := p.repo.SyncDs18b20()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) ReplaceDs18b20(c *gin.Context) {

	var ds18b20 []entities.Ds18b20

	if err := c.ShouldBind(&ds18b20); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceDs18b20(ds18b20)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}
