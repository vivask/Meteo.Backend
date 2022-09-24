package repo

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/utils"
)

func (p proxyService) AddManualToVpn(host *entities.ToVpnManual) error {
	host.ID = utils.HashNow32()
	err := p.db.Create(host).Error
	if err != nil {
		return fmt.Errorf("error insert tovpnManual: %w", err)
	}
	return nil
}

func (p proxyService) GetAllAutoToVpn(pageable dto.Pageable) (*[]entities.ToVpnAuto, error) {
	hosts := new([]entities.ToVpnAuto)
	err := p.db.Order("createdat DESC").Find(&hosts).Error
	if err != nil {
		return nil, fmt.Errorf("error read tovpnAuto: %w", err)
	}
	return hosts, nil
}

func (p proxyService) GetAllIgnore(pageable dto.Pageable) (*[]entities.ToVpnIgnore, error) {
	hosts := new([]entities.ToVpnIgnore)
	err := p.db.Find(&hosts).Error
	if err != nil {
		return nil, fmt.Errorf("error read tovpnIgnore: %w", err)
	}
	return hosts, nil
}

func (p proxyService) GetManualToVpn(id uint32) (*entities.ToVpnManual, error) {
	host := new(entities.ToVpnManual)
	err := p.db.Where("id = ?", id).First(host).Error
	if err != nil {
		return nil, fmt.Errorf("error read tovpnManual: %w", err)
	}
	return host, nil

}

func (p proxyService) DelManualToVpn(id uint32) error {
	host := entities.ToVpnManual{ID: id}
	err := p.db.Delete(&host).Error
	if err != nil {
		return fmt.Errorf("error delete tovpnManual: %w", err)
	}
	return nil
}

func (p proxyService) AddAutoToVpn(host *entities.ToVpnAuto) error {
	err := p.db.Create(host).Error
	if err != nil {
		return fmt.Errorf("error insert tovpnAuto: %w", err)
	}
	return nil
}

func (p proxyService) DelAutoToVpn(id string) error {
	err := p.db.Where("hostname = ?", id).Delete(&entities.ToVpnAuto{}).Error
	if err != nil {
		return fmt.Errorf("error delete tovpnAuto: %w", err)
	}
	return nil
}

func (p proxyService) AddIgnoreToVpn(id string) error {
	tx := p.db.Begin()
	err := tx.Create(&entities.ToVpnIgnore{ID: id}).Error
	if err != nil {
		return fmt.Errorf("error create tovpnIgnore: %w", err)
	}
	err = tx.Where("hostname = ?", id).Delete(&entities.ToVpnAuto{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete tovpnAuto: %w", err)
	}
	tx.Commit()
	return nil
}

func (p proxyService) DelIgnoreToVpn(id string) error {
	err := p.db.Where("hostname = ?", id).Delete(&entities.ToVpnIgnore{}).Error
	if err != nil {
		return fmt.Errorf("error delete tovpnIgnore: %w", err)
	}
	return nil
}

func (p proxyService) RestoreAutoToVpn(id string) error {
	tx := p.db.Begin()
	err := tx.Create(&entities.ToVpnAuto{ID: id}).Error
	if err != nil {
		return fmt.Errorf("error create tovpnAuto: %w", err)
	}
	err = tx.Where("hostname = ?", id).Delete(&entities.ToVpnIgnore{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete tovpnIgnore: %w", err)
	}
	tx.Commit()
	return nil
}
