package proxy

import "github.com/gin-gonic/gin"

func (p proxyAPI) RegisterAPIV1(router *gin.RouterGroup) {
	proxy := router.Group("/proxy")
	proxy.GET("/status/get", p.GetProxyesState)

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
