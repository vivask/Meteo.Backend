package mikrotik

import "github.com/gin-gonic/gin"

func (p mikrotikAPI) RegisterAPIV1(router *gin.RouterGroup) {
	mikrotik := router.Group("/mikrotik")
	mikrotik.PUT("/ppp/check", p.CheckBYFLY)
	mikrotik.POST("/router/zones/update", p.RouterSyncDNS)
	mikrotik.PUT("/tovpn/manual/add", p.AddManualHostToVpn)
	mikrotik.PUT("/tovpn/manual/del", p.RemoveManualHostFromVpn)
	mikrotik.PUT("/tovpn/auto/add", p.AddAutoHostToVpn)
	mikrotik.PUT("/tovpn/auto/del", p.RemoveAutoHostFromVpn)
}
