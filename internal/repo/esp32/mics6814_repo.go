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

func (p esp32Service) AddMics6814(co, no2, nh3 float64) error {
	if isLockedMics6814() {
		return nil
	}

	mics6814 := entities.Mics6814{ID: utils.HashTime(time.Now()), Co: co, No2: no2, Nh3: nh3}
	err := p.db.Create(&mics6814).Error
	if err != nil {
		return fmt.Errorf("insert mics6814: %v, error: %w", mics6814, err)
	}
	settings, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}
	message := fmt.Sprintf("MICS6814: Превышена концентрация CO: %f ppm", co)
	if co > settings.Max6814Co && !settings.Max6814CoAlarm {
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
	message = fmt.Sprintf("MICS6814: Превышена концентрация NO2: %f ppm", no2)
	if no2 > settings.Max6814No2 && !settings.Max6814No2Alarm {
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
	message = fmt.Sprintf("MICS6814: Превышена концентрация NH3: %f ppm", nh3)
	if nh3 > settings.Max6814Nh3 && !settings.Max6814Nh3Alarm {
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

func (p esp32Service) GetMics6814MinByHours(period dto.Period) ([]entities.Mics6814, error) {

	var mics6814 []entities.Mics6814
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	MIN(no2) no2, MIN(nh3) nh3, MIN(co) co
	FROM mics6814
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&mics6814).Error
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return mics6814, err
}

func (p esp32Service) GetMics6814MaxByHours(period dto.Period) ([]entities.Mics6814, error) {

	var mics6814 []entities.Mics6814
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	MAX(no2) no2, MAX(nh3) nh3, MAX(co) co
	FROM mics6814
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&mics6814).Error
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return mics6814, err
}

func (p esp32Service) GetMics6814AvgByHours(period dto.Period) ([]entities.Mics6814, error) {

	var mics6814 []entities.Mics6814
	query := `SELECT DATE_TRUNC('hour', date_time) as gdate,
	AVG(no2) no2, AVG(nh3) nh3, AVG(co) co
	FROM mics6814
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&mics6814).Error
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return mics6814, err
}

func (p esp32Service) GetMics6814MinByDays(period dto.Period) ([]entities.Mics6814, error) {

	var mics6814 []entities.Mics6814
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	MIN(no2) no2, MIN(nh3) nh3, MIN(co) co
	FROM mics6814
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&mics6814).Error
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return mics6814, err
}

func (p esp32Service) GetMics6814MaxByDays(period dto.Period) ([]entities.Mics6814, error) {

	var mics6814 []entities.Mics6814
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	MAX(no2) no2, MAX(nh3) nh3, MAX(co) co
	FROM mics6814
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&mics6814).Error
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return mics6814, err
}

func (p esp32Service) GetMics6814AvgByDays(period dto.Period) ([]entities.Mics6814, error) {

	var mics6814 []entities.Mics6814
	query := `SELECT DATE_TRUNC('day', date_time) as gdate,
	AVG(no2) no2, AVG(nh3) nh3, AVG(co) co
	FROM mics6814
	WHERE date_time > ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&mics6814).Error
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return mics6814, err
}

func (p esp32Service) GetMics6814MinByMonths(period dto.Period) ([]entities.Mics6814, error) {

	var mics6814 []entities.Mics6814
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	MIN(no2) no2, MIN(nh3) nh3, MIN(co) co
	FROM mics6814
	WHERE date_time >= ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&mics6814).Error
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return mics6814, err
}

func (p esp32Service) GetMics6814MaxByMonths(period dto.Period) ([]entities.Mics6814, error) {

	var mics6814 []entities.Mics6814
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	MAX(no2) no2, MAX(nh3) nh3, MAX(co) co
	FROM mics6814
	WHERE date_time >= ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&mics6814).Error
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return mics6814, err
}

func (p esp32Service) GetMics6814AvgByMonths(period dto.Period) ([]entities.Mics6814, error) {

	var mics6814 []entities.Mics6814
	query := `SELECT DATE_TRUNC('month', date_time) as gdate,
	AVG(no2) no2, AVG(nh3) nh3, AVG(co) co
	FROM mics6814
	WHERE date_time >= ? AND date_time <= ?
	GROUP BY gdate
	ORDER BY gdate`
	err := p.db.Raw(query, period.Begin, period.End).Scan(&mics6814).Error
	if err != nil {
		return nil, fmt.Errorf("error read mics6814: %w", err)
	}
	return mics6814, err
}
