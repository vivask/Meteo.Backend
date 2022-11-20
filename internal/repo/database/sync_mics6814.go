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

const _MICS6814_ = "mics6814"

func (p databaseService) GetNotSyncMics6814() ([]entities.Mics6814, error) {
	table, err := p.GetSyncTable(_MICS6814_)
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	var mics6814 []entities.Mics6814
	if table.SyncedAt.IsZero() {
		err = p.db.Find(&mics6814).Error
	} else {
		err = p.db.Where("date_time >= ?", table.SyncedAt).Find(&mics6814).Error
	}
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return mics6814, err
}

func (p databaseService) AddSyncMics6814(mics6814 []entities.Mics6814) error {
	m.AutoSyncOff(_MICS6814_)
	defer m.AutoSyncOn(_MICS6814_)

	tx := p.db.Begin()
	count := 0
	for _, v := range mics6814 {
		err := tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insert error: %w", err)
		}
		count++
	}

	err := p.UpdatedAtSynTable(_MICS6814_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}
	log.Infof("Received and insert [%d] records to mics6814", count)
	tx.Commit()
	return nil
}

func NotInMics6814(id string, set []entities.Mics6814) bool {
	for _, v := range set {
		if v.ID == id {
			return false
		}
	}
	return true
}

func (p databaseService) SyncMics6814() error {

	err := lock.LockMics6814(true)
	if err != nil {
		return fmt.Errorf("LockMics6814 error: %w", err)
	}
	defer lock.UnlockBmx280(true)

	m.AutoSyncOff(_MICS6814_)
	defer m.AutoSyncOn(_MICS6814_)

	body, err := kit.GetExt("/esp32/database/mics6814/get")
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

	// Search not exist external
	var wg sync.WaitGroup
	var newExt []entities.Mics6814
	go func(arr *[]entities.Mics6814) {
		for _, v := range intMics6814 {
			if NotInMics6814(v.ID, extMics6814) {
				*arr = append(*arr, v)
			}
		}
		wg.Done()
	}(&newExt)
	var newInt []entities.Mics6814
	go func(arr *[]entities.Mics6814) {
		for _, v := range extMics6814 {
			if NotInMics6814(v.ID, intMics6814) {
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

	err = p.UpdatedAtSynTable(_MICS6814_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}

	log.Infof("Received and insert [%d] records from %s to mics6814", count)

	tx.Commit()
	return nil
}

func (p databaseService) ReplaceMics6814(readings []entities.Mics6814) error {
	m.AutoSyncOff(_MICS6814_)
	defer m.AutoSyncOn(_MICS6814_)

	tx := p.db.Begin()
	err := tx.Delete(&entities.Mics6814{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete mics6814 error: %w", err)
	}
	for _, v := range readings {
		err = tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insert mics6814 error: %w", err)
		}
	}
	err = p.UpdatedAtSynTable(_MICS6814_)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdatedAtSynTable error: %w", err)
	}
	tx.Commit()
	return nil
}
