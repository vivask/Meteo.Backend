package entities

import "time"

type Esp32DateTime struct {
	Esp32DateTime    string    `gorm:"esp32_date_time" json:"esp32_date_time"`
	Esp32DateTimeNow time.Time `gorm:"esp32_date_time_now" json:"esp32_date_time_now"`
	DateTime         time.Time `gorm:"date_time" json:"date_time"`
}

type Settings struct {
	ID                    uint32    `gorm:"column:id;primaryKey" json:"id"`
	SetupMode             bool      `gorm:"column:setup_mode;not null;default:false" json:"setup_mode"`
	SetupStatus           bool      `gorm:"column:setup_status;not null;default:true" json:"setup_status"`
	Reboot                bool      `gorm:"column:reboot;not null;default:false" json:"reboot"`
	Rebooted              bool      `gorm:"column:rebooted;not null;default:true" json:"rebooted"`
	MaxCh2o               float64   `gorm:"column:max_ch2o;not null;default:150" json:"max_ch2o"`
	MaxCh2oAlarm          bool      `gorm:"column:max_ch2o_alarm;not null;default:false" json:"max_ch2o_alarm"`
	MaxDs18b20            float64   `gorm:"column:max_ds18b20;not null;default:30.0" json:"max_ds18b20"`
	MinDs18b20            float64   `gorm:"column:min_ds18b20;not null;default:9.0" json:"min_ds18b20"`
	MaxDs18b20Alarm       bool      `gorm:"column:max_ds18b20_alarm;not null;default:false" json:"max_ds18b20_alarm"`
	MinDs18b20Alarm       bool      `gorm:"column:min_ds18b20_alarm;not null;default:false" json:"min_ds18b20_alarm"`
	Max6814Nh3            float64   `gorm:"column:max_6814_nh3;not null;default:5.0" json:"max_6814_nh3"`
	Max6814Co             float64   `gorm:"column:max_6814_co;not null;default:5.0" json:"max_6814_co"`
	Max6814No2            float64   `gorm:"column:max_6814_no2;not null;default:5.0" json:"max_6814_no2"`
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
	DigisparkReboot       bool      `gorm:"column:digispark_reboot;not null;default:false" json:"digispark_reboot"`
	Esp32DateTimeNow      time.Time `gorm:"column:esp32_date_time_now;not null;default:CURRENT_TIMESTAMP" json:"esp32_date_time_now"`
	UpdatedAt             time.Time `gorm:"column:updatedat;not null;default:CURRENT_TIMESTAMP" json:"date_time"`
	Bmx280Lock            bool      `gorm:"column:bmx280_lock;not null;default:false" json:"bmx280_lock"`
	Ds18b20Lock           bool      `gorm:"column:ds18b20_lock;not null;default:false" json:"ds18b20_lock"`
	Mics6814Lock          bool      `gorm:"column:mics6814_lock;not null;default:false" json:"mics6814_lock"`
	RadsensLock           bool      `gorm:"column:radsens_lock;not null;default:false" json:"radsens_lock"`
	Ze08Lock              bool      `gorm:"column:ze08_lock;not null;default:false" json:"ze08_lock"`
	Aht25Lock             bool      `gorm:"column:aht25_lock;not null;default:false" json:"aht25_lock"`
	Firmware              string    `gorm:"column:firmware;not null;default:''" json:"firmware"`
	UpgradeStatus         int       `gorm:"column:upgrade_status;not null;default:0" json:"upgrade_status"`
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
	RadsensPulse        float64   `json:"radsens_pulse"`
	MaxRadStatAlarm     bool      `json:"max_rad_stat_alarm"`
	MaxRadDynAlarm      bool      `json:"max_rad_dyn_alarm"`
	RadsensHVState      bool      `json:"radsens_hv_state"`
	RadsensSens         int       `json:"radsens_sens"`
	RadsensCreatedAt    time.Time `json:"radsens_created"`
	Ze08Ch2o            float64   `json:"ze08_ch2o"`
	MaxCh2oAlarm        bool      `json:"max_ch2o_alarm"`
	Ze08CreatedAt       time.Time `json:"ze08_created"`
	Aht25Tempr          float64   `json:"aht25_tempr"`
	Aht25Hum            float64   `json:"aht25_hum"`
	Aht25CreatedAt      time.Time `json:"aht25_created"`
	Esp32DateTimeNow    time.Time `json:"esp32_date_time_now"`
	Bmx280Lock          bool      `json:"bmx280_lock"`
	Ds18b20Lock         bool      `json:"ds18b20_lock"`
	Mics6814Lock        bool      `json:"mics6814_lock"`
	RadsensLock         bool      `json:"radsens_lock"`
	Ze08Lock            bool      `json:"ze08_lock"`
	Aht25Lock           bool      `json:"aht25_lock"`
}

type Sensors struct {
	Bmx280Lock       bool      `json:"bmx280_lock"`
	Ds18b20Lock      bool      `json:"ds18b20_lock"`
	Mics6814Lock     bool      `json:"mics6814_lock"`
	RadsensLock      bool      `json:"radsens_lock"`
	Ze08Lock         bool      `json:"ze08_lock"`
	Aht25Lock        bool      `json:"aht25_lock"`
	Esp32DateTimeNow time.Time `json:"esp32_date_time_now"`
}

type SingleModeData struct {
	Order string `json:"order"`
	Body  string `json:"body"`
}

type MeasureResult struct {
	Bme280_status       int8    `json:"bmp280_status"`
	Bme280_pressure     float64 `json:"bmp280_pressure"`
	Bme280_temperature  float64 `json:"bmp280_temperature"`
	Bme280_humidity     float64 `json:"bmp280_humidity"`
	Ds18b20_status      int8    `json:"ds18b20_status"`
	Ds18b20_address     string  `json:"ds18b20_address"`
	Ds18b20_temperature float64 `json:"ds18b20_temperature"`
	Radsens_status      int8    `json:"radsens_status"`
	Radsens_i_static    float64 `json:"radsens_i_static"`
	Radsens_i_dynamic   float64 `json:"radsens_i_dynamic"`
	Radsens_pulse       uint32  `json:"radsens_pulse"`
	Radsens_hv_state    uint8   `json:"radsens_hv_state"`
	Radsens_sens        uint8   `json:"radsens_sens"`
	Ze08_status         int8    `json:"ze08_status"`
	Ze08_ch2o           uint16  `json:"ze08_ch2o"`
	Mics6814_status     int8    `json:"mics6814_status"`
	Mics6814_co         float64 `json:"mics6814_co"`
	Mics6814_nh3        float64 `json:"mics6814_nh3"`
	Mics6814_no2        float64 `json:"mics6814_no2"`
	Aht25_status        int8    `json:"aht25_status"`
	Aht25_temperature   float64 `json:"aht25_temperature"`
	Aht25_humidity      float64 `json:"aht25_humidity"`
}
