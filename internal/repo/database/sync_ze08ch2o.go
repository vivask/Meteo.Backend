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

const _ZE08CH2O_ = "ze08ch2o"

func (p databaseService) GetNotSyncZe08ch2o() ([]entities.Ze08ch2o, error) {
	table, err := p.GetSyncTable(_ZE08CH2O_)
	if err != nil {
		return nil, fmt.Errorf("error read tasks: %w", err)
	}
	var ze08ch2o []entities.Ze08ch2o
	if table.SyncedAt.IsZero() {
		err = p.db.Find(&ze08ch2o).Error
	} else {
		err = p.db.Where("date_time >= ?", table.SyncedAt).Find(&ze08ch2o).Error
	}
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return ze08ch2o, err
}

func (p databaseService) AddSyncZe08ch2o(ze08ch2o []entities.Ze08ch2o) error {
	m.AutoSyncOff(_ZE08CH2O_)
	defer m.AutoSyncOn(_ZE08CH2O_)

	tx := p.db.Begin()
	count := 0
	for _, v := range ze08ch2o {
		err := tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insert error: %w", err)
		}
		count++
	}

	err := p.UpdatedAtSynTable(_ZE08CH2O_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to ze08ch2o", count)
	tx.Commit()
	return nil
}

func NotInZe08ch2o(id string, set []entities.Ze08ch2o) bool {
	for _, v := range set {
		if v.ID == id {
			return false
		}
	}
	return true
}

func (p databaseService) SyncZe08ch2o() error {

	err := lock.LockZe08ch2o(true)
	if err != nil {
		return fmt.Errorf("LockZe08ch2o error: %w", err)
	}
	defer lock.UnlockZe08ch2o(true)

	m.AutoSyncOff(_ZE08CH2O_)
	defer m.AutoSyncOn(_ZE08CH2O_)

	body, err := kit.GetExt("/esp32/database/ze08ch2o/get")
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

	// Search not exist external
	var wg sync.WaitGroup
	var newExt []entities.Ze08ch2o
	go func(arr *[]entities.Ze08ch2o) {
		for _, v := range intZe08ch2o {
			if NotInZe08ch2o(v.ID, extZe08ch2o) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newExt)
	// Search not exist internal
	var newInt []entities.Ze08ch2o
	go func(arr *[]entities.Ze08ch2o) {
		for _, v := range extZe08ch2o {
			if NotInZe08ch2o(v.ID, intZe08ch2o) {
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

	tx := p.db.Begin()
	count := 0
	for _, v := range newInt {
		err := tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insert error: %w", err)
		}
		count++
	}

	err = p.UpdatedAtSynTable(_ZE08CH2O_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records to ze08ch2o", count)

	tx.Commit()
	return nil
}

func (p databaseService) ReplaceZe08ch2o(readings []entities.Ze08ch2o) error {
	m.AutoSyncOff(_ZE08CH2O_)
	defer m.AutoSyncOn(_ZE08CH2O_)

	tx := p.db.Begin()
	err := tx.Delete(&entities.Ze08ch2o{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete ze08ch2o error: %w", err)
	}

	for _, v := range readings {
		err = tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insert ze08ch2o error: %w", err)
		}
	}

	err = p.UpdatedAtSynTable(_ZE08CH2O_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	tx.Commit()
	return nil
}
