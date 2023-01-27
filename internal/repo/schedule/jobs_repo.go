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
		Preload("Period").Preload("Task.Params").Find(&jobs).Error
	if err != nil {
		return nil, fmt.Errorf("error read jobs: %w", err)
	}
	return jobs, err
}

func (p scheduleService) GetJobByID(id uint32) (*entities.Jobs, error) {
	job := entities.Jobs{ID: id}
	err := p.db.Preload("Params").Preload("Executor").Preload("Period").
		Preload("Task.Params").First(&job).Error
	if err != nil {
		return nil, fmt.Errorf("error read jobs: %w", err)
	}
	return &job, err
}

func (p scheduleService) AddJob(job entities.Jobs) (uint32, bool, error) {
	if job.Period.ID == "one" {
		dt, err := utils.GetDateTime(job.Date, job.Time)
		if err != nil {
			return 0, false, fmt.Errorf("GetDateTime error: %w", err)
		}
		if dt.Unix()-time.Now().Unix() < 0 {
			return 0, false, fmt.Errorf("task can't be run at this time: %v", job.Time)
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
		return 0, false, fmt.Errorf("error create job: %w", err)
	}
	tx.Commit()
	// try execute job
	_, err = kit.PutInt("/schedule/jobs/reload", nil)
	if err != nil {
		log.Warningf("Reload jobs error: %v", err)
		err := p.db.Model(&entities.Jobs{}).Where("id = ?", job.ID).Update("active", false).Error
		if err != nil {
			return job.ID, false, fmt.Errorf("update jobs error: %w", err)
		}
		return job.ID, false, fmt.Errorf("reload jobs error")
	}
	return job.ID, true, nil
}

func (p scheduleService) EditJob(job entities.Jobs) (bool, error) {
	if job.Period.ID == "one" {
		dt, err := utils.GetDateTime(job.Date, job.Time)
		if err != nil {
			return false, fmt.Errorf("GetDateTime error: %w", err)
		}
		if dt.Unix()-time.Now().Unix() < 0 {
			return false, fmt.Errorf("task can't be run at this time: %v", job.Time)
		}

	}
	tx := p.db.Begin()
	err := tx.Where("job_id = ?", job.ID).Delete(&entities.JobParams{}).Error
	if err != nil {
		return false, fmt.Errorf("remove job_params error: %w", err)
	}
	for _, param := range job.Params {
		param.ID = utils.HashNow32()
		param.JobID = job.ID
	}
	err = tx.Save(&job).Error
	if err != nil {
		tx.Rollback()
		return false, fmt.Errorf("error update jobs: %w", err)
	}
	tx.Commit()
	// try execute job
	_, err = kit.PutInt("/schedule/jobs/reload", nil)
	if err != nil {
		log.Warningf("Reload jobs error: %v", err)
		err := p.db.Model(&entities.Jobs{}).Where("id = ?", job.ID).Update("active", false).Error
		if err != nil {
			return false, fmt.Errorf("update jobs error: %w", err)
		}
		return false, fmt.Errorf("reload jobs error")
	}
	return true, nil
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
		Preload("Executor").Preload("Period").Preload("Task.Params").Find(&jobs).Error
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

func (p scheduleService) GetAllExecutors(pageable dto.Pageable) ([]entities.Executors, error) {
	var executors []entities.Executors
	err := p.db.Order("idx asc").Find(&executors).Error
	if err != nil {
		return nil, fmt.Errorf("error read executors: %w", err)
	}
	return executors, nil
}
