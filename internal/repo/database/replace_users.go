package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"
)

const _USERS_ = "users"

func (p databaseService) ReplaceUser(readings []entities.User) error {

	tx := p.db.Begin()
	err := tx.Where("id IS NOT NULL").Delete(&entities.User{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete users error: %w", err)
	}

	err = tx.Create(readings).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert users error: %w", err)
	}
	err = p.UpdatedAtSynTable(_TASKS_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to users", len(readings))

	tx.Commit()
	return nil
}
