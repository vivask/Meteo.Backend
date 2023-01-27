package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"
)

const _TOVPN_MANUALS_ = "tovpn_manuals"

func (p databaseService) ReplaceToVpnManual(readings []entities.ToVpnManual) error {

	tx := p.db.Begin()
	err := tx.Where("id IS NOT NULL").Delete(&entities.ToVpnManual{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete tovpn_manuals error: %w", err)
	}

	err = tx.Create(readings).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert tovpn_manuals error: %w", err)
	}
	err = p.UpdatedAtSynTable(_TOVPN_MANUALS_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to tovpn_manuals", len(readings))

	tx.Commit()
	return nil
}
