package server

import (
	"github.com/gin-gonic/gin"
)

func (p serverAPI) RegisterAPIV1(router *gin.RouterGroup) {
	server := router.Group("server")
	server.PUT("/radius/stop", p.RadiusStart)
	server.PUT("/radius/start", p.RadiusStop)
	server.PUT("/samba/start", p.SambaStart)
	server.PUT("/samba/stop", p.SambaStop)
	server.PUT("/samba/restart", p.SambaRestart)
	server.PUT("/transmission/start", p.TransmissionStart)
	server.PUT("/transmission/stop", p.TransmissionStop)
	server.PUT("/transmission/restart", p.TransmissionRestart)
	server.PUT("/transmission/jobs/start", p.TransmissionStartJobs)
	server.PUT("/transmission/jobs/stop", p.TransmissionStopJobs)
	server.PUT("/storage/mount", p.StorageMount)
	server.PUT("/storage/umount", p.StorageUmount)
	server.PUT("/storage/remount", p.StorageRemount)
	server.GET("/health", p.GetHealth)
}
