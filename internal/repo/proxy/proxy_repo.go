package repo

import (
	"meteo/internal/dto"
	"meteo/internal/entities"

	"gorm.io/gorm"
)

// ProxyService api controller of produces
type ProxyService interface {
	GetAllBlockHosts(pageable dto.Pageable) (*[]entities.Blocklist, error)
	ClearBlocklist() error
	AddBlockHost(host entities.Blocklist) error
	GetAllAutoToVpn(pageable dto.Pageable) (*[]entities.ToVpnAuto, error)
	GetAllIgnore(pageable dto.Pageable) (*[]entities.ToVpnIgnore, error)
	GetManualToVpn(id uint32) (*entities.ToVpnManual, error)
	AddManualToVpn(host *entities.ToVpnManual) error
	DelManualToVpn(id uint32) error
	AddAutoToVpn(host *entities.ToVpnAuto) error
	DelAutoToVpn(id string) error
	AddIgnoreToVpn(id string) error
	DelIgnoreToVpn(id string) error
	RestoreAutoToVpn(id string) error
	AddHomeZoneHost(host entities.Homezone) error
	GetAllHomeZoneHosts() (*[]entities.Homezone, error)
}

type proxyService struct {
	db *gorm.DB
}

// NewProductService get product service instance
func NewProxyService(db *gorm.DB) ProxyService {
	return &proxyService{db}
}
