package repo

import (
	"meteo/internal/dto"
	"meteo/internal/entities"

	"gorm.io/gorm"
)

// DatabaseService interface
type DatabaseService interface {
	GetAllTables(pageable dto.Pageable) ([]entities.SyncTables, error)
	GetAllSTypes(pageable dto.Pageable) ([]entities.SyncTypes, error)
	AddTable(table entities.SyncTables) error
	EditTable(table entities.SyncTables) error
	DelTable(id string) error
	DelTables([]entities.SyncTables) error
	GetTableById(id string) (*entities.SyncTables, error)
	ImportTableContent(id string) error
	ImportTablesContent(tables []entities.SyncTables) error
	SaveTableContent(id string) error
	SaveTablesContent(tables []entities.SyncTables) error
	ExecRaw(cb entities.Callback) error
}

type databaseService struct {
	db *gorm.DB
}

// NewDatabaseService get database service instance
func NewDatabaseService(db *gorm.DB) DatabaseService {
	return &databaseService{db}
}
