package repo

import (
	"gorm.io/gorm"
)

// ServerService interface
type ServerService interface {
}

type serverService struct {
	db *gorm.DB
}

// NewServerService get server service instance
func NewServerService(db *gorm.DB) ServerService {
	return &serverService{db}
}
