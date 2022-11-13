package repo

import (
	"fmt"
	"meteo/internal/entities"
)

func (p esp32Service) SuccessUpgrade() error {
	set := entities.Settings{}
	err := p.db.First(&set).Error
	if err != nil {
		return fmt.Errorf("read settings error: %w", err)
	}
	set.Firmware = "_EMPTY_"
	set.UpgradeStatus = 1
	err = p.db.Save(&set).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) TerminateUpgrade() error {
	set := entities.Settings{}
	err := p.db.First(&set).Error
	if err != nil {
		return fmt.Errorf("read settings error: %w", err)
	}
	set.Firmware = "_EMPTY_"
	set.UpgradeStatus = -1
	err = p.db.Save(&set).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) UpgradeEsp32(fName string) error {
	set := entities.Settings{}
	err := p.db.First(&set).Error
	if err != nil {
		return fmt.Errorf("read settings error: %w", err)
	}
	set.Firmware = fName
	set.UpgradeStatus = 0
	err = p.db.Save(&set).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) GetUpgradeStatus() (*entities.Settings, error) {
	return p.GetSettings()
}
