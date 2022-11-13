package v1

import (
	repo "meteo/internal/repo/schedule"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
)

// ScheduleAPI api interface
type ScheduleAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
	StartCron()
	StopCron()
	JobsReload(c *gin.Context)
	JobRun(c *gin.Context)
	GetCronJobs(c *gin.Context)
}

type scheduleAPI struct {
	repo repo.ScheduleService
	cron *gocron.Scheduler
	jobs map[uint32]*gocron.Job
}

// NewScheduletAPI get sshclient service instance
func NewScheduleAPI(db *gorm.DB) ScheduleAPI {
	return &scheduleAPI{
		repo: repo.NewScheduleService(db),
		cron: gocron.NewScheduler(time.Local),
		jobs: map[uint32]*gocron.Job{},
	}
}
