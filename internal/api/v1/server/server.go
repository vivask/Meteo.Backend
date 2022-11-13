package server

import (
	repo "meteo/internal/repo/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ServerAPI api interface
type ServerAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
	StartRadius() error
	StopRadius() error
	StartSamba() error
	StopSamba() error
	StartTransmission() error
	StopTransmission() error
	MountStorage() error
	UmountStorage() error
}

type serverAPI struct {
	repo repo.ServerService
}

// NewServerAPI get server service instance
func NewServerAPI(db *gorm.DB) ServerAPI {
	return &serverAPI{repo: repo.NewServerService(db)}
}
