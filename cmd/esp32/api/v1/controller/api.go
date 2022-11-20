package controller

import (
	repo "meteo/internal/repo/esp32"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Esp32API api controller of produces
type Esp32API interface {
	RegisterAPIV1(router *gin.RouterGroup) *gin.RouterGroup
	Middleware() gin.HandlerFunc
	Handler(*gin.Context)
	UploadFirmware(c *gin.Context)
}

type esp32API struct {
	repo repo.Esp32Service
}

// NewEsp32API get product service instance
func NewEsp32API(db *gorm.DB) Esp32API {
	return &esp32API{repo: repo.NewEsp32Service(db)}
}

func (p esp32API) Middleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
	})
}
