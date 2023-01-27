package v1

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/errors"
	"meteo/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	UPLOAD                  = "./firmware"
	WATCHDOG_FIRMWARE_RETRY = 60
	WATCHDOG_FIRMWARE_TIMER = 2
)

var firmwareFile string

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

func (p esp32API) Handler(c *gin.Context) {

	msg := map[string]interface{}{}
	if err := c.ShouldBind(&msg); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	Esp32MAC := msg["DEVICE"].(string)

	if !config.Default.Esp32.Check || Esp32MAC == config.Default.Esp32.Mac {
		switch msg["ORDER"] {
		case "LOGING":
			err := p.repo.AddLoging(msg["MESSAGE"],
				msg["TYPE"], msg["DATE_TIME"])
			if err != nil {
				log.Errorf("add loging error: %v", err)
			}
		case "DS18B20":
			err := p.repo.AddDs18b20(msg["TEMPR"], msg["DATE_TIME"])
			if err != nil {
				log.Errorf("add ds18b20 error: %v", err)
			}
		case "BME280":
			err := p.repo.AddBme280(msg["PRESS"], msg["TEMPR"], msg["HUM"], msg["DATE_TIME"])
			if err != nil {
				log.Errorf("add bmx280 error: %v", err)
			}
		case "RADSENS":
			err := p.repo.AddRadsens(msg["RID"], msg["RIS"], msg["PULSE"], msg["DATE_TIME"])
			if err != nil {
				log.Errorf("add radsens error: %v", err)
			}
		case "ZE08CH2O":
			err := p.repo.AddZe08ch2o(msg["CH2O"], msg["DATE_TIME"])
			if err != nil {
				log.Errorf("add ze08ch2o error: %v", err)
			}
		case "6814":
			err := p.repo.AddMics6814(msg["CO"], msg["NO2"], msg["NH3"], msg["DATE_TIME"])
			if err != nil {
				log.Errorf("add mics6814 error: %v", err)
			}
		case "GET_SETTINGS":
			if res, err := p.repo.SetEsp32Settings(msg["CPU0"], msg["CPU1"], msg["DATE_TIME"]); err != nil {
				log.Errorf("get settings error: %v", err)
			} else {
				var data struct {
					ValveState         uint8   `json:"valve_state"`
					CCS811Baseline     uint8   `json:"ccs811_baseline"`
					MinTempn           float64 `json:"min_temp"`
					MaxTemp            float64 `json:"max_temp"`
					ValveDisable       uint8   `json:"valve_disable"`
					SetupMode          uint8   `json:"setup_mode"`
					Reboot             uint8   `json:"reboot"`
					RadsensHVMode      uint8   `json:"radsens_hv_mode"`
					RadsensHVState     uint8   `json:"radsens_hv_state"`
					RadsensSensitivity uint8   `json:"radsens_sensitivity"`
					Firmware           string  `json:"firmware"`
					ClearJournalEsp32  uint8   `json:"clear_journal_esp32"`
				}
				data.ValveState = toUint8(res.ValveState)
				data.CCS811Baseline = uint8(res.CCS811Baseline)
				data.MinTempn = res.MinTemp
				data.MaxTemp = res.MaxTemp
				data.ValveDisable = toUint8(res.ValveDisable)
				data.Firmware = res.Firmware
				data.SetupMode = toUint8(res.SetupMode)
				data.Reboot = toUint8(res.Reboot)
				data.RadsensHVMode = toUint8(res.RadsensHVMode)
				data.RadsensHVState = toUint8(res.RadsensHVState)
				data.RadsensSensitivity = uint8(res.RadsensSensitivity)
				data.ClearJournalEsp32 = toUint8(res.ClearJournalEsp32)
				c.Header("Content-Type", "application/json")
				c.JSON(http.StatusOK, &data)
			}
		case "RADSENS_HV_SET":
			err := p.repo.SetHVRadsens(msg["STATE"])
			if err != nil {
				log.Errorf("set radsens HV error: %v", err)
			}
		case "RADSENS_SENSITIVITY_SET":
			err := p.repo.SetSensRadsens(msg["SENSITIVITY"])
			if err != nil {
				log.Errorf("set radsens sensitivity error: %v", err)
			}
		case "AP_MODE_ON":
			err := p.repo.SetSTAMode()
			if err != nil {
				log.Errorf("esp32 set acceess point mode error: %v", err)
			}
		case "REBOOTED":
			err := p.repo.Esp32Rebooted()
			if err != nil {
				log.Errorf("esp32 reboot error: %v", err)
			}
		case "UPGRADE_FAIL":
			err := p.repo.TerminateUpgrade()
			if err != nil {
				log.Errorf("esp32 upgrade fail: %v", err)
			}
			log.Warningf("Esp32 upgrade fail")
			/*err = os.Remove(firmwareFile)
			if err != nil {
				log.Error(err)
			}*/
		case "UPGRADE_SUCCESS":
			err := p.repo.SuccessUpgrade()
			if err != nil {
				log.Error(err)
			}
		case "JOURNAL_CLEARED":
			err := p.repo.SetJournaCleared()
			if err != nil {
				log.Errorf("esp32 journal clear fail: %v", err)
			}
		default:
			log.Errorf("Unknown ESP32 order: %s", msg["ORDER"])
		}
	} else {
		log.Errorf("Unknown device: %s", msg["DEVICE"])
	}
}
