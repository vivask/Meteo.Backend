package repo

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
)

func (p proxyService) GetAllBlockHosts(pageable dto.Pageable) (*[]entities.Blocklist, error) {
	// do stuff
	lists := new([]entities.Blocklist)
	err := p.db.Find(lists).Error
	if err != nil {
		return nil, fmt.Errorf("error read blocklist: %w", err)
	}
	return lists, nil
}

func (p proxyService) ClearBlocklist() error {
	err := p.db.Delete(&entities.Blocklist{}).Error
	if err != nil {
		return fmt.Errorf("error delete blocklist: %w", err)
	}
	return nil
}

func (p proxyService) AddBlockHost(host entities.Blocklist) error {
	err := p.db.Create(&host).Error
	if err != nil {
		return fmt.Errorf("error insert blocklist: %w", err)
	}
	return nil
}
