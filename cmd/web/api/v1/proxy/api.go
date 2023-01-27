package proxy

import (
	repo "meteo/internal/repo/proxy"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProxyAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
}

type proxyAPI struct {
	repo repo.ProxyService
}

// NewProxyAPI get product service instance
func NewProxyAPI(db *gorm.DB) ProxyAPI {
	return &proxyAPI{repo: repo.NewProxyService(db)}
}
