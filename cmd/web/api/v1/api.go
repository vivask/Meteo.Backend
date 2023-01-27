package v1

import (
	"meteo/cmd/web/api/v1/backup"
	"meteo/cmd/web/api/v1/database"
	"meteo/cmd/web/api/v1/esp32"
	"meteo/cmd/web/api/v1/master"
	"meteo/cmd/web/api/v1/media"
	"meteo/cmd/web/api/v1/nut"
	"meteo/cmd/web/api/v1/proxy"
	"meteo/cmd/web/api/v1/radius"
	"meteo/cmd/web/api/v1/schedule"
	"meteo/cmd/web/api/v1/sshclient"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WebAPI interface {
	RegisterProtectedAPIV1(router *gin.RouterGroup)
	RegisterPublicAPIV1(router *gin.RouterGroup)
}

type webAPI struct {
	proxy     proxy.ProxyAPI
	schedule  schedule.ScheduleAPI
	sshclient sshclient.SshClientAPI
	database  database.DatabaseAPI
	esp32     esp32.Esp32API
	radius    radius.RadiusAPI
	media     media.MediaAPI
	main      master.MainAPI
	backup    backup.BackupAPI
	nut       nut.NutAPI
}

// NewWebAPI get product service instance
func NewWebAPI(db *gorm.DB) WebAPI {
	return &webAPI{
		proxy:     proxy.NewProxyAPI(db),
		schedule:  schedule.NewScheduleAPI(db),
		sshclient: sshclient.NewSshClientAPI(db),
		database:  database.NewDatabaseAPI(db),
		esp32:     esp32.NewEsp32API(db),
		radius:    radius.NewRadiusAPI(db),
		media:     media.NewMediaAPI(),
		main:      master.NewMainAPI(),
		backup:    backup.NewBackupAPI(),
		nut:       nut.NewNutAPI(),
	}
}
