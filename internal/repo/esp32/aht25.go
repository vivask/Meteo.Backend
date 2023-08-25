package repo

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/utils"
	"time"

	"gorm.io/gorm"
)

func (p esp32Service) GetLastAht25() (*entities.Aht25, error) {
	var aht25 entities.Aht25
	err := p.db.Order("date_time desc").Last(&aht25).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entities.Aht25{}, nil
		}
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	return &aht25, err
}

func (p esp32Service) AddAht25(temperature, humidity float64) error {
	if isLockedAht25() {
		return nil
	}

	aht25 := entities.Aht25{ID: utils.HashTime(time.Now()), Tempr: temperature, Hum: humidity}
	err := p.db.Create(&aht25).Error
	if err != nil {
		return fmt.Errorf("insert aht25: %v, error: %w", aht25, err)
	}

	return err
}

func (p esp32Service) GetAht25MinByHours(period dto.Period) ([]entities.Aht25, error) {

	var aht25 []entities.Aht25
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	MIN(tempr) tempr, MIN(hum) hum
	FROM aht25
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&aht25).Error
	if err != nil {
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	return aht25, err
}

func (p esp32Service) GetAht25MaxByHours(period dto.Period) ([]entities.Aht25, error) {

	var aht25 []entities.Aht25
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	MAX(tempr) tempr, MAX(hum) hum
	FROM aht25
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&aht25).Error
	if err != nil {
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	return aht25, err
}

func (p esp32Service) GetAht25AvgByHours(period dto.Period) ([]entities.Aht25, error) {

	var aht25 []entities.Aht25
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	AVG(tempr) tempr, AVG(hum) hum
	FROM aht25
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&aht25).Error
	if err != nil {
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	return aht25, err
}

func (p esp32Service) GetAht25MinByDays(period dto.Period) ([]entities.Aht25, error) {

	var aht25 []entities.Aht25
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	MIN(tempr) tempr, MIN(hum) hum
	FROM aht25
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&aht25).Error
	if err != nil {
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	return aht25, err
}

func (p esp32Service) GetAht25MaxByDays(period dto.Period) ([]entities.Aht25, error) {

	var aht25 []entities.Aht25
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	MAX(tempr) tempr, MAX(hum) hum
	FROM aht25
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&aht25).Error
	if err != nil {
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	return aht25, err
}

func (p esp32Service) GetAht25AvgByDays(period dto.Period) ([]entities.Aht25, error) {

	var aht25 []entities.Aht25
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	AVG(tempr) tempr, AVG(hum) hum
	FROM aht25
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&aht25).Error
	if err != nil {
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	return aht25, err
}

func (p esp32Service) GetAht25MinByMonths(period dto.Period) ([]entities.Aht25, error) {

	var aht25 []entities.Aht25
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	MIN(tempr) tempr, MIN(hum) hum
	FROM aht25
	WHERE date_time >= ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&aht25).Error
	if err != nil {
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	return aht25, err
}

func (p esp32Service) GetAht25MaxByMonths(period dto.Period) ([]entities.Aht25, error) {

	var aht25 []entities.Aht25
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	MAX(tempr) tempr, MAX(hum) hum
	FROM aht25
	WHERE date_time >= ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&aht25).Error
	if err != nil {
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	return aht25, err
}

func (p esp32Service) GetAht25AvgByMonths(period dto.Period) ([]entities.Aht25, error) {

	var aht25 []entities.Aht25
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	AVG(tempr) tempr, AVG(hum) hum
	FROM aht25
	WHERE date_time >= ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&aht25).Error
	if err != nil {
		return nil, fmt.Errorf("error read aht25: %w", err)
	}
	return aht25, err
}
