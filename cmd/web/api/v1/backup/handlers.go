package backup

import (
	"encoding/json"
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p backupAPI) GetServicesState(c *gin.Context) {

	var (
		state entities.Services
		body  []byte
		err   error
	)

	if kit.IsBackupHealthy("/sshclient/health") {
		body, err = kit.GetBackup("/sshclient/backup/state")
	} else {
		body, err = kit.GetBackup("/cluster/backup/state")
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
