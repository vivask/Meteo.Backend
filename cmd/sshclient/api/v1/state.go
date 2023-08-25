package v1

import (
	"meteo/internal/config"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p sshclientAPI) GetMainState(c *gin.Context) {
	var state entities.Services

	state.ClusterService = kit.IsMainHealthy("/cluster/health")
	state.ClusterLog = kit.IsMainLogEmpty("/cluster/logging/empty")
	state.MessangerService = kit.IsMainHealthy("/messanger/health")
	state.MessangerLog = kit.IsMainLogEmpty("/messanger/logging/empty")
	state.ProxyService = kit.IsMainHealthy("/proxy/health")
	state.ProxyLog = kit.IsMainLogEmpty("/proxy/logging/empty")
	state.SshclientService = kit.IsMainHealthy("/sshclient/health")
	state.SshclientLog = kit.IsMainLogEmpty("/sshclient/logging/empty")
	state.ScheduleService = kit.IsMainHealthy("/schedule/health")
	state.ScheduleLog = kit.IsMainLogEmpty("/schedule/logging/empty")
	state.WebService = kit.IsMainHealthy("/web/health")
	state.WebLog = kit.IsMainLogEmpty("/web/logging/empty")
	state.Esp32Service = kit.IsMainHealthy("/esp32/health")
	state.Esp32Log = kit.IsMainLogEmpty("/esp32/logging/empty")
	state.RadiusService = kit.IsMainHealthy("/radius/health")
	state.RadiusLog = kit.IsMainLogEmpty("/radius/logging/empty")
	state.MediaService = kit.IsMainHealthy("/media/health")
	state.MediaLog = kit.IsMainLogEmpty("/media/logging/empty")
	state.NutService = kit.IsMainHealthy("/nut/health")
	state.NutLog = kit.IsMainLogEmpty("/nut/logging/empty")
	state.SambaService = kit.IsMainHealthy("/media/health/samba")
	state.StorageService = kit.IsMainHealthy("/media/health/storage")
	state.TransmissionService = kit.IsMainHealthy("/media/health/transmission")
	state.GogsService = utils.RawConnect(config.Default.Client.Local, []string{"3000"})
	state.PostgresService = utils.RawConnect(config.Default.Client.Local, []string{"5432"})

	c.JSON(http.StatusOK, state)
}

func (p sshclientAPI) GetBackupState(c *gin.Context) {
	var state entities.Services

	state.ClusterService = kit.IsBackupHealthy("/cluster/health")
	state.ClusterLog = kit.IsBackupLogEmpty("/cluster/logging/empty")
	state.MessangerService = kit.IsBackupHealthy("/messanger/health")
	state.MessangerLog = kit.IsBackupLogEmpty("/messanger/logging/empty")
	state.ProxyService = kit.IsBackupHealthy("/proxy/health")
	state.ProxyLog = kit.IsBackupLogEmpty("/proxy/logging/empty")
	state.SshclientService = kit.IsBackupHealthy("/sshclient/health")
	state.SshclientLog = kit.IsBackupLogEmpty("/sshclient/logging/empty")
	state.ScheduleService = kit.IsBackupHealthy("/schedule/health")
	state.ScheduleLog = kit.IsBackupLogEmpty("/schedule/logging/empty")
	state.WebService = kit.IsBackupHealthy("/web/health")
	state.WebLog = kit.IsBackupLogEmpty("/web/logging/empty")
	state.Esp32Service = kit.IsBackupHealthy("/esp32/health")
	state.Esp32Log = kit.IsBackupLogEmpty("/esp32/logging/empty")
	state.RadiusService = kit.IsBackupHealthy("/radius/health")
	state.RadiusLog = kit.IsBackupLogEmpty("/radius/logging/empty")
	state.StorageService = false //kit.IsBackupHealthy("/sshclient/server/backup/storage/helath")
	state.PostgresService = utils.RawConnect(config.Default.Client.Local, []string{"5432"})

	c.JSON(http.StatusOK, state)
}
