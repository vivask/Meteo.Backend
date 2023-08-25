package repo

import (
	"fmt"
	"meteo/internal/entities"
	m "meteo/internal/entities/migration"
	"meteo/internal/log"
)

const _TOVPN_MANUALS_ = "tovpn_manuals"

func (p databaseService) GetAllToVpnManual() ([]entities.ToVpnManual, error) {
	var hosts []entities.ToVpnManual
	err := p.db.Preload("AccesList").Find(&hosts).Error
	if err != nil {
		return nil, fmt.Errorf("error read tovpn_manuals: %w", err)
	}
	return hosts, err
}

func (p databaseService) ReplaceToVpnManual(readings []entities.ToVpnManual) error {
	m.AutoSyncOff(_TOVPN_MANUALS_)
	defer m.AutoSyncOn(_TOVPN_MANUALS_)

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
