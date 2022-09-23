package entities

import (
	"meteo/internal/log"

	"gorm.io/gorm"
)

// AutoMigrate for migrate database schema
func AutoMigrate(db *gorm.DB) {
	log.Info("Migrating model")
	if err := db.Transaction(func(tx *gorm.DB) error {
		return tx.AutoMigrate(&User{}, &Product{}, &ProductProps{})
	}); err != nil {
		log.Error(err)
	}
}
