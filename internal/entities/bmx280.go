package entities

import "time"

type Bmx280 struct {
	ID        string    `gorm:"column:id;not null;primaryKey;unique;index;size:45" json:"id"`
	Press     float64   `gorm:"column:press;not null" json:"press"`
	Tempr     float64   `gorm:"column:tempr;not null" json:"tempr"`
	Hum       float64   `gorm:"column:hum;not null" json:"hum"`
	CreatedAt time.Time `gorm:"column:date_time;not null;unique;index;;default:CURRENT_TIMESTAMP" json:"date_time"`
}

func (Bmx280) TableName() string {
	return "bmx280"
}
