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
	GetAllManualToVpn(pageable dto.Pageable) (*[]entities.ToVpnManual, error)
	GetManualToVpnByID(id uint32) (*entities.ToVpnManual, error)
	AddManualToVpn(host entities.ToVpnManual) (uint32, error)
	EditManualToVpn(host entities.ToVpnManual) error
	DelManualFromVpn(id uint32) error
	GetAccessLists(pageable dto.Pageable) (*[]entities.AccesList, error)
	AddAutoToVpn(host entities.ToVpnAuto) error
	DelAutoFromVpn(hosts []entities.ToVpnAuto) error
	GetAllIgnoreAutoToVpn(pageable dto.Pageable) (*[]entities.ToVpnIgnore, error)
	IgnoreAutoToVpn(hosts []entities.ToVpnAuto) error
	RestoreAutoToVpn(hosts []entities.ToVpnIgnore) error
	DelIgnoreAutoToVpn(hosts []entities.ToVpnIgnore) error
	GetAllHomeZoneHosts() (*[]entities.Homezone, error)
	AddHomeZoneHost(host entities.Homezone) (uint32, error)
	EditHomeZoneHost(host entities.Homezone) error
	DelHomeZoneHost(id uint32) error
}

type proxyService struct {
	db *gorm.DB
}

// NewProductService get product service instance
func NewProxyService(db *gorm.DB) ProxyService {
	return &proxyService{db}
}
