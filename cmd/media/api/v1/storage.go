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

const RETRY = 3

var count = 0
var umounted = false

func (p mediaAPI) MountStorage() error {
	cmd := fmt.Sprintf("mount %s %s", config.Default.Media.Storage.Device, config.Default.Media.Storage.MountPoint)
	err, _, _ := utils.NewShell(cmd).Run(60)
	if err != nil {
		matched, _ := regexp.MatchString("exit status 255", err.Error())
		if !matched {
			return fmt.Errorf("mount storage error: %w", err)
		}
	}
	cmd = fmt.Sprintf("chown -R smbuser:users %s", config.Default.Media.Storage.MountPoint)
	err, _, _ = utils.NewShell(cmd).Run(20)
	if err != nil {
		return fmt.Errorf("chown storage error: %w", err)
	}
	return nil
}

func (p mediaAPI) UmountStorage() error {
	cmd := fmt.Sprintf("umount %s", config.Default.Media.Storage.MountPoint)
	err, _, _ := utils.NewShell(cmd).Run(20)
	if err != nil {
		return fmt.Errorf("umount storage error: %w", err)
	}
	return nil
}

func (p mediaAPI) IsMounted() bool {
	if count > RETRY {
		count = 0
		return false
	} else {
		cmd := fmt.Sprintf("df -P | grep '%s' | grep '%s' || echo 1", config.Default.Media.Storage.Device, config.Default.Media.Storage.MountPoint)
		err, out, _ := utils.NewShell(cmd).Run(1)
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

func (p mediaAPI) HealthStorage() error {
	if !umounted {
		if !p.IsMounted() {
			return fmt.Errorf("storage unavailable")
		}
	}
	return nil
}

func (p mediaAPI) StorageMount(c *gin.Context) {
	err := p.MountStorage()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	umounted = false
	c.Status(http.StatusOK)
}

func (p mediaAPI) StorageUmount(c *gin.Context) {
	err := p.UmountStorage()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	umounted = true
	c.Status(http.StatusOK)
}

func (p mediaAPI) StorageRemount(c *gin.Context) {
	err := p.UmountStorage()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	umounted = true

	err = p.MountStorage()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	umounted = false
	c.Status(http.StatusOK)
}
