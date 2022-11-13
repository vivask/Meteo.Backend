package repo

import (
	"fmt"
	"meteo/internal/entities"
)

func (p esp32Service) GetSettings() (*entities.Settings, error) {
	settings := new(entities.Settings)
	err := p.db.First(settings).Error
	if err != nil {
		return nil, fmt.Errorf("error read settings: %w", err)
	}
	return settings, err
}

func (p esp32Service) SetSettings(s *entities.Settings) error {
	err := p.db.Save(s).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}

func (p esp32Service) SetEsp32Settings(cpu0L, cpu1L, dti interface{}) (*entities.Settings, error) {
	dt, err := toTime(dti)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	cpu0, err := toFloat(cpu0L)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	cpu1, err := toFloat(cpu1L)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	var settings entities.Settings
	err = p.db.First(&settings).Error
	if err != nil {
		return nil, fmt.Errorf("read settings error: %w", err)
	}
	settings.Esp32DateTimeNow = dt
	settings.Cpu0Load = cpu0
	settings.Cpu1Load = cpu1
	err = p.db.Save(&settings).Error
	if err != nil {
		return nil, fmt.Errorf("update settings error: %w", err)
	}
	err = p.db.First(&settings).Error
	if err != nil {
		return nil, fmt.Errorf("error read settings: %w", err)
	}
	return &settings, nil
}
