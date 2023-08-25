package repo

import (
	"encoding/json"
	"fmt"
	"meteo/internal/entities"
	m "meteo/internal/entities/migration"
	"meteo/internal/kit"
	"meteo/internal/log"
	lock "meteo/internal/repo/esp32"
	"sync"
	"unsafe"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const _ZE08CH2O_ = "ze08ch2o"

func (p databaseService) GetAllZe08ch2o() ([]entities.Ze08ch2o, error) {
	var data []entities.Ze08ch2o
	err := p.db.Order("id").Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return data, err
}

func (p databaseService) GetNotSyncZe08ch2o() ([]entities.Ze08ch2o, error) {
	table, err := p.GetSyncTable(_ZE08CH2O_)
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	var ze08ch2o []entities.Ze08ch2o
	if table.SyncedAt.IsZero() {
		err = p.db.Order("id").Find(&ze08ch2o).Error
	} else {
		err = p.db.Order("id").Where("date_time >= ?", table.SyncedAt).Find(&ze08ch2o).Error
	}
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return ze08ch2o, err
}

func batchCreateZe08ch2o(data []entities.Ze08ch2o, tx *gorm.DB) error {

	chunkSize := int(65534 / unsafe.Sizeof(entities.Ze08ch2o{}))

	for {
		if len(data) == 0 {
			break
		}

		if len(data) < chunkSize {
			chunkSize = len(data)
		}

		chunk := data[0:chunkSize]
		err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&chunk).Error
		if err != nil {
			return fmt.Errorf("insert error: %w", err)
		}
		data = data[chunkSize:]
	}

	return nil
}

func (p databaseService) AddSyncZe08ch2o(ze08ch2o []entities.Ze08ch2o) error {
	m.AutoSyncOff(_ZE08CH2O_)
	defer m.AutoSyncOn(_ZE08CH2O_)

	tx := p.db.Begin()
	err := batchCreateZe08ch2o(ze08ch2o, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert error: %w", err)
	}

	err = p.UpdatedAtSynTable(_ZE08CH2O_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to ze08ch2o", len(ze08ch2o))
	tx.Commit()
	return nil
}

func notInZe08ch2o(id string, set []entities.Ze08ch2o) bool {

	low := 0
	high := len(set) - 1

	for low <= high {
		median := (low + high) / 2

		if set[median].ID < id {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(set) || set[low].ID != id {
		return true
	}

	return false
}

func (p databaseService) SyncZe08ch2o() error {

	err := lock.LockZe08ch2o(true)
	if err != nil {
		return fmt.Errorf("LockZe08ch2o error: %w", err)
	}

	body, err := kit.GetExt("/esp32/database/ze08ch2o/notsync")
	if err != nil {
		return fmt.Errorf("error GET: %w", err)
	}

	var extZe08ch2o []entities.Ze08ch2o
	err = json.Unmarshal(body, &extZe08ch2o)
	if err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	intZe08ch2o, err := p.GetNotSyncZe08ch2o()
	if err != nil {
		return fmt.Errorf("error read ze08ch2o: %w", err)
	}
	lock.UnlockZe08ch2o(true)

	// Search not exist external
	var wg sync.WaitGroup
	var newExt []entities.Ze08ch2o
	go func(arr *[]entities.Ze08ch2o) {
		for _, v := range intZe08ch2o {
			if notInZe08ch2o(v.ID, extZe08ch2o) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newExt)

	// Search not exist internal
	var newInt []entities.Ze08ch2o
	go func(arr *[]entities.Ze08ch2o) {
		for _, v := range extZe08ch2o {
			if notInZe08ch2o(v.ID, intZe08ch2o) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newInt)
	wg.Add(2)
	wg.Wait()

	_, err = kit.PostExt("/esp32/database/ze08ch2o", newExt)
	if err != nil {
		return fmt.Errorf("error POST: %w", err)
	}

	return p.AddSyncZe08ch2o(newInt)
}

func (p databaseService) ReplaceZe08ch2o(readings []entities.Ze08ch2o) error {
	m.AutoSyncOff(_ZE08CH2O_)
	defer m.AutoSyncOn(_ZE08CH2O_)

	err := lock.LockZe08ch2o(true)
	if err != nil {
		return fmt.Errorf("LockZe08ch2o error: %w", err)
	}
	defer lock.UnlockZe08ch2o(true)

	tx := p.db.Begin()
	err = tx.Where("id IS NOT NULL").Delete(&entities.Ze08ch2o{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete ze08ch2o error: %w", err)
	}

	err = batchCreateZe08ch2o(readings, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert ze08ch2o error: %w", err)
	}

	err = p.UpdatedAtSynTable(_ZE08CH2O_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	tx.Commit()
	return nil
}
