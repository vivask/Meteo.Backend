package v1

import (
	"fmt"
	"meteo/internal/errors"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var sambaStopped = false

func (p mediaAPI) StartSamba() error {
	cmd := "nmbd --configfile=/etc/samba/smb.conf"
	shell := utils.NewShell(cmd)
	err, out, _ := shell.Run(1)
	if err != nil {
		return fmt.Errorf("nmbd start error: %w", err)
	}
	log.Debugf("Smbd start out: %s", out)
	cmd = "smbd --configfile=/etc/samba/smb.conf --no-process-group"
	shell = utils.NewShell(cmd)
	err, out, _ = shell.Run(1)
	if err != nil {
		return fmt.Errorf("smbd start error: %w", err)
	}
	log.Debugf("Smbd start out: %s", out)

	return nil
}

func (p mediaAPI) StopSamba() error {
	cmd := "pidof smbd"
	shell := utils.NewShell(cmd)
	err, pid_smbd, _ := shell.Run(1)
	if err != nil {
		return fmt.Errorf("%s error: %w", cmd, err)
	}
	cmd = "pidof nmbd"
	shell = utils.NewShell(cmd)
	err, pid_nmbd, _ := shell.Run(1)
	if err != nil {
		return fmt.Errorf("%s error: %w", cmd, err)
	}
	cmd = fmt.Sprintf("kill -9 %s %s", pid_nmbd, pid_smbd)
	shell = utils.NewShell(cmd)
	err, _, _ = shell.Run(1)
	if err != nil {
		return fmt.Errorf("%s error: %w", cmd, err)
	}
	return nil
}

func (p mediaAPI) HealthSamba() error {
	if !sambaStopped {
		cmd := "smbclient -L '\\localhost\\' -U 'guest%' -m SMB3"
		shell := utils.NewShell(cmd)
		err, _, _ := shell.Run(1)
		if err != nil {
			return fmt.Errorf("smbd health error: %w", err)
		}
	}
	return nil
}

func (p mediaAPI) SambaStart(c *gin.Context) {
	err := p.StartSamba()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	sambaStopped = false
	c.Status(http.StatusOK)
}

func (p mediaAPI) SambaStop(c *gin.Context) {
	err := p.StopSamba()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	sambaStopped = true
	c.Status(http.StatusOK)
}

func (p mediaAPI) SambaRestart(c *gin.Context) {
	err := p.StopSamba()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	sambaStopped = true

	err = p.StartSamba()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	sambaStopped = false
	c.Status(http.StatusOK)
}

func (p mediaAPI) SambaLogRotate() error {
	const (
		path        = "/var/log/samba"
		nmbd        = "log.nmbd"
		smbd        = "log.smbd"
		maxLength   = 1024 * 1024
		maxLogFiles = 4
	)

	logName := fmt.Sprintf("%s/%s", path, smbd)

	err := log.Rotate(logName, maxLength)
	if err != nil {
		return fmt.Errorf("rotate [%s] error: %w", logName, err)
	}

	err = log.RemoveOutdated(logName, maxLogFiles)
	if err != nil {
		return fmt.Errorf("remove outdated error: %w", err)
	}

	logName = fmt.Sprintf("%s/%s", path, nmbd)

	err = log.Rotate(logName, maxLength)
	if err != nil {
		return fmt.Errorf("rotate [%s] error: %w", logName, err)
	}

	err = log.RemoveOutdated(logName, maxLogFiles)
	if err != nil {
		return fmt.Errorf("remove outdated error: %w", err)
	}

	return nil
}
