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

func (p databaseAPI) ReplaceRadsens(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtRadsens(c)
	} else {
		err = extToIntRadsens(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtRadsens(c *gin.Context) error {

	var radsens []entities.Radsens

	body, err := kit.GetInt("/esp32/database/replace/radsens")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &radsens)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/radsens", radsens)
	if err != nil {
		return err
	}

	log.Infof("Radsens replased [%d] records", len(radsens))

	return nil
}

func extToIntRadsens(c *gin.Context) error {

	var radsens []entities.Radsens

	body, err := kit.GetExt("/esp32/database/replace/radsens")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &radsens)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/radsens", radsens)
	if err != nil {
		return err
	}

	log.Infof("Radsens replased [%d] records", len(radsens))

	return nil
}
