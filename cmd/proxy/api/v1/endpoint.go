package v1

import "github.com/gin-gonic/gin"

func (p proxyAPI) RegisterAPIV1(router *gin.RouterGroup) {
	proxy := router.Group("/proxy")
	proxy.Use(p.Middleware())
	proxy.GET("/status/get", p.GetState)
	proxy.PUT("/server/start", p.SetStart)
	proxy.PUT("/server/stop", p.SetStop)
	proxy.PUT("/adblock/on", p.SetAdBlockOn)
	proxy.PUT("/adblock/off", p.SetAdBlockOff)
	proxy.PUT("/cache/on", p.SetCacheOn)
	proxy.PUT("/cache/off", p.SetCacheOff)
	proxy.PUT("/unlock/on", p.SetUnlockOn)
	proxy.PUT("/unlock/off", p.SetUnlockOff)
	proxy.PUT("/adblock/load", p.AdBlockLoad)
	proxy.POST("/adblock/load", p.AdBlockLoad)
	proxy.POST("/adblock/update", p.AdBlockUpdate)
	proxy.PUT("/zones/update", p.ReloadZones)
	proxy.PUT("/unlocker/reload", p.ReloadUnlocker)
	proxy.GET("/health", p.GetHealth)
	proxy.GET("/logging", p.GetLogging)
	proxy.PUT("/logging", p.ClearLogging)
}
