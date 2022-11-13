package repo

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/utils"

	"gorm.io/gorm"
)

func (p esp32Service) GetLastRadsens() (*entities.Radsens, error) {
	var radsens entities.Radsens
	err := p.db.Order("date_time desc").Last(&radsens).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entities.Radsens{}, nil
		}
		return nil, fmt.Errorf("error read radsens: %w", err)
	}
	return &radsens, err
}

func (p esp32Service) AddRadsens(dyn, stat, pl, dts interface{}) error {
	if isLockedRadsens() {
		return nil
	}

	static, err := toFloat(stat)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}
	dynamic, err := toFloat(dyn)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}
	pulse, err := toInt(pl)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}
	dt, err := toTime(dts)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}

	radsens := entities.Radsens{ID: utils.HashTime(dt), Dynamic: dynamic, Static: static, Pulse: pulse}
	err = p.db.Create(&radsens).Error
	if err != nil {
		return fmt.Errorf("insert radsens: %v, error: %w", radsens, err)
	}
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	message := fmt.Sprintf("RadSens: Превышена динамическая интенсивность излучения: %f мкР/ч", dynamic)
	if dynamic > settings.MaxRadDyn && !settings.MaxRadDynAlarm {
		settings.MaxRadDynAlarm = true
		err = p.db.Save(&settings).Error
		if err != nil {
			return fmt.Errorf("update settings error: %w", err)
		}
		_, err := kit.PostInt("/messanger/telegram", message)
		if err != nil {
			return fmt.Errorf("can't send telegram message: %w", err)
		}
	}
	message = fmt.Sprintf("RadSens: Превышена статическая интенсивность излучения: %f мкР/ч", static)
	if static > settings.MaxRadStat && !settings.MaxRadStatAlarm {
		settings.MaxRadStatAlarm = true
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

func (p esp32Service) SetHVRadsens(state interface{}) error {
	need_state, err := toBool(state)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}
	set := entities.Settings{}
	err = p.db.First(&set).Error
	if err != nil {
		return fmt.Errorf("read settings error: %w", err)
	}
	set.RadsensHVMode = false
	set.RadsensHVState = need_state
	err = p.db.Save(&set).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) SetSensRadsens(sens interface{}) error {
	s, err := toInt(sens)
	if err != nil {
		return fmt.Errorf("convert error: %w", err)
	}
	set := entities.Settings{}
	err = p.db.First(&set).Error
	if err != nil {
		return fmt.Errorf("read settings error: %w", err)
	}
	set.RadsensSensitivity = s
	err = p.db.Save(&set).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) RadsensStaticChk() error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	settings.MaxRadStatAlarm = false
	err = p.db.Save(&settings).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) RadsensDynamicChk() error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	settings.MaxRadDynAlarm = false
	err = p.db.Save(&settings).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}

func (p esp32Service) RadsensHVSet() error {
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	settings.RadsensHVState = !settings.RadsensHVState
	err = p.db.Save(&settings).Error
	if err != nil {
		return fmt.Errorf("update settings error: %w", err)
	}
	return nil
}
