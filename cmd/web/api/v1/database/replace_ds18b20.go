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

func (p databaseAPI) ReplaceDs18b20(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtDs18b20(c)
	} else {
		err = extToIntDs18b20(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtDs18b20(c *gin.Context) error {

	var ds18b20 []entities.Ds18b20

	body, err := kit.GetInt("/esp32/database/ds18b20")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &ds18b20)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/ds18b20", ds18b20)
	if err != nil {
		return err
	}

	log.Infof("Ds18b20 replased [%d] records", len(ds18b20))

	return nil
}

func extToIntDs18b20(c *gin.Context) error {

	var ds18b20 []entities.Ds18b20

	body, err := kit.GetExt("/esp32/database/ds18b20")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &ds18b20)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/ds18b20", ds18b20)
	if err != nil {
		return err
	}

	log.Infof("Ds18b20 replased [%d] records", len(ds18b20))

	return nil
}
