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

func (p databaseAPI) ReplaceRadverified(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtRadverified(c)
	} else {
		err = extToIntRadverified(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtRadverified(c *gin.Context) error {

	var radverified []entities.Radverified

	body, err := kit.GetInt("/esp32/database/radverified")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &radverified)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/radverified", radverified)
	if err != nil {
		return err
	}

	log.Infof("Radverified replased [%d] records", len(radverified))

	return nil
}

func extToIntRadverified(c *gin.Context) error {

	var radverified []entities.Radverified

	body, err := kit.GetExt("/esp32/database/radverified")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &radverified)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/radverified", radverified)
	if err != nil {
		return err
	}

	log.Infof("Radverified replased [%d] records", len(radverified))

	return nil
}
