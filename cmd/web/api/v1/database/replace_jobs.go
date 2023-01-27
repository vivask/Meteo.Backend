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

func (p databaseAPI) ReplaceJobs(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtJobs(c)
	} else {
		err = extToIntJobs(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtJobs(c *gin.Context) error {

	var jobs []entities.Jobs

	body, err := kit.GetInt("/esp32/database/jobs")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &jobs)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/jobs", jobs)
	if err != nil {
		return err
	}

	log.Info("Jobs replased [%d] records", len(jobs))

	return nil
}

func extToIntJobs(c *gin.Context) error {

	var jobs []entities.Jobs

	body, err := kit.GetExt("/esp32/database/jobs")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &jobs)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/jobs", jobs)
	if err != nil {
		return err
	}

	log.Info("Jobs replased [%d] records", len(jobs))

	return nil
}
