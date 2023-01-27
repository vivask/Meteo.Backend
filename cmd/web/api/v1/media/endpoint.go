package media

import "github.com/gin-gonic/gin"

func (p mediaAPI) RegisterAPIV1(router *gin.RouterGroup) {
	media := router.Group("/media")
	media.PUT("/transmission/jobs/start", p.TransmissionStartJobs)
	media.PUT("/transmission/jobs/stop", p.TransmissionStopJobs)
	media.PUT("/transmission/start", p.TransmissionStart)
	media.PUT("/transmission/stop", p.TransmissionStop)
	media.PUT("/transmission/restart", p.TransmissionRestart)
	media.PUT("/samba/start", p.SambaStart)
	media.PUT("/samba/stop", p.SambaStop)
	media.PUT("/samba/restart", p.SambaRestart)
	media.PUT("/storage/mount", p.StorageMount)
	media.PUT("/storage/umount", p.StorageUmount)
	media.PUT("/storage/remount", p.StorageRemount)
}
