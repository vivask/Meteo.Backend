package schedule

import (
	repo "meteo/internal/repo/schedule"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ScheduleAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
	GetAllTasks(c *gin.Context)
	AddTask(c *gin.Context)
	EditTask(c *gin.Context)
	DelTask(c *gin.Context)
	GetAllJobs(c *gin.Context)
	AddJob(c *gin.Context)
	GetAllPeriods(c *gin.Context)
	GetAllExecutors(c *gin.Context)
	GetCronJobs(c *gin.Context)
	ActivateJob(c *gin.Context)
	DeactivateJob(c *gin.Context)
	RunJob(c *gin.Context)
}

type scheduleAPI struct {
	repo repo.ScheduleService
}

// NewScheduleAPI get product service instance
func NewScheduleAPI(db *gorm.DB) ScheduleAPI {
	return &scheduleAPI{repo: repo.NewScheduleService(db)}
}
