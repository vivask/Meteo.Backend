package server

import "github.com/gin-gonic/gin"

func (p serverAPI) RegisterMainAPIV1(router *gin.RouterGroup) {
	main := router.Group("/main")
	main.GET("/state/get", p.GetMainServicesState)
	main.PUT("/transmission/jobs/start", p.TransmissionStartJobs)
	main.PUT("/transmission/jobs/stop", p.TransmissionStopJobs)
	main.PUT("/transmission/start", p.TransmissionStart)
	main.PUT("/transmission/stop", p.TransmissionStop)
	main.PUT("/transmission/restart", p.TransmissionRestart)
	main.PUT("/samba/start", p.SambaStart)
	main.PUT("/samba/stop", p.SambaStop)
	main.PUT("/samba/restart", p.SambaRestart)
	main.PUT("/storage/mount", p.MainStorageMount)
	main.PUT("/storage/umount", p.MainStorageUmount)
	main.PUT("/storage/remount", p.MainStorageRemount)
	main.PUT("/restart/:id", p.RestartMainServerCont)
	main.PUT("/stop/:id", p.StopMainServerCont)
	main.PUT("/start/:id", p.StartMainServerCont)
	main.PUT("/reboot", p.MainReboot)
	main.PUT("/shutdown", p.MainShutdown)
}

func (p serverAPI) RegisterBackupAPIV1(router *gin.RouterGroup) {
	backup := router.Group("/backup")
	backup.GET("/state/get", p.GetBackupServicesState)
	backup.PUT("/kodi/restart", p.RestarKodi)
	backup.PUT("/storage/restart", p.RestarStorageKodi)
	backup.PUT("/storage/stop", p.StopStorageKodi)
	backup.PUT("/storage/start", p.StartStorageKodi)
	backup.PUT("/restart/:id", p.RestartBackupServerCont)
	backup.PUT("/stop/:id", p.StopBackupServerCont)
	backup.PUT("/start/:id", p.StartBackupServerCont)
	backup.PUT("/reboot", p.BackupReboot)
	backup.PUT("/shutdown", p.BackupShutdown)
}
