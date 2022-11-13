package repo

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"time"

	"github.com/gofrs/uuid"
)

func (p esp32Service) AddLoging(msg, t, dts interface{}) error {
	message, ok := msg.(string)
	if !ok {
		return fmt.Errorf("convert interface to string: %v", msg)
	}
	msgType, ok := t.(string)
	if !ok {
		return fmt.Errorf("convert interface to string: %v", t)
	}
	dtstring, ok := dts.(string)
	if !ok {
		return fmt.Errorf("convert interface to string: %v", dts)
	}
	dt, err := time.ParseInLocation("2006-01-02 15:04:05", dtstring, time.Local)
	if err != nil {
		return fmt.Errorf("parse time error: %w", err)
	}
	id, _ := uuid.NewV4()
	logging := entities.Logging{ID: id.String(), Message: message, Type: msgType, CreatedAt: dt}
	err = p.db.Create(&logging).Error
	if err != nil {
		return fmt.Errorf("error insert logging: %w", err)
	}
	return nil
}

func (p esp32Service) JournalClear() error {
	err := p.db.Exec("DELETE FROM logging").Error
	if err != nil {
		return fmt.Errorf("error delete loging: %w", err)
	}

	var settings entities.Settings
	err = p.db.First(&settings).Error
	if err != nil {
		return fmt.Errorf("read settings error: %w", err)
	}
	settings.ClearJournalEsp32 = true
	err = p.db.Save(&settings).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) SetJournaCleared() error {
	set := entities.Settings{}
	err := p.db.First(&set).Error
	if err != nil {
		return fmt.Errorf("read settings error: %w", err)
	}
	set.ClearJournalEsp32 = false
	err = p.db.Save(&set).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) GetAllLoging(pageable dto.Pageable) ([]entities.Logging, error) {
	var loging []entities.Logging
	err := p.db.Order("date_time DESC").Find(&loging).Error
	if err != nil {
		return nil, fmt.Errorf("error read loging: %w", err)
	}
	return loging, err
}
