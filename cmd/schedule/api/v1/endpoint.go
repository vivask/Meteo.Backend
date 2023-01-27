package v1

import "github.com/gin-gonic/gin"

func (p scheduleAPI) RegisterAPIV1(router *gin.RouterGroup) {
	schedule := router.Group("/schedule")
	schedule.PUT("/jobs/reload", p.JobsReload)
	schedule.PUT("/job/run/:id", p.JobRun)
	schedule.GET("/cron/get", p.GetCronJobs)
	schedule.GET("/health", p.GetHealth)
	schedule.GET("/logging", p.GetLogging)
	schedule.PUT("/logging", p.ClearLogging)
}
