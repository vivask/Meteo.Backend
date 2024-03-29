package repo

import (
	"fmt"
	"meteo/internal/entities"
	m "meteo/internal/entities/migration"
	"meteo/internal/log"
)

const _RADVERIFIED_ = "radverified"

func (p databaseService) GetAllRadverified() ([]entities.Radverified, error) {
	var users []entities.Radverified
	err := p.db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("error read radverified: %w", err)
	}
	return users, err
}

func (p databaseService) ReplaceRadverified(readings []entities.Radverified) error {
	m.AutoSyncOff(_RADVERIFIED_)
	defer m.AutoSyncOn(_RADVERIFIED_)

	tx := p.db.Begin()
	err := tx.Where("id IS NOT NULL").Delete(&entities.Radverified{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete radverified error: %w", err)
	}

	err = tx.Create(readings).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert radverified error: %w", err)
	}
	err = p.UpdatedAtSynTable(_RADVERIFIED_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to radverified", len(readings))

	tx.Commit()
	return nil
}
