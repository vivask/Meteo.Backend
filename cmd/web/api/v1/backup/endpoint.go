package backup

import "github.com/gin-gonic/gin"

func (p backupAPI) RegisterAPIV1(router *gin.RouterGroup) {
	backup := router.Group("/backup")
	backup.GET("/state", p.GetServicesState)
	backup.PUT("/kodi/restart", p.RestarKodi)
	backup.PUT("/storage/restart", p.RestarStorageKodi)
	backup.PUT("/storage/stop", p.StopStorageKodi)
	backup.PUT("/storage/start", p.StartStorageKodi)
	backup.PUT("/restart/:id", p.RestartServerCont)
	backup.PUT("/stop/:id", p.StopServerCont)
	backup.PUT("/start/:id", p.StartServerCont)
	backup.PUT("/reboot", p.Reboot)
	backup.PUT("/shutdown", p.Shutdown)
	backup.GET("/logging/:id", p.GetLogging)
	backup.PUT("/logging/:id", p.ClearLogging)
}
