package v1

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/utils"
	"regexp"
)

var nutStopped = false

func (p nutAPI) StartUpsDriver() error {
	/*const retry = 5
	var out string
	var err error
	var try int = 0

	for {
		err, out, _ = utils.NewShell("/usr/sbin/upsdrvctl -u root start").Run(5)
		if err == nil {
			break
		} else if try >= retry {
			return fmt.Errorf("upsdrive start error: %w, OUT: %s", err, out)
		}
		try++
		utils.NewShell("/usr/sbin/upsdrvctl -u root shutdown").Run(5)
	}
	log.Debugf("Upsdriver start out: %s", out)

	err, out, _ = utils.NewShell("/usr/sbin/upsd -u $USER").Run(1)
	if err != nil {
		return fmt.Errorf("upsd start error: %w, OUT: %s", err, out)
	}
	log.Debugf("Upsd start out: %s", out)

	err, out, _ = utils.NewShell("/usr/sbin/upsmon").Run(1)
	if err != nil {
		return fmt.Errorf("upsmon start error: %w, OUT: %s", err, out)
	}
	log.Debugf("Upsmon start out: %s", out)*/

	return nil
}

func (p nutAPI) HealthUpsDriver() error {
	if !nutStopped {
		cmd := fmt.Sprintf("upsc %s@localhost:%d", config.Default.Nut.Driver, config.Default.Nut.Port)
		err, out, _ := utils.NewShell(cmd).Run(1)
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
