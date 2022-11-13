package repo

import (
	"meteo/internal/dto"
	"meteo/internal/entities"

	"gorm.io/gorm"
)

// ScheduleService api controller of produces
type ScheduleService interface {
	GetTaskByID(id string) (*entities.Tasks, error)
	GetAllTasks(pageable dto.Pageable) ([]entities.Tasks, error)
	AddTask(task entities.Tasks) error
	EditTask(task entities.Tasks) error
	DelTask(id string) error
	GetAllJobs(pageable dto.Pageable) ([]entities.Jobs, error)
	GetJobByID(id uint32) (*entities.Jobs, error)
	AddJob(job entities.Jobs) error
	EditJob(job entities.Jobs) error
	ActivateJob(id uint32) error
	DeactivateJob(id uint32, off bool) error
	RunJob(id uint32) error
	DeleteJob(id uint32) error
	GetAllActiveJobs() ([]entities.Jobs, error)
	GetAllPeriods(pageable dto.Pageable) ([]entities.Periods, error)
	GetAllDays(pageable dto.Pageable) ([]entities.Days, error)
	GetAllExecutors(pageable dto.Pageable) ([]entities.Executors, error)
}

type scheduleService struct {
	db *gorm.DB
}

// NewScheduleService get schedule service instance
func NewScheduleService(db *gorm.DB) ScheduleService {
	return &scheduleService{db}
}
