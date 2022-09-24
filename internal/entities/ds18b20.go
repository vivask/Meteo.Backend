package entities

import "time"

type Ds18b20 struct {
	ID        string    `gorm:"column:id;not null;primaryKey;unique;index;size:45" json:"id"`
	Tempr     float64   `gorm:"column:tempr;not null" json:"tempr"`
	CreatedAt time.Time `gorm:"column:date_time;not null;unique;index;;default:CURRENT_TIMESTAMP" json:"date_time"`
}

func (Ds18b20) TableName() string {
	return "ds18b20"
}
