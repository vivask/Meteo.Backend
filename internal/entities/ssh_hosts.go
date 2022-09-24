package entities

import "time"

type SshHosts struct {
	ID                uint32    `gorm:"column:id;not null;primaryKey;unique;index" json:"id"`
	Host              string    `gorm:"column:host;not null;unique;index;size:45" json:"host"`
	Finger            string    `gorm:"column:finger;not null;type:TEXT" json:"finger"`
	CreatedAt         time.Time `gorm:"column:created;not null;default:CURRENT_TIMESTAMP" json:"created"`
	UpdatedAt         time.Time `gorm:"column:used" json:"used"`
	SSHKeysID         uint32    `gorm:"column:ssh_keys_id" json:"ssh_keys_id"`
	HasRecentActivity bool      `gorm:"-" sql:"activity"`
	ShortFinger       string    `gorm:"-" sql:"short_finger"`
}

func (SshHosts) TableName() string {
	return "ssh_hosts"
}
