package database

import (
	"meteo/internal/errors"
	"meteo/internal/kit"
	"meteo/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) SyncBmx280(c *gin.Context) {
	direction := c.Param("direction")

	var err error
	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		_, err = kit.PutInt("/esp32/database/sync/bmx280", nil)
	} else {
		_, err = kit.PutExt("/esp32/database/sync/bmx280", nil)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) SyncMics6814(c *gin.Context) {
	direction := c.Param("direction")

	var err error
	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		_, err = kit.PutInt("/esp32/database/sync/mics6814", nil)
	} else {
		_, err = kit.PutExt("/esp32/database/sync/mics6814", nil)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) SyncRadsens(c *gin.Context) {
	direction := c.Param("direction")

	var err error
	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		_, err = kit.PutInt("/esp32/database/sync/radsens", nil)
	} else {
		_, err = kit.PutExt("/esp32/database/sync/radsens", nil)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) SyncZe08ch2o(c *gin.Context) {
	direction := c.Param("direction")

	var err error
	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		_, err = kit.PutInt("/esp32/database/sync/ze08ch2o", nil)
	} else {
		_, err = kit.PutExt("/esp32/database/sync/ze08ch2o", nil)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) SyncDs18b20(c *gin.Context) {
	direction := c.Param("direction")

	var err error
	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		_, err = kit.PutInt("/esp32/database/sync/ds18b20", nil)
	} else {
		_, err = kit.PutExt("/esp32/database/sync/ds18b20", nil)
	}

	if err != nil {
		log.Error(err)
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}
