package entities

import "time"

type Esp32DateTime struct {
	Esp32DateTime    string    `gorm:"esp32_date_time" json:"esp32_date_time"`
	Esp32DateTimeNow time.Time `gorm:"esp32_date_time_now" json:"esp32_date_time_now"`
	Cpu0Load         float64   `gorm:"column:cpu0_load" json:"cpu0_load"`
	Cpu1Load         float64   `gorm:"column:cpu1_load" json:"cpu1_load"`
	DateTime         time.Time `gorm:"date_time" json:"date_time"`
}

type Settings struct {
	ID                    uint32    `gorm:"column:id;primaryKey" json:"id"`
	ValveState            bool      `gorm:"column:valve_state;not null;default:false" json:"valve_state"`
	ValveDisable          bool      `gorm:"column:valve_disable;not null;default:false" json:"valve_disable"`
	MinTempn              float64   `gorm:"column:min_temp;not null;default:9.0" json:"min_temp"`
	MaxTemp               float64   `gorm:"column:max_temp;not null;default:12.0" json:"max_temp"`
	CCS811Baseline        int       `gorm:"column:ccs811_baseline;not null;default:0" json:"CCS811_baseline"`
	Firmware              string    `gorm:"column:firmware;not null;default:'_EMPTY_'" json:"firmware"`
	UpgradeStatus         int       `gorm:"column:upgrade_status;not null;default:1" json:"upgrade_status"`
	SetupMode             bool      `gorm:"column:setup_mode;not null;default:false" json:"setup_mode"`
	SetupStatus           bool      `gorm:"column:setup_status;not null;default:true" json:"setup_status"`
	Reboot                bool      `gorm:"column:reboot;not null;default:false" json:"reboot"`
	Rebooted              bool      `gorm:"column:rebooted;not null;default:true" json:"rebooted"`
	MaxCh2o               int       `gorm:"column:max_ch2o;not null;default:150" json:"max_ch2o"`
	MaxCh2oAlarm          bool      `gorm:"column:max_ch2o_alarm;not null;default:false" json:"max_ch2o_alarm"`
	MaxDs18b20            float64   `gorm:"column:max_ds18b20;not null;default:30.0" json:"max_ds18b20"`
	MinDs18b20            float64   `gorm:"column:min_ds18b20;not null;default:8.0" json:"min_ds18b20"`
	MaxDs18b20Alarm       bool      `gorm:"column:max_ds18b20_alarm;not null;default:false" json:"max_ds18b20_alarm"`
	MinDs18b20Alarm       bool      `gorm:"column:min_ds18b20_alarm;not null;default:false" json:"min_ds18b20_alarm"`
	Max6814Nh3            float64   `gorm:"column:max_6814_nh3;not null;default:5.0" json:"max_6814_nh3"`
	Max6814Co             float64   `gorm:"column:max_6814_co;not null;default:2000.0" json:"max_6814_co"`
	Max6814No2            float64   `gorm:"column:max_6814_no2;not null;default:20.0" json:"max_6814_no2"`
	Max6814Nh3Alarm       bool      `gorm:"column:max_6814_nh3_alarm;not null;default:false" json:"max_6814_nh3_alarm"`
	Max6814CoAlarm        bool      `gorm:"column:max_6814_co_alarm;not null;default:false" json:"max_6814_co_alarm"`
	Max6814No2Alarm       bool      `gorm:"column:max_6814_no2_alarm;not null;default:false" json:"max_6814_no2_alarm"`
	MaxRadStat            float64   `gorm:"column:max_rad_stat;not null;default:30.0" json:"max_rad_stat"`
	MaxRadDyn             float64   `gorm:"column:max_rad_dyn;not null;default:30.0" json:"max_rad_dyn"`
	MaxRadStatAlarm       bool      `gorm:"column:max_rad_stat_alarm;not null;default:false" json:"max_rad_stat_alarm"`
	MaxRadDynAlarm        bool      `gorm:"column:max_rad_dyn_alarm;not null;default:false" json:"max_rad_dyn_alarm"`
	MaxBmx280Tempr        float64   `gorm:"column:max_bmx280_tempr;not null;default:30.0" json:"max_bmx280_tempr"`
	MinBmx280Tempr        float64   `gorm:"column:min_bmx280_tempr;not null;default:-20.0" json:"min_bmx280_tempr"`
	MaxBmx280TemprAlarm   bool      `gorm:"column:max_bmx280_tempr_alarm;not null;default:false" json:"max_bmx280_tempr_alarm"`
	MinBmx280TemprAlarm   bool      `gorm:"column:min_bmx280_tempr_alarm;not null;default:false" json:"min_bmx280_tempr_alarm"`
	RadsensHVState        bool      `gorm:"column:radsens_hv_state;not null;default:false" json:"radsens_hv_state"`
	RadsensHVMode         bool      `gorm:"column:radsens_hv_mode;not null;default:true" json:"radsens_hv_mode"`
	RadsensSensitivity    int       `gorm:"column:radsens_sensitivity;not null;default:105" json:"radsens_sensitivity"`
	RadsensSensitivitySet bool      `gorm:"column:radsens_sensitivity_set;not null;default:false" json:"radsens_sensitivity_set"`
	ClearJournalEsp32     bool      `gorm:"column:clear_journal_esp32;not null;default:false" json:"clear_journal_esp32"`
	Cpu0Load              float64   `gorm:"column:cpu0_load" json:"cpu0_load"`
	Cpu1Load              float64   `gorm:"column:cpu1_load" json:"cpu1_load"`
	Esp32DateTimeNow      time.Time `gorm:"column:esp32_date_time_now;not null;default:CURRENT_TIMESTAMP" json:"esp32_date_time_now"`
	UpdatedAt             time.Time `gorm:"column:updatedat;not null;default:CURRENT_TIMESTAMP" json:"date_time"`
}

