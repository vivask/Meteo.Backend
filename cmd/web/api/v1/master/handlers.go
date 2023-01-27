package master

import (
	"encoding/json"
	"fmt"
	"meteo/internal/config"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/kit"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p mainAPI) GetServicesState(c *gin.Context) {

	var state entities.Services

	state.ClusterService = kit.IsMainHealthy("/cluster/health")
	state.MessangerService = kit.IsMainHealthy("/messanger/health")
	state.ProxyService = kit.IsMainHealthy("/proxy/health")
	state.SshclientService = kit.IsMainHealthy("/sshclient/health")
	state.ScheduleService = kit.IsMainHealthy("/schedule/health")
	state.WebService = kit.IsMainHealthy("/web/health")
	state.Esp32Service = kit.IsMainHealthy("/esp32/health")
	state.RadiusService = kit.IsMainHealthy("/radius/health")
	state.MediaService = kit.IsMainHealthy("/media/health")
	state.NutService = kit.IsMainHealthy("/nut/health")
	state.SambaService = kit.IsMainHealthy("/media/health/samba")
	state.StorageService = kit.IsMainHealthy("/media/health/storage")
	state.TransmissionService = kit.IsMainHealthy("/media/health/transmission")

	state.GogsService = utils.RawConnect(config.Default.Client.Local, []string{"2222", "3000"})

	state.PostgresService = utils.RawConnect(config.Default.Client.Local, []string{"5432"})

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
