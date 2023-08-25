package repo

import (
	"fmt"
	"meteo/internal/entities"
	m "meteo/internal/entities/migration"
	"meteo/internal/log"
)

const _JOBS_ = "jobs"

func (p databaseService) GetAllJobs() ([]entities.Jobs, error) {
	var jobs []entities.Jobs
	err := p.db.Preload("Params").Preload("Executor").
		Preload("Period").Preload("Task.Params").Find(&jobs).Error
	if err != nil {
		return nil, fmt.Errorf("error read jobs: %w", err)
	}
	return jobs, err
}

func (p databaseService) ReplaceJobs(readings []entities.Jobs) error {
	m.AutoSyncOff(_JOBS_)
	defer m.AutoSyncOn(_JOBS_)

	tx := p.db.Begin()
	err := tx.Where("id IS NOT NULL").Delete(&entities.Jobs{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete jobs error: %w", err)
	}

	err = tx.Create(readings).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert jobs error: %w", err)
	}
	err = p.UpdatedAtSynTable(_JOBS_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to jobs", len(readings))

	tx.Commit()
	return nil
}
