package server

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/kit"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p serverAPI) RadiusStart(c *gin.Context) {
	err := p.StartRadius()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) RadiusStop(c *gin.Context) {
	err := p.StopRadius()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) StartRadius() error {
	cmd := "/usr/sbin/radiusd -d /etc/raddb"
	if config.Default.Server.Radius.DebugMode {
		cmd = "/usr/sbin/radiusd -X -d /etc/raddb"
	}
	shell := utils.NewShell(cmd)
	err, out, _ := shell.Run(1)
	log.Debugf("Radius start out: %s", out)
	return err
}

func (p serverAPI) StopRadius() error {
	cmd := "pidof radiusd"
	shell := utils.NewShell(cmd)
	err, pid, _ := shell.Run(1)
	if err != nil {
		return fmt.Errorf("%s error: %w", cmd, err)
	}
	cmd = fmt.Sprintf("kill %s", pid)
	shell = utils.NewShell(cmd)
	err, _, _ = shell.Run(1)
	if err != nil {
		return fmt.Errorf("%s error: %w", cmd, err)
	}
	return nil
}

func (p serverAPI) HealthRadius() error {
	if kit.IsLeader() {
		cmd := fmt.Sprintf("radtest -4 -t mschap %s %s localhost:%d 10 %s",
			config.Default.Server.Radius.HealthUser,
			config.Default.Server.Radius.HealthPasswd,
			config.Default.Server.Radius.HealthPort,
			config.Default.Server.Radius.HealthKey)
		shell := utils.NewShell(cmd)
		err, _, _ := shell.Run(2)
		if err != nil {
			return fmt.Errorf("radius health error: %w", err)
		}
	}
	return nil
}
