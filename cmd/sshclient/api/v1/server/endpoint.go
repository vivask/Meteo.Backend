package server

import "github.com/gin-gonic/gin"

func (p serverAPI) RegisterAPIV1(router *gin.RouterGroup) {
	server := router.Group("/server")
	server.PUT("/backup/kodi/restart", p.RestarKodi)
	server.PUT("/backup/kodi/storage/restart", p.RestarStorageKodi)
	server.PUT("/backup/kodi/storage/stop", p.StopStorageKodi)
	server.PUT("/backup/kodi/storage/start", p.StartStorageKodi)
	server.PUT("/backup/restart/:id", p.RestartBackupCont)
	server.PUT("/backup/stop/:id", p.StopBackupCont)
	server.PUT("/backup/start/:id", p.StartBackupCont)
	server.PUT("/backup/reboot", p.BackupReboot)
	server.PUT("/backup/shutdown", p.BackupShutdown)

	server.PUT("/main/restart/:id", p.RestartMainCont)
	server.PUT("/main/stop/:id", p.StopMainCont)
	server.PUT("/main/start/:id", p.StartMainCont)
	server.PUT("/main/reboot", p.MainReboot)
	server.PUT("/main/shutdown", p.MainShutdown)

}
