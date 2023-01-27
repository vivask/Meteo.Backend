package database

import (
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) SyncEsp32Tables(c *gin.Context) {
	err := p.repo.SyncBmx280()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	err = p.repo.SyncDs18b20()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	err = p.repo.SyncMics6814()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	err = p.repo.SyncZe08ch2o()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	err = p.repo.SyncRadsens()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}
