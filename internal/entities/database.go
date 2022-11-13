package entities

import (
	"time"

	"gorm.io/gorm"
)

type NullTime struct {
	time.Time
	Valid bool
}

type SyncTypes struct {
	ID   string `gorm:"column:id;primaryKey;size:20" json:"id"`
	Note string `gorm:"column:note;not null" json:"note"`
}

func (SyncTypes) TableName() string {
	return "sync_types"
}

type SyncParams struct {
	ID          uint32    `gorm:"column:id;primaryKey" json:"id"`
	SyncTypeID  string    `json:"-"`
	SyncTypes   SyncTypes `gorm:"foreignkey:SyncTypeID" json:"stype"`
	SyncTableID string    `gorm:"column:table_id;not null" json:"table_id"`
}

func (SyncParams) TableName() string {
	return "sync_params"
}

type SyncTables struct {
	ID       string       `gorm:"column:name;size:45;primaryKey" json:"name"`
	Note     string       `gorm:"column:note" json:"note"`
	SyncedAt time.Time    `gorm:"column:syncedat" json:"syncedat"`
	Params   []SyncParams `gorm:"foreignKey:SyncTableID;references:ID" json:"params"`
	IsImport bool         `gorm:"-" json:"import"`
}

func (SyncTables) SyncTables() string {
	return "sync_tables"
}

func (t *SyncTables) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Delete(&SyncParams{}, "table_id = ?", t.ID).Error
}

type Callback struct {
	Query  string        `json:"query"`
	Params []interface{} `json:"params"`
}
