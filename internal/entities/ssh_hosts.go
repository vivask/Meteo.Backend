package entities

import "time"

type SshHosts struct {
	ID                uint32    `gorm:"column:id;primaryKey" json:"id"`
	Host              string    `gorm:"column:host;not null;unique;index;size:45" json:"host"`
	Finger            string    `gorm:"column:finger;not null;type:TEXT" json:"finger"`
	CreatedAt         time.Time `gorm:"column:created;not null;default:Current_timestamp" json:"created"`
	UpdatedAt         time.Time `gorm:"column:used;autoUpdateTime" json:"used"`
	SshKeyID          uint32    `json:"-"`
	SshKeys           SshKeys   `gorm:"foreignkey:SshKeyID" json:"ssh_key"`
	HasRecentActivity bool      `gorm:"-" sql:"activity" json:"activity"`
	ShortFinger       string    `gorm:"-" sql:"short_finger" json:"short_finger"`
}

func (SshHosts) TableName() string {
	return "ssh_hosts"
}

type Touch struct {
	User string `json:"user"`
	Host string `json:"host"`
}
