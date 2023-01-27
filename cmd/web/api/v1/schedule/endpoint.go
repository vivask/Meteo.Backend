package schedule

import "github.com/gin-gonic/gin"

func (p scheduleAPI) RegisterAPIV1(router *gin.RouterGroup) {
	schedule := router.Group("/schedule")
	schedule.GET("/tasks", p.GetAllTasks)
	schedule.PUT("/tasks", p.AddTask)
	schedule.POST("/tasks", p.EditTask)
	schedule.DELETE("/tasks/:id", p.DelTask)
	schedule.PUT("/task/run", p.RunTask)

	schedule.GET("/jobs", p.GetAllJobs)
	schedule.PUT("/jobs", p.AddJob)
	schedule.POST("/jobs", p.EditJob)
	schedule.DELETE("/jobs/:id", p.DelJob)
	schedule.GET("/periods", p.GetAllPeriods)
	schedule.GET("/executors", p.GetAllExecutors)

	schedule.GET("/cron", p.GetCronJobs)

	schedule.PUT("/job/activate/:id", p.ActivateJob)
	schedule.PUT("/job/deactivate/:id", p.DeactivateJob)
	schedule.PUT("/job/run/:id", p.RunJob)
}
