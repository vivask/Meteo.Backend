package backup

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

func (p backupAPI) GetServicesState(c *gin.Context) {

	var state entities.Services

	state.ClusterService = kit.IsBackupHealthy("/cluster/health")
	state.MessangerService = kit.IsBackupHealthy("/messanger/health")
	state.ProxyService = kit.IsBackupHealthy("/proxy/health")
	state.SshclientService = kit.IsBackupHealthy("/sshclient/health")
	state.ScheduleService = kit.IsBackupHealthy("/schedule/health")
	state.WebService = kit.IsBackupHealthy("/web/health")
	state.Esp32Service = kit.IsBackupHealthy("/esp32/health")
	state.RadiusService = kit.IsBackupHealthy("/radius/health")
	state.StorageService = kit.IsBackupHealthy("/sshclient/server/backup/storage/helath")

	state.PostgresService = utils.RawConnect(config.Default.Client.Local, []string{"5432"})

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": state})
}

func (p backupAPI) RestarKodi(c *gin.Context) {

	_, err := kit.PutBackup("/sshclient/server/backup/kodi/restart", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p backupAPI) RestarStorageKodi(c *gin.Context) {

	_, err := kit.PutBackup("/sshclient/server/backup/kodi/storage/restart", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p backupAPI) StopStorageKodi(c *gin.Context) {

	_, err := kit.PutBackup("/sshclient/server/backup/kodi/storage/stop", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p backupAPI) StartStorageKodi(c *gin.Context) {

	_, err := kit.PutBackup("/sshclient/server/backup/kodi/storage/start", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p backupAPI) RestartServerCont(c *gin.Context) {

	url := fmt.Sprintf("/sshclient/server/backup/restart/%s", c.Param("id"))
	_, err := kit.PutBackup(url, nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p backupAPI) StopServerCont(c *gin.Context) {

	url := fmt.Sprintf("/sshclient/server/backup/stop/%s", c.Param("id"))
	_, err := kit.PutBackup(url, nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p backupAPI) StartServerCont(c *gin.Context) {

	url := fmt.Sprintf("/sshclient/server/backup/start/%s", c.Param("id"))
	_, err := kit.PutBackup(url, nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p backupAPI) Reboot(c *gin.Context) {

	_, err := kit.PutBackup("/sshclient/server/backup/reboot", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p backupAPI) Shutdown(c *gin.Context) {

	_, err := kit.PutBackup("/sshclient/server/backup/shutdown", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p backupAPI) GetLogging(c *gin.Context) {

	url := fmt.Sprintf("/%s/logging", c.Param("id"))
	body, err := kit.GetBackup(url)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	var lines []string
	json.Unmarshal(body, &lines)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": lines})
}

func (p backupAPI) ClearLogging(c *gin.Context) {

	url := fmt.Sprintf("/%s/logging", c.Param("id"))
	_, err := kit.PutBackup(url, nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}
