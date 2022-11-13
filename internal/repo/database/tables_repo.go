package repo

import (
	"errors"
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/log"
	"meteo/internal/utils"
	"os"
)

const DELIMITER = "|"

func (p databaseService) GetAllTables(pageable dto.Pageable) ([]entities.SyncTables, error) {
	var tables []entities.SyncTables

	err := p.db.Preload("Params").Preload("Params.SyncTypes").Find(&tables).Error
	if err != nil {
		return nil, fmt.Errorf("error read sync_tables: %w", err)
	}
	for idx := range tables {
		csv := fmt.Sprintf("/import/%s.csv", tables[idx].ID)
		if _, err := os.Stat(csv); !errors.Is(err, os.ErrNotExist) {
			tables[idx].IsImport = true
		}
	}
	return tables, err
}

func (p databaseService) GetAllSTypes(pageable dto.Pageable) ([]entities.SyncTypes, error) {
	var stypes []entities.SyncTypes

	err := p.db.Find(&stypes).Error
	if err != nil {
		return nil, fmt.Errorf("error read sync_types: %w", err)
	}
	return stypes, err
}

func (p databaseService) AddTable(table entities.SyncTables) error {
	for _, param := range table.Params {
		param.ID = utils.HashNow32()
		param.SyncTableID = table.ID
	}
	tx := p.db.Begin()
	err := tx.Create(&table).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error create sync_tables: %w", err)
	}
	tx.Commit()
	return nil
}

func (p databaseService) EditTable(table entities.SyncTables) error {
	tx := p.db.Begin()
	err := tx.Where("table_id = ?", table.ID).Delete(&entities.SyncParams{}).Error
	if err != nil {
		return fmt.Errorf("remove sync_params error: %w", err)
	}
	for _, param := range table.Params {
		param.ID = utils.HashNow32()
		param.SyncTableID = table.ID
	}
	err = tx.Save(&table).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error create sync_tables: %w", err)
	}
	tx.Commit()
	return nil
}

func (p databaseService) DelTable(id string) error {
	tx := p.db.Begin()
	err := tx.Delete(&entities.SyncTables{ID: id}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete sync_tables: %w", err)
	}
	tx.Commit()
	return nil
}

func (p databaseService) DelTables(tables []entities.SyncTables) error {
	tx := p.db.Begin()
	log.Infof("TABLES: %v", tables)
	err := tx.Delete(&tables).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete sync_tables: %w", err)
	}
	tx.Commit()
	return nil
}

func (p databaseService) GetTableById(id string) (*entities.SyncTables, error) {
	table := &entities.SyncTables{}

	err := p.db.Where("name = ?", id).Preload("Params").Preload("Params.SyncTypes").First(table).Error
	if err != nil {
		return nil, fmt.Errorf("error read sync_tables: %w", err)
	}
	return table, err
}

func (p databaseService) ImportTableContent(id string) error {
	table, err := p.GetTableById(id)
	if err != nil {
		return fmt.Errorf("can't read table by id [%s]: %w", id, err)
	}
	csv := fmt.Sprintf("/import/%s.csv", table.ID)
	tx := p.db.Begin()
	query := fmt.Sprintf("DELETE FROM %s;", table.ID)
	err = tx.Exec(query).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("can't import table: %w", err)
	}
	query = fmt.Sprintf("COPY %s FROM '%s' DELIMITER '%s' CSV HEADER;", table.ID, csv, DELIMITER)
	err = tx.Exec(query).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("can't import table: %w", err)
	}
	tx.Commit()
	return nil
}

func (p databaseService) ImportTablesContent(tables []entities.SyncTables) error {
	tx := p.db.Begin()
	for idx := range tables {
		id := tables[idx].ID
		table, err := p.GetTableById(id)
		if err != nil {
			return fmt.Errorf("can't read table by id [%s]: %w", id, err)
		}
		csv := fmt.Sprintf("/import/%s.csv", table.ID)
		query := fmt.Sprintf("DELETE FROM %s;", table.ID)
		err = tx.Exec(query).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("can't import table: %w", err)
		}
		query = fmt.Sprintf("COPY %s FROM '%s' DELIMITER '%s' CSV HEADER;", table.ID, csv, DELIMITER)
		err = tx.Exec(query).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("can't import table: %w", err)
		}
	}
	tx.Commit()
	return nil
}

func (p databaseService) SaveTableContent(id string) error {
	table, err := p.GetTableById(id)
	if err != nil {
		return fmt.Errorf("can't read table by id [%s]: %w", id, err)
	}
	csv := fmt.Sprintf("/import/%s.csv", table.ID)
	query := fmt.Sprintf("COPY %s TO '%s' WITH DELIMITER '%s' CSV HEADER;", table.ID, csv, DELIMITER)
	err = p.db.Exec(query).Error
	if err != nil {
		return fmt.Errorf("can't save table: %w", err)
	}
	return nil
}

func (p databaseService) SaveTablesContent(tables []entities.SyncTables) error {
	tx := p.db.Begin()
	for idx := range tables {
		id := tables[idx].ID
		table, err := p.GetTableById(id)
		if err != nil {
			return fmt.Errorf("can't read table by id [%s]: %w", id, err)
		}
		csv := fmt.Sprintf("/import/%s.csv", table.ID)
		query := fmt.Sprintf("COPY %s TO '%s' WITH DELIMITER '%s' CSV HEADER;", table.ID, csv, DELIMITER)
		err = tx.Exec(query).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("can't save table: %w", err)
		}
	}
	tx.Commit()
	return nil
}
