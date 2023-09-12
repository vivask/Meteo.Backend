package repo

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/utils"
)

func (p radiusService) GetAllUsers(pageable dto.Pageable) ([]entities.Radcheck, error) {
	var users []entities.Radcheck
	err := p.db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("error read tasks: %w", err)
	}
	return users, nil
}

func (p radiusService) AddUser(user entities.Radcheck) (uint32, error) {
	user.Id = utils.HashNow32()
	err := p.db.Create(&user).Error
	if err != nil {
		return 0, fmt.Errorf("error create radius user: %w", err)
	}
	return user.Id, nil
}

func (p radiusService) EditUser(user entities.Radcheck) error {
	err := p.db.Save(&user).Error
	if err != nil {
		return fmt.Errorf("error update radius user: %w", err)
	}
	return nil
}

func (p radiusService) DelUser(id uint32) error {
	tx := p.db.Begin()
	user := &entities.Radcheck{Id: id}
	err := tx.First(user).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error read radcheck: %w", err)
	}
	if user.UserName == config.Default.Radius.HealthUser {
		tx.Rollback()
		return fmt.Errorf("can't delete system user %s", config.Default.Radius.HealthUser)
	}
	err = tx.Delete(user).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete radius user: %w", err)
	}
	tx.Commit()
	return nil
}
