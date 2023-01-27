package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"
)

const _RADCHECK_ = "hadcheck"

func (p databaseService) ReplaceRadcheck(readings []entities.Radcheck) error {

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
