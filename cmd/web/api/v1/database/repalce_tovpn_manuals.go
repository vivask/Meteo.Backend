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

func (p databaseAPI) ReplaceToVpnManual(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && kit.IsBackup()) {
		err = intToExtToVpnManual(c)
	} else {
		err = extToIntToVpnManual(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtToVpnManual(c *gin.Context) error {

	var tovpn_manuals []entities.ToVpnManual

	body, err := kit.GetInt("/esp32/database/tovpn_manuals")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &tovpn_manuals)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/tovpn_manuals", tovpn_manuals)
	if err != nil {
		return err
	}

	log.Infof("ToVpnManual replased [%d] records", len(tovpn_manuals))

	return nil
}

func extToIntToVpnManual(c *gin.Context) error {

	var tovpn_manuals []entities.ToVpnManual

	body, err := kit.GetExt("/esp32/database/tovpn_manuals")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &tovpn_manuals)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/tovpn_manuals", tovpn_manuals)
	if err != nil {
		return err
	}

	log.Infof("ToVpnManual replased [%d] records", len(tovpn_manuals))

	return nil
}
