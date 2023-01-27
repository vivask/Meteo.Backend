package repo

import (
	"fmt"
	"meteo/internal/entities"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// AuthService api controller of produces
type AuthService interface {
	Create(user entities.User) error
	GetUserByID(userID string) (*entities.User, error)
	GetUserByName(name string) (*entities.User, error)
	UpdateUsername(user *entities.User) error
	UpdatePassword(userID string, password string) error
}

type authService struct {
	db *gorm.DB
}

// NewProductService get product service instance
func NewAuthService(db *gorm.DB) AuthService {
	return &authService{db}
}

func (p authService) Create(user entities.User) error {
	id, _ := uuid.NewV4()
	user.ID = id.String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Aproved = false

	err := p.db.Create(&user).Error
	if err != nil {
		return fmt.Errorf("error insert user: %w", err)
	}
	return nil
}

func (p authService) GetUserByID(userID string) (*entities.User, error) {
	user := new(entities.User)
	err := p.db.Where("id = ?", userID).First(user).Error
	if err != nil {
		return nil, fmt.Errorf("error read users id %s, error: %w", userID, err)
	}
	return user, nil
}

func (p authService) GetUserByName(name string) (*entities.User, error) {
	user := new(entities.User)
	err := p.db.Where("username = ?", name).First(user).Error
	if err != nil {
		return nil, fmt.Errorf("error read users %s, error: %w", name, err)
	}
	return user, nil
}

func (p authService) UpdateUsername(user *entities.User) error {
	user.UpdatedAt = time.Now()
	err := p.db.Where("id = ?", user.ID).Save(user).Error
	return err
}

func (p authService) UpdatePassword(userID string, password string) error {
	user := new(entities.User)
	err := p.db.Where("id = ?", user.ID).First(&user).Error
	if err != nil {
		return err
	}
	user.Password = password
	user.UpdatedAt = time.Now()
	err = p.db.Where("id = ?", userID).Save(user).Error
	return err
}
