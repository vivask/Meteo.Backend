package proxy

import "github.com/gin-gonic/gin"

func (p proxyAPI) RegisterAPIV1(router *gin.RouterGroup) {
	proxy := router.Group("/proxy")
	proxy.GET("/status/get", p.GetProxyesState)

	proxy.PUT("/master/server/start", p.SetMasterProxyStart)
	proxy.PUT("/master/server/stop", p.SetMasterProxyStop)
	proxy.PUT("/master/adblock/on", p.SetMasterProxyAdBlockOn)
	proxy.PUT("/master/adblock/off", p.SetMasterProxyAdBlockOff)
	proxy.PUT("/master/cache/on", p.SetMasterProxyCacheOn)
	proxy.PUT("/master/cache/off", p.SetMasterProxyCacheOff)
	proxy.PUT("/master/unlock/on", p.SetMasterProxyUnlockOn)
	proxy.PUT("/master/unlock/off", p.SetMasterProxyUnlockOff)

	proxy.PUT("/slave/server/start", p.SetSlaveProxyStart)
	proxy.PUT("/slave/server/stop", p.SetSlaveProxyStop)
	proxy.PUT("/slave/adblock/on", p.SetSlaveProxyAdBlockOn)
	proxy.PUT("/slave/adblock/off", p.SetSlaveProxyAdBlockOff)
	proxy.PUT("/slave/cache/on", p.SetSlaveProxyCacheOn)
	proxy.PUT("/slave/cache/off", p.SetSlaveProxyCacheOff)
	proxy.PUT("/slave/unlock/on", p.SetSlaveProxyUnlockOn)
	proxy.PUT("/slave/unlock/off", p.SetSlaveProxyUnlockOff)

	proxy.GET("/zones/get", p.GetAllZones)
	proxy.POST("/zones/add", p.AddZone)
	proxy.POST("/zones/edit", p.EditZone)
	proxy.DELETE("/zones/:id", p.DelZone)

	proxy.GET("/vpnlists/get", p.GetAccessLists)
	proxy.GET("/manualvpn/get", p.GetAllManualToVpn)
	proxy.POST("/manualvpn/add", p.AddManualToVpn)
	proxy.POST("/manualvpn/edit", p.EditManualToVpn)
	proxy.DELETE("/manualvpn/:id", p.DelManualFromVpn)

	proxy.GET("/autovpn/get", p.GetAllAutoToVpn)
	proxy.PUT("/autovpn/ignore/", p.IgnoreAutoToVpn)
	proxy.PUT("/autovpn/delete", p.DelAutoFromVpn)

	proxy.GET("/ignorevpn/get", p.GetAllIgnoreAutoToVpn)
	proxy.PUT("/ignorevpn/restore", p.RestoreAutoToVpn)
	proxy.PUT("/ignorevpn/delete", p.DelIgnoreAutoToVpn)
}
