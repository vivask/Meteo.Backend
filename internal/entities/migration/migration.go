package entities

import (
	"meteo/internal/config"
	"meteo/internal/entities"
	"meteo/internal/log"

	"gorm.io/gorm"
)

var ignoreSync map[string]struct{}

func prepareIgnoreMap() {
	ignoreSync = make(map[string]struct{})
	for _, ignore := range config.Default.Database.Exclude {
		ignoreSync[ignore] = struct{}{}
	}
}

func AutoSyncOff(table string) {
	ignoreSync[table] = struct{}{}
}

func AutoSyncOn(table string) {
	delete(ignoreSync, table)
}

// AutoMigrate for migrate database schema
func AutoMigrate(db *gorm.DB, migrate bool) {

	prepareIgnoreMap()

	if !migrate {
		goto end
	}

	log.Info("Migrating model")
	if err := db.Transaction(func(tx *gorm.DB) error {
		// Migrate child tables
		err := tx.AutoMigrate(&entities.SyncTypes{}, &entities.SyncParams{},
			&entities.TaskParams{}, &entities.JobParams{}, &entities.Periods{},
			&entities.Executors{}, &entities.JobParams{},
			&entities.GitService{}, &entities.Settings{}, &entities.AccesList{})
		if err != nil {
			return err
		}
		log.Debugf("Migrate child tables success")

		return nil

	}); err != nil {
		log.Fatal(err)
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		// Create predefined fields
		err := CreateSettings(db)
		if err != nil {
			return err
		}
		err = CreateExecutors(db)
		if err != nil {
			return err
		}
		err = CreatePeriods(db)
		if err != nil {
			return err
		}
		err = CreateAccesList(db)
		if err != nil {
			return err
		}
		err = CreateGitService(db)
		if err != nil {
			return err
		}
		err = CreateSyncTypes(db)
		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		log.Fatal(err)
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		// Migrate other tables
		err := tx.AutoMigrate(&entities.User{}, &entities.Bmx280{}, &entities.SyncTables{}, &entities.Ds18b20{},
			&entities.GitUsers{}, &entities.Logging{}, &entities.Mics6814{}, &entities.Blocklist{}, &entities.Homezone{},
			&entities.ToVpnManual{}, &entities.ToVpnAuto{}, &entities.ToVpnIgnore{}, &entities.Radsens{},
			&entities.Tasks{}, &entities.Jobs{}, &entities.SshHosts{}, &entities.SshKeys{}, &entities.Ze08ch2o{}, &entities.Aht25{},
			&entities.Radacct{}, &entities.Radcheck{}, &entities.Radgroupcheck{}, &entities.Radgroupreply{}, &entities.Radreply{},
			&entities.Radusergroup{}, &entities.Radpostauth{}, &entities.Nas{}, &entities.Radverified{})
		if err != nil {
			return err
		}
		log.Debugf("Migrate other tables success")

		return nil

	}); err != nil {
		log.Fatal(err)
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		// Create predefined fields
		err := CreateTasks(db)
		if err != nil {
			return err
		}
		err = CreateSyncTables(db)
		if err != nil {
			return err
		}
		err = CreateHealthRadiusUser(db)
		if err != nil {
			return err
		}
		err = CreateAdminUser(db)
		if err != nil {
			return err
		}
		return nil

	}); err != nil {
		log.Fatal(err)
	}

end:

	if config.Default.Database.Sync {
		db.Callback().Create().After("gorm:create").Register("sync_create", syncCreate)
		db.Callback().Update().After("gorm:update").Register("sync_update", syncUpdate)
		db.Callback().Delete().After("gorm:delete").Register("sync_delete", syncDelete)
	}
}
