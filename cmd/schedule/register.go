package main

import (
	api "meteo/cmd/schedule/api/v1"
	"meteo/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAPIV1(router *gin.RouterGroup, db *gorm.DB) {
	scheduleAPI := api.NewScheduleAPI(db)
	if config.Default.Schedule.Active {
		scheduleAPI.StartCron()
	}
	scheduleAPI.RegisterAPIV1(router)
}
