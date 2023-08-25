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

const _RADSENS_ = "radsens"

func (p databaseService) GetAllRadsens() ([]entities.Radsens, error) {
	var data []entities.Radsens
	err := p.db.Order("id").Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read radsens: %w", err)
	}
	return data, err
}

func (p databaseService) GetNotSyncRadsens() ([]entities.Radsens, error) {
	table, err := p.GetSyncTable(_RADSENS_)
	if err != nil {
		return nil, fmt.Errorf("error read radsens: %w", err)
	}
	var radsens []entities.Radsens
	if table.SyncedAt.IsZero() {
		err = p.db.Order("id").Find(&radsens).Error
	} else {
		err = p.db.Order("id").Where("date_time >= ?", table.SyncedAt).Find(&radsens).Error
	}
	if err != nil {
		return nil, fmt.Errorf("error read radsens: %w", err)
	}
	return radsens, err
}

func batchCreateRadsens(data []entities.Radsens, tx *gorm.DB) error {

	chunkSize := int(65534 / unsafe.Sizeof(entities.Radsens{}))

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

func (p databaseService) AddSyncRadsens(radsens []entities.Radsens) error {
	m.AutoSyncOff(_RADSENS_)
	defer m.AutoSyncOn(_RADSENS_)

	tx := p.db.Begin()
	err := batchCreateRadsens(radsens, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("batch create error: %w", err)
	}

	err = p.UpdatedAtSynTable(_RADSENS_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to radsens", len(radsens))

	tx.Commit()
	return nil
}

func notInRadsens(id string, set []entities.Radsens) bool {

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

func (p databaseService) SyncRadsens() error {

	err := lock.LockRadsens(true)
	if err != nil {
		return fmt.Errorf("LockRadsens error: %w", err)
	}

	body, err := kit.GetExt("/esp32/database/radsens/notsync")
	if err != nil {
		return fmt.Errorf("error GET: %w", err)
	}

	var extRadsens []entities.Radsens
	err = json.Unmarshal(body, &extRadsens)
	if err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	intRadsens, err := p.GetNotSyncRadsens()
	if err != nil {
		return fmt.Errorf("error read radsens: %w", err)
	}
	lock.UnlockRadsens(true)

	// Search not exist external and send
	var wg sync.WaitGroup
	var newExt []entities.Radsens
	go func(arr *[]entities.Radsens) {
		for _, val := range intRadsens {
			if notInRadsens(val.ID, extRadsens) {
				*arr = append(*arr, val)
			}
		}
		wg.Done()
	}(&newExt)

	// Search not exist internal
	var newInt []entities.Radsens
	go func(arr *[]entities.Radsens) {
		for _, val := range extRadsens {
			if notInRadsens(val.ID, intRadsens) {
				*arr = append(*arr, val)
			}
		}
		wg.Done()
	}(&newInt)
	wg.Add(2)
	wg.Wait()

	_, err = kit.PostExt("/esp32/database/radsens", newExt)
	if err != nil {
		return fmt.Errorf("error POST: %w", err)
	}

	return p.AddSyncRadsens(newInt)
}

func (p databaseService) ReplaceRadsens(readings []entities.Radsens) error {
	m.AutoSyncOff(_RADSENS_)
	defer m.AutoSyncOn(_RADSENS_)

	err := lock.LockRadsens(true)
	if err != nil {
		return fmt.Errorf("LockRadsens error: %w", err)
	}
	defer lock.UnlockRadsens(true)

	tx := p.db.Begin()
	err = tx.Where("id IS NOT NULL").Delete(&entities.Radsens{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete radsens error: %w", err)
	}

	err = batchCreateRadsens(readings, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert radsens error: %w", err)
	}
	err = p.UpdatedAtSynTable(_RADSENS_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to radsens", len(readings))

	tx.Commit()
	return nil
}
