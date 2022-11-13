package schedule

import "github.com/gin-gonic/gin"

func (p scheduleAPI) RegisterAPIV1(router *gin.RouterGroup) {
	schedule := router.Group("/schedule")
	schedule.GET("/tasks/get", p.GetAllTasks)
	schedule.POST("/task/add", p.AddTask)
	schedule.POST("/task/edit", p.EditTask)
	schedule.DELETE("/task/:id", p.DelTask)

	schedule.GET("/jobs/get", p.GetAllJobs)
	schedule.POST("/job/add", p.AddJob)
	schedule.POST("/job/edit", p.EditJob)
	schedule.DELETE("/job/:id", p.DelJob)
	schedule.GET("/periods/get", p.GetAllPeriods)
	schedule.GET("/days/get", p.GetAllDays)
	schedule.GET("/executors/get", p.GetAllExecutors)

	schedule.GET("/cron/get", p.GetCronJobs)

	schedule.PUT("/job/activate/:id", p.ActivateJob)
	schedule.PUT("/job/deactivate/:id", p.DeactivateJob)
	schedule.PUT("/job/run/:id", p.RunJob)

}
