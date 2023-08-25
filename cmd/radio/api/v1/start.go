package v1

import (
	"meteo/internal/log"
)

func (p radioAPI) StartService() {

	err := p.StartDump1090()
	if err != nil {
		log.Fatalf("Fail start dump1090: %v", err)
	} else {
		log.Info("Dump1090 start success")
	}
}
