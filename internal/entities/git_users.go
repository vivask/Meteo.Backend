package entities

import "time"

type GitUsers struct {
	ID                uint32    `gorm:"column:id;not null;primaryKey;unique;index" json:"id"`
	Username          string    `gorm:"column:username;not null;size:45" json:"username"`
	Password          string    `gorm:"column:password;not null;size:45" json:"password"`
	Service           string    `gorm:"column:service;not null;unique;index;size:45" json:"service"`
	CreatedAt         time.Time `gorm:"column:created;not null;default:CURRENT_TIMESTAMP" json:"created"`
	UpdatedAt         time.Time `gorm:"column:used" json:"used"`
	HasRecentActivity bool      `gorm:"-" sql:"activity"`
}

func (GitUsers) TableName() string {
	return "git_users"
}
