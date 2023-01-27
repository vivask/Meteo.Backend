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

const _DS18B20_ = "ds18b20"

func (p databaseService) GetAllDs18b20() ([]entities.Ds18b20, error) {
	var data []entities.Ds18b20
	err := p.db.Order("id").Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return data, err
}

func (p databaseService) GetNotSyncDs18b20() ([]entities.Ds18b20, error) {

	table, err := p.GetSyncTable(_DS18B20_)
	if err != nil {
		return nil, fmt.Errorf("error read tasks: %w", err)
	}
	var ds18b20 []entities.Ds18b20
	if table.SyncedAt.IsZero() {
		err = p.db.Order("id").Find(&ds18b20).Error
	} else {
		err = p.db.Order("id").Where("date_time >= ?", table.SyncedAt).Find(&ds18b20).Error
	}
	if err != nil {
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return ds18b20, err
}

func batchCreateDs18b20(data []entities.Ds18b20, tx *gorm.DB) error {

	chunkSize := int(65534 / unsafe.Sizeof(&entities.Ds18b20{}))

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

func (p databaseService) AddSyncDs18b20(ds18b20 []entities.Ds18b20) error {
	m.AutoSyncOff(_DS18B20_)
	defer m.AutoSyncOn(_DS18B20_)

	tx := p.db.Begin()
	err := batchCreateDs18b20(ds18b20, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("batch create error: %w", err)
	}

	err = p.UpdatedAtSynTable(_DS18B20_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to ds18b20", len(ds18b20))

	tx.Commit()
	return nil
}

func notInDs18b20(id string, set []entities.Ds18b20) bool {

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

func (p databaseService) SyncDs18b20() error {

	err := lock.LockDs18b20(true)
	if err != nil {
		return fmt.Errorf("LockDs18b20 error: %w", err)
	}
	defer lock.UnlockDs18b20(true)

	body, err := kit.GetExt("/esp32/database/ds18b20/notsync")
	if err != nil {
		return fmt.Errorf("error GET: %w", err)
	}

	var extDs18b20 []entities.Ds18b20
	err = json.Unmarshal(body, &extDs18b20)
	if err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	intDs18b20, err := p.GetNotSyncDs18b20()
	if err != nil {
		return fmt.Errorf("error read ds18b20: %w", err)
	}

	// Search not exist external and send
	var wg sync.WaitGroup
	var newExt []entities.Ds18b20
	go func(arr *[]entities.Ds18b20) {
		for _, val := range intDs18b20 {
			if notInDs18b20(val.ID, extDs18b20) {
				*arr = append(*arr, val)
			}
		}
		wg.Done()
	}(&newExt)

	// Search not exist internal
	var newInt []entities.Ds18b20
	go func(arr *[]entities.Ds18b20) {
		for _, val := range extDs18b20 {
			if notInDs18b20(val.ID, intDs18b20) {
				*arr = append(*arr, val)
			}
		}
		wg.Done()
	}(&newInt)
	wg.Add(2)
	wg.Wait()

	if len(newExt) > 0 {
		_, err = kit.PostExt("/esp32/database/ds18b20", newExt)
		if err != nil {
			return fmt.Errorf("error POST: %w", err)
		}
	}

	return p.AddSyncDs18b20(newInt)
}

func (p databaseService) ReplaceDs18b20(readings []entities.Ds18b20) error {
	m.AutoSyncOff(_DS18B20_)
	defer m.AutoSyncOn(_DS18B20_)

	err := lock.LockDs18b20(true)
	if err != nil {
		return fmt.Errorf("LockDs18b20 error: %w", err)
	}
	defer lock.UnlockDs18b20(true)

	tx := p.db.Begin()
	err = tx.Where("id IS NOT NULL").Delete(&entities.Ds18b20{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete ds18b20 error: %w", err)
	}

	err = batchCreateDs18b20(readings, tx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert ds18b20 error: %w", err)
	}
	err = p.UpdatedAtSynTable(_DS18B20_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to ds18b20", len(readings))

	tx.Commit()
	return nil
}
