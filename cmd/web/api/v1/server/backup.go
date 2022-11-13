package server

import (
	"fmt"
	"meteo/internal/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p serverAPI) GetBackupServicesState(c *gin.Context) {

	var state Services

	req, err := kit.GetBackup("/cluster/health")
	state.ClusterService = (err == nil) && (len(req) == 0)

	req, err = kit.GetBackup("/messanger/health")
	state.MessangerService = (err == nil) && (len(req) == 0)

	req, err = kit.GetBackup("/server/health")
	state.ServerService = (err == nil) && (len(req) == 0)

	req, err = kit.GetBackup("/esp32/health")
	state.Esp32Service = (err == nil) && (len(req) == 0)

	req, err = kit.GetBackup("/proxy/health")
	state.ProxyService = (err == nil) && (len(req) == 0)

	req, err = kit.GetBackup("/sshclient/health")
	state.SshclientService = (err == nil) && (len(req) == 0)

	req, err = kit.GetBackup("/schedule/health")
	state.ScheduleService = (err == nil) && (len(req) == 0)

	req, err = kit.GetBackup("/web/health")
	state.WebService = (err == nil) && (len(req) == 0)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": state})
}

func (p serverAPI) RestarKodi(c *gin.Context) {

	_, err := kit.PutBackup("/sshclient/server/backup/kodi/restart", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) RestarStorageKodi(c *gin.Context) {

	_, err := kit.PutBackup("/sshclient/server/backup/kodi/storage/restart", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) StopStorageKodi(c *gin.Context) {

	_, err := kit.PutBackup("/sshclient/server/backup/kodi/storage/stop", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) StartStorageKodi(c *gin.Context) {

	_, err := kit.PutBackup("/sshclient/server/backup/kodi/storage/start", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) RestartBackupServerCont(c *gin.Context) {

	url := fmt.Sprintf("/sshclient/server/backup/restart/%s", c.Param("id"))
	_, err := kit.PutMain(url, nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) StopBackupServerCont(c *gin.Context) {

	url := fmt.Sprintf("/sshclient/server/backup/stop/%s", c.Param("id"))
	_, err := kit.PutMain(url, nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) StartBackupServerCont(c *gin.Context) {

	url := fmt.Sprintf("/sshclient/server/backup/start/%s", c.Param("id"))
	_, err := kit.PutMain(url, nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) BackupReboot(c *gin.Context) {

	_, err := kit.PutBackup("/sshclient/server/backup/reboot", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) BackupShutdown(c *gin.Context) {

	_, err := kit.PutBackup("/sshclient/server/backup/shutdown", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
