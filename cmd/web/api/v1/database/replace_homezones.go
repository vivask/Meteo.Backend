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

func (p databaseAPI) ReplaceHomezone(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtHomezone(c)
	} else {
		err = extToIntHomezone(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtHomezone(c *gin.Context) error {

	var homezones []entities.Homezone

	body, err := kit.GetInt("/esp32/database/homezones")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &homezones)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/homezones", homezones)
	if err != nil {
		return err
	}

	log.Info("Homezone replased [%d] records", len(homezones))

	return nil
}

func extToIntHomezone(c *gin.Context) error {

	var homezones []entities.Homezone

	body, err := kit.GetExt("/esp32/database/homezones")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &homezones)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/homezones", homezones)
	if err != nil {
		return err
	}

	log.Info("Homezone replased [%d] records", len(homezones))

	return nil
}
