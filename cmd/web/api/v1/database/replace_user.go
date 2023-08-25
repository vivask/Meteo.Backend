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

func (p databaseAPI) ReplaceUser(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtUser(c)
	} else {
		err = extToIntUser(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtUser(c *gin.Context) error {

	var users []entities.User

	body, err := kit.GetInt("/esp32/database/users")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &users)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/users", users)
	if err != nil {
		return err
	}

	log.Infof("User replased [%d] records", len(users))

	return nil
}

func extToIntUser(c *gin.Context) error {

	var users []entities.User

	body, err := kit.GetExt("/esp32/database/users")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &users)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/users", users)
	if err != nil {
		return err
	}

	log.Infof("User replased [%d] records", len(users))

	return nil
}