func (Settings) TableName() string {
	return "settings"
}

type HomePage struct {
	Bmx280Press         float64   `json:"bmx280_press"`
	Bmx280Tempr         float64   `json:"bmx280_tempr"`
	Bmx280Hum           float64   `json:"bmx280_hum"`
	MaxBmx280TemprAlarm bool      `json:"max_bmx280_tempr_alarm"`
	MinBmx280TemprAlarm bool      `json:"min_bmx280_tempr_alarm"`
	Bmx280CreatedAt     time.Time `json:"bmx280_created"`
	Ds18b20Tempr        float64   `json:"ds18b20_tempr"`
	Ds18b20CreatedAt    time.Time `json:"ds18b20_created"`
	MaxDs18b20Alarm     bool      `json:"max_ds18b20_alarm"`
	MinDs18b20Alarm     bool      `json:"min_ds18b20_alarm"`
	Mics6814No2         float64   `json:"mics6814_no2"`
	Mics6814Nh3         float64   `json:"mics6814_nh3"`
	Mics6814Co          float64   `json:"mics6814_co"`
	Max6814Nh3Alarm     bool      `json:"max_6814_nh3_alarm"`
	Max6814CoAlarm      bool      `json:"max_6814_co_alarm"`
	Max6814No2Alarm     bool      `json:"max_6814_no2_alarm"`
	Mics6814CreatedAt   time.Time `json:"mics6814_created"`
	RadsensDynamic      float64   `json:"radsens_dynamic"`
	RadsensStatic       float64   `json:"radsens_static"`
	RadsensPulse        int       `json:"radsens_pulse"`
	MaxRadStatAlarm     bool      `json:"max_rad_stat_alarm"`
	MaxRadDynAlarm      bool      `json:"max_rad_dyn_alarm"`
	RadsensHVState      bool      `json:"radsens_hv_state"`
	RadsensCreatedAt    time.Time `json:"radsens_created"`
	Ze08Ch2o            int       `json:"ze08_ch2o"`
	MaxCh2oAlarm        bool      `json:"max_ch2o_alarm"`
	Ze08CreatedAt       time.Time `json:"ze08_created"`
	Esp32DateTimeNow    time.Time `json:"esp32_date_time_now"`
}
