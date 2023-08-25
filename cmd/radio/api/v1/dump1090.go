package v1

import (
	"fmt"
	"meteo/internal/log"
	"meteo/internal/utils"
)

func (p radioAPI) StartDump1090() error {

	err, out, _ := utils.NewShell("dump1090 --quiet --net --adaptive-range --no-fix --write-json /var/log/backend").Run(3)
	if err != nil && err.Error() != "command timed out" {
		log.Infof("Dump1090 start out: %s", out)
		return fmt.Errorf("dump1090 start error: %w", err)
	}
	log.Debugf("Dump1090 start out: %s", out)

	return nil
}
