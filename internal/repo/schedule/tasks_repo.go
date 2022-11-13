package repo

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/utils"
)

func (p scheduleService) GetTaskByID(id string) (*entities.Tasks, error) {
	task := entities.Tasks{ID: id}
	err := p.db.Preload("Params").First(&task).Error
	if err != nil {
		return nil, fmt.Errorf("error read tasks: %w", err)
	}
	return &task, nil
}

func (p scheduleService) GetAllTasks(pageable dto.Pageable) ([]entities.Tasks, error) {
	var tasks []entities.Tasks
	err := p.db.Order("name desc").Preload("Params").Find(&tasks).Error
	if err != nil {
		return nil, fmt.Errorf("error read tasks: %w", err)
	}
	return tasks, nil
}

func (p scheduleService) AddTask(task entities.Tasks) error {
	for _, param := range task.Params {
		param.ID = utils.HashNow32()
		param.TaskID = task.ID
	}
	tx := p.db.Begin()
	err := tx.Create(&task).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error create tasks: %w", err)
	}
	tx.Commit()
	return nil
}

func (p scheduleService) EditTask(task entities.Tasks) error {
	tx := p.db.Begin()
	err := tx.Where("task_id = ?", task.ID).Delete(&entities.TaskParams{}).Error
	if err != nil {
		return fmt.Errorf("error delete task_params: %w", err)
	}
	for _, param := range task.Params {
		param.ID = utils.HashNow32()
		param.TaskID = task.ID
	}
	err = tx.Save(&task).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error update tasks: %w", err)
	}
	tx.Commit()

	return nil
}

func (p scheduleService) DelTask(id string) error {
	tx := p.db.Begin()
	err := tx.Delete(&entities.Tasks{ID: id}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete tasks: %w", err)
	}
	tx.Commit()
	return nil
}
