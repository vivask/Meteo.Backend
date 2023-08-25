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

func (p databaseAPI) ReplaceToVpnAuto(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtToVpnAuto(c)
	} else {
		err = extToIntToVpnAuto(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtToVpnAuto(c *gin.Context) error {

	var tovpn_autos []entities.ToVpnAuto

	body, err := kit.GetInt("/esp32/database/tovpn_autos")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &tovpn_autos)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/tovpn_autos", tovpn_autos)
	if err != nil {
		return err
	}

	log.Infof("ToVpnAuto replased [%d] records", len(tovpn_autos))

	return nil
}

func extToIntToVpnAuto(c *gin.Context) error {

	var tovpn_autos []entities.ToVpnAuto

	body, err := kit.GetExt("/esp32/database/tovpn_autos")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &tovpn_autos)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/tovpn_autos", tovpn_autos)
	if err != nil {
		return err
	}

	log.Infof("ToVpnAuto replased [%d] records", len(tovpn_autos))

	return nil
}
