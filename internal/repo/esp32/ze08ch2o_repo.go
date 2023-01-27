package repo

import (
	"fmt"
	"meteo/internal/dto"
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

	value, err := toFloat(ch2o)
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
	message := fmt.Sprintf("ZE08CH2O: Превышена концентрация CH2O: %f ppm", value)
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

func (p esp32Service) GetZe08MinByHours(period dto.Period) ([]entities.Ze08ch2o, error) {

	var ze08 []entities.Ze08ch2o
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	MIN(ch2o) ch2o
	FROM ze08ch2o
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&ze08).Error
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return ze08, err
}

func (p esp32Service) GetZe08MaxByHours(period dto.Period) ([]entities.Ze08ch2o, error) {

	var ze08 []entities.Ze08ch2o
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	MAX(ch2o) ch2o
	FROM ze08ch2o
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&ze08).Error
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return ze08, err
}

func (p esp32Service) GetZe08AvgByHours(period dto.Period) ([]entities.Ze08ch2o, error) {

	var ze08 []entities.Ze08ch2o
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	AVG(ch2o) ch2o
	FROM ze08ch2o
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&ze08).Error
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return ze08, err
}

func (p esp32Service) GetZe08MinByDays(period dto.Period) ([]entities.Ze08ch2o, error) {

	var ze08 []entities.Ze08ch2o
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	MIN(ch2o) ch2o
	FROM ze08ch2o
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&ze08).Error
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return ze08, err
}

func (p esp32Service) GetZe08MaxByDays(period dto.Period) ([]entities.Ze08ch2o, error) {

	var ze08 []entities.Ze08ch2o
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	MAX(ch2o) ch2o
	FROM ze08ch2o
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&ze08).Error
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return ze08, err
}

func (p esp32Service) GetZe08AvgByDays(period dto.Period) ([]entities.Ze08ch2o, error) {

	var ze08 []entities.Ze08ch2o
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	AVG(ch2o) ch2o
	FROM ze08ch2o
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&ze08).Error
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return ze08, err
}

func (p esp32Service) GetZe08MinByMonths(period dto.Period) ([]entities.Ze08ch2o, error) {

	var ze08 []entities.Ze08ch2o
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	MIN(ch2o) ch2o
	FROM ze08ch2o
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&ze08).Error
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return ze08, err
}

func (p esp32Service) GetZe08MaxByMonths(period dto.Period) ([]entities.Ze08ch2o, error) {

	var ze08 []entities.Ze08ch2o
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	MAX(ch2o) ch2o
	FROM ze08ch2o
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&ze08).Error
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return ze08, err
}

func (p esp32Service) GetZe08AvgByMonths(period dto.Period) ([]entities.Ze08ch2o, error) {

	var ze08 []entities.Ze08ch2o
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	AVG(ch2o) ch2o
	FROM ze08ch2o
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&ze08).Error
	if err != nil {
		return nil, fmt.Errorf("error read ze08ch2o: %w", err)
	}
	return ze08, err
}
