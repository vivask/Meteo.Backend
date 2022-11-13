package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/utils"

	"gorm.io/gorm"
)

func (p esp32Service) GetLastDs18b20() (*entities.Ds18b20, error) {
	var ds18b20 entities.Ds18b20
	err := p.db.Order("date_time desc").Last(&ds18b20).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entities.Ds18b20{}, nil
		}
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return &ds18b20, err
}

func (p esp32Service) AddDs18b20(tempr, dts interface{}) error {
	if isLockedDs18b20() {
		return nil
	}

	temperature, err := toFloat(tempr)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}

	dt, err := toTime(dts)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}

	ds18b20 := entities.Ds18b20{ID: utils.HashTime(dt), Tempr: temperature}
	err = p.db.Create(&ds18b20).Error
	if err != nil {
		return fmt.Errorf("insert ds18b20: %v, error: %w", ds18b20, err)
	}
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	message := fmt.Sprintf("DS18B20: Температура воздуха: %f°C", temperature)
	if temperature < settings.MinDs18b20 {
		settings.MinDs18b20Alarm = true
		err = p.db.Save(&settings).Error
		if err != nil {
			return fmt.Errorf("update settings error: %w", err)
		}
		_, err := kit.PostInt("/messanger/telegram", message)
		if err != nil {
			return fmt.Errorf("can't send telegram message: %w", err)
		}
	}
	if temperature > settings.MaxDs18b20 && !settings.MaxDs18b20Alarm {
		settings.MaxDs18b20Alarm = true
		err = p.db.Save(&settings).Error
		if err != nil {
			return fmt.Errorf("update settings error: %w", err)
		}
		_, err := kit.PostInt("/messanger/telegram", message)
		if err != nil {
			return fmt.Errorf("can't send telegram message: %w", err)
		}
	}
	return err
}

func (p esp32Service) Ds18b20TemperatureChk() error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	settings.MinDs18b20Alarm = false
	settings.MaxDs18b20Alarm = false
	err = p.db.Save(&settings).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}
