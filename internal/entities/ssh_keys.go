package entities

import "time"

type SshKeys struct {
	ID                uint32    `gorm:"column:id;not null;primaryKey;unique;index" json:"id"`
	Finger            string    `gorm:"column:finger;not null;type:TEXT" json:"finger"`
	Owner             string    `gorm:"column:owner;not null;unique;index;size:45" json:"owner"`
	CreatedAt         time.Time `gorm:"column:created;not null;default:CURRENT_TIMESTAMP" json:"created"`
	UpdatedAt         time.Time `gorm:"column:used" json:"used"`
	HasRecentActivity bool      `gorm:"-" sql:"activity"`
	ShortFinger       string    `gorm:"-" sql:"short_finger"`
}

func (SshKeys) TableName() string {
	return "ssh_keys"
}
