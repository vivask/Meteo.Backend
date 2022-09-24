package proxy

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"
	"meteo/internal/utils"

	"gorm.io/gorm"
)

// ProxyService api controller of produces
type ProxyService interface {
	AddManualToVpn(host *entities.ToVpnManual) error
}

type proxyService struct {
	db *gorm.DB
}

// NewProductService get product service instance
func NewProxyService(db *gorm.DB) ProxyService {
	return &proxyService{db}
}

func (p proxyService) AddManualToVpn(host *entities.ToVpnManual) error {
	host.ID = utils.HashNow32()
	err := p.db.Create(host).Error
	if err != nil {
		return fmt.Errorf("error insert tovpnManual: %w", err)
	}
	log.Debugf("save tovpnManual: %v", host)
	return nil
}
