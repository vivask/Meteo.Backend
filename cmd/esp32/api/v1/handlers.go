package v1

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	UPLOAD                  = "./firmware"
	WATCHDOG_FIRMWARE_RETRY = 60
	WATCHDOG_FIRMWARE_TIMER = 2
	OK                      = 0
)

var (
	RADSENS_HV_STATE    uint8 = 0
	RADSENS_SENSITIVITY uint8 = 0
)

func (p esp32API) UploadFirmware(c *gin.Context) {

	file, err := c.FormFile("firmware")
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	// Upload the file to specific dst.
	dst := fmt.Sprintf("%s/%s", UPLOAD, file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p esp32API) GetSettings(c *gin.Context) {
	settings, err := p.repo.GetSettings()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, &settings)
}

func (p esp32API) set_measure(data *entities.MeasureResult, c *gin.Context) {

	state, err := p.repo.GetSensorsState()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	// log.Infof("DATA: %v", data)
	if data.Bme280_status == OK && !state.Bmx280Lock {
		if err := p.repo.AddBme280(data.Bme280_pressure, data.Bme280_temperature, data.Bme280_humidity); err != nil {
			c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
			return
		}
	}

	if data.Ds18b20_status == OK && !state.Ds18b20Lock && data.Ds18b20_temperature != 0.0 {
		if err := p.repo.AddDs18b20(data.Ds18b20_temperature); err != nil {
			c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
			return
		}
	}

	if data.Radsens_status == OK && !state.RadsensLock {
		if RADSENS_HV_STATE != data.Radsens_hv_state {
			RADSENS_HV_STATE = data.Radsens_hv_state
			if err := p.repo.ResetRadsensHV(RADSENS_HV_STATE); err != nil {
				c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
				return
			}
		}

		if RADSENS_SENSITIVITY != data.Radsens_sens {
			RADSENS_SENSITIVITY = data.Radsens_sens
			if err := p.repo.ResetRadsensSens(RADSENS_SENSITIVITY); err != nil {
				c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
				return
			}
		}

		if err := p.repo.AddRadsens(data.Radsens_i_dynamic, data.Radsens_i_static, data.Radsens_pulse); err != nil {
			c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
			return
		}
	}

	if data.Ze08_status == OK && !state.Ze08Lock {
		if err := p.repo.AddZe08ch2o(data.Ze08_ch2o); err != nil {
			c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
			return
		}
	}

	if data.Mics6814_status == OK && !state.Mics6814Lock {
		if err := p.repo.AddMics6814(data.Mics6814_co, data.Mics6814_no2, data.Mics6814_nh3); err != nil {
			c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
			return
		}
	}

	if data.Aht25_status == OK && !state.Aht25Lock {
		if err := p.repo.AddAht25(data.Aht25_temperature, data.Aht25_humidity); err != nil {
			c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
			return
		}
	}

	c.Status(http.StatusOK)
}

func (p esp32API) SetMeashure(c *gin.Context) {

	data := entities.MeasureResult{}
	if err := c.ShouldBind(&data); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	p.set_measure(&data, c)
}

func (p esp32API) GetOrders(c *gin.Context) {
	settings, err := p.repo.GetSettings()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	type Orders struct {
		SetupMode          uint8  `json:"ap"`
		Reboot             uint8  `json:"reboot"`
		RadsensHVMode      uint8  `json:"r_mode"`
		RadsensHVState     uint8  `json:"r_state"`
		RadsensSensitivity uint8  `json:"r_sens"`
		ClearJournalEsp32  uint8  `json:"clear_j"`
		DigisparkReboot    uint8  `json:"dg_reboot"`
		Firmware           string `json:"firmware"`
	}

	orders := &Orders{
		SetupMode:          toUint8(settings.SetupMode),
		Reboot:             toUint8(settings.Reboot),
		RadsensHVMode:      toUint8(settings.RadsensHVMode),
		RadsensHVState:     toUint8(settings.RadsensHVState),
		RadsensSensitivity: uint8(settings.RadsensSensitivity),
		ClearJournalEsp32:  toUint8(settings.ClearJournalEsp32),
		DigisparkReboot:    toUint8(settings.DigisparkReboot),
		Firmware:           settings.Firmware,
	}

	err = p.repo.ResetOrders()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (p esp32API) DelayedReset() {
	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	err := p.repo.ResetOrders()
	// 	if err != nil {
	// 		log.Error(err)
	// 	}
	// }()
}

func (p esp32API) ResetAccessPoint(c *gin.Context) {
	err := p.repo.ResetAccessPoint()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	p.DelayedReset()
	c.Status(http.StatusOK)
}

func (p esp32API) ResetStm32(c *gin.Context) {
	err := p.repo.ResetStm32()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	p.DelayedReset()
	c.Status(http.StatusOK)
}

func (p esp32API) ResetRadsens(c *gin.Context) {
	err := p.repo.ResetRadsens()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	p.DelayedReset()
	c.Status(http.StatusOK)
}

func (p esp32API) ResetJournal(c *gin.Context) {
	err := p.repo.ResetJournal()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	p.DelayedReset()
	c.Status(http.StatusOK)
}

func (p esp32API) ResetAvr(c *gin.Context) {
	err := p.repo.ResetAvr(false)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	p.DelayedReset()
	c.Status(http.StatusOK)
}

func (p esp32API) SetLoggingEsp32(c *gin.Context) {
	data := entities.Logging{}
	if err := c.ShouldBind(&data); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	err := p.repo.AddLoging(data.Message, data.Type, data.CreatedAt)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}
