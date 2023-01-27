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
	DeafultLeader = NewLeader(
		config.Default.Client.Local,
		config.Default.Client.Remote,
		config.Default.App.Server)
}

func IsMain() bool { return config.Default.App.Server == "main" }

func IsBackup() bool { return config.Default.App.Server == "backup" }

func getClusterState() (*entities.Cluster, error) {
	body, err := GetInt("/cluster/leader/get")

	if err != nil {
		return nil, fmt.Errorf("local Cluster Service not responding: %w", err)
	}

	c := &entities.Cluster{}
	err = json.Unmarshal(body, c)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error on cluster: %w", err)
	}
	log.Debugf("Cluster state Leader: %v, Alive: %v", c.Leader, c.AliveRemote)
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
	const PASSPHRASE = "aA2Xh41FiC4Wtj3e5b2LbytMdn6on7P0"
	resp, err := UdpSend(PASSPHRASE)
	return (err == nil && resp == PASSPHRASE)
}

func LocalIP() string { return DeafultLeader.LocalIP() }

func RemoreIP() string { return DeafultLeader.RemoreIP() }
