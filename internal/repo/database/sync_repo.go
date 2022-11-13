package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"
)

func (p databaseService) ExecRaw(cb entities.Callback) error {
	err := p.db.Exec(cb.Query, cb.Params...).Error
	if err != nil {
		log.Debug(cb.Query)
		log.Debug(cb.Params...)
		return fmt.Errorf("exec error: %w", err)
	}
	return nil
}
