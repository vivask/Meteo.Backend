package repo

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/utils"
)

func (p sshclientService) AddGitUser(gitUser entities.GitUsers) error {
	gitUser.ID = utils.HashNow32()
	err := p.db.Omit("UpdatedAt").Create(&gitUser).Error
	if err != nil {
		return fmt.Errorf("error insert git_users: %w", err)
	}
	return nil
}

func (p sshclientService) EditGitUser(user entities.GitUsers) error {
	err := p.db.Save(&user).Error
	if err != nil {
		return fmt.Errorf("error update git_users: %w", err)
	}
	return nil
}

func (p sshclientService) DelGitUser(id uint32) error {
	gitUsers := entities.GitUsers{ID: id}
	err := p.db.Delete(&gitUsers).Error
	if err != nil {
		return fmt.Errorf("error delete git_users: %w", err)
	}
	return nil
}

func (p sshclientService) GetAllGitUsers(pageable dto.Pageable) ([]entities.GitUsers, error) {
	var users []entities.GitUsers
	err := p.db.Preload("SshKeys").Find(&users).Order("created desc").Error
	if err != nil {
		return nil, fmt.Errorf("error read git_users: %w", err)
	}
	for i, user := range users {
		users[i].HasRecentActivity = !user.UpdatedAt.IsZero()
	}
	return users, err
}

func (p sshclientService) UpTimeGitUsers(service string) error {
	var gitUsers entities.GitUsers
	err := p.db.Model(&gitUsers).Where("service = ?", service).Update("used", "CURRENT_TIMESTAMP").Error
	if err != nil {
		return fmt.Errorf("update git_users error: %w", err)
	}
	return nil
}

func (p sshclientService) GetUserKeyByService(service string) (*entities.GitUsers, error) {
	var gitUsers entities.GitUsers
	err := p.db.Where("service = ?", service).First(&gitUsers).Error
	if err != nil {
		return nil, fmt.Errorf("error read git_users: %w", err)
	}
	return &gitUsers, err
}

func (p sshclientService) GetAllGitServices(pageable dto.Pageable) ([]entities.GitService, error) {
	var services []entities.GitService
	err := p.db.Order("name").Find(&services).Error
	if err != nil {
		return nil, fmt.Errorf("error read git_services: %w", err)
	}
	return services, err
}
