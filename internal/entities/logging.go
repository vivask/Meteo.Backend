package entities

import "time"

type Logging struct {
	ID        string    `gorm:"column:id;not null;primaryKey;unique;index;size:45" json:"id"`
	Message   string    `gorm:"column:message;not null;size:128" json:"message"`
	Type      string    `gorm:"column:type;not null;size:2" json:"type"`
	CreatedAt time.Time `gorm:"column:date_time;not null;autoCreateTime:false;default:CURRENT_TIMESTAMP" json:"date_time"`
	Date      string    `gorm:"-" sql:"date" json:"date"`
	Time      string    `gorm:"-" sql:"time" json:"time"`
}

func (Logging) TableName() string {
	return "logging"
}
