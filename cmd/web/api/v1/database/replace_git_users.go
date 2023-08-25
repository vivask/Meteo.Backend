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

func (p databaseAPI) ReplaceGitUsers(c *gin.Context) {
	direction := c.Param("direction")

	var err error

	if (direction == "forward" && kit.IsMain()) || (direction == "back" && !kit.IsMain()) {
		err = intToExtGitUsers(c)
	} else {
		err = extToIntGitUsers(c)
	}

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func intToExtGitUsers(c *gin.Context) error {

	var git_users []entities.GitUsers

	body, err := kit.GetInt("/esp32/database/git_users")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &git_users)
	if err != nil {
		return err
	}

	_, err = kit.PutExt("/esp32/database/replace/git_users", git_users)
	if err != nil {
		return err
	}

	log.Infof("GitUsers replased [%d] records", len(git_users))

	return nil
}

func extToIntGitUsers(c *gin.Context) error {

	var git_users []entities.GitUsers

	body, err := kit.GetExt("/esp32/database/git_users")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &git_users)
	if err != nil {
		return err
	}

	_, err = kit.PutInt("/esp32/database/replace/git_users", git_users)
	if err != nil {
		return err
	}

	log.Infof("GitUsers replased [%d] records", len(git_users))

	return nil
}
