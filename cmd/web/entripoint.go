package main

import (
	web "meteo/cmd/web/api/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var webAPI web.WebAPI

func RegisterPublicAPIV1(public *gin.RouterGroup, db *gorm.DB) {
	webAPI = web.NewWebAPI(db)
	webAPI.RegisterPublicAPIV1(public)
}

func RegisterProtectedAPIV1(protected *gin.RouterGroup, db *gorm.DB) {
	webAPI.RegisterProtectedAPIV1(protected)
}
