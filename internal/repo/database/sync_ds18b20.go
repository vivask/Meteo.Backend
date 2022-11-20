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
)

const _DS18B20_ = "ds18b20"

func (p databaseService) GetNotSyncDs18b20() ([]entities.Ds18b20, error) {
	table, err := p.GetSyncTable(_DS18B20_)
	if err != nil {
		return nil, fmt.Errorf("error read tasks: %w", err)
	}
	var ds18b20 []entities.Ds18b20
	if table.SyncedAt.IsZero() {
		err = p.db.Find(&ds18b20).Error
	} else {
		err = p.db.Where("date_time >= ?", table.SyncedAt).Find(&ds18b20).Error
	}
	if err != nil {
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return ds18b20, err
}

func (p databaseService) AddSyncDs18b20(ds18b20 []entities.Ds18b20) error {
	m.AutoSyncOff(_DS18B20_)
	defer m.AutoSyncOn(_DS18B20_)

	tx := p.db.Begin()
	count := 0
	for _, v := range ds18b20 {
		err := tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insert error: %w", err)
		}
		count++
	}

	err := p.UpdatedAtSynTable(_DS18B20_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to ds18b20", count)

	tx.Commit()
	return nil
}

func NotInDs18b20(id string, set []entities.Ds18b20) bool {
	for _, v := range set {
		if v.ID == id {
			return false
		}
	}
	return true
}

func (p databaseService) SyncDs18b20() error {

	err := lock.LockDs18b20(true)
	if err != nil {
		return fmt.Errorf("LockZe08ch2o error: %w", err)
	}
	defer lock.UnlockDs18b20(true)

	m.AutoSyncOff(_DS18B20_)
	defer m.AutoSyncOn(_DS18B20_)

	body, err := kit.GetExt("/esp32/database/ds18b20/get")
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
		for _, v := range intDs18b20 {
			if NotInDs18b20(v.ID, extDs18b20) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newExt)
	var newInt []entities.Ds18b20
	go func(arr *[]entities.Ds18b20) {
		for _, v := range extDs18b20 {
			if NotInDs18b20(v.ID, intDs18b20) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newInt)
	wg.Add(2)
	wg.Wait()

	_, err = kit.PostExt("/esp32/database/ds18b20", newExt)
	if err != nil {
		return fmt.Errorf("error POST: %w", err)
	}

	tx := p.db.Begin()
	// Search not exist internal and insert
	count := 0
	for _, v := range newInt {
		err := tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insert error: %w", err)
		}
		count++
	}

	err = p.UpdatedAtSynTable(_DS18B20_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}
	log.Infof("Received and insert [%d] records to ds18b20", count)

	tx.Commit()
	return nil
}

func (p databaseService) ReplaceDs18b20(readings []entities.Ds18b20) error {
	m.AutoSyncOff(_DS18B20_)
	defer m.AutoSyncOn(_DS18B20_)

	tx := p.db.Begin()
	err := tx.Delete(&entities.Ds18b20{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete ds18b20 error: %w", err)
	}
	for _, v := range readings {
		err = tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insert ds18b20 error: %w", err)
		}
	}
	err = p.UpdatedAtSynTable(_DS18B20_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}
	tx.Commit()
	return nil
}
