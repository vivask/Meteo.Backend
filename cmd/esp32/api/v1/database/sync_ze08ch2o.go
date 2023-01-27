package database

import (
	"meteo/internal/entities"
	"meteo/internal/errors"
	"net/http"

	lock "meteo/internal/repo/esp32"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) GetNotSyncZe08ch2o(c *gin.Context) {
	bmx280, err := p.repo.GetNotSyncZe08ch2o()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, bmx280)
}

func (p databaseAPI) GetAllZe08ch2o(c *gin.Context) {
	bmx280, err := p.repo.GetAllZe08ch2o()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, bmx280)
}

func (p databaseAPI) AddSyncZe08ch2o(c *gin.Context) {

	var bmx280 []entities.Ze08ch2o

	if err := c.ShouldBind(&bmx280); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.AddSyncZe08ch2o(bmx280)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) LockZe08ch2o(c *gin.Context) {
	err := lock.LockZe08ch2o(false)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) UnlockZe08ch2o(c *gin.Context) {
	err := lock.UnlockZe08ch2o(false)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) SyncZe08ch2o(c *gin.Context) {
	err := p.repo.SyncZe08ch2o()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) ReplaceZe08ch2o(c *gin.Context) {

	var bmx280 []entities.Ze08ch2o

	if err := c.ShouldBind(&bmx280); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceZe08ch2o(bmx280)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}
