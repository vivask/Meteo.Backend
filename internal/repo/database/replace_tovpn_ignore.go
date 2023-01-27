package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"
)

const _TOVPN_IGNORE_ = "tovpn_ignores"

func (p databaseService) ReplaceToVpnIgnore(readings []entities.ToVpnIgnore) error {

	tx := p.db.Begin()
	err := tx.Where("id IS NOT NULL").Delete(&entities.ToVpnIgnore{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete tovpn_ignores error: %w", err)
	}

	err = tx.Create(readings).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert tovpn_ignores error: %w", err)
	}
	err = p.UpdatedAtSynTable(_TOVPN_IGNORE_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to tovpn_ignores", len(readings))

	tx.Commit()
	return nil
}
