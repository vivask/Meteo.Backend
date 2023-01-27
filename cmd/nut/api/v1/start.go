package v1

import (
	"meteo/internal/config"
	"meteo/internal/log"
)

func (p nutAPI) StartService() {
	if config.Default.App.Server == "main" {
		err := p.StartUpsDriver()
		if err != nil {
			log.Fatalf("Fail start upsdriver: %v", err)
		} else {
			log.Info("Upsdriver start success")
		}
	}
}
