package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"
)

const _GIT_USERS_ = "git_users"

func (p databaseService) ReplaceGitUsers(readings []entities.GitUsers) error {

	tx := p.db.Begin()

	err := tx.Where("id IS NOT NULL").Delete(&entities.GitUsers{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete git_users error: %w", err)
	}

	err = tx.Create(readings).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert git_users error: %w", err)
	}

	err = p.UpdatedAtSynTable(_GIT_USERS_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to git_users", len(readings))

	tx.Commit()
	return nil
}
