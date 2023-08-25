package repo

import (
	"fmt"
	"meteo/internal/entities"
	m "meteo/internal/entities/migration"
	"meteo/internal/log"
)

const _TOVPN_AUTO_ = "tovpn_autos"

func (p databaseService) GetAllToVpnAuto() ([]entities.ToVpnAuto, error) {
	var hosts []entities.ToVpnAuto
	err := p.db.Find(&hosts).Error
	if err != nil {
		return nil, fmt.Errorf("error read tovpn_autos: %w", err)
	}
	return hosts, err
}

func (p databaseService) ReplaceToVpnAuto(readings []entities.ToVpnAuto) error {
	m.AutoSyncOff(_TOVPN_AUTO_)
	defer m.AutoSyncOn(_TOVPN_AUTO_)

	tx := p.db.Begin()
	err := tx.Where("hostname IS NOT NULL").Delete(&entities.ToVpnAuto{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete tovpn_autos error: %w", err)
	}

	err = tx.Create(readings).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert tovpn_autos error: %w", err)
	}
	err = p.UpdatedAtSynTable(_TOVPN_AUTO_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to tovpn_autos", len(readings))

	tx.Commit()
	return nil
}
