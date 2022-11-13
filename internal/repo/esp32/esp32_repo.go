package repo

import (
	"meteo/internal/dto"
	"meteo/internal/entities"

	"gorm.io/gorm"
)

// Esp32Service api controller of produces
type Esp32Service interface {
	GetSettings() (*entities.Settings, error)
	SetSettings(s *entities.Settings) error
	AddLoging(msg, t, dts interface{}) error
	AddDs18b20(tempr, dts interface{}) error
	AddBme280(press, tempr, hum, dts interface{}) error
	AddRadsens(dyn, stat, pl, dts interface{}) error
	AddZe08ch2o(ch2o, dts interface{}) error
	AddMics6814(co, no2, nh3, dts interface{}) error
	SetEsp32Settings(cpu0L, cpu1L, dti interface{}) (*entities.Settings, error)
	SetHVRadsens(state interface{}) error
	SetSensRadsens(sens interface{}) error
	JournalClear() error
	SetJournaCleared() error
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
	Ds18b20TemperatureChk() error
	Ze08ch2oChk() error
}

type esp32Service struct {
	db *gorm.DB
}

// NewProductService get product service instance
func NewEsp32Service(db *gorm.DB) Esp32Service {
	return &esp32Service{db}
}
