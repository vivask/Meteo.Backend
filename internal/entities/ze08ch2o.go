package entities

import "time"

type Ze08ch2o struct {
	ID        string    `gorm:"column:id;not null;primaryKey;unique;index;size:45" json:"id"`
	Ch2o      int       `gorm:"column:ch2o;not null" json:"ch2o"`
	CreatedAt time.Time `gorm:"column:date_time;not null;unique;index;;default:CURRENT_TIMESTAMP" json:"date_time"`
}

func (Ze08ch2o) TableName() string {
	return "ze08ch2o"
}
