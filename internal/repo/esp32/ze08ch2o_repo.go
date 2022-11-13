package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/utils"

	"gorm.io/gorm"
)

func (p esp32Service) GetLastZe08ch2o() (*entities.Ze08ch2o, error) {
	var ze08ch2o entities.Ze08ch2o
	err := p.db.Order("date_time desc").Last(&ze08ch2o).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entities.Ze08ch2o{}, nil
		}
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return &ze08ch2o, err
}

func (p esp32Service) AddZe08ch2o(ch2o, dts interface{}) error {
	if isLockedZe08ch2o() {
		return nil
	}

	value, err := toInt(ch2o)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}
	dt, err := toTime(dts)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}

	ze08ch2o := entities.Ze08ch2o{ID: utils.HashTime(dt), Ch2o: value}
	err = p.db.Create(&ze08ch2o).Error
	if err != nil {
		return fmt.Errorf("insert ze08ch2o: %v, error: %w", ze08ch2o, err)
	}
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	message := fmt.Sprintf("ZE08CH2O: Превышена концентрация CH2O: %d ppm", value)
	if value > settings.MaxCh2o && !settings.MaxCh2oAlarm {
		settings.MaxCh2oAlarm = true
		err = p.db.Save(&settings).Error
		if err != nil {
			return fmt.Errorf("update settings error: %w", err)
		}
		_, err := kit.PostInt("/messanger/telegram", message)
		if err != nil {
			return fmt.Errorf("can't send telegram message: %w", err)
		}
	}
	return nil
}

func (p esp32Service) Ze08ch2oChk() error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	settings.MaxCh2oAlarm = false
	err = p.db.Save(&settings).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}
