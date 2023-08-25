package entities

import "time"

type Bmx280 struct {
	ID        string    `gorm:"column:id;primaryKey;size:45" json:"id"`
	Press     float64   `gorm:"column:press;not null" json:"press"`
	Tempr     float64   `gorm:"column:tempr;not null" json:"tempr"`
	Hum       float64   `gorm:"column:hum;not null" json:"hum"`
	CreatedAt time.Time `gorm:"column:date_time;uniqueIndex:,sort:desc;not null;default:CURRENT_TIMESTAMP" json:"date_time"`
	Gdate     time.Time `gorm:"->;-:migration" json:"gdate"`
}

func (Bmx280) TableName() string {
	return "bmx280"
}
