package database

import (
	"meteo/internal/entities"
	"meteo/internal/errors"
	"net/http"

	lock "meteo/internal/repo/esp32"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) GetNotSyncRadsens(c *gin.Context) {
	bmx280, err := p.repo.GetNotSyncRadsens()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, bmx280)
}

func (p databaseAPI) GetAllRadsens(c *gin.Context) {
	bmx280, err := p.repo.GetAllRadsens()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, bmx280)
}

func (p databaseAPI) AddSyncRadsens(c *gin.Context) {

	var bmx280 []entities.Radsens

	if err := c.ShouldBind(&bmx280); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.AddSyncRadsens(bmx280)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) LockRadsens(c *gin.Context) {
	err := lock.LockRadsens(false)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) UnlockRadsens(c *gin.Context) {
	err := lock.UnlockRadsens(false)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) SyncRadsens(c *gin.Context) {
	err := p.repo.SyncRadsens()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) ReplaceRadsens(c *gin.Context) {

	var radsens []entities.Radsens

	if err := c.ShouldBind(&radsens); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceRadsens(radsens)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}
