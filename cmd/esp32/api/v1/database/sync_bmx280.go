package database

import (
	"meteo/internal/entities"
	"meteo/internal/errors"
	"net/http"

	lock "meteo/internal/repo/esp32"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) GetNotSyncBmx280(c *gin.Context) {
	bmx280, err := p.repo.GetNotSyncBmx280()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, bmx280)
}

func (p databaseAPI) GetAllBmx280(c *gin.Context) {
	bmx280, err := p.repo.GetAllBmx280()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, bmx280)
}

func (p databaseAPI) AddSyncBmx280(c *gin.Context) {

	var bmx280 []entities.Bmx280

	if err := c.ShouldBind(&bmx280); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.AddSyncBmx280(bmx280)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) LockBmx280(c *gin.Context) {
	err := lock.LockBmx280(false)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) UnlockBmx280(c *gin.Context) {
	err := lock.UnlockBmx280(false)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) SyncBmx280(c *gin.Context) {
	err := p.repo.SyncBmx280()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) ReplaceBmx280(c *gin.Context) {

	var bmx280 []entities.Bmx280

	if err := c.ShouldBind(&bmx280); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceBmx280(bmx280)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}
