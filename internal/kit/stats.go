package kit

import (
	"meteo/internal/config"
	"meteo/internal/entities"
	"meteo/internal/utils"
)

func GetMainState() *entities.Services {
	var state entities.Services

	state.ClusterService = IsMainHealthy("/cluster/health")
	//state.ClusterLog = IsMainLogNotEmpty("/cluster/health")
	state.MessangerService = IsMainHealthy("/messanger/health")
	state.ProxyService = IsMainHealthy("/proxy/health")
	state.SshclientService = IsMainHealthy("/sshclient/health")
	state.ScheduleService = IsMainHealthy("/schedule/health")
	state.WebService = IsMainHealthy("/web/health")
	state.Esp32Service = IsMainHealthy("/esp32/health")
	state.RadiusService = IsMainHealthy("/radius/health")
	state.MediaService = IsMainHealthy("/media/health")
	state.NutService = IsMainHealthy("/nut/health")
	state.SambaService = IsMainHealthy("/media/health/samba")
	state.StorageService = IsMainHealthy("/media/health/storage")
	state.TransmissionService = IsMainHealthy("/media/health/transmission")
	state.GogsService = utils.RawConnect(config.Default.Client.Local, []string{"3000"})
	state.PostgresService = utils.RawConnect(config.Default.Client.Local, []string{"5432"})

	return &state
}
