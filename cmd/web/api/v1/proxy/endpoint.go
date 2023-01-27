package proxy

import "github.com/gin-gonic/gin"

func (p proxyAPI) RegisterAPIV1(router *gin.RouterGroup) {
	proxy := router.Group("/proxy")
	proxy.GET("/status", p.GetProxyesState)

	proxy.PUT("/main/server/start", p.SetMainProxyStart)
	proxy.PUT("/main/server/stop", p.SetMainProxyStop)
	proxy.PUT("/main/adblock/on", p.SetMainProxyAdBlockOn)
	proxy.PUT("/main/adblock/off", p.SetMainProxyAdBlockOff)
	proxy.PUT("/main/cache/on", p.SetMainProxyCacheOn)
	proxy.PUT("/main/cache/off", p.SetMainProxyCacheOff)
	proxy.PUT("/main/unlock/on", p.SetMainProxyUnlockOn)
	proxy.PUT("/main/unlock/off", p.SetMainProxyUnlockOff)

	proxy.PUT("/backup/server/start", p.SetBackupProxyStart)
	proxy.PUT("/backup/server/stop", p.SetBackupProxyStop)
	proxy.PUT("/backup/adblock/on", p.SetBackupProxyAdBlockOn)
	proxy.PUT("/backup/adblock/off", p.SetBackupProxyAdBlockOff)
	proxy.PUT("/backup/cache/on", p.SetBackupProxyCacheOn)
	proxy.PUT("/backup/cache/off", p.SetBackupProxyCacheOff)
	proxy.PUT("/backup/unlock/on", p.SetBackupProxyUnlockOn)
	proxy.PUT("/backup/unlock/off", p.SetBackupProxyUnlockOff)

	proxy.GET("/zones", p.GetAllZones)
	proxy.PUT("/zones", p.AddZone)
	proxy.POST("/zones", p.EditZone)
	proxy.DELETE("/zones/:id", p.DelZone)

	proxy.GET("/vpnlists", p.GetAccessLists)
	proxy.GET("/manualvpn", p.GetAllManualToVpn)
	proxy.PUT("/manualvpn", p.AddManualToVpn)
	proxy.POST("/manualvpn", p.EditManualToVpn)
	proxy.DELETE("/manualvpn/:id", p.DelManualFromVpn)

	proxy.GET("/autovpn", p.GetAllAutoToVpn)
	proxy.PUT("/autovpn", p.IgnoreAutoToVpn)
	proxy.POST("/autovpn/delete", p.DelAutoFromVpn)

	proxy.GET("/ignorevpn", p.GetAllIgnoreAutoToVpn)
	proxy.PUT("/ignorevpn", p.RestoreAutoToVpn)
	proxy.POST("/ignorevpn/delete", p.DelIgnoreAutoToVpn)
}
