package v1

import (
	"meteo/internal/log"
)

func (p radiusAPI) StartService() {

	err := p.StartRadius()
	if err != nil {
		log.Fatalf("Fail start radius: %v", err)
	} else {
		log.Info("Radius start success")
	}

}
