package repo

import (
	"fmt"
	"meteo/internal/entities"
	m "meteo/internal/entities/migration"
	"meteo/internal/log"
)

const _RADCHECK_ = "radcheck"

func (p databaseService) GetAllRadcheck() ([]entities.Radcheck, error) {
	var users []entities.Radcheck
	err := p.db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("error read radcheck: %w", err)
	}
	return users, err
}
func (p databaseService) ReplaceRadcheck(readings []entities.Radcheck) error {
	m.AutoSyncOff(_RADCHECK_)
	defer m.AutoSyncOn(_RADCHECK_)

	tx := p.db.Begin()
	err := tx.Where("id IS NOT NULL").Delete(&entities.Radcheck{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete hadcheck error: %w", err)
	}

	err = tx.Create(readings).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert hadcheck error: %w", err)
	}
	err = p.UpdatedAtSynTable(_RADCHECK_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to hadcheck", len(readings))

	tx.Commit()
	return nil
}
