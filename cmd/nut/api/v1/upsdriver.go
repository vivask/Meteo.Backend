package v1

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/log"
	"meteo/internal/utils"
	"regexp"
)

var nutStopped = false

func (p nutAPI) StartUpsDriver() error {
	cmd := "/usr/sbin/upsdrvctl -u root start"
	shell := utils.NewShell(cmd)
	err, out, _ := shell.Run(1)
	if err != nil {
		return fmt.Errorf("upsdrive start error: %w", err)
	}
	log.Debugf("Upsdriver start out: %s", out)

	cmd = "/usr/sbin/upsd -u $USER"
	shell = utils.NewShell(cmd)
	err, out, _ = shell.Run(1)
	if err != nil {
		return fmt.Errorf("upsd start error: %w", err)
	}
	log.Debugf("Upsd start out: %s", out)

	cmd = "/usr/sbin/upsmon"
	shell = utils.NewShell(cmd)
	err, out, _ = shell.Run(1)
	if err != nil {
		return fmt.Errorf("upsmon start error: %w", err)
	}
	log.Debugf("Upsmon start out: %s", out)

	return nil
}

func (p nutAPI) HealthUpsDriver() error {
	if !nutStopped {
		cmd := fmt.Sprintf("upsc %s@localhost:%d", config.Default.Nut.Driver, config.Default.Nut.Port)
		shell := utils.NewShell(cmd)
		err, out, _ := shell.Run(1)
		if err != nil {
			return fmt.Errorf("upsdriver health error: %w", err)
		}
		matched, _ := regexp.MatchString("ups.test.result: Done and passed", out)
		if !matched {
			return fmt.Errorf("upsdriver not healthy: %s", out)
		}
	}
	return nil
}
