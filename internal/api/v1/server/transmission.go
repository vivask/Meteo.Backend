package server

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var transmissionStoped = false

func (p serverAPI) StartTransmission() error {
	cmd := fmt.Sprintf("transmission-daemon --log-info --config-dir /etc/transmission-daemon --username %s --password %s --logfile /var/log/transmission/transmission.log",
		config.Default.Server.Transmission.User,
		config.Default.Server.Transmission.Password)
	shell := utils.NewShell(cmd)
	err, out, _ := shell.Run(1)
	log.Debugf("Transmission start out: %s", out)
	return err
}

func (p serverAPI) StopTransmission() error {
	cmd := "pidof transmission-daemon"
	shell := utils.NewShell(cmd)
	err, pid, _ := shell.Run(1)
	if err != nil {
		return fmt.Errorf("%s error: %v", cmd, err)
	}
	cmd = fmt.Sprintf("kill -9 %s", pid)
	shell = utils.NewShell(cmd)
	err, _, _ = shell.Run(1)
	if err != nil {
		return fmt.Errorf("%s error: %v", cmd, err)
	}
	return nil
}

func (p serverAPI) StartTransmissionJobs() error {
	cmd := "transmission-remote localhost -tall --start"
	shell := utils.NewShell(cmd)
	err, out, _ := shell.Run(1)
	matched, _ := regexp.MatchString("success", out)
	if err != nil {
		return fmt.Errorf("start transmission jobs error: %w", err)
	}
	if !matched {
		return fmt.Errorf("start transmission jobs error: %s", out)
	}
	return nil
}

func (p serverAPI) StopTransmissionJobs() error {
	cmd := "transmission-remote localhost -tall --stop"
	shell := utils.NewShell(cmd)
	err, out, _ := shell.Run(1)
	if err != nil {
		return fmt.Errorf("stop transmission jobs error: %w", err)
	}
	matched, _ := regexp.MatchString("success", out)
	if !matched {
		return fmt.Errorf("stop transmission jobs error: %s", out)
	}
	return nil
}

func (p serverAPI) HealthTransmission() error {
	if !transmissionStoped {
		cmd := "transmission-remote --authenv --session-info"
		shell := utils.NewShell(cmd)
		err, _, _ := shell.Run(1)
		if err != nil {
			return fmt.Errorf("smbd health error: %v", err)
		}
	}
	return nil
}

func (p serverAPI) TransmissionStart(c *gin.Context) {
	err := p.StartTransmission()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	transmissionStoped = false
	c.Status(http.StatusOK)
}

func (p serverAPI) TransmissionStop(c *gin.Context) {
	err := p.StopTransmission()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	transmissionStoped = true
	c.Status(http.StatusOK)
}

func (p serverAPI) TransmissionRestart(c *gin.Context) {
	err := p.StopTransmission()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	transmissionStoped = true

	err = p.StartTransmission()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	transmissionStoped = false
	c.Status(http.StatusOK)
}

func (p serverAPI) TransmissionStartJobs(c *gin.Context) {
	err := p.StartTransmissionJobs()
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

func (p serverAPI) TransmissionStopJobs(c *gin.Context) {
	err := p.StopTransmissionJobs()
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
