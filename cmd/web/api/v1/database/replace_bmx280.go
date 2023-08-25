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

func (p databaseAPI) ReplaceBmx280(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtBmx280(c)
	} else {
		err = extToIntBmx280(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtBmx280(c *gin.Context) error {

	var bmx280 []entities.Bmx280

	body, err := kit.GetInt("/esp32/database/bmx280")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &bmx280)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/bmx280", bmx280)
	if err != nil {
		return err
	}

	log.Infof("Bmx280 replased [%d] records", len(bmx280))

	return nil
}

func extToIntBmx280(c *gin.Context) error {

	var bmx280 []entities.Bmx280

	body, err := kit.GetExt("/esp32/database/bmx280")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &bmx280)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/bmx280", bmx280)
	if err != nil {
		return err
	}

	log.Infof("Bmx280 replased [%d] records", len(bmx280))

	return nil
}
