package repo

import (
	"fmt"
	"meteo/internal/entities"
	"time"
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
	var settings entities.Settings
	err = p.db.First(&settings).Error
	if err != nil {
		return nil, fmt.Errorf("read settings error: %w", err)
	}
	settings.Esp32DateTimeNow = dt
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

func (p esp32Service) ResetOrders() error {
	settings := new(entities.Settings)
	err := p.db.First(settings).Error
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	settings.SetupMode = false
	settings.Reboot = false
	settings.RadsensHVMode = false
	settings.ClearJournalEsp32 = false
	settings.DigisparkReboot = false
	settings.Firmware = ""
	settings.UpgradeStatus = 0
	settings.Esp32DateTimeNow = time.Now()
	err = p.db.Save(&settings).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}

	return nil
}

func (p esp32Service) GetSensorsState() (*entities.Sensors, error) {
	settings, err := p.GetSettings()
	if err != nil {
		return nil, fmt.Errorf("error read settings: %w", err)
	}

	return &entities.Sensors{
		Bmx280Lock:       settings.Bmx280Lock,
		Ds18b20Lock:      settings.Ds18b20Lock,
		Mics6814Lock:     settings.Mics6814Lock,
		RadsensLock:      settings.RadsensLock,
		Ze08Lock:         settings.Ze08Lock,
		Aht25Lock:        settings.Aht25Lock,
		Esp32DateTimeNow: settings.Esp32DateTimeNow,
	}, nil
}

func (p esp32Service) LockBmx280(lock bool) error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	settings.Bmx280Lock = lock
	err = p.db.Save(settings).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}

func (p esp32Service) LockDs18b20(lock bool) error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	settings.Ds18b20Lock = lock
	err = p.db.Save(settings).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}

func (p esp32Service) LockRadsens(lock bool) error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	settings.RadsensLock = lock
	err = p.db.Save(settings).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}

func (p esp32Service) LockMics6814(lock bool) error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	settings.Mics6814Lock = lock
	err = p.db.Save(settings).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}

func (p esp32Service) LockZe08(lock bool) error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	settings.Ze08Lock = lock
	err = p.db.Save(settings).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}

func (p esp32Service) LockAht25(lock bool) error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	settings.Aht25Lock = lock
	err = p.db.Save(settings).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}
