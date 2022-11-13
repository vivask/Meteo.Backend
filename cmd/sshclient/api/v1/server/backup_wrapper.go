package server

import (
	"meteo/internal/config"
	"meteo/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p serverAPI) RestarKodi(c *gin.Context) {

	err := p.KodiRestart()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SSHERR",
				"message": err.Error()})
		return
	}
	log.Info("Restart Kodi success")
	c.Status(http.StatusOK)
}

func (p serverAPI) RestarStorageKodi(c *gin.Context) {

	err := p.KodiStorageRestart()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SSHERR",
				"message": err.Error()})
		return
	}
	log.Info("Restart Storage Kodi success")
	c.Status(http.StatusOK)
}

func (p serverAPI) StopStorageKodi(c *gin.Context) {

	err := p.KodiStorageStop()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SSHERR",
				"message": err.Error()})
		return
	}
	log.Info("Stop Storage Kodi success")
	c.Status(http.StatusOK)
}

func (p serverAPI) StartStorageKodi(c *gin.Context) {

	err := p.KodiStorageStart()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SSHERR",
				"message": err.Error()})
		return
	}
	log.Info("Start Storage Kodi success")
	c.Status(http.StatusOK)
}

func (p serverAPI) RestartBackupCont(c *gin.Context) {

	adderess := config.Default.SshClient.Backup.Host
	port := config.Default.SshClient.Backup.Port
	username := config.Default.SshClient.Backup.User

	err := p.ContainerReboot(adderess, port, username, c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SSHERR",
				"message": err.Error()})
		return
	}
	log.Info("Restart Backup Server container success")
	c.Status(http.StatusOK)
}

func (p serverAPI) StopBackupCont(c *gin.Context) {

	adderess := config.Default.SshClient.Backup.Host
	port := config.Default.SshClient.Backup.Port
	username := config.Default.SshClient.Backup.User

	err := p.ContainerShutdown(adderess, port, username, c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SSHERR",
				"message": err.Error()})
		return
	}
	log.Info("Stop Backup Server container success")
	c.Status(http.StatusOK)
}

func (p serverAPI) StartBackupCont(c *gin.Context) {

	adderess := config.Default.SshClient.Backup.Host
	port := config.Default.SshClient.Backup.Port
	username := config.Default.SshClient.Backup.User

	err := p.ContainerStart(adderess, port, username, c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SSHERR",
				"message": err.Error()})
		return
	}
	log.Info("Start Backup Server container success")
	c.Status(http.StatusOK)
}

func (p serverAPI) BackupReboot(c *gin.Context) {

	adderess := config.Default.SshClient.Backup.Host
	port := config.Default.SshClient.Backup.Port
	username := config.Default.SshClient.Backup.User

	err := p.ServerReboot(adderess, port, username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SSHERR",
				"message": err.Error()})
		return
	}
	log.Info("Backup Server reboot success")
	c.Status(http.StatusOK)
}

func (p serverAPI) BackupShutdown(c *gin.Context) {

	adderess := config.Default.SshClient.Backup.Host
	port := config.Default.SshClient.Backup.Port
	username := config.Default.SshClient.Backup.User

	err := p.ServerShutdown(adderess, port, username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SSHERR",
				"message": err.Error()})
		return
	}
	log.Info("Backup Server shutdown success")
	c.Status(http.StatusOK)
}
