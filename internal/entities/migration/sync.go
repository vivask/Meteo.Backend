package entities

import (
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/log"
	"meteo/internal/utils"
	"time"

	"gorm.io/gorm"
)

func syncCreate(db *gorm.DB) {
	//log.Infof("SQL: %v", db.Statement.SQL.String())
	//log.Infof("Params: %v", db.Statement.Vars...)

	_, tblName, err := utils.ParseQuery(db.Statement.SQL.String())

	if err != nil {
		log.Error(err)
		return
	}

	if _, ok := ignoreSync[tblName]; ok {
		return
	}

	if kit.IsAliveRemote() {
		_, err = kit.PostExt("/cluster/database/exec", entities.Callback{Query: db.Statement.SQL.String(), Params: params(db.Statement.Vars)})

		if err != nil {
			log.Error(err)
		}
	}

}

func syncUpdate(db *gorm.DB) {
	//log.Infof("SQL: %v", db.Statement.SQL.String())
	//log.Infof("Params: %v", db.Statement.Vars...)

	_, tblName, err := utils.ParseQuery(db.Statement.SQL.String())

	if err != nil {
		log.Error(err)
		return
	}

	if _, ok := ignoreSync[tblName]; ok {
		return
	}

	if kit.IsAliveRemote() {

		_, err = kit.PostExt("/cluster/database/exec", entities.Callback{Query: db.Statement.SQL.String(), Params: params(db.Statement.Vars)})

		if err != nil {
			log.Error(err)
		}
	}

}

func syncDelete(db *gorm.DB) {
	//log.Infof("SQL: %v", db.Statement.SQL.String())
	//log.Infof("Params: %v", db.Statement.Vars...)

	_, tblName, err := utils.ParseQuery(db.Statement.SQL.String())

	if err != nil {
		log.Error(err)
		return
	}

	if _, ok := ignoreSync[tblName]; ok {
		return
	}

	if kit.IsAliveRemote() {
		_, err = kit.PostExt("/cluster/database/exec", entities.Callback{Query: db.Statement.SQL.String(), Params: params(db.Statement.Vars)})

		if err != nil {
			log.Error(err)
		}
	}

}

func params(params []interface{}) []interface{} {
	prep := []interface{}{}
	for _, param := range params {
		if dt, ok := param.(time.Time); ok {
			dts := dt.Format("2006-01-02 15:04:05")
			prep = append(prep, dts)
		} else {
			prep = append(prep, param)
		}
	}
	return prep
}
