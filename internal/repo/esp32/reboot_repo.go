package repo

import (
	"fmt"
	"meteo/internal/entities"
)

func (p esp32Service) Esp32Reboot() error {
	set := entities.Settings{}
	err := p.db.First(&set).Error
	if err != nil {
		return fmt.Errorf("read settings error: %w", err)
	}
	set.Reboot = true
	set.Rebooted = false
	err = p.db.Save(&set).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) Esp32Rebooted() error {
	set := entities.Settings{}
	err := p.db.First(&set).Error
	if err != nil {
		return fmt.Errorf("read settings error: %w", err)
	}
	set.Reboot = false
	set.Rebooted = true
	err = p.db.Save(&set).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) GetRebootStatus() (*entities.Settings, error) {
	return p.GetSettings()
}
