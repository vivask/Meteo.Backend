package repo

import (
	"fmt"
	"meteo/internal/entities"
	m "meteo/internal/entities/migration"
	"meteo/internal/log"
)

const _TASKS_ = "tasks"

func (p databaseService) GetAllTasks() ([]entities.Tasks, error) {
	var tasks []entities.Tasks
	err := p.db.Preload("Params").Find(&tasks).Error
	if err != nil {
		return nil, fmt.Errorf("error read tasks: %w", err)
	}
	return tasks, err
}

func (p databaseService) ReplaceTasks(readings []entities.Tasks) error {
	m.AutoSyncOff(_TASKS_)
	defer m.AutoSyncOn(_TASKS_)

	tx := p.db.Begin()
	err := tx.Where("id IS NOT NULL").Delete(&entities.Tasks{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete tasks error: %w", err)
	}

	err = tx.Create(readings).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert tasks error: %w", err)
	}
	err = p.UpdatedAtSynTable(_TASKS_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to tasks", len(readings))

	tx.Commit()
	return nil
}
