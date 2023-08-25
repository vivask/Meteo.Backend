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

const _MICS6814_ = "mics6814"

func (p databaseService) GetAllMics6814() ([]entities.Mics6814, error) {
	var data []entities.Mics6814
	err := p.db.Order("id").Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return data, err
}

func (p databaseService) GetNotSyncMics6814() ([]entities.Mics6814, error) {
	table, err := p.GetSyncTable(_MICS6814_)
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	var mics6814 []entities.Mics6814
	if table.SyncedAt.IsZero() {
		err = p.db.Order("id").Find(&mics6814).Error
	} else {
		err = p.db.Order("id").Where("date_time >= ?", table.SyncedAt).Find(&mics6814).Error
	}
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return mics6814, err
}

func batchCreateMics6814(data []entities.Mics6814, tx *gorm.DB) error {

	chunkSize := int(65534 / unsafe.Sizeof(entities.Mics6814{}))

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

func (p databaseService) AddSyncMics6814(data []entities.Mics6814) error {
	m.AutoSyncOff(_MICS6814_)
	defer m.AutoSyncOn(_MICS6814_)

	tx := p.db.Begin()
	err := batchCreateMics6814(data, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert error: %w", err)
	}

	err = p.UpdatedAtSynTable(_MICS6814_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}
	log.Infof("Received and insert [%d] records to mics6814", len(data))
	tx.Commit()
	return nil
}

func notInMics6814(id string, set []entities.Mics6814) bool {

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

func (p databaseService) SyncMics6814() error {

	err := lock.LockMics6814(true)
	if err != nil {
		return fmt.Errorf("LockMics6814 error: %w", err)
	}

	body, err := kit.GetExt("/esp32/database/mics6814/notsync")
	if err != nil {
		return fmt.Errorf("internal error: %w", err)
	}

	var extMics6814 []entities.Mics6814
	err = json.Unmarshal(body, &extMics6814)
	if err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	intMics6814, err := p.GetNotSyncMics6814()
	if err != nil {
		return fmt.Errorf("error read mics6814: %w", err)
	}
	lock.UnlockMics6814(true)

	// Search not exist external
	var wg sync.WaitGroup
	var newExt []entities.Mics6814
	go func(arr *[]entities.Mics6814) {
		for _, v := range intMics6814 {
			if notInMics6814(v.ID, extMics6814) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newExt)

	// Search not exist internal
	var newInt []entities.Mics6814
	go func(arr *[]entities.Mics6814) {
		for _, v := range extMics6814 {
			if notInMics6814(v.ID, intMics6814) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newInt)
	wg.Add(2)
	wg.Wait()

	_, err = kit.PostExt("/esp32/database/mics6814", newExt)
	if err != nil {
		return fmt.Errorf("internal error: %w", err)
	}

	return p.AddSyncMics6814(newInt)
}

func (p databaseService) ReplaceMics6814(readings []entities.Mics6814) error {
	m.AutoSyncOff(_MICS6814_)
	defer m.AutoSyncOn(_MICS6814_)

	err := lock.LockDs18b20(true)
	if err != nil {
		return fmt.Errorf("LockDs18b20 error: %w", err)
	}
	defer lock.UnlockDs18b20(true)

	tx := p.db.Begin()
	err = tx.Where("id IS NOT NULL").Delete(&entities.Mics6814{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete mics6814 error: %w", err)
	}

	err = batchCreateMics6814(readings, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert mics6814 error: %w", err)
	}

	err = p.UpdatedAtSynTable(_MICS6814_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}
	tx.Commit()
	return nil
}
