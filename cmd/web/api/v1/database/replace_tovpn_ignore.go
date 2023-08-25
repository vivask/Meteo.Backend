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

func (p databaseAPI) ReplaceToVpnIgnore(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtToVpnIgnore(c)
	} else {
		err = extToIntToVpnIgnore(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtToVpnIgnore(c *gin.Context) error {

	var tovpn_ignores []entities.ToVpnIgnore

	body, err := kit.GetInt("/esp32/database/tovpn_ignores")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &tovpn_ignores)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/tovpn_ignores", tovpn_ignores)
	if err != nil {
		return err
	}

	log.Infof("ToVpnIgnore replased [%d] records", len(tovpn_ignores))

	return nil
}

func extToIntToVpnIgnore(c *gin.Context) error {

	var tovpn_ignores []entities.ToVpnIgnore

	body, err := kit.GetExt("/esp32/database/tovpn_ignores")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &tovpn_ignores)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/tovpn_ignores", tovpn_ignores)
	if err != nil {
		return err
	}

	log.Infof("ToVpnIgnore replased [%d] records", len(tovpn_ignores))

	return nil
}
