package entities

import "time"

type GitService struct {
	ID   string `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name;not null;unique" json:"name"`
}

func (GitService) TableName() string {
	return "git_service"
}

type GitUsers struct {
	ID                uint32    `gorm:"column:id;primaryKey" json:"id"`
	Username          string    `gorm:"column:username;not null;size:45" json:"username"`
	Password          string    `gorm:"column:password;not null;size:45" json:"password"`
	SshKeyID          uint32    `json:"-"`
	SshKeys           SshKeys   `gorm:"foreignkey:SshKeyID" json:"ssh_key"`
	CreatedAt         time.Time `gorm:"column:created;not null;default:Current_timestamp" json:"created"`
	UpdatedAt         time.Time `gorm:"column:used" json:"used"`
	HasRecentActivity bool      `gorm:"-" sql:"activity" json:"activity"`
}

func (GitUsers) TableName() string {
	return "git_users"
}
