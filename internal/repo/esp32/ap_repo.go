package repo

import (
	"fmt"
	"meteo/internal/entities"
)

func (p esp32Service) SetAccesPointMode() error {
	set := entities.Settings{}
	err := p.db.First(&set).Error
	if err != nil {
		return fmt.Errorf("read settings error: %w", err)
	}
	set.SetupMode = true
	set.SetupStatus = false
	err = p.db.Save(&set).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) SetSTAMode() error {
	set := entities.Settings{}
	err := p.db.First(&set).Error
	if err != nil {
		return fmt.Errorf("read settings error: %w", err)
	}
	set.SetupMode = false
	set.SetupStatus = true
	err = p.db.Save(&set).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) GetStatusAccesPoint() (*entities.Settings, error) {
	return p.GetSettings()
}
