package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"
	"time"
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

func (p databaseService) GetSyncTable(id string) (*entities.SyncTables, error) {
	table := &entities.SyncTables{}
	err := p.db.Where("name = ?", id).Preload("Params").First(table).Error
	if err != nil {
		return nil, fmt.Errorf("error read sync_tables: %w", err)
	}
	return table, nil
}

func (p databaseService) UpdatedAtSynTable(name string) error {
	table, err := p.GetSyncTable(name)
	if err != nil {
		return fmt.Errorf("GetSyncTable error: %w", err)
	}
	table.SyncedAt = time.Now()
	err = p.db.Model(table).Where("name = ?", name).Updates(table).Error
	if err != nil {
		return fmt.Errorf("update sync_tables error: %w", err)
	}
	return nil
}
