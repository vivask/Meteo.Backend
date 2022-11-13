package entities

import (
	"time"
)

type Blocklist struct {
	ID string `gorm:"column:hostname;primaryKey;size:100" json:"id"`
}

func (Blocklist) TableName() string {
	return "blocklists"
}

type ToVpnAuto struct {
	ID        string    `gorm:"column:hostname;primaryKey;size:100" json:"id"`
	CreatedAt time.Time `gorm:"column:created;not null;default:Current_timestamp" json:"created"`
}

func (ToVpnAuto) TableName() string {
	return "tovpn_autos"
}

type AccesList struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
}

func (AccesList) TableName() string {
	return "access_lists"
}

type ToVpnManual struct {
	ID        uint32    `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:hostname;not null;unique;index;size:100" json:"name"`
	Note      string    `gorm:"column:note" json:"note"`
	ListID    string    `json:"-"`
	AccesList AccesList `gorm:"foreignkey:ListID" json:"list"`
}

func (ToVpnManual) TableName() string {
	return "tovpn_manuals"
}

type ToVpnIgnore struct {
	ID        string    `gorm:"column:hostname;primaryKey;size:100" json:"id"`
	UpdatedAt time.Time `gorm:"column:updated" json:"updated"`
}

func (ToVpnIgnore) TableName() string {
	return "tovpn_ignores"
}

type ProxyState struct {
	Active  bool `json:"active"`
	Master  bool `json:"master"`
	AdBlock bool `json:"adblock"`
	Cache   bool `json:"cache"`
	Unlock  bool `json:"unlock"`
}
