package repo

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/log"

	"gorm.io/gorm"
)

func (p radiusService) GetAllAccounting(pageable dto.Pageable) ([]entities.Radacct, error) {

	var accounts []entities.Radacct
	err := p.db.Table("radacct").
		Select("radacct.*, radverified.username AS valid, radverified.callingstationid AS verified").
		Joins("left join radverified on radacct.callingstationid = radverified.callingstationid").
		Order("acctstarttime DESC").
		Find(&accounts).Error
	if err != nil {
		return nil, fmt.Errorf("error read radacct: %w", err)
	}

	if len(accounts) > pageable.Limit {
		accounts = accounts[pageable.Offset:pageable.Limit]
	}

	return accounts, nil
}

func (p radiusService) GetVerifiedAccounting(pageable dto.Pageable) ([]entities.Radacct, error) {

	var accounts []entities.Radacct
	err := p.db.Table("radacct").
		Select("radacct.*, radverified.username AS valid, radverified.callingstationid AS verified").
		Joins("left join radverified on radacct.callingstationid = radverified.callingstationid").
		Where("radverified.callingstationid IS NOT NULL").
		Order("acctstarttime DESC").
		Find(&accounts).Error
	if err != nil {
		return nil, fmt.Errorf("error read radacct: %w", err)
	}

	if len(accounts) > pageable.Limit {
		accounts = accounts[pageable.Offset:pageable.Limit]
	}

	return accounts, nil
}

func (p radiusService) GetNotVerifiedAccounting(pageable dto.Pageable) ([]entities.Radacct, error) {

	var accounts []entities.Radacct
	err := p.db.Table("radacct").
		Select("radacct.*, radverified.username AS valid, radverified.callingstationid AS verified").
		Joins("left join radverified on radacct.callingstationid = radverified.callingstationid").
		Where("radverified.callingstationid IS NULL").
		Order("acctstarttime DESC").
		Find(&accounts).Error
	if err != nil {
		return nil, fmt.Errorf("error read radacct: %w", err)
	}

	if len(accounts) > pageable.Limit {
		accounts = accounts[pageable.Offset:pageable.Limit]
	}

	return accounts, nil
}

func (p radiusService) GetAlarmAccounting(pageable dto.Pageable) ([]entities.Radacct, error) {

	var accounts []entities.Radacct
	err := p.db.Table("radacct").
		Select("radacct.*, radverified.username AS valid, radverified.callingstationid AS verified").
		Joins("left join radverified on radacct.callingstationid = radverified.callingstationid").
		Where("radverified.callingstationid IS NOT NULL AND radverified.username != radacct.username").
		Order("acctstarttime DESC").
		Find(&accounts).Error
	if err != nil {
		return nil, fmt.Errorf("error read radacct: %w", err)
	}

	if len(accounts) > pageable.Limit {
		accounts = accounts[pageable.Offset:pageable.Limit]
	}

	return accounts, nil
}

func (p radiusService) GetAllVerified(pageable dto.Pageable) ([]entities.Radverified, error) {

	var verified []entities.Radverified
	query := `SELECT radverified.*,
	 (SELECT acctupdatetime
		FROM radacct
	  WHERE callingstationid=radverified.callingstationid
		ORDER BY acctupdatetime DESC
		LIMIT 1)
		FROM radverified;`
	err := p.db.Raw(query).Scan(&verified).Error
	if err != nil {
		return nil, fmt.Errorf("error read radacct: %w", err)
	}
	return verified, nil
}

func (p radiusService) Verify(id int) error {
	var account entities.Radacct
	account.RadAcctId = id
	err := p.db.First(&account).Error
	if err != nil {
		return fmt.Errorf("error read account: %w", err)
	}

	var verified entities.Radverified
	err = p.db.Where("callingstationid = ?", account.CallingStationId).First(&verified).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		verified.CallingStationId = account.CallingStationId
		verified.UserName = account.UserName
		log.Infof("Verified: %v", verified)
		err = p.db.Create(&verified).Error
		if err != nil {
			return fmt.Errorf("error create verified: %w", err)
		}
	} else {
		return fmt.Errorf("error read verified: %w", err)
	}

	if err == nil && verified.UserName != account.UserName {
		verified.UserName = account.UserName
		err = p.db.Save(&verified).Error
		if err != nil {
			return fmt.Errorf("error update verified: %w", err)
		}
	}

	return nil
}

func (p radiusService) ExcludeUser(id int) error {
	tx := p.db.Begin()
	verified := &entities.Radverified{Id: id}
	err := tx.Delete(verified).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete radius verified: %w", err)
	}
	tx.Commit()
	return nil
}

func (p radiusService) DeleteAccounting(id string) error {
	var rows []entities.Radacct
	var err error
	if id == "suspect" {
		rows, err = p.GetAlarmAccounting(dto.Pageable{})
	} else if id == "unverified" {
		rows, err = p.GetNotVerifiedAccounting(dto.Pageable{})
	} else if id == "verified" {
		rows, err = p.GetVerifiedAccounting(dto.Pageable{})
	} else {
		query := "DELETE FROM radacct"
		err = p.db.Exec(query).Error
		if err != nil {
			return fmt.Errorf("error delete all radacct: %w", err)
		}
		return nil
	}

	if err != nil {
		return fmt.Errorf("error get radacct: %w", err)
	}

	tx := p.db.Begin()
	err = tx.Where("radacctid IS NOT NULL").Delete(&rows).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete radacct: %w", err)
	}
	tx.Commit()
	return nil
}
