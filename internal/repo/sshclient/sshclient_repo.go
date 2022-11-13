package repo

import (
	"meteo/internal/dto"
	"meteo/internal/entities"

	"gorm.io/gorm"
)

// SshClientService interface
type SshClientService interface {
	AddSshKey(sshKey entities.SshKeys) error
	DelSshKey(id uint32) error
	GetAllSshKeys(pageable dto.Pageable) ([]entities.SshKeys, error)
	GetSshKeysByHost(host string) ([]entities.SshKeys, error)
	DelSshHost(id uint32) error
	GetAllSshHosts(pageable dto.Pageable) ([]entities.SshHosts, error)
	AddSshHost(host entities.SshHosts) error
	EditSshHost(host entities.SshHosts) error
	UpTimeSshHosts(host string) error
	AddGitUser(gitUser entities.GitUsers) error
	EditGitUser(user entities.GitUsers) error
	DelGitUser(id uint32) error
	GetAllGitUsers(pageable dto.Pageable) ([]entities.GitUsers, error)
	UpTimeGitUsers(service string) error
	GetUserKeyByService(service string) (*entities.GitUsers, error)
	GetAllGitServices(pageable dto.Pageable) ([]entities.GitService, error)
	UpTimeSshKey(owner string) error
}

type sshclientService struct {
	db *gorm.DB
}

// NewSshClientService get sshclient service instance
func NewSshClientService(db *gorm.DB) SshClientService {
	return &sshclientService{db}
}
