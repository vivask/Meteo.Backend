package repo

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/utils"
	"time"

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

// func (p esp32Service) AddDs18b20(tempr, dts interface{}) error {
// 	if isLockedDs18b20() {
// 		return nil
// 	}

// 	temperature, err := toFloat(tempr)
// 	if err != nil {
// 		return fmt.Errorf("convert error: %w", err)
// 	}

// 	dt, err := toTime(dts)
// 	if err != nil {
// 		return fmt.Errorf("convert error: %w", err)
// 	}

// 	ds18b20 := entities.Ds18b20{ID: utils.HashTime(dt), Tempr: temperature}
// 	err = p.db.Create(&ds18b20).Error
// 	if err != nil {
// 		return fmt.Errorf("insert ds18b20: %v, error: %w", ds18b20, err)
// 	}
// 	settings, err := p.GetSettings()
// 	if err != nil {
// 		return fmt.Errorf("error read settings: %w", err)
// 	}
// 	message := fmt.Sprintf("DS18B20: Температура воздуха: %f°C", temperature)
// 	if temperature < settings.MinDs18b20 {
// 		settings.MinDs18b20Alarm = true
// 		err = p.db.Save(&settings).Error
// 		if err != nil {
// 			return fmt.Errorf("update settings error: %w", err)
// 		}
// 		_, err := kit.PostInt("/messanger/telegram", message)
// 		if err != nil {
// 			return fmt.Errorf("can't send telegram message: %w", err)
// 		}
// 	}
// 	if temperature > settings.MaxDs18b20 && !settings.MaxDs18b20Alarm {
// 		settings.MaxDs18b20Alarm = true
// 		err = p.db.Save(&settings).Error
// 		if err != nil {
// 			return fmt.Errorf("update settings error: %w", err)
// 		}
// 		_, err := kit.PostInt("/messanger/telegram", message)
// 		if err != nil {
// 			return fmt.Errorf("can't send telegram message: %w", err)
// 		}
// 	}
// 	return err
// }

func (p esp32Service) AddDs18b20(temperature float64) error {
	if isLockedDs18b20() {
		return nil
	}

	ds18b20 := entities.Ds18b20{ID: utils.HashTime(time.Now()), Tempr: temperature}
	err := p.db.Create(&ds18b20).Error
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

func (p esp32Service) GetDs18b20MinByHours(period dto.Period) ([]entities.Ds18b20, error) {

	var data []entities.Ds18b20
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	MIN(tempr) tempr
	FROM ds18b20
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return data, err
}

func (p esp32Service) GetDs18b20MaxByHours(period dto.Period) ([]entities.Ds18b20, error) {

	var data []entities.Ds18b20
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	MAX(tempr) tempr
	FROM ds18b20
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return data, err
}

func (p esp32Service) GetDs18b20AvgByHours(period dto.Period) ([]entities.Ds18b20, error) {

	var data []entities.Ds18b20
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	AVG(tempr) tempr
	FROM ds18b20
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return data, err
}

func (p esp32Service) GetDs18b20MinByDays(period dto.Period) ([]entities.Ds18b20, error) {

	var data []entities.Ds18b20
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	MIN(tempr) tempr
	FROM ds18b20
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return data, err
}

func (p esp32Service) GetDs18b20MaxByDays(period dto.Period) ([]entities.Ds18b20, error) {

	var data []entities.Ds18b20
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	MAX(tempr) tempr
	FROM ds18b20
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return data, err
}

func (p esp32Service) GetDs18b20AvgByDays(period dto.Period) ([]entities.Ds18b20, error) {

	var data []entities.Ds18b20
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	AVG(tempr) tempr
	FROM ds18b20
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return data, err
}

func (p esp32Service) GetDs18b20MinByMonths(period dto.Period) ([]entities.Ds18b20, error) {

	var data []entities.Ds18b20
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	MIN(tempr) tempr
	FROM ds18b20
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return data, err
}

func (p esp32Service) GetDs18b20MaxByMonths(period dto.Period) ([]entities.Ds18b20, error) {

	var data []entities.Ds18b20
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	MAX(tempr) tempr
	FROM ds18b20
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return data, err
}

func (p esp32Service) GetDs18b20AvgByMonths(period dto.Period) ([]entities.Ds18b20, error) {

	var data []entities.Ds18b20
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	AVG(tempr) tempr
	FROM ds18b20
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error read ds18b20: %w", err)
	}
	return data, err
}
