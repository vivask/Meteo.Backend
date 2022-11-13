package main

import (
	api "meteo/internal/api/v1/server"
	"meteo/internal/config"
	"meteo/internal/log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAPIV1(router *gin.RouterGroup, db *gorm.DB) {
	serverAPI := api.NewServerAPI(db)
	serverAPI.RegisterAPIV1(router)

	err := serverAPI.StartRadius()
	if err != nil {
		log.Errorf("Fail start radius: %v", err)
	} else {
		log.Info("Radius start success")
	}

	if config.Default.App.Server == "main" {
		err := serverAPI.MountStorage()
		if err != nil {
			log.Errorf("Fail mount storage: %v", err)
		} else {
			log.Info("Storage mount success")
		}

		err = serverAPI.StartSamba()
		if err != nil {
			log.Errorf("Fail start samba: %v", err)
		} else {
			log.Info("Samba start success")
		}

		err = serverAPI.StartTransmission()
		if err != nil {
			log.Errorf("Fail start transmission: %v", err)
		} else {
			log.Info("Transmission start success")
		}
	}
}
