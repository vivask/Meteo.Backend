package entities

import (
	"time"

	"gorm.io/gorm"
)

type NullTime struct {
	time.Time
	Valid bool
}

type SyncParams struct {
	ID          uint32 `gorm:"column:id;not null;unique;index" json:"id"`
	SyncType    string `gorm:"column:sync_type;not null;size:20" json:"sync_type"`
	SyncTableID string `gorm:"column:table_id;not null" json:"table_id"`
}

func (SyncParams) TableName() string {
	return "sync_params"
}

type SyncTables struct {
	ID   string `gorm:"column:name;not null;size:45;primaryKey;unique;index" json:"name"`
	Note string `gorm:"column:note;not null" json:"note"`
	//	SyncedAt time.Time    `gorm:"column:syncedat" json:"syncedat"`
	Params []SyncParams `gorm:"foreignKey:LocationID;references:ID" json:"params"`
}

func (SyncTables) SyncTables() string {
	return "sync_tables"
}

func (t *SyncTables) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Delete(&SyncParams{}, "table_id = ?", t.ID).Error
}
