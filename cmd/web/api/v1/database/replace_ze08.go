package database

import (
	"encoding/json"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/kit"
	"meteo/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) ReplaceZe08ch2o(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtZe08ch2o(c)
	} else {
		err = extToIntZe08ch2o(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtZe08ch2o(c *gin.Context) error {

	var ze08ch2o []entities.Ze08ch2o

	body, err := kit.GetInt("/esp32/database/ze08ch2o")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &ze08ch2o)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/ze08ch2o", ze08ch2o)
	if err != nil {
		return err
	}

	log.Info("Ze08ch2o replased [%d] records", len(ze08ch2o))

	return nil
}

func extToIntZe08ch2o(c *gin.Context) error {

	var ze08ch2o []entities.Ze08ch2o

	body, err := kit.GetExt("/esp32/database/ze08ch2o")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &ze08ch2o)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/ze08ch2o", ze08ch2o)
	if err != nil {
		return err
	}

	log.Info("Ze08ch2o replased [%d] records", len(ze08ch2o))

	return nil
}
