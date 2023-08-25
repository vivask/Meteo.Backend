package entities

import "time"

type Ze08ch2o struct {
	ID        string    `gorm:"column:id;primaryKey;size:45" json:"id"`
	Ch2o      float64   `gorm:"column:ch2o;not null" json:"ch2o"`
	CreatedAt time.Time `gorm:"column:date_time;uniqueIndex:,sort:desc;not null;default:CURRENT_TIMESTAMP" json:"date_time"`
	Gdate     time.Time `gorm:"->;-:migration" json:"gdate"`
}

func (Ze08ch2o) TableName() string {
	return "ze08ch2o"
}
