package entities

import "time"

type Ds18b20 struct {
	ID        string    `gorm:"column:id;primaryKey;size:45" json:"id"`
	Tempr     float64   `gorm:"column:tempr;not null" json:"tempr"`
	CreatedAt time.Time `gorm:"column:date_time;not null;default:CURRENT_TIMESTAMP" json:"date_time"`
	Gdate     time.Time `gorm:"->;-:migration" json:"gdate"`
}

func (Ds18b20) TableName() string {
	return "ds18b20"
}
