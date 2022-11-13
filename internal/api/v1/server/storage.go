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

const RETRY = 3

var count = 0
var umounted = false

func (p serverAPI) MountStorage() error {
	cmd := fmt.Sprintf("mount %s %s", config.Default.Server.Storage.Device, config.Default.Server.Storage.MountPoint)
	shell := utils.NewShell(cmd)
	err, _, _ := shell.Run(20)
	if err != nil {
		matched, _ := regexp.MatchString("exit status 255", err.Error())
		if !matched {
			return fmt.Errorf("mount storage error: %w", err)
		}
	}
	return nil
}

func (p serverAPI) UmountStorage() error {
	cmd := fmt.Sprintf("umount %s", config.Default.Server.Storage.MountPoint)
	shell := utils.NewShell(cmd)
	err, _, _ := shell.Run(20)
	if err != nil {
		return fmt.Errorf("umount storage error: %w", err)
	}
	return nil
}

func (p serverAPI) IsMounted() bool {
	if count > RETRY {
		count = 0
		return false
	} else {
		cmd := fmt.Sprintf("df -P | grep '%s' | grep '%s' || echo 1", config.Default.Server.Storage.Device, config.Default.Server.Storage.MountPoint)
		shell := utils.NewShell(cmd)
		err, out, _ := shell.Run(1)
		if err != nil {
			log.Errorf("mount storage error: %v", err)
			return false
		}
		if out == "1" {
			p.MountStorage()
			count++
			return p.IsMounted()
		}
	}
	return true
}

func (p serverAPI) HealthStorage() error {
	if !umounted {
		if !p.IsMounted() {
			return fmt.Errorf("storage unavailable")
		}
	}
	return nil
}

func (p serverAPI) StorageMount(c *gin.Context) {
	err := p.MountStorage()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	umounted = false
	c.Status(http.StatusOK)
}

func (p serverAPI) StorageUmount(c *gin.Context) {
	err := p.UmountStorage()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	umounted = true
	c.Status(http.StatusOK)
}

func (p serverAPI) StorageRemount(c *gin.Context) {
	err := p.UmountStorage()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	umounted = true

	err = p.MountStorage()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	umounted = false
	c.Status(http.StatusOK)
}
