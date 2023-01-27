package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"
)

const _SSH_KEYS_ = "ssh_keys"

func (p databaseService) ReplaceSshKeys(readings []entities.SshKeys) error {

	tx := p.db.Begin()
	err := tx.Where("id IS NOT NULL").Delete(&entities.SshKeys{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete ssh_keys error: %w", err)
	}

	err = tx.Create(readings).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert ssh_keys error: %w", err)
	}
	err = p.UpdatedAtSynTable(_SSH_KEYS_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to ssh_keys", len(readings))

	tx.Commit()
	return nil
}
