package master

import (
	"encoding/json"
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p mainAPI) GetServicesState(c *gin.Context) {

	var (
		state entities.Services
		body  []byte
		err   error
	)

	if kit.IsMainHealthy("/sshclient/health") {
		body, err = kit.GetMain("/sshclient/main/state")
	} else {
		body, err = kit.GetMain("/cluster/main/state")
	}
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	err = json.Unmarshal(body, &state)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": state})
}

func (p mainAPI) RestartServerCont(c *gin.Context) {

	url := fmt.Sprintf("/sshclient/server/main/restart/%s", c.Param("id"))
	_, err := kit.PutMain(url, nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mainAPI) StopServerCont(c *gin.Context) {

	url := fmt.Sprintf("/sshclient/server/main/stop/%s", c.Param("id"))
	_, err := kit.PutMain(url, nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mainAPI) StartServerCont(c *gin.Context) {

	url := fmt.Sprintf("/sshclient/server/main/start/%s", c.Param("id"))
	_, err := kit.PutMain(url, nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mainAPI) Reboot(c *gin.Context) {

	_, err := kit.PutMain("/sshclient/server/main/reboot", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mainAPI) Shutdown(c *gin.Context) {

	_, err := kit.PutMain("/sshclient/server/main/shutdown", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mainAPI) GetLogging(c *gin.Context) {

	url := fmt.Sprintf("/%s/logging", c.Param("id"))
	body, err := kit.GetMain(url)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	var lines []string
	json.Unmarshal(body, &lines)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": lines})
}

func (p mainAPI) ClearLogging(c *gin.Context) {

	url := fmt.Sprintf("/%s/logging", c.Param("id"))
	_, err := kit.PutMain(url, nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}
