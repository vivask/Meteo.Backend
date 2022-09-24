package entities

import (
	"meteo/internal/log"

	"gorm.io/gorm"
)

// AutoMigrate for migrate database schema
func AutoMigrate(db *gorm.DB) {
	log.Info("Migrating model")
	if err := db.Transaction(func(tx *gorm.DB) error {
		tx.AutoMigrate(&SyncParams{})
		return tx.AutoMigrate(&SyncTables{})

		/*		return tx.AutoMigrate(&Bmx280{}, &SyncParams{}, &SyncTables{}, &Ds18b20{},
				&GitKeys{}, &GitUsers{}, &Logging{}, &Mics6814{}, &Blocklist{}, &Homezone{},
				&AccesList{}, &ToVpnManual{}, &ToVpnAuto{}, &ToVpnIgnore{}, &Radsens{}, &TaskParams{},
				&Tasks{}, &Periods{}, &Days{}, &Executors{}, &JobParams{}, &Jobs{},
				&Settings{}, &SshHosts{}, &SshKeys{}, &Ze08ch2o{})*/
	}); err != nil {
		log.Fatal(err)
	}
}
