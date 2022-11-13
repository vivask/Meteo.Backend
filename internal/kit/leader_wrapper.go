package kit

import (
	"encoding/json"
	"fmt"
	"meteo/internal/config"
	"meteo/internal/entities"
	"meteo/internal/log"
)

var DeafultLeader *Leader

func InitLeader() {
	DeafultLeader = NewLeader(config.Default.App.Master, config.Default.Client.Local,
		config.Default.Client.Remote, config.Default.App.Server)
}

func IsMain() bool { return config.Default.App.Server == "main" }

func IsBackup() bool { return config.Default.App.Server == "backup" }

func getClusterState() (*entities.Cluster, error) {
	body, err := GetInt("/cluster/leader/get")

	if err != nil {
		//message := fmt.Sprintf("Cluster servise on [%s] is down, error: %v", config.Default.App.Server, err)
		//PostInt("/messanger/telegram", message)
		return nil, fmt.Errorf("local Cluster Service not responding: %w", err)
	}

	c := &entities.Cluster{}
	err = json.Unmarshal(body, c)
	if err != nil {
		//message := fmt.Sprintf("Cluster servise on [%s] is down, error: %v", config.Default.App.Server, err)
		//PostInt("/messanger/telegram", message)
		return nil, fmt.Errorf("unmarshal error on cluster: %w", err)
	}
	//log.Infof("Cluster state Leader: %v, Alive: %v", c.Leader, c.AliveRemote)
	return c, nil
}

func IsLeader() bool {
	c, err := getClusterState()
	if err != nil {
		log.Error(err)
		return false
	}
	return c.Leader
}

func IsAliveRemote() bool {
	c, err := getClusterState()
	if err != nil {
		log.Error(err)
		return false
	}
	return c.AliveRemote
}

func LocalIP() string { return DeafultLeader.LocalIP() }

func RemoreIP() string { return DeafultLeader.RemoreIP() }
