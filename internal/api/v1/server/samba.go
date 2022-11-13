package server

import (
	"fmt"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var sambaStopped = false

func (p serverAPI) StartSamba() error {
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

func (p serverAPI) StopSamba() error {
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

func (p serverAPI) HealthSamba() error {
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

func (p serverAPI) SambaStart(c *gin.Context) {
	err := p.StartSamba()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	sambaStopped = false
	c.Status(http.StatusOK)
}

func (p serverAPI) SambaStop(c *gin.Context) {
	err := p.StopSamba()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	sambaStopped = true
	c.Status(http.StatusOK)
}

func (p serverAPI) SambaRestart(c *gin.Context) {
	err := p.StopSamba()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	sambaStopped = true

	err = p.StartSamba()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	sambaStopped = false
	c.Status(http.StatusOK)
}
