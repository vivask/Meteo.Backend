package repo

import (
	"fmt"
	"meteo/internal/entities"
	m "meteo/internal/entities/migration"
	"meteo/internal/log"
)

const _HOMEZONES_ = "homezones"

func (p databaseService) GetAllHomezones() ([]entities.Homezone, error) {
	var hosts []entities.Homezone
	err := p.db.Find(&hosts).Error
	if err != nil {
		return nil, fmt.Errorf("error read homezones: %w", err)
	}
	return hosts, err
}

func (p databaseService) ReplaceHomezone(readings []entities.Homezone) error {
	m.AutoSyncOff(_HOMEZONES_)
	defer m.AutoSyncOn(_HOMEZONES_)

	tx := p.db.Begin()
	err := tx.Where("id IS NOT NULL").Delete(&entities.Homezone{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete homezones error: %w", err)
	}

	err = tx.Create(readings).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert homezones error: %w", err)
	}
	err = p.UpdatedAtSynTable(_HOMEZONES_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to homezones", len(readings))

	tx.Commit()
	return nil
}
