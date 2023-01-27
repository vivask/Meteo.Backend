package entities

import "time"

type Mics6814 struct {
	ID        string    `gorm:"column:id;primaryKey;size:45" json:"id"`
	No2       float64   `gorm:"column:no2;not null" json:"no2"`
	Nh3       float64   `gorm:"column:nh3;not null" json:"nh3"`
	Co        float64   `gorm:"column:co;not null" json:"co"`
	CreatedAt time.Time `gorm:"column:date_time;not null;default:CURRENT_TIMESTAMP" json:"date_time"`
	Gdate     time.Time `gorm:"-" json:"gdate"`
}

func (Mics6814) TableName() string {
	return "mics6814"
}
