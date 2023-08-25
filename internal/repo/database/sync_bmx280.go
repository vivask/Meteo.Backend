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

const _BMX280_ = "bmx280"

func (p databaseService) GetAllBmx280() ([]entities.Bmx280, error) {
	var bmx280 []entities.Bmx280
	err := p.db.Order("id").Find(&bmx280).Error
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return bmx280, err
}

func (p databaseService) GetNotSyncBmx280() ([]entities.Bmx280, error) {
	table, err := p.GetSyncTable(_BMX280_)
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	var bmx280 []entities.Bmx280
	if table.SyncedAt.IsZero() {
		err = p.db.Order("id").Find(&bmx280).Error
	} else {
		err = p.db.Order("id").Where("date_time >= ?", table.SyncedAt).Find(&bmx280).Error
	}
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return bmx280, err
}

func batchCreateBmx280(data []entities.Bmx280, tx *gorm.DB) error {

	chunkSize := int(65534 / unsafe.Sizeof(entities.Bmx280{}))

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

func (p databaseService) AddSyncBmx280(data []entities.Bmx280) error {
	m.AutoSyncOff(_BMX280_)
	defer m.AutoSyncOn(_BMX280_)

	tx := p.db.Begin()
	err := batchCreateBmx280(data, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert error: %w", err)
	}

	err = p.UpdatedAtSynTable(_BMX280_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}
	log.Infof("Received and insert [%d] records to bmx280", len(data))
	tx.Commit()
	return nil
}

func notInBmx280(id string, set []entities.Bmx280) bool {

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

func (p databaseService) SyncBmx280() error {

	err := lock.LockBmx280(true)
	if err != nil {
		return fmt.Errorf("LockBmx280 error: %w", err)
	}

	body, err := kit.GetExt("/esp32/database/bmx280/notsync")
	if err != nil {
		return fmt.Errorf("internal error: %w", err)
	}

	var extBmx280 []entities.Bmx280
	err = json.Unmarshal(body, &extBmx280)
	if err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	intBmx280, err := p.GetNotSyncBmx280()
	if err != nil {
		return fmt.Errorf("error read bmx280: %w", err)
	}
	lock.UnlockBmx280(true)

	// Search not exist external
	var wg sync.WaitGroup
	var newExt []entities.Bmx280
	go func(arr *[]entities.Bmx280) {
		for _, v := range intBmx280 {
			if notInBmx280(v.ID, extBmx280) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newExt)

	// Search not exist internal
	var newInt []entities.Bmx280
	go func(arr *[]entities.Bmx280) {
		for _, v := range extBmx280 {
			if notInBmx280(v.ID, intBmx280) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newInt)
	wg.Add(2)
	wg.Wait()

	_, err = kit.PostExt("/esp32/database/bmx280", newExt)
	if err != nil {
		return fmt.Errorf("internal error: %w", err)
	}

	return p.AddSyncBmx280(newInt)
}

func (p databaseService) ReplaceBmx280(readings []entities.Bmx280) error {
	m.AutoSyncOff(_BMX280_)
	defer m.AutoSyncOn(_BMX280_)

	err := lock.LockBmx280(true)
	if err != nil {
		return fmt.Errorf("LockBmx280 error: %w", err)
	}
	defer lock.UnlockBmx280(true)

	tx := p.db.Begin()
	err = tx.Where("id IS NOT NULL").Delete(&entities.Bmx280{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete bmx280 error: %w", err)
	}

	err = batchCreateBmx280(readings, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert bmx280 error: %w", err)
	}

	err = p.UpdatedAtSynTable(_BMX280_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}
	tx.Commit()
	return nil
}
