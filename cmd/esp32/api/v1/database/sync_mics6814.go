package database

import (
	"meteo/internal/entities"
	"meteo/internal/errors"
	"net/http"

	lock "meteo/internal/repo/esp32"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) GetNotSyncMics6814(c *gin.Context) {
	bmx280, err := p.repo.GetNotSyncMics6814()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, bmx280)
}

func (p databaseAPI) GetAllMics6814(c *gin.Context) {
	bmx280, err := p.repo.GetAllMics6814()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, bmx280)
}

func (p databaseAPI) AddSyncMics6814(c *gin.Context) {

	var bmx280 []entities.Mics6814

	if err := c.ShouldBind(&bmx280); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.AddSyncMics6814(bmx280)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) LockMics6814(c *gin.Context) {
	err := lock.LockMics6814(false)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) UnlockMics6814(c *gin.Context) {
	err := lock.UnlockMics6814(false)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) SyncMics6814(c *gin.Context) {
	err := p.repo.SyncMics6814()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) ReplaceMics6814(c *gin.Context) {

	var bmx280 []entities.Mics6814

	if err := c.ShouldBind(&bmx280); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceMics6814(bmx280)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}
