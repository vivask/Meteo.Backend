package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/utils"

	"gorm.io/gorm"
)

func (p esp32Service) GetLastMics6814() (*entities.Mics6814, error) {
	var mics6814 entities.Mics6814
	err := p.db.Order("date_time desc").Last(&mics6814).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entities.Mics6814{}, nil
		}
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return &mics6814, err
}

func (p esp32Service) AddMics6814(co, no2, nh3, dts interface{}) error {
	if isLockedMics6814() {
		return nil
	}

	val_co, err := toFloat(co)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}
	val_no2, err := toFloat(no2)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}
	val_nh3, err := toFloat(nh3)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}
	dt, err := toTime(dts)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}

	mics6814 := entities.Mics6814{ID: utils.HashTime(dt), Co: val_co, No2: val_no2, Nh3: val_nh3}
	err = p.db.Create(&mics6814).Error
	if err != nil {
		return fmt.Errorf("insert mics6814: %v, error: %w", mics6814, err)
	}
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	message := fmt.Sprintf("MICS6814: Превышена концентрация CO: %f ppm", val_co)
	if val_co > settings.Max6814Co && !settings.Max6814CoAlarm {
		settings.Max6814CoAlarm = true
		err = p.db.Save(&settings).Error
		if err != nil {
			return fmt.Errorf("update settings error: %w", err)
		}
		_, err := kit.PostInt("/messanger/telegram", message)
		if err != nil {
			return fmt.Errorf("can't send telegram message: %w", err)
		}
	}
	message = fmt.Sprintf("MICS6814: Превышена концентрация NO2: %f ppm", val_no2)
	if val_no2 > settings.Max6814No2 && !settings.Max6814No2Alarm {
		settings.Max6814No2Alarm = true
		err = p.db.Save(&settings).Error
		if err != nil {
			return fmt.Errorf("update settings error: %w", err)
		}
		_, err := kit.PostInt("/messanger/telegram", message)
		if err != nil {
			return fmt.Errorf("can't send telegram message: %w", err)
		}
	}
	message = fmt.Sprintf("MICS6814: Превышена концентрация NH3: %f ppm", val_nh3)
	if val_nh3 > settings.Max6814Nh3 && !settings.Max6814Nh3Alarm {
		settings.Max6814Nh3Alarm = true
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

func (p esp32Service) Mics6814CoChk() error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	settings.Max6814CoAlarm = false
	err = p.db.Save(&settings).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) Mics6814No2Chk() error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	settings.Max6814No2Alarm = false
	err = p.db.Save(&settings).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) Mics6814Nh3Chk() error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	settings.Max6814Nh3Alarm = false
	err = p.db.Save(&settings).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}
