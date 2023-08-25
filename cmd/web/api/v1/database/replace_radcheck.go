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

func (p databaseAPI) ReplaceRadcheck(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtRadcheck(c)
	} else {
		err = extToIntRadcheck(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtRadcheck(c *gin.Context) error {

	var radcheck []entities.Radcheck

	body, err := kit.GetInt("/esp32/database/radcheck")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &radcheck)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/radcheck", radcheck)
	if err != nil {
		return err
	}

	log.Infof("Radcheck replased [%d] records", len(radcheck))

	return nil
}

func extToIntRadcheck(c *gin.Context) error {

	var radcheck []entities.Radcheck

	body, err := kit.GetExt("/esp32/database/radcheck")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &radcheck)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/radcheck", radcheck)
	if err != nil {
		return err
	}

	log.Infof("Radcheck replased [%d] records", len(radcheck))

	return nil
}
