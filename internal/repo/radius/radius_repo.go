package repo

import (
	"meteo/internal/dto"
	"meteo/internal/entities"

	"gorm.io/gorm"
)

// RadiusService api controller of produces
type RadiusService interface {
	GetAllUsers(pageable dto.Pageable) ([]entities.Radcheck, error)
	AddUser(user entities.Radcheck) (uint32, error)
	EditUser(user entities.Radcheck) error
	DelUser(id uint32) error
	GetAllAccounting(pageable dto.Pageable) ([]entities.Radacct, error)
	GetVerifiedAccounting(pageable dto.Pageable) ([]entities.Radacct, error)
	GetNotVerifiedAccounting(pageable dto.Pageable) ([]entities.Radacct, error)
	GetAlarmAccounting(pageable dto.Pageable) ([]entities.Radacct, error)
	GetAllVerified(pageable dto.Pageable) ([]entities.Radverified, error)
	Verify(id int) error
	ExcludeUser(id int) error
	DeleteAccounting(id string) error
}

type radiusService struct {
	db *gorm.DB
}

// NewRadiusService get schedule service instance
func NewRadiusService(db *gorm.DB) RadiusService {
	return &radiusService{db}
}
