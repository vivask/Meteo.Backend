package repo

import (
	"fmt"
	"meteo/internal/entities"
)

func (p esp32Service) GetHomePageData() (*entities.HomePage, error) {
	bme280, err := p.GetLastBmx280()
	if err != nil {
		return nil, fmt.Errorf("can't read bme280: %w", err)
	}
	ds18b20, err := p.GetLastDs18b20()
	if err != nil {
		return nil, fmt.Errorf("can't read ds18b20: %w", err)
	}
	mics6814, err := p.GetLastMics6814()
	if err != nil {
		return nil, fmt.Errorf("can't read mics6814: %w", err)
	}
	radsens, err := p.GetLastRadsens()
	if err != nil {
		return nil, fmt.Errorf("can't read radsens: %w", err)
	}
	ze08ch2o, err := p.GetLastZe08ch2o()
	if err != nil {
		return nil, fmt.Errorf("can't read radsens: %w", err)
	}
	aht25, err := p.GetLastAht25()
	if err != nil {
		return nil, fmt.Errorf("can't read radsens: %w", err)
	}
	set, err := p.GetSettings()
	if err != nil {
		return nil, fmt.Errorf("can't read radsens: %w", err)
	}
	hp := &entities.HomePage{
		Bmx280Press:         bme280.Press / 133,
		Bmx280Tempr:         bme280.Tempr,
		Bmx280Hum:           bme280.Hum,
		MaxBmx280TemprAlarm: set.MaxBmx280TemprAlarm,
		MinBmx280TemprAlarm: set.MinBmx280TemprAlarm,
		Bmx280CreatedAt:     bme280.CreatedAt,
		Ds18b20Tempr:        ds18b20.Tempr,
		MaxDs18b20Alarm:     set.MaxDs18b20Alarm,
		MinDs18b20Alarm:     set.MinDs18b20Alarm,
		Ds18b20CreatedAt:    ds18b20.CreatedAt,
		Mics6814No2:         mics6814.No2,
		Mics6814Nh3:         mics6814.Nh3,
		Mics6814Co:          mics6814.Co,
		Max6814Nh3Alarm:     set.Max6814Nh3Alarm,
		Max6814CoAlarm:      set.Max6814CoAlarm,
		Max6814No2Alarm:     set.Max6814No2Alarm,
		Mics6814CreatedAt:   mics6814.CreatedAt,
		RadsensDynamic:      radsens.Dynamic,
		RadsensStatic:       radsens.Static,
		RadsensPulse:        radsens.Pulse,
		MaxRadStatAlarm:     set.MaxRadStatAlarm,
		MaxRadDynAlarm:      set.MaxRadDynAlarm,
		RadsensHVState:      set.RadsensHVState,
		RadsensSens:         set.RadsensSensitivity,
		RadsensCreatedAt:    radsens.CreatedAt,
		Ze08Ch2o:            ze08ch2o.Ch2o,
		MaxCh2oAlarm:        set.MaxCh2oAlarm,
		Ze08CreatedAt:       ze08ch2o.CreatedAt,
		Aht25Tempr:          aht25.Tempr,
		Aht25Hum:            aht25.Hum,
		Aht25CreatedAt:      aht25.CreatedAt,
		Esp32DateTimeNow:    set.Esp32DateTimeNow,
		Bmx280Lock:          set.Bmx280Lock,
		Ds18b20Lock:         set.Ds18b20Lock,
		Mics6814Lock:        set.Mics6814Lock,
		RadsensLock:         set.RadsensLock,
		Ze08Lock:            set.Ze08Lock,
		Aht25Lock:           set.Aht25Lock,
	}
	return hp, nil
}
