package entities

import "time"

type Radsens struct {
	ID        string    `gorm:"column:id;not null;primaryKey;unique;index;size:45" json:"id"`
	Dynamic   float64   `gorm:"column:dynamic;not null" json:"dynamic"`
	Static    float64   `gorm:"column:static;not null" json:"static"`
	Pulse     int       `gorm:"column:pulse;not null" json:"pulse"`
	CreatedAt time.Time `gorm:"column:date_time;not null;unique;index;;default:CURRENT_TIMESTAMP" json:"date_time"`
}

func (Radsens) TableName() string {
	return "radsens"
}
