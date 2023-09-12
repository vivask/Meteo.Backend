package radius

import (
	repo "meteo/internal/repo/radius"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RadiusAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
}

type radiusAPI struct {
	repo repo.RadiusService
}

// NewRadiusAPI get radius service instance
func NewRadiusAPI(db *gorm.DB) RadiusAPI {
	return &radiusAPI{repo: repo.NewRadiusService(db)}
}
