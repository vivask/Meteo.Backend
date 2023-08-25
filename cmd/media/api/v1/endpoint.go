package v1

import "github.com/gin-gonic/gin"

func (p mediaAPI) RegisterAPIV1(router *gin.RouterGroup) {
	media := router.Group("/media")
	media.GET("/health", p.GetHealth)
	media.GET("/health/samba", p.GetHealthSamba)
	media.GET("/health/storage", p.GetHealthStore)
	media.GET("/health/transmission", p.GetHealthTransmission)
	media.GET("/logging", p.GetLogging)
	media.PUT("/logging", p.ClearLogging)
	media.GET("/logging/empty", p.IsEmptyLog)
	media.PUT("/samba/start", p.SambaStart)
	media.PUT("/samba/stop", p.SambaStop)
	media.PUT("/samba/restart", p.SambaRestart)
	media.PUT("/transmission/start", p.TransmissionStart)
	media.PUT("/transmission/stop", p.TransmissionStop)
	media.PUT("/transmission/restart", p.TransmissionRestart)
	media.PUT("/transmission/jobs/start", p.TransmissionStartJobs)
	media.PUT("/transmission/jobs/stop", p.TransmissionStopJobs)
	media.PUT("/storage/mount", p.StorageMount)
	media.PUT("/storage/umount", p.StorageUmount)
	media.PUT("/storage/remount", p.StorageRemount)
	media.PUT("/rotate", p.RotateMediaLogs)
}
