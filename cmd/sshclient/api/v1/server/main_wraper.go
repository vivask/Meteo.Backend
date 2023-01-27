package server

import (
	"meteo/internal/config"
	"meteo/internal/errors"
	"meteo/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p serverAPI) RestartMainCont(c *gin.Context) {

	adderess := config.Default.SshClient.Main.Host
	port := config.Default.SshClient.Main.Port
	username := config.Default.SshClient.Main.User

	err := p.ContainerReboot(adderess, port, username, c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	log.Info("Restart Main Server container success")
	c.Status(http.StatusOK)
}

func (p serverAPI) StopMainCont(c *gin.Context) {

	adderess := config.Default.SshClient.Main.Host
	port := config.Default.SshClient.Main.Port
	username := config.Default.SshClient.Main.User

	err := p.ContainerShutdown(adderess, port, username, c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	log.Info("Stop server container success")
	c.Status(http.StatusOK)
}

func (p serverAPI) StartMainCont(c *gin.Context) {

	adderess := config.Default.SshClient.Main.Host
	port := config.Default.SshClient.Main.Port
	username := config.Default.SshClient.Main.User

	err := p.ContainerStart(adderess, port, username, c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	log.Info("Start server container success")
	c.Status(http.StatusOK)
}

func (p serverAPI) MainReboot(c *gin.Context) {

	adderess := config.Default.SshClient.Main.Host
	port := config.Default.SshClient.Main.Port
	username := config.Default.SshClient.Main.User

	err := p.ServerShutdown(adderess, port, username, "sudo systemctl reboot")
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	log.Info("Main Server reboot success")
	c.Status(http.StatusOK)
}

func (p serverAPI) MainShutdown(c *gin.Context) {

	adderess := config.Default.SshClient.Main.Host
	port := config.Default.SshClient.Main.Port
	username := config.Default.SshClient.Main.User

	err := p.ServerShutdown(adderess, port, username, "sudo poweroff")
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	log.Info("Main Server shutdown success")
	c.Status(http.StatusOK)
}
