package v1

import (
	repo "meteo/internal/repo/proxy"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProxyAPI api controller of produces
type ProxyAPI interface {
	AddManualToVpn(*gin.Context)
	/*EditManualToVpn(*gin.Context)
	DelManualToVpn(*gin.Context)
	GetManualToVpn(*gin.Context)
	ChangeMasterState(*gin.Context)
	ChangeSlaveState(*gin.Context)
	ChangeMasterCache(*gin.Context)
	ChangeSlaveCache(*gin.Context)
	ChangeMasterBlk(*gin.Context)
	ChangeSlaveBlk(*gin.Context)
	ChangeMasterUnlocker(*gin.Context)
	ChangeSlaveUnlocker(*gin.Context)
	DelAutoToVpn(*gin.Context)
	AddIgnoreToVpn(*gin.Context)
	DelIgnoreToVpn(*gin.Context)
	RestoreIgnoreToVpn(*gin.Context)
	GetState(*gin.Context)
	SyncProxyZones(*gin.Context)
	ReloadProxyZones(*gin.Context)
	ChangeState(*gin.Context)
	ChangeCache(*gin.Context)
	ChangeBlk(*gin.Context)
	ChangeUnlocker(*gin.Context)
	BlklistUpdate(*gin.Context)
	ArpPing(*gin.Context)
	AddHomeZoneHost(*gin.Context)
	EditHomeZoneHost(*gin.Context)
	DelHomeZoneHost(*gin.Context)
	AddHost(*gin.Context)
	EditHost(*gin.Context)
	DelHost(*gin.Context)*/
}

type proxyAPI struct {
	repo repo.ProxyService
}

// NewProxyAPI get product service instance
func NewProxyAPI(db *gorm.DB) ProxyAPI {
	return &proxyAPI{repo: repo.NewProxyService(db)}
}

func (p proxyAPI) AddManualToVpn(ctx *gin.Context) {

	ctx.Status(http.StatusAccepted)
}
