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

func (p databaseAPI) ReplaceSshHosts(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtSshHosts(c)
	} else {
		err = extToIntSshHosts(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtSshHosts(c *gin.Context) error {

	var ssh_hosts []entities.SshHosts

	body, err := kit.GetInt("/esp32/database/ssh_hosts")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &ssh_hosts)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/ssh_hosts", ssh_hosts)
	if err != nil {
		return err
	}

	log.Infof("SshHosts replased [%d] records", len(ssh_hosts))

	return nil
}

func extToIntSshHosts(c *gin.Context) error {

	var ssh_hosts []entities.SshHosts

	body, err := kit.GetExt("/esp32/database/ssh_hosts")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &ssh_hosts)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/ssh_hosts", ssh_hosts)
	if err != nil {
		return err
	}

	log.Infof("SshHosts replased [%d] records", len(ssh_hosts))

	return nil
}
