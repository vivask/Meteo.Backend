package entities

import (
	"fmt"
	"meteo/internal/log"

	"gorm.io/gorm"
)

// AutoMigrate for migrate database schema
func AutoMigrate(db *gorm.DB) {
	log.Info("Migrating model")
	if err := db.Transaction(func(tx *gorm.DB) error {
		// Migrate child tables
		err := tx.AutoMigrate(&SyncParams{}, &TaskParams{}, &JobParams{}, &Periods{},
			&Days{}, &Executors{}, &JobParams{})
		if err != nil {
			return err
		}
		log.Debugf("Migrate child tables success")

		return err

	}); err != nil {
		log.Fatal(err)
	}
	if err := db.Transaction(func(tx *gorm.DB) error {
		// Migrate other tables
		err := tx.AutoMigrate(&Bmx280{}, &SyncTables{}, &Ds18b20{},
			&GitKeys{}, &GitUsers{}, &Logging{}, &Mics6814{}, &Blocklist{}, &Homezone{},
			&AccesList{}, &ToVpnManual{}, &ToVpnAuto{}, &ToVpnIgnore{}, &Radsens{},
			&Tasks{}, &Jobs{}, &Settings{}, &SshHosts{}, &SshKeys{}, &Ze08ch2o{},
			&Radacct{}, &Radcheck{}, &Radgroupcheck{}, &Radgroupreply{}, &Radreply{},
			&Radusergroup{}, &Radpostauth{}, &Nas{})
		if err != nil {
			return err
		}
		log.Debugf("Migrate other tables success")

		return err

	}); err != nil {
		log.Fatal(err)
	}

	// Create predefined fields
	if err := db.Transaction(func(tx *gorm.DB) error {
		log.Info("Create predefined fields")
		if tx.First(&Settings{}).Error == gorm.ErrRecordNotFound {
			err := tx.Create(&Settings{ID: 1}).Error
			if err != nil {
				return fmt.Errorf("insert settings error: %w", err)
			}
			log.Debugf("save settings: initialize")
		}
		if tx.First(&Executors{}).Error == gorm.ErrRecordNotFound {
			executors := []Executors{
				{ID: "Master"},
				{ID: "Leader"},
				{ID: "Slave"},
				{ID: "All"},
			}
			for _, executor := range executors {
				err := tx.Create(&executor).Error
				if err != nil {
					return fmt.Errorf("insert executors error: %w", err)
				}
			}
			log.Debugf("save executors: initialize")
		}
		if tx.First(&Periods{}).Error == gorm.ErrRecordNotFound {
			periods := []Periods{
				{ID: "one", Name: "Once", Idx: 1},
				{ID: "sec", Name: "Second", Idx: 2},
				{ID: "min", Name: "Minute", Idx: 3},
				{ID: "hour", Name: "Hour", Idx: 4},
				{ID: "day", Name: "Day", Idx: 5},
				{ID: "week", Name: "Week", Idx: 6},
				{ID: "day_of_week", Name: "Day of week", Idx: 7},
				{ID: "month", Name: "Month", Idx: 8},
				{ID: "year", Name: "Year", Idx: 9},
			}
			for _, period := range periods {
				err := tx.Create(&period).Error
				if err != nil {
					return fmt.Errorf("insert periods error: %w", err)
				}
			}
			log.Debugf("save periods: initialize")
		}
		if tx.First(&Days{}).Error == gorm.ErrRecordNotFound {
			days := []Days{
				{ID: 1, Name: "Monday"},
				{ID: 2, Name: "Tuesday"},
				{ID: 3, Name: "Wednesday"},
				{ID: 4, Name: "Thursday"},
				{ID: 5, Name: "Friday"},
				{ID: 6, Name: "Saturday"},
				{ID: 7, Name: "Sunday"},
			}
			for _, day := range days {
				err := tx.Create(&day).Error
				if err != nil {
					return fmt.Errorf("insert days error: %w", err)
				}
			}
			log.Debugf("save days: initialize")
		}
		if tx.First(&AccesList{}).Error == gorm.ErrRecordNotFound {
			lists := []AccesList{
				{ID: "tovpn"},
				{ID: "local"},
			}
			for _, list := range lists {
				err := tx.Create(&list).Error
				if err != nil {
					return fmt.Errorf("insert access_lists error: %w", err)
				}
			}
			log.Debugf("save access_lists: initialize")
		}

		return nil

	}); err != nil {
		log.Fatal(err)
	}

	//db.Callback().Create().After("gorm:after_create").Register("sync:sync_create", r.syncCreate)
	//db.Callback().Update().After("gorm:after_update").Register("sync:sync_update", r.syncUpdate)
	//db.Callback().Delete().After("gorm:after_delete").Register("sync:sync_delete", r.syncDelete)

}
