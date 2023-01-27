package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"
)

const _RADVERIFIED_ = "radverified"

func (p databaseService) ReplaceRadverified(readings []entities.Radverified) error {

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
