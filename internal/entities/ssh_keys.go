package entities

import "time"

type SshKeys struct {
	ID                uint32    `gorm:"column:id;primaryKey" json:"id"`
	Finger            string    `gorm:"column:finger;not null;type:TEXT" json:"finger"`
	Owner             string    `gorm:"column:owner;not null;unique;size:45" json:"owner"`
	CreatedAt         time.Time `gorm:"column:created;not null;default:Current_timestamp" json:"created"`
	UpdatedAt         time.Time `gorm:"column:used;autoUpdateTime" json:"used"`
	HasRecentActivity bool      `gorm:"-" sql:"activity" json:"activity"`
	ShortFinger       string    `gorm:"-" sql:"short_finger" json:"short_finger"`
}

func (SshKeys) TableName() string {
	return "ssh_keys"
}
