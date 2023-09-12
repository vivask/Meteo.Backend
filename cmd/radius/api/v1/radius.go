package v1

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/errors"
	"meteo/internal/kit"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p radiusAPI) RadiusStart(c *gin.Context) {
	err := p.StartRadius()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p radiusAPI) RadiusStop(c *gin.Context) {
	err := p.StopRadius()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p radiusAPI) StartRadius() error {
	cmd := "chown -R radius:radius /var/log/radius"
	shell := utils.NewShell(cmd)
	err, _, _ := shell.Run(1)
	if err != nil {
		return err
	}

	cmd = "/usr/sbin/radiusd -d /etc/raddb"
	if config.Default.Radius.DebugMode {
		cmd = "/usr/sbin/radiusd -X -d /etc/raddb"
	}
	shell = utils.NewShell(cmd)
	err, out, _ := shell.Run(1)
	log.Debugf("Radius start out: %s", out)
	return err
}

func (p radiusAPI) StopRadius() error {
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

func (p radiusAPI) HealthRadius() error {
	if kit.IsLeader() {
		cmd := fmt.Sprintf("radtest -4 -t mschap %s %s localhost:%d 10 %s",
			config.Default.Radius.HealthUser,
			config.Default.Radius.HealthPasswd,
			config.Default.Radius.HealthPort,
			config.Default.Radius.HealthKey)
		shell := utils.NewShell(cmd)
		err, _, _ := shell.Run(2)
		if err != nil {
			return fmt.Errorf("radius health error: %w", err)
		}
	}
	return nil
}

func (p radiusAPI) RadiusLogRotate() error {
	const (
		path        = "/var/log/radius"
		fName       = "radius.log"
		maxLength   = 1024 * 1024
		maxLogFiles = 4
	)

	logName := fmt.Sprintf("%s/%s", path, fName)

	err := log.Rotate(logName, maxLength)
	if err != nil {
		return fmt.Errorf("rotate [%s] error: %w", logName, err)
	}

	err = log.RemoveOutdated(logName, maxLogFiles)
	if err != nil {
		return fmt.Errorf("remove outdated error: %w", err)
	}
	return nil
}
