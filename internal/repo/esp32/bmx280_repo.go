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

// func (p esp32Service) AddBme280(press, tempr, hum, dts interface{}) error {
// 	if isLockedBmx280() {
// 		return nil
// 	}

// 	temperature, err := toFloat(tempr)
// 	if err != nil {
// 		return fmt.Errorf("convert error: %w", err)
// 	}
// 	pressure, err := toFloat(press)
// 	if err != nil {
// 		return fmt.Errorf("convert error: %w", err)
// 	}
// 	humidity, err := toFloat(hum)
// 	if err != nil {
// 		return fmt.Errorf("convert error: %w", err)
// 	}
// 	dt, err := toTime(dts)
// 	if err != nil {
// 		return fmt.Errorf("convert error: %w", err)
// 	}

// 	bmx280 := entities.Bmx280{ID: utils.HashTime(dt), Press: pressure, Tempr: temperature, Hum: humidity}
// 	err = p.db.Create(&bmx280).Error
// 	if err != nil {
// 		return fmt.Errorf("insert bmx280: %v, error: %w", bmx280, err)
// 	}
// 	settings, err := p.GetSettings()
// 	if err != nil {
// 		return fmt.Errorf("error read settings: %w", err)
// 	}
// 	message := fmt.Sprintf("BMX280: Температура воздуха: %f°C", temperature)
// 	if temperature < settings.MinBmx280Tempr && !settings.MinBmx280TemprAlarm {
// 		settings.MinBmx280TemprAlarm = true
// 		err = p.db.Save(&settings).Error
// 		if err != nil {
// 			return fmt.Errorf("update settings error: %w", err)
// 		}
// 		_, err := kit.PostInt("/messanger/telegram", message)
// 		if err != nil {
// 			return fmt.Errorf("can't send telegram message: %w", err)
// 		}
// 	}
// 	if temperature > settings.MaxBmx280Tempr && !settings.MinBmx280TemprAlarm {
// 		settings.MaxBmx280TemprAlarm = true
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

func (p esp32Service) AddBme280(pressure, temperature, humidity float64) error {
	if isLockedBmx280() {
		return nil
	}

	bmx280 := entities.Bmx280{ID: utils.HashTime(time.Now()), Press: pressure, Tempr: temperature, Hum: humidity}
	err := p.db.Create(&bmx280).Error
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
	if temperature > settings.MaxBmx280Tempr && !settings.MaxBmx280TemprAlarm {
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

func (p esp32Service) GetBmx280MinByHours(period dto.Period) ([]entities.Bmx280, error) {

	var bmx280 []entities.Bmx280
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	MIN(press) press, MIN(tempr) tempr, MIN(hum) hum
	FROM bmx280
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&bmx280).Error
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return bmx280, err
}

func (p esp32Service) GetBmx280MaxByHours(period dto.Period) ([]entities.Bmx280, error) {

	var bmx280 []entities.Bmx280
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	MAX(press) press, MAX(tempr) tempr, MAX(hum) hum
	FROM bmx280
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&bmx280).Error
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return bmx280, err
}

func (p esp32Service) GetBmx280AvgByHours(period dto.Period) ([]entities.Bmx280, error) {

	var bmx280 []entities.Bmx280
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	AVG(press) press, AVG(tempr) tempr, AVG(hum) hum
	FROM bmx280
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&bmx280).Error
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return bmx280, err
}

func (p esp32Service) GetBmx280MinByDays(period dto.Period) ([]entities.Bmx280, error) {

	var bmx280 []entities.Bmx280
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	MIN(press) press, MIN(tempr) tempr, MIN(hum) hum
	FROM bmx280
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&bmx280).Error
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return bmx280, err
}

func (p esp32Service) GetBmx280MaxByDays(period dto.Period) ([]entities.Bmx280, error) {

	var bmx280 []entities.Bmx280
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	MAX(press) press, MAX(tempr) tempr, MAX(hum) hum
	FROM bmx280
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&bmx280).Error
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return bmx280, err
}

func (p esp32Service) GetBmx280AvgByDays(period dto.Period) ([]entities.Bmx280, error) {

	var bmx280 []entities.Bmx280
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	AVG(press) press, AVG(tempr) tempr, AVG(hum) hum
	FROM bmx280
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&bmx280).Error
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return bmx280, err
}

func (p esp32Service) GetBmx280MinByMonths(period dto.Period) ([]entities.Bmx280, error) {

	var bmx280 []entities.Bmx280
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	MIN(press) press, MIN(tempr) tempr, MIN(hum) hum
	FROM bmx280
	WHERE date_time >= ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&bmx280).Error
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return bmx280, err
}

func (p esp32Service) GetBmx280MaxByMonths(period dto.Period) ([]entities.Bmx280, error) {

	var bmx280 []entities.Bmx280
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	MAX(press) press, MAX(tempr) tempr, MAX(hum) hum
	FROM bmx280
	WHERE date_time >= ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&bmx280).Error
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return bmx280, err
}

func (p esp32Service) GetBmx280AvgByMonths(period dto.Period) ([]entities.Bmx280, error) {

	var bmx280 []entities.Bmx280
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	AVG(press) press, AVG(tempr) tempr, AVG(hum) hum
	FROM bmx280
	WHERE date_time >= ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&bmx280).Error
	if err != nil {
		return nil, fmt.Errorf("error read bmx280: %w", err)
	}
	return bmx280, err
}
