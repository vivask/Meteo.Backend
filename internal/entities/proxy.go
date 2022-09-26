package entities

import "time"

type Blocklist struct {
	ID string `gorm:"column:hostname;not null;primaryKey;unique;index;size:100" json:"id"`
}

func (Blocklist) TableName() string {
	return "blocklists"
}

type Homezone struct {
	ID         uint32 `gorm:"column:id;not null;primaryKey;unique;index" json:"id"`
	DomainName string `gorm:"column:domain_name;not null;unique;index;size:100" json:"domain_name"`
	IPv4       string `gorm:"column:ip;not null;size:20" json:"ip"`
	Mac        string `gorm:"column:mac;size:20" json:"mac"`
	Note       string `gorm:"column:note" json:"note"`
	Active     bool   `json:"active"`
}

func (Homezone) TableName() string {
	return "homezones"
}

type ToVpnAuto struct {
	ID        string    `gorm:"column:hostname;not null;primaryKey;unique;index;size:100" json:"id"`
	CreatedAt time.Time `gorm:"column:createdat;not null;default:CURRENT_TIMESTAMP" json:"createdat"`
}

func (ToVpnAuto) TableName() string {
	return "tovpn_autos"
}

type AccesList struct {
	ID string `gorm:"column:id;not null;primaryKey;unique;index" json:"id"`
}

func (AccesList) TableName() string {
	return "access_lists"
}

type ToVpnManual struct {
	ID     uint32 `gorm:"column:id;not null;primaryKey;unique;index" json:"id"`
	Name   string `gorm:"column:hostname;not null;unique;index;size:100" json:"name"`
	Note   string `gorm:"column:note" json:"note"`
	ListID string `gorm:"column:list_id" json:"list_id"`
}

func (ToVpnManual) TableName() string {
	return "tovpn_manuals"
}

type ToVpnIgnore struct {
	ID        string    `gorm:"column:hostname;not null;primaryKey;unique;index;size:100" json:"id"`
	UpdatedAt time.Time `gorm:"column:updatedat" json:"updatedat"`
}

func (ToVpnIgnore) TableName() string {
	return "tovpn_ignores"
}

type ProxyState struct {
	Active     bool `json:"active"`
	BlkListOn  bool `json:"blkliston"`
	CacheOn    bool `json:"cacheon"`
	UnlockerOn bool `json:"unlockeron"`
}
