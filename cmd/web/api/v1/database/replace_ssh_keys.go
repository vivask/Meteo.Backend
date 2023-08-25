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

func (p databaseAPI) ReplaceSshKeys(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtSshKeys(c)
	} else {
		err = extToIntSshKeys(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtSshKeys(c *gin.Context) error {

	var ssh_keys []entities.SshKeys

	body, err := kit.GetInt("/esp32/database/ssh_keys")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &ssh_keys)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/ssh_keys", ssh_keys)
	if err != nil {
		return err
	}

	log.Infof("SshKeys replased [%d] records", len(ssh_keys))

	return nil
}

func extToIntSshKeys(c *gin.Context) error {

	var ssh_keys []entities.SshKeys

	body, err := kit.GetExt("/esp32/database/ssh_keys")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &ssh_keys)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/ssh_keys", ssh_keys)
	if err != nil {
		return err
	}

	log.Infof("SshKeys replased [%d] records", len(ssh_keys))

	return nil
}
