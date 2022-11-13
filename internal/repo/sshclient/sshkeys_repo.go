package repo

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/utils"
	"time"
)

func (p sshclientService) AddSshKey(sshKey entities.SshKeys) error {
	sshKey.ID = utils.HashNow32()
	err := p.db.Omit("UpdatedAt").Create(&sshKey).Error
	if err != nil {
		return fmt.Errorf("error insert ssh_keys: %w", err)
	}
	return nil
}

func (p sshclientService) EditSshKey(sshKey entities.SshKeys) error {
	sshKey.ID = utils.HashNow32()
	err := p.db.Omit("UpdatedAt").Create(&sshKey).Error
	if err != nil {
		return fmt.Errorf("error insert ssh_keys: %w", err)
	}
	return nil
}

func (p sshclientService) DelSshKey(id uint32) error {
	sshKeys := entities.SshKeys{ID: id}
	err := p.db.Delete(&sshKeys).Error
	if err != nil {
		return fmt.Errorf("error delete ssh_keys: %w", err)
	}
	return nil
}

func (p sshclientService) GetAllSshKeys(pageable dto.Pageable) ([]entities.SshKeys, error) {
	var keys []entities.SshKeys
	err := p.db.Order("created desc").Find(&keys).Error
	if err != nil {
		return nil, fmt.Errorf("error read ssh_keys: %w", err)
	}
	for i, key := range keys {
		if len(key.Finger) > 52 {
			keys[i].ShortFinger = key.Finger[36:52]
		} else {
			keys[i].ShortFinger = key.Finger
		}
		keys[i].HasRecentActivity = !key.UpdatedAt.IsZero()
	}
	return keys, err
}

func (p sshclientService) GetSshKeysByHost(host string) ([]entities.SshKeys, error) {
	var keys []entities.SshKeys
	err := p.db.Model(&entities.SshKeys{}).Select("ssh_keys.*").
		Joins("left join ssh_hosts on ssh_hosts.ssh_key_id = ssh_keys.id").
		Where("ssh_hosts.host = ?", host).
		Find(&keys).Error

	if err != nil {
		return nil, fmt.Errorf("error read ssh_keys: %w", err)
	}
	return keys, err
}

func (p sshclientService) UpTimeSshKey(owner string) error {
	key := entities.SshKeys{}
	err := p.db.Where("owner = ?", owner).First(&key).Error
	if err != nil {
		return fmt.Errorf("read sshhosts error: %w", err)
	}
	key.UpdatedAt = time.Now().Local()
	err = p.db.Save(&key).Error
	if err != nil {
		return fmt.Errorf("update sshhosts error: %w", err)
	}
	return nil
}
