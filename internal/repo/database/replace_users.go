package repo

import (
	"fmt"
	"meteo/internal/entities"
	m "meteo/internal/entities/migration"
	"meteo/internal/log"
)

const _USERS_ = "users"

func (p databaseService) GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	err := p.db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("error read users: %w", err)
	}
	return users, err
}

func (p databaseService) ReplaceUser(readings []entities.User) error {
	m.AutoSyncOff(_USERS_)
	defer m.AutoSyncOn(_USERS_)

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
	err = p.UpdatedAtSynTable(_USERS_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to users", len(readings))

	tx.Commit()
	return nil
}
