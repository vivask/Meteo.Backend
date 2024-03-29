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

	SyncBmx280() error
	ReplaceBmx280(readings []entities.Bmx280) error
	AddSyncBmx280(bmx280 []entities.Bmx280) error
	GetNotSyncBmx280() ([]entities.Bmx280, error)
	GetAllBmx280() ([]entities.Bmx280, error)

	SyncDs18b20() error
	ReplaceDs18b20(readings []entities.Ds18b20) error
	AddSyncDs18b20(ds18b20 []entities.Ds18b20) error
	GetNotSyncDs18b20() ([]entities.Ds18b20, error)
	GetAllDs18b20() ([]entities.Ds18b20, error)

	SyncZe08ch2o() error
	ReplaceZe08ch2o(readings []entities.Ze08ch2o) error
	AddSyncZe08ch2o(ze08ch2o []entities.Ze08ch2o) error
	GetNotSyncZe08ch2o() ([]entities.Ze08ch2o, error)
	GetAllZe08ch2o() ([]entities.Ze08ch2o, error)

	SyncMics6814() error
	ReplaceMics6814(readings []entities.Mics6814) error
	AddSyncMics6814(mics6814 []entities.Mics6814) error
	GetNotSyncMics6814() ([]entities.Mics6814, error)
	GetAllMics6814() ([]entities.Mics6814, error)

	SyncRadsens() error
	ReplaceRadsens(readings []entities.Radsens) error
	AddSyncRadsens(mics6814 []entities.Radsens) error
	GetNotSyncRadsens() ([]entities.Radsens, error)
	GetAllRadsens() ([]entities.Radsens, error)

	GetAllAht25() ([]entities.Aht25, error)
	GetNotSyncAht25() ([]entities.Aht25, error)
	AddSyncAht25(data []entities.Aht25) error
	SyncAht25() error
	ReplaceAht25(readings []entities.Aht25) error

	DropTables(tables []entities.SyncTables) error

	GetAllGitUsers() ([]entities.GitUsers, error)
	GetAllHomezones() ([]entities.Homezone, error)
	GetAllJobs() ([]entities.Jobs, error)
	GetAllRadcheck() ([]entities.Radcheck, error)
	GetAllRadverified() ([]entities.Radverified, error)
	GetAllSshHosts() ([]entities.SshHosts, error)
	GetAllSshKeys() ([]entities.SshKeys, error)
	GetAllTasks() ([]entities.Tasks, error)
	GetAllToVpnAuto() ([]entities.ToVpnAuto, error)
	GetAllToVpnIgnore() ([]entities.ToVpnIgnore, error)
	GetAllToVpnManual() ([]entities.ToVpnManual, error)
	GetAllUsers() ([]entities.User, error)

	ReplaceSshKeys(readings []entities.SshKeys) error
	ReplaceSshHosts(readings []entities.SshHosts) error
	ReplaceGitUsers(readings []entities.GitUsers) error
	ReplaceHomezone(readings []entities.Homezone) error
	ReplaceToVpnManual(readings []entities.ToVpnManual) error
	ReplaceToVpnAuto(readings []entities.ToVpnAuto) error
	ReplaceToVpnIgnore(readings []entities.ToVpnIgnore) error
	ReplaceTasks(readings []entities.Tasks) error
	ReplaceJobs(readings []entities.Jobs) error
	ReplaceRadcheck(readings []entities.Radcheck) error
	ReplaceUser(readings []entities.User) error
	ReplaceRadverified(readings []entities.Radverified) error
}

type databaseService struct {
	db *gorm.DB
}

// NewDatabaseService get database service instance
func NewDatabaseService(db *gorm.DB) DatabaseService {
	return &databaseService{db}
}
