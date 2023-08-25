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

const _AHT25_ = "aht25"

func (p databaseService) GetAllAht25() ([]entities.Aht25, error) {
	var aht25 []entities.Aht25
	err := p.db.Order("id").Find(&aht25).Error
	if err != nil {
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	return aht25, err
}

func (p databaseService) GetNotSyncAht25() ([]entities.Aht25, error) {
	table, err := p.GetSyncTable(_AHT25_)
	if err != nil {
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	var aht25 []entities.Aht25
	if table.SyncedAt.IsZero() {
		err = p.db.Order("id").Find(&aht25).Error
	} else {
		err = p.db.Order("id").Where("date_time >= ?", table.SyncedAt).Find(&aht25).Error
	}
	if err != nil {
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	return aht25, err
}

func batchCreateAht25(data []entities.Aht25, tx *gorm.DB) error {

	chunkSize := int(65534 / unsafe.Sizeof(entities.Aht25{}))

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

func (p databaseService) AddSyncAht25(data []entities.Aht25) error {
	m.AutoSyncOff(_AHT25_)
	defer m.AutoSyncOn(_AHT25_)

	tx := p.db.Begin()
	err := batchCreateAht25(data, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert error: %w", err)
	}

	err = p.UpdatedAtSynTable(_AHT25_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}
	log.Infof("Received and insert [%d] records to aht25", len(data))
	tx.Commit()
	return nil
}

func notInAht25(id string, set []entities.Aht25) bool {

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

func (p databaseService) SyncAht25() error {

	err := lock.LockAht25(true)
	if err != nil {
		return fmt.Errorf("LockAht25 error: %w", err)
	}

	body, err := kit.GetExt("/esp32/database/aht25/notsync")
	if err != nil {
		return fmt.Errorf("internal error: %w", err)
	}

	var extAht25 []entities.Aht25
	err = json.Unmarshal(body, &extAht25)
	if err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	intAht25, err := p.GetNotSyncAht25()
	if err != nil {
		return fmt.Errorf("error read aht25: %w", err)
	}
	lock.UnlockAht25(true)

	// Search not exist external
	var wg sync.WaitGroup
	var newExt []entities.Aht25
	go func(arr *[]entities.Aht25) {
		for _, v := range intAht25 {
			if notInAht25(v.ID, extAht25) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newExt)

	// Search not exist internal
	var newInt []entities.Aht25
	go func(arr *[]entities.Aht25) {
		for _, v := range extAht25 {
			if notInAht25(v.ID, intAht25) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newInt)
	wg.Add(2)
	wg.Wait()

	_, err = kit.PostExt("/esp32/database/aht25", newExt)
	if err != nil {
		return fmt.Errorf("internal error: %w", err)
	}

	return p.AddSyncAht25(newInt)
}

func (p databaseService) ReplaceAht25(readings []entities.Aht25) error {
	m.AutoSyncOff(_AHT25_)
	defer m.AutoSyncOn(_AHT25_)

	err := lock.LockAht25(true)
	if err != nil {
		return fmt.Errorf("LockAht25 error: %w", err)
	}
	defer lock.UnlockAht25(true)

	tx := p.db.Begin()
	err = tx.Where("id IS NOT NULL").Delete(&entities.Aht25{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete aht25 error: %w", err)
	}

	err = batchCreateAht25(readings, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert aht25 error: %w", err)
	}

	err = p.UpdatedAtSynTable(_AHT25_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}
	tx.Commit()
	return nil
}
