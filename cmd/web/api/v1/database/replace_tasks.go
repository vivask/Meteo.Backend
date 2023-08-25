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

func (p databaseAPI) ReplaceTasks(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtTasks(c)
	} else {
		err = extToIntTasks(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtTasks(c *gin.Context) error {

	var tasks []entities.Tasks

	body, err := kit.GetInt("/esp32/database/tasks")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &tasks)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/tasks", tasks)
	if err != nil {
		return err
	}

	log.Infof("Tasks replased [%d] records", len(tasks))

	return nil
}

func extToIntTasks(c *gin.Context) error {

	var tasks []entities.Tasks

	body, err := kit.GetExt("/esp32/database/tasks")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &tasks)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/tasks", tasks)
	if err != nil {
		return err
	}

	log.Infof("Tasks replased [%d] records", len(tasks))

	return nil
}
