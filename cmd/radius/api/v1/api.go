package v1

import (
	repo "meteo/internal/repo/radius"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RadiusAPI api interface
type RadiusAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
	StartService()
}

type radiusAPI struct {
	repo repo.RadiusService
}

// NewRadiusAPI get server service instance
func NewRadiusAPI(db *gorm.DB) RadiusAPI {
	return &radiusAPI{repo: repo.NewRadiusService(db)}
}
