package repo

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/log"
	"meteo/internal/utils"
	"time"
)

func (p scheduleService) GetAllJobs(pageable dto.Pageable) ([]entities.Jobs, error) {
	var jobs []entities.Jobs

	err := p.db.Order("note asc").Preload("Params").Preload("Executor").
		Preload("Period").Preload("Task.Params").Preload("Day").Find(&jobs).Error
	if err != nil {
		return nil, fmt.Errorf("error read jobs: %w", err)
	}
	return jobs, err
}

func (p scheduleService) GetJobByID(id uint32) (*entities.Jobs, error) {
	job := entities.Jobs{ID: id}
	err := p.db.Preload("Params").Preload("Executor").Preload("Period").
		Preload("Task.Params").Preload("Day").First(&job).Error
	if err != nil {
		return nil, fmt.Errorf("error read jobs: %w", err)
	}
	return &job, err
}

func (p scheduleService) AddJob(job entities.Jobs) error {
	if job.Period.ID == "one" {
		dt, err := utils.GetDateTime(job.Date, job.Time)
		if err != nil {
			return fmt.Errorf("GetDateTime error: %w", err)
		}
		if dt.Unix()-time.Now().Unix() < 0 {
			return fmt.Errorf("task can't be run at this time: %v", job.Time)
		}
	}
	tx := p.db.Begin()
	job.ID = utils.HashNow32()
	for _, param := range job.Params {
		param.ID = utils.HashNow32()
		param.JobID = job.ID
	}
	job.Active = true
	err := tx.Create(&job).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error create job: %w", err)
	}
	tx.Commit()
	// try execute job
	_, err = kit.PutInt("/schedule/jobs/reload", nil)
	if err != nil {
		log.Warningf("Reload jobs error: %v", err)
		err := p.db.Model(&entities.Jobs{}).Where("id = ?", job.ID).Update("active", false).Error
		if err != nil {
			return fmt.Errorf("update jobs error: %w", err)
		}
		return fmt.Errorf("reload jobs error")
	}
	return nil
}

func (p scheduleService) EditJob(job entities.Jobs) error {
	if job.Period.ID == "one" {
		dt, err := utils.GetDateTime(job.Date, job.Time)
		if err != nil {
			return fmt.Errorf("GetDateTime error: %w", err)
		}
		if dt.Unix()-time.Now().Unix() < 0 {
			return fmt.Errorf("task can't be run at this time: %v", job.Time)
		}

	}
	tx := p.db.Begin()
	err := tx.Where("job_id = ?", job.ID).Delete(&entities.JobParams{}).Error
	if err != nil {
		return fmt.Errorf("remove job_params error: %w", err)
	}
	for _, param := range job.Params {
		param.ID = utils.HashNow32()
		param.JobID = job.ID
	}
	err = tx.Save(&job).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error update jobs: %w", err)
	}
	tx.Commit()
	// try execute job
	_, err = kit.PutInt("/schedule/jobs/reload", nil)
	if err != nil {
		log.Warningf("Reload jobs error: %v", err)
		err := p.db.Model(&entities.Jobs{}).Where("id = ?", job.ID).Update("active", false).Error
		if err != nil {
			return fmt.Errorf("update jobs error: %w", err)
		}
		return fmt.Errorf("reload jobs error")
	}
	return nil
}

func (p scheduleService) ActivateJob(id uint32) error {
	job, err := p.GetJobByID(id)
	if err != nil {
		return fmt.Errorf("error GetJob: %w", err)
	}
	if job.Period.ID == "one" {
		dt, err := utils.GetDateTime(job.Date, job.Time)
		if err != nil {
			return fmt.Errorf("GetDateTime error: %w", err)
		}
		if dt.Unix()-time.Now().Unix() < 0 {
			return fmt.Errorf("task can't be run at this time: %v", job.Time)
		}
	}
	err = p.db.Model(&entities.Jobs{}).Where("id = ?", id).Update("active", true).Error
	if err != nil {
		return fmt.Errorf("update jobs error: %w", err)
	}
	// try execute job
	_, err = kit.PutInt("/schedule/jobs/reload", nil)
	if err != nil {
		log.Warningf("Reload jobs error: %v", err)
		err := p.db.Model(&entities.Jobs{}).Where("id = ?", id).Update("active", false).Error
		if err != nil {
			return fmt.Errorf("update jobs error: %w", err)
		}
		return fmt.Errorf("reload jobs error")
	}
	return nil
}

func (p scheduleService) DeactivateJob(id uint32, off bool) error {
	err := p.db.Model(&entities.Jobs{}).Where("id = ?", id).Update("active", false).Error
	if err != nil {
		return fmt.Errorf("update jobs error: %w", err)
	}
	if off {
		_, err = kit.PutInt("/schedule/jobs/reload", nil)
		if err != nil {
			log.Warningf("Reload jobs error: %v", err)
			return fmt.Errorf("reload jobs error")
		}
	}
	return nil
}

func (p scheduleService) RunJob(id uint32) error {
	url := fmt.Sprintf("/schedule/job/run/%v", id)
	_, err := kit.PutInt(url, nil)
	if err != nil {
		return fmt.Errorf("run job error: %w", err)
	}
	return nil
}

func (p scheduleService) DeleteJob(id uint32) error {
	err := p.db.Delete(&entities.Jobs{ID: id}).Error
	if err != nil {
		return fmt.Errorf("error delete jobs: %w", err)
	}
	_, err = kit.PutInt("/schedule/jobs/reload", nil)
	if err != nil {
		return fmt.Errorf("reload jobs error")
	}
	return nil
}

func (p scheduleService) GetAllActiveJobs() ([]entities.Jobs, error) {
	var jobs []entities.Jobs
	err := p.db.Order("note asc").Where("active = true").Preload("Params").
		Preload("Executor").Preload("Period").Preload("Task.Params").Preload("Day").Find(&jobs).Error
	if err != nil {
		return nil, fmt.Errorf("error read jobs: %w", err)
	}
	return jobs, err
}

func (p scheduleService) GetAllPeriods(pageable dto.Pageable) ([]entities.Periods, error) {
	var periods []entities.Periods
	err := p.db.Order("idx asc").Find(&periods).Error
	if err != nil {
		return nil, fmt.Errorf("error read periods: %w", err)
	}
	return periods, nil
}

func (p scheduleService) GetAllDays(pageable dto.Pageable) ([]entities.Days, error) {
	var days []entities.Days
	err := p.db.Find(&days).Error
	if err != nil {
		return nil, fmt.Errorf("error read days: %w", err)
	}
	return days, nil
}

func (p scheduleService) GetAllExecutors(pageable dto.Pageable) ([]entities.Executors, error) {
	var executors []entities.Executors
	err := p.db.Find(&executors).Error
	if err != nil {
		return nil, fmt.Errorf("error read executors: %w", err)
	}
	return executors, nil
}
