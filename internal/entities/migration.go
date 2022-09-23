package entities

import (
	"meteo/internal/log"

	"gorm.io/gorm"
)

// AutoMigrate for migrate database schema
func AutoMigrate(db *gorm.DB) {
	log.Info("Migrating model")
	err := db.AutoMigrate(&User{}, &Product{}, &ProductProps{}).Error()
	if len(err) != 0 {
		log.Errorf("Can't automigrate schema %v", err)
	}
}
