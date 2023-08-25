package repo

import (
	"meteo/internal/dto"
	"meteo/internal/entities"
	"time"

	"gorm.io/gorm"
)

// Esp32Service api controller of produces
type Esp32Service interface {
	GetSettings() (*entities.Settings, error)
	SetSettings(s *entities.Settings) error
	AddLoging(message, msgType string, time time.Time) error
	AddDs18b20(temperature float64) error
	AddBme280(pressure, temperature, humidity float64) error
	AddRadsens(dynamic, static float64, pulse uint32) error
	AddZe08ch2o(ch2o uint16) error
	AddMics6814(co, no2, nh3 float64) error
	SetEsp32Settings(cpu0L, cpu1L, dti interface{}) (*entities.Settings, error)
	SetHVRadsens(state interface{}) error
	SetSensRadsens(sens interface{}) error
	JournalClear() error
	// SetJournaCleared() error
	GetAllLoging(pageable dto.Pageable) ([]entities.Logging, error)
	SetAccesPointMode() error
	SetSTAMode() error
	Esp32Reboot() error
	Esp32Rebooted() error

	UpgradeEsp32(fName string) error
	GetUpgradeStatus() (*entities.Settings, error)
	SuccessUpgrade() error
	TerminateUpgrade() error

	GetLastBmx280() (*entities.Bmx280, error)
	GetLastDs18b20() (*entities.Ds18b20, error)
	GetLastMics6814() (*entities.Mics6814, error)
	GetLastRadsens() (*entities.Radsens, error)
	GetLastZe08ch2o() (*entities.Ze08ch2o, error)
	GetHomePageData() (*entities.HomePage, error)
	Mics6814CoChk() error
	Mics6814No2Chk() error
	Mics6814Nh3Chk() error
	Bme280TemperatureChk() error
	RadsensStaticChk() error
	RadsensDynamicChk() error
	RadsensHVSet() error
	RadsensSetSens(val uint) error
	Ds18b20TemperatureChk() error
	Ze08ch2oChk() error

	GetBmx280MinByHours(period dto.Period) ([]entities.Bmx280, error)
	GetBmx280MaxByHours(period dto.Period) ([]entities.Bmx280, error)
	GetBmx280AvgByHours(period dto.Period) ([]entities.Bmx280, error)
	GetBmx280MinByDays(period dto.Period) ([]entities.Bmx280, error)
	GetBmx280MaxByDays(period dto.Period) ([]entities.Bmx280, error)
	GetBmx280AvgByDays(period dto.Period) ([]entities.Bmx280, error)
	GetBmx280MinByMonths(period dto.Period) ([]entities.Bmx280, error)
	GetBmx280MaxByMonths(period dto.Period) ([]entities.Bmx280, error)
	GetBmx280AvgByMonths(period dto.Period) ([]entities.Bmx280, error)

	GetZe08MinByHours(period dto.Period) ([]entities.Ze08ch2o, error)
	GetZe08MaxByHours(period dto.Period) ([]entities.Ze08ch2o, error)
	GetZe08AvgByHours(period dto.Period) ([]entities.Ze08ch2o, error)
	GetZe08MinByDays(period dto.Period) ([]entities.Ze08ch2o, error)
	GetZe08MaxByDays(period dto.Period) ([]entities.Ze08ch2o, error)
	GetZe08AvgByDays(period dto.Period) ([]entities.Ze08ch2o, error)
	GetZe08MinByMonths(period dto.Period) ([]entities.Ze08ch2o, error)
	GetZe08MaxByMonths(period dto.Period) ([]entities.Ze08ch2o, error)
	GetZe08AvgByMonths(period dto.Period) ([]entities.Ze08ch2o, error)

	GetDs18b20MinByHours(period dto.Period) ([]entities.Ds18b20, error)
	GetDs18b20MaxByHours(period dto.Period) ([]entities.Ds18b20, error)
	GetDs18b20AvgByHours(period dto.Period) ([]entities.Ds18b20, error)
	GetDs18b20MinByDays(period dto.Period) ([]entities.Ds18b20, error)
	GetDs18b20MaxByDays(period dto.Period) ([]entities.Ds18b20, error)
	GetDs18b20AvgByDays(period dto.Period) ([]entities.Ds18b20, error)
	GetDs18b20MinByMonths(period dto.Period) ([]entities.Ds18b20, error)
	GetDs18b20MaxByMonths(period dto.Period) ([]entities.Ds18b20, error)
	GetDs18b20AvgByMonths(period dto.Period) ([]entities.Ds18b20, error)

	GetMics6814MinByHours(period dto.Period) ([]entities.Mics6814, error)
	GetMics6814MaxByHours(period dto.Period) ([]entities.Mics6814, error)
	GetMics6814AvgByHours(period dto.Period) ([]entities.Mics6814, error)
	GetMics6814MinByDays(period dto.Period) ([]entities.Mics6814, error)
	GetMics6814MaxByDays(period dto.Period) ([]entities.Mics6814, error)
	GetMics6814AvgByDays(period dto.Period) ([]entities.Mics6814, error)
	GetMics6814MinByMonths(period dto.Period) ([]entities.Mics6814, error)
	GetMics6814MaxByMonths(period dto.Period) ([]entities.Mics6814, error)
	GetMics6814AvgByMonths(period dto.Period) ([]entities.Mics6814, error)

	GetRadsensMinByHours(period dto.Period) ([]entities.Radsens, error)
	GetRadsensMaxByHours(period dto.Period) ([]entities.Radsens, error)
	GetRadsensAvgByHours(period dto.Period) ([]entities.Radsens, error)
	GetRadsensMinByDays(period dto.Period) ([]entities.Radsens, error)
	GetRadsensMaxByDays(period dto.Period) ([]entities.Radsens, error)
	GetRadsensAvgByDays(period dto.Period) ([]entities.Radsens, error)
	GetRadsensMinByMonths(period dto.Period) ([]entities.Radsens, error)
	GetRadsensMaxByMonths(period dto.Period) ([]entities.Radsens, error)
	GetRadsensAvgByMonths(period dto.Period) ([]entities.Radsens, error)

	GetLastAht25() (*entities.Aht25, error)
	AddAht25(temperature, humidity float64) error
	GetAht25MinByHours(period dto.Period) ([]entities.Aht25, error)
	GetAht25MaxByHours(period dto.Period) ([]entities.Aht25, error)
	GetAht25AvgByHours(period dto.Period) ([]entities.Aht25, error)
	GetAht25MinByDays(period dto.Period) ([]entities.Aht25, error)
	GetAht25MaxByDays(period dto.Period) ([]entities.Aht25, error)
	GetAht25AvgByDays(period dto.Period) ([]entities.Aht25, error)
	GetAht25MinByMonths(period dto.Period) ([]entities.Aht25, error)
	GetAht25MaxByMonths(period dto.Period) ([]entities.Aht25, error)
	GetAht25AvgByMonths(period dto.Period) ([]entities.Aht25, error)

	ResetAccessPoint() error
	ResetStm32() error
	ResetRadsens() error
	ResetRadsensHV(val uint8) error
	ResetRadsensSens(val uint8) error
	ResetJournal() error
	ResetAvr(val bool) error
	ResetOrders() error

	GetSensorsState() (*entities.Sensors, error)
	LockBmx280(lock bool) error
	LockDs18b20(lock bool) error
	LockRadsens(lock bool) error
	LockMics6814(lock bool) error
	LockZe08(lock bool) error
	LockAht25(lock bool) error
}

type esp32Service struct {
	db *gorm.DB
}

// NewProductService get product service instance
func NewEsp32Service(db *gorm.DB) Esp32Service {
	return &esp32Service{db}
}
