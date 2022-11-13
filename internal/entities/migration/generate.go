package entities

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/entities"
	"meteo/internal/log"
	"meteo/internal/utils"

	"gorm.io/gorm"
)

func CreateSettings(db *gorm.DB) error {

	if db.First(&entities.Settings{}).Error == gorm.ErrRecordNotFound {
		err := db.Create(&entities.Settings{ID: 1}).Error
		if err != nil {
			return fmt.Errorf("insert settings error: %w", err)
		}
		log.Debugf("save settings: initialize")
	}

	return nil
}

func CreateExecutors(db *gorm.DB) error {
	if db.First(&entities.Executors{}).Error == gorm.ErrRecordNotFound {
		executors := []entities.Executors{
			{ID: "Master"},
			{ID: "Leader"},
			{ID: "Slave"},
			{ID: "All"},
		}
		err := db.Create(&executors).Error
		if err != nil {
			return fmt.Errorf("insert executors error: %w", err)
		}
		log.Debugf("save executors: initialize")
	}
	return nil
}

func CreatePeriods(db *gorm.DB) error {

	if db.First(&entities.Periods{}).Error == gorm.ErrRecordNotFound {
		periods := []entities.Periods{
			{ID: "once", Name: "Once", Idx: 1},
			{ID: "second", Name: "Second", Idx: 2},
			{ID: "minute", Name: "Minute", Idx: 3},
			{ID: "hour", Name: "Hour", Idx: 4},
			{ID: "day", Name: "Day", Idx: 5},
			{ID: "week", Name: "Week", Idx: 6},
			{ID: "day_of_week", Name: "Day of week", Idx: 7},
			{ID: "month", Name: "Month", Idx: 8},
			{ID: "year", Name: "Year", Idx: 9},
		}
		err := db.Create(&periods).Error
		if err != nil {
			return fmt.Errorf("insert periods error: %w", err)
		}
		log.Debugf("save periods: initialize")
	}
	return nil

}

func CreateDays(db *gorm.DB) error {

	if db.First(&entities.Days{}).Error == gorm.ErrRecordNotFound {
		days := []entities.Days{
			{ID: 1, Name: "Monday"},
			{ID: 2, Name: "Tuesday"},
			{ID: 3, Name: "Wednesday"},
			{ID: 4, Name: "Thursday"},
			{ID: 5, Name: "Friday"},
			{ID: 6, Name: "Saturday"},
			{ID: 7, Name: "Sunday"},
		}
		err := db.Create(&days).Error
		if err != nil {
			return fmt.Errorf("insert days error: %w", err)
		}
		log.Debugf("save days: initialize")
	}
	return nil

}

func CreateAccesList(db *gorm.DB) error {

	if db.First(&entities.AccesList{}).Error == gorm.ErrRecordNotFound {
		lists := []entities.AccesList{
			{ID: "tovpn"},
			{ID: "local"},
		}
		err := db.Create(&lists).Error
		if err != nil {
			return fmt.Errorf("insert access_lists error: %w", err)
		}
		log.Debugf("save access_lists: initialize")
	}
	return nil

}

func CreateGitService(db *gorm.DB) error {

	if db.First(&entities.GitService{}).Error == gorm.ErrRecordNotFound {
		lists := []entities.GitService{
			{ID: "backup_mikrotik", Name: "Backup mikrotik routers"},
		}
		err := db.Create(&lists).Error
		if err != nil {
			return fmt.Errorf("insert git_service error: %w", err)
		}
		log.Debugf("save git_service: initialize")
	}
	return nil

}

func CreateSyncTypes(db *gorm.DB) error {

	if db.First(&entities.SyncTypes{}).Error == gorm.ErrRecordNotFound {
		types := []entities.SyncTypes{
			{ID: "sync", Note: "Synchronization"},
			{ID: "replace", Note: "Replace"},
		}
		err := db.Create(&types).Error
		if err != nil {
			return fmt.Errorf("insert git_service error: %w", err)
		}
		log.Debugf("save git_service: initialize")
	}

	return nil

}

func CreateTasks(db *gorm.DB) error {
	if db.First(&entities.Tasks{}).Error == gorm.ErrRecordNotFound {
		tasks := []entities.Tasks{
			{
				ID:     "telegram",
				Name:   "Telegram messanger",
				Api:    "/messanger/telegram/schedule",
				Note:   "Бот отправки сообщений в телеграм",
				Params: []entities.TaskParams{{ID: 1, Name: "msg", Note: "Message body", TaskID: "telegram"}},
			},
			{
				ID:   "syncdb",
				Name: "Synchronyze ESP32 tables",
				Api:  "/database/sync",
				Note: "Синхронизация таблиц контроллера",
			},
			{
				ID:   "storMount",
				Name: "Storage mount",
				Api:  "/xu4/storage/mount",
				Note: "Проверка подключения хранилища",
			},
			{
				ID:   "mikrotiks",
				Name: "Mikrotiks Backup",
				Api:  "/xu4/mikrotiks/backup",
				Note: "Бэкап конфигураций и образов роутеров mikrotik",
			},
			{
				ID:   "powercom",
				Name: "Powercom health",
				Api:  "/xu4/powercom/health",
				Note: "Проверка состояния драйвера ИБП",
			},
			{
				ID:   "adblock",
				Name: "AdBlock Load",
				Api:  "/proxy/adblock/load",
				Note: "Загрузка и обновление списков блокировки",
			},
		}
		err := db.Create(&tasks).Error
		if err != nil {
			return fmt.Errorf("insert git_service error: %w", err)
		}
		log.Debugf("save git_service: initialize")
	}
	return nil
}

func CreateHealthRadiusUser(db *gorm.DB) error {

	if db.First(&entities.Radcheck{}).Error == gorm.ErrRecordNotFound {
		types := []entities.Radcheck{
			{
				Id:        utils.HashNow32(),
				UserName:  config.Default.Server.Radius.HealthUser,
				Attribute: "Cleartext-Password",
				Op:        ":=",
				Value:     config.Default.Server.Radius.HealthPasswd},
		}
		err := db.Create(&types).Error
		if err != nil {
			return fmt.Errorf("insert radcheck error: %w", err)
		}
		log.Debugf("save radcheck: initialize")
	}

	return nil

}
