package entities

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"

	"gorm.io/gorm"
)

func CreateSyncTables(db *gorm.DB) error {
	if db.First(&entities.SyncTables{}).Error == gorm.ErrRecordNotFound {
		tables := []entities.SyncTables{
			{
				ID: "bmx280",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "sync"}},
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "ds18b20",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "sync"}},
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "mics6814",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "sync"}},
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "radsens",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "sync"}},
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "ze08ch2o",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "sync"}},
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "ssh_keys",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "ssh_hosts",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "git_users",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "homezones",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "tovpn_manuals",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "tovpn_autos",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "tovpn_ignores",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "tasks",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "jobs",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "radcheck",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
			{
				ID: "users",
				Params: []entities.SyncParams{
					{SyncTypes: entities.SyncTypes{ID: "replace"}},
				},
			},
		}
		err := db.Create(&tables).Error
		if err != nil {
			return fmt.Errorf("insert git_service error: %w", err)
		}
		log.Debugf("save git_service: initialize")
	}
	return nil
}
