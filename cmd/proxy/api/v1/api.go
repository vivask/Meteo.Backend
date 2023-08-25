package v1

import (
	"meteo/cmd/proxy/api/v1/tools"
	repo "meteo/internal/repo/proxy"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProxyAPI api interface
type ProxyAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
	Start() error
	//Middleware() gin.HandlerFunc
	GetState(c *gin.Context)
	SetStart(c *gin.Context)
	SetStop(c *gin.Context)
	SetAdBlockOn(c *gin.Context)
	SetAdBlockOff(c *gin.Context)
	SetCacheOn(c *gin.Context)
	SetCacheOff(c *gin.Context)
	SetUnlockOn(c *gin.Context)
	SetUnlockOff(c *gin.Context)
}

type proxyAPI struct {
	repo repo.ProxyService
	dns  *tools.Server
}

// NewProxyAPI get proxy service instance
func NewProxyAPI(db *gorm.DB) ProxyAPI {
	return &proxyAPI{
		repo: repo.NewProxyService(db),
		dns:  tools.NewServer(),
	}
}

/*func (p proxyAPI) Middleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
	})
}*/
