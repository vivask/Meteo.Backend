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

func (p databaseAPI) ReplaceMics6814(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtMics6814(c)
	} else {
		err = extToIntMics6814(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtMics6814(c *gin.Context) error {

	var mics6814 []entities.Mics6814

	body, err := kit.GetInt("/esp32/database/mics6814")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &mics6814)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/mics6814", mics6814)
	if err != nil {
		return err
	}

	log.Infof("Mics6814 replased [%d] records", len(mics6814))

	return nil
}

func extToIntMics6814(c *gin.Context) error {

	var mics6814 []entities.Mics6814

	body, err := kit.GetExt("/esp32/database/mics6814")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &mics6814)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/mics6814", mics6814)
	if err != nil {
		return err
	}

	log.Infof("Mics6814 replased [%d] records", len(mics6814))

	return nil
}
