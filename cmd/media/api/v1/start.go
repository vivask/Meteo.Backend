package v1

import (
	"meteo/internal/config"
	"meteo/internal/log"
)

func (p mediaAPI) StartService() {
	if config.Default.App.Server == "main" {
		err := p.MountStorage()
		if err != nil {
			log.Fatalf("Fail mount storage: %v", err)
		} else {
			log.Info("Storage mount success")
		}

		err = p.StartSamba()
		if err != nil {
			log.Fatalf("Fail start samba: %v", err)
		} else {
			log.Info("Samba start success")
		}

		err = p.StartTransmission()
		if err != nil {
			log.Errorf("Fail start transmission: %v", err)
		} else {
			log.Info("Transmission start success")
		}
	}
}
