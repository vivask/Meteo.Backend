package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"
)

const _SSH_HOSTS_ = "ssh_hosts"

func (p databaseService) ReplaceSshHosts(readings []entities.SshHosts) error {

	tx := p.db.Begin()
	err := tx.Where("id IS NOT NULL").Delete(&entities.SshHosts{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete ssh_hosts error: %w", err)
	}

	err = tx.Create(readings).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert ssh_hosts error: %w", err)
	}
	err = p.UpdatedAtSynTable(_SSH_HOSTS_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to ssh_hosts", len(readings))

	tx.Commit()
	return nil
}
