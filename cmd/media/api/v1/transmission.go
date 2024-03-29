package v1

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/errors"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var transmissionStoped = false

func (p mediaAPI) StartTransmission() error {
	err, _, _ := utils.NewShell("sysctl -p").Run(1)
	if err != nil {
		return err
	}

	cmd := fmt.Sprintf("transmission-daemon --log-info --config-dir /etc/transmission-daemon --username %s --password %s --logfile /var/log/transmission/transmission.log",
		config.Default.Media.Transmission.User,
		config.Default.Media.Transmission.Password)
	err, out, _ := utils.NewShell(cmd).Run(1)
	log.Debugf("Transmission start out: %s", out)
	return err
}

func (p mediaAPI) StopTransmission() error {
	err, pid, _ := utils.NewShell("pidof transmission-daemon").Run(1)
	if err != nil {
		return err
	}

	err, _, _ = utils.NewShell(fmt.Sprintf("kill -9 %s", pid)).Run(1)
	if err != nil {
		return err
	}
	return nil
}

func (p mediaAPI) StartTransmissionJobs() error {
	err, out, _ := utils.NewShell("transmission-remote localhost -tall --start").Run(1)
	matched, _ := regexp.MatchString("success", out)
	if err != nil {
		return fmt.Errorf("start transmission jobs error: %w", err)
	}
	if !matched {
		return fmt.Errorf("start transmission jobs error: %s", out)
	}
	return nil
}

func (p mediaAPI) StopTransmissionJobs() error {
	err, out, _ := utils.NewShell("transmission-remote localhost -tall --stop").Run(1)
	if err != nil {
		return fmt.Errorf("stop transmission jobs error: %w", err)
	}
	matched, _ := regexp.MatchString("success", out)
	if !matched {
		return fmt.Errorf("stop transmission jobs error: %s", out)
	}
	return nil
}

func (p mediaAPI) HealthTransmission() error {
	if !transmissionStoped {
		err, _, _ := utils.NewShell("transmission-remote --authenv --session-info").Run(1)
		if err != nil {
			return fmt.Errorf("transmission health error: %v", err)
		}
	}
	return nil
}

func (p mediaAPI) TransmissionStart(c *gin.Context) {
	err := p.StartTransmission()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	transmissionStoped = false
	c.Status(http.StatusOK)
}

func (p mediaAPI) TransmissionStop(c *gin.Context) {
	err := p.StopTransmission()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	transmissionStoped = true
	c.Status(http.StatusOK)
}

func (p mediaAPI) TransmissionRestart(c *gin.Context) {
	err := p.StopTransmission()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	transmissionStoped = true

	err = p.StartTransmission()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	transmissionStoped = false
	c.Status(http.StatusOK)
}

func (p mediaAPI) TransmissionStartJobs(c *gin.Context) {
	err := p.StartTransmissionJobs()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mediaAPI) TransmissionStopJobs(c *gin.Context) {
	err := p.StopTransmissionJobs()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mediaAPI) TransmissionLogRotate() error {
	const (
		path        = "/var/log/transmission"
		fName       = "transmission.log"
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
