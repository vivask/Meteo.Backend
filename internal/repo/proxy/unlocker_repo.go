package repo

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/utils"
	"regexp"
)

const NOTFOUND = "not found host in access list"

func (p proxyService) GetAllManualToVpn(pageable dto.Pageable) (*[]entities.ToVpnManual, error) {
	hosts := new([]entities.ToVpnManual)
	err := p.db.Preload("AccesList").Find(&hosts).Error
	if err != nil {
		return nil, fmt.Errorf("error read tovpnIgnore: %w", err)
	}
	return hosts, nil
}

func (p proxyService) AddManualToVpn(host entities.ToVpnManual) error {
	host.ID = utils.HashNow32()
	tx := p.db.Begin()
	err := tx.Create(&host).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error insert tovpnManual: %w", err)
	}
	_, err = kit.PutInt("/sshclient/mikrotik/tovpn/manual/add", host)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("sshclient internal error: %w", err)
	}
	tx.Commit()

	return ReloadUnlocker()
}

func (p proxyService) EditManualToVpn(host entities.ToVpnManual) error {

	_, err := kit.PutInt("/sshclient/mikrotik/tovpn/manual/del", host)
	if err != nil {
		return fmt.Errorf("sshclient internal error: %w", err)
	}

	tx := p.db.Begin()
	err = tx.Save(&host).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error update tovpnManual: %w", err)
	}
	_, err = kit.PutInt("/sshclient/mikrotik/tovpn/manual/add", host)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("sshclient internal error: %w", err)
	}
	tx.Commit()

	return ReloadUnlocker()
}

func (p proxyService) DelManualFromVpn(id uint32) error {

	host := entities.ToVpnManual{}
	tx := p.db.Begin()
	err := tx.Preload("AccesList").Where("id = ?", id).First(&host).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error read tovpnManual: %w", err)
	}

	err = tx.Delete(&host).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete tovpnManual: %w", err)
	}

	_, err = kit.PutInt("/sshclient/mikrotik/tovpn/manual/del", host)
	if err != nil {
		matched, _ := regexp.MatchString(NOTFOUND, err.Error())
		if !matched {
			tx.Rollback()
			return fmt.Errorf("sshclient internal error: %w", err)
		}
	}
	tx.Commit()

	return ReloadUnlocker()
}

func (p proxyService) GetAllAutoToVpn(pageable dto.Pageable) (*[]entities.ToVpnAuto, error) {
	hosts := new([]entities.ToVpnAuto)
	err := p.db.Order("created DESC").Find(&hosts).Error
	if err != nil {
		return nil, fmt.Errorf("error read tovpnAuto: %w", err)
	}
	return hosts, nil
}

func (p proxyService) GetManualToVpnByID(id uint32) (*entities.ToVpnManual, error) {
	host := new(entities.ToVpnManual)
	err := p.db.Where("id = ?", id).First(host).Error
	if err != nil {
		return nil, fmt.Errorf("error read tovpnManual: %w", err)
	}
	return host, nil
}

func (p proxyService) AddAutoToVpn(host entities.ToVpnAuto) error {
	tx := p.db.Begin()
	_, err := kit.PutInt("/sshclient/mikrotik/tovpn/auto/add", host)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("sshclient internal error: %w", err)
	}
	err = tx.Create(&host).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error insert tovpnAuto: %w", err)
	}
	tx.Commit()

	return ReloadUnlocker()
}

func (p proxyService) DelAutoFromVpn(hosts []entities.ToVpnAuto) error {
	re, _ := regexp.Compile(NOTFOUND)
	tx := p.db.Begin()
	for _, host := range hosts {
		err := tx.Delete(&host).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error delete tovpnAuto: %w", err)
		}
		_, err = kit.PutInt("/sshclient/mikrotik/tovpn/auto/del", host)
		if err != nil {
			matched := re.MatchString(err.Error())
			if !matched {
				tx.Rollback()
				return fmt.Errorf("sshclient internal error: %w", err)
			}
		}
	}
	tx.Commit()

	return ReloadUnlocker()
}

func (p proxyService) GetAllIgnoreAutoToVpn(pageable dto.Pageable) (*[]entities.ToVpnIgnore, error) {
	hosts := new([]entities.ToVpnIgnore)
	err := p.db.Order("updated DESC").Find(&hosts).Error
	if err != nil {
		return nil, fmt.Errorf("error read tovpnIgnore: %w", err)
	}
	return hosts, nil
}

func (p proxyService) IgnoreAutoToVpn(hosts []entities.ToVpnAuto) error {
	re, _ := regexp.Compile(NOTFOUND)
	tx := p.db.Begin()
	for _, host := range hosts {
		err := tx.Delete(&host).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error delete tovpnAuto: %w", err)
		}
		_, err = kit.PutInt("/sshclient/mikrotik/tovpn/auto/del", host)
		if err != nil {
			matched := re.MatchString(err.Error())
			if !matched {
				tx.Rollback()
				return fmt.Errorf("sshclient internal error: %w", err)
			}
		}
		err = tx.Create(&entities.ToVpnIgnore{ID: host.ID, UpdatedAt: host.CreatedAt}).Error
		if err != nil {
			return fmt.Errorf("error create tovpnIgnore: %w", err)
		}
	}
	tx.Commit()

	return ReloadUnlocker()
}

func (p proxyService) RestoreAutoToVpn(hosts []entities.ToVpnIgnore) error {
	tx := p.db.Begin()
	for _, host := range hosts {
		err := tx.Create(&entities.ToVpnAuto{ID: host.ID, CreatedAt: host.UpdatedAt}).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error insert tovpnAuto: %w", err)
		}
		_, err = kit.PutInt("/sshclient/mikrotik/tovpn/auto/add", host)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("[sshclient internal error: %w", err)
		}
		err = tx.Delete(&host).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error delete tovpnIgnore: %w", err)
		}
	}
	tx.Commit()

	return ReloadUnlocker()
}

func (p proxyService) DelIgnoreAutoToVpn(hosts []entities.ToVpnIgnore) error {
	tx := p.db.Begin()
	for _, host := range hosts {
		err := tx.Delete(&host).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error delete tovpnIgnore: %w", err)
		}
	}
	tx.Commit()

	return ReloadUnlocker()
}

func (p proxyService) GetAccessLists(pageable dto.Pageable) (*[]entities.AccesList, error) {
	lists := new([]entities.AccesList)
	err := p.db.Find(&lists).Error
	if err != nil {
		return nil, fmt.Errorf("error read access_lists: %w", err)
	}
	return lists, nil
}

func ReloadUnlocker() error {

	_, err := kit.PutInt("/proxy/unlocker/reload", nil)
	if err != nil {
		return fmt.Errorf("sshclient internal error: %w", err)
	}

	_, err = kit.PutExt("/proxy/unlocker/reload", nil)
	if err != nil {
		return fmt.Errorf("sshclient internal error: %w", err)
	}
	return nil
}
