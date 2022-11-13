package repo

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/log"
	"meteo/internal/utils"
	"time"
)

func (p sshclientService) GetAllSshHosts(pageable dto.Pageable) ([]entities.SshHosts, error) {
	var hosts []entities.SshHosts
	err := p.db.Preload("SshKeys").Find(&hosts).Order("created desc").Error
	if err != nil {
		return nil, fmt.Errorf("error read sshhosts: %w", err)
	}
	for i, host := range hosts {
		if len(host.Finger) > 52 {
			hosts[i].ShortFinger = host.Finger[36:52]
		} else {
			hosts[i].ShortFinger = host.Finger
		}
		hosts[i].HasRecentActivity = !host.UpdatedAt.IsZero()
	}
	return hosts, err
}

func (p sshclientService) AddSshHost(host entities.SshHosts) error {
	host.ID = utils.HashString32(host.Host)
	err := p.db.Omit("UpdatedAt").Create(&host).Error
	if err != nil {
		log.Errorf("error insert sshhosts: %v", err)
		return fmt.Errorf("error insert sshhosts: %w", err)
	}
	return err
}

func (p sshclientService) EditSshHost(host entities.SshHosts) error {
	err := p.db.Save(&host).Error
	if err != nil {
		return fmt.Errorf("error update ssh_hosts: %w", err)
	}
	return nil
}

func (p sshclientService) DelSshHost(id uint32) error {
	host := entities.SshHosts{ID: id}
	err := p.db.Delete(&host).Error
	if err != nil {
		return fmt.Errorf("error delete sshhost: %w", err)
	}
	return nil
}

func (p sshclientService) UpTimeSshHosts(host string) error {
	knowhost := entities.SshHosts{}
	err := p.db.Where("host = ?", host).First(&knowhost).Error
	if err != nil {
		return fmt.Errorf("read sshhosts error: %w", err)
	}
	knowhost.UpdatedAt = time.Now().Local()
	err = p.db.Save(&knowhost).Error
	if err != nil {
		return fmt.Errorf("update sshhosts error: %w", err)
	}
	return nil
}
