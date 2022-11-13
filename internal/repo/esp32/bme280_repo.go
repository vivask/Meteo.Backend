package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/utils"

	"gorm.io/gorm"
)

func (p esp32Service) GetLastBmx280() (*entities.Bmx280, error) {
	var bmx280 entities.Bmx280
	err := p.db.Order("date_time desc").Last(&bmx280).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entities.Bmx280{}, nil
		}
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return &bmx280, err
}

func (p esp32Service) AddBme280(press, tempr, hum, dts interface{}) error {
	if isLockedBmx280() {
		return nil
	}

	temperature, err := toFloat(tempr)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}
	pressure, err := toFloat(press)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}
	humidity, err := toFloat(hum)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}
	dt, err := toTime(dts)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}

	bmx280 := entities.Bmx280{ID: utils.HashTime(dt), Press: pressure, Tempr: temperature, Hum: humidity}
	err = p.db.Create(&bmx280).Error
	if err != nil {
		return fmt.Errorf("insert bmx280: %v, error: %w", bmx280, err)
	}
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	message := fmt.Sprintf("BMX280: Температура воздуха: %f°C", temperature)
	if temperature < settings.MinBmx280Tempr && !settings.MinBmx280TemprAlarm {
		settings.MinBmx280TemprAlarm = true
		err = p.db.Save(&settings).Error
		if err != nil {
			return fmt.Errorf("update settings error: %w", err)
		}
		_, err := kit.PostInt("/messanger/telegram", message)
		if err != nil {
			return fmt.Errorf("can't send telegram message: %w", err)
		}
	}
	if temperature > settings.MaxBmx280Tempr && !settings.MinBmx280TemprAlarm {
		settings.MaxBmx280TemprAlarm = true
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

func (p esp32Service) Bme280TemperatureChk() error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	settings.MinBmx280TemprAlarm = false
	settings.MaxBmx280TemprAlarm = false
	err = p.db.Save(&settings).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}
