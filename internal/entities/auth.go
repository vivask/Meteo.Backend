package entities

import "time"

type User struct {
	ID        string    `gorm:"column:id;not null;primaryKey;size:45" json:"id"`
	Username  string    `gorm:"column:username;unique;not null;size:45" json:"username"`
	Email     string    `gorm:"column:email;unique;not null;size:255" json:"email"`
	Password  string    `gorm:"column:password;not null;size:255" json:"password"`
	CreatedAt time.Time `gorm:"column:createdat;not null;default:CURRENT_TIMESTAMP" json:"createdat"`
	UpdatedAt time.Time `gorm:"column:updatedat" json:"updatedat"`
	Aproved   bool      `gorm:"column:aproved" json:"aproved"`
}

func (User) TableName() string {
	return "users"
}
