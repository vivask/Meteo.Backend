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

const _BMX280_ = "bmx280"

func (p databaseService) GetNotSyncBmx280() ([]entities.Bmx280, error) {
	table, err := p.GetSyncTable(_BMX280_)
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	var bmx280 []entities.Bmx280
	if table.SyncedAt.IsZero() {
		err = p.db.Find(&bmx280).Error
	} else {
		err = p.db.Where("date_time >= ?", table.SyncedAt).Find(&bmx280).Error
	}
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return bmx280, err
}

func (p databaseService) AddSyncBmx280(bmx280 []entities.Bmx280) error {
	m.AutoSyncOff(_BMX280_)
	defer m.AutoSyncOn(_BMX280_)

	tx := p.db.Begin()
	count := 0
	for _, v := range bmx280 {
		err := tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insert error: %w", err)
		}
		count++
	}

	err := p.UpdatedAtSynTable(_BMX280_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}
	log.Infof("Received and insert [%d] records to bmx280", count)
	tx.Commit()
	return nil
}

func NotInBmx280(id string, set []entities.Bmx280) bool {
	for _, v := range set {
		if v.ID == id {
			return false
		}
	}
	return true
}

func (p databaseService) SyncBmx280() error {

	err := lock.LockBmx280(true)
	if err != nil {
		return fmt.Errorf("LockBmx280 error: %w", err)
	}
	defer lock.UnlockBmx280(true)

	m.AutoSyncOff(_BMX280_)
	defer m.AutoSyncOn(_BMX280_)

	body, err := kit.GetExt("/esp32/database/bmx280/get")
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

	// Search not exist external
	var wg sync.WaitGroup
	var newExt []entities.Bmx280
	go func(arr *[]entities.Bmx280) {
		for _, v := range intBmx280 {
			if NotInBmx280(v.ID, extBmx280) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newExt)
	var newInt []entities.Bmx280
	go func(arr *[]entities.Bmx280) {
		for _, v := range extBmx280 {
			if NotInBmx280(v.ID, intBmx280) {
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

	err = p.UpdatedAtSynTable(_BMX280_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to bmx280", count)

	tx.Commit()
	return nil
}

func (p databaseService) ReplaceBmx280(readings []entities.Bmx280) error {
	m.AutoSyncOff(_BMX280_)
	defer m.AutoSyncOn(_BMX280_)

	tx := p.db.Begin()
	err := tx.Delete(&entities.Bmx280{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete bmx280 error: %w", err)
	}
	for _, v := range readings {
		err = tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insert bmx280 error: %w", err)
		}
	}
	err = p.UpdatedAtSynTable(_BMX280_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}
	tx.Commit()
	return nil
}
