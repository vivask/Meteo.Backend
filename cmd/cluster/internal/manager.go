package vip

import (
	"encoding/json"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/log"
	"regexp"
	"sync"
	"time"
)

const PASSPHRASE = "aA2Xh41FiC4Wtj3e5b2LbytMdn6on7P0"

var (
	leader      bool
	aliveRemote bool
	healthy     bool = true
	m           sync.Mutex
)

type VIPHandler func(alive, leader bool)

type Params struct {
	Remote  string
	Handler VIPHandler
}

type VIPManager struct {
	stop                chan bool
	finished            chan bool
	networkConfigurator NetworkConfigurator
	port                string
	remoteAlive         bool
	handler             VIPHandler
}

func NewVIPManager(networkConfigurator NetworkConfigurator, p *Params) *VIPManager {
	return &VIPManager{
		networkConfigurator: networkConfigurator,
		port:                p.Remote,
		handler:             p.Handler,
	}
}

func (manager *VIPManager) addIP(verbose bool) {
	if error := manager.networkConfigurator.AddIP(); error != nil {
		log.Errorf("Could not set ip=%v interface=%v, error:  %v", manager.networkConfigurator.IP(), manager.networkConfigurator.Interface(), error)
	} else if verbose {
		log.Infof("Added IP %v on %v", manager.networkConfigurator.IP(), manager.networkConfigurator.Interface())
	}
}

func (manager *VIPManager) deleteIP(verbose bool) {
	if error := manager.networkConfigurator.DeleteIP(); error != nil {
		log.Errorf("Could not delete ip=%v interface=%v, error:  %v", manager.networkConfigurator.IP(), manager.networkConfigurator.Interface(), error)
	} else if verbose {
		log.Infof("Delete IP %v on %v", manager.networkConfigurator.IP(), manager.networkConfigurator.Interface())
	}
}

func (manager *VIPManager) Start() error {
	manager.stop = make(chan bool, 1)
	manager.finished = make(chan bool, 1)
	ticker := time.NewTicker(time.Second)

	setLeader(false)

	manager.deleteIP(false)

	go func() {
		for {
			select {
			case <-ticker.C:
				resp, err := UdpSend(PASSPHRASE)
				remoteAlive := (err == nil && resp == PASSPHRASE)
				//err := UdptSend(PASSPHRASE)
				//alive := (err == nil && isAlive())
				if manager.remoteAlive != remoteAlive {
					if remoteAlive {
						log.Info("Remote server is alive")
					} else {
						log.Info("Remote server is dead")
					}
					manager.remoteAlive = remoteAlive
					SetAliveRemote(remoteAlive)
				}
				if manager.isRemoteDead(remoteAlive) {
					log.Debug("RemoteDead")
					continue
				}
				if manager.isSelfMain() {
					log.Debug("isSelfMain")
					continue
				}
				if manager.isRemoteNotLeader() {
					log.Debug("RemoteNotLeader")
					continue
				}
				manager.setFollowing()

			case <-manager.stop:
				log.Debug("Virtual IP Manager Stopping")

				if IsLeader() {
					manager.deleteIP(true)
				}

				close(manager.finished)

				return
			}
		}
	}()

	log.Debug("Virtual IP Manager Started")

	return nil
}

func (manager *VIPManager) isRemoteDead(remoteAlive bool) bool {
	if !remoteAlive {
		manager.setLeader()
		return true
	}
	return false
}

func (manager *VIPManager) isSelfMain() bool {
	isServicesHealthy()
	if kit.IsMain() && healthy {
		manager.setLeader()
		return true
	}
	if kit.IsMain() && !healthy {
		manager.setFollowing()
		return true
	}
	return false
}

func (manager *VIPManager) isRemoteNotLeader() bool {
	if !isRemoteLeader() {
		manager.setLeader()
		return true
	}
	return false
}

func (manager *VIPManager) setLeader() {
	if !IsLeader() {
		setLeader(true)
		log.Info("LEADING")
		manager.addIP(true)
		manager.handler(manager.remoteAlive, true)
	}
}

func (manager *VIPManager) setFollowing() {
	if IsLeader() {
		setLeader(false)
		log.Info("FOLLOWING")
		manager.deleteIP(true)
		manager.handler(manager.remoteAlive, false)
	}
}

func (manager *VIPManager) Stop() {
	close(manager.stop)

	<-manager.finished

	log.Debug("Virtual IP Manager stopped")
}

func IsLeader() bool {
	m.Lock()
	defer m.Unlock()
	return leader
}

func setLeader(v bool) {
	m.Lock()
	leader = v
	m.Unlock()
}

func SetAliveRemote(v bool) {
	m.Lock()
	aliveRemote = v
	m.Unlock()
}

func IsAliveRemote() bool {
	m.Lock()
	defer m.Unlock()
	return aliveRemote
}

func isServicesHealthy() bool {
	body, err := kit.GetInt("/esp32/health")
	if err != nil {
		logHelathy(err.Error())
		return false
	}
	if len(body) > 0 {
		logHelathy(string(body))
		return false
	}

	body, err = kit.GetInt("/server/health")
	if err != nil {
		logHelathy(err.Error())
		return false
	}
	if len(body) > 0 {
		logHelathy(string(body))
		return false
	}
	healthy = true
	return true
}

func logHelathy(msg string) {
	if healthy {
		log.Error(msg)
	}
	healthy = false
}

func isRemoteLeader() bool {
	body, err := kit.GetExt("/cluster/leader/get")
	if err != nil {
		matched, _ := regexp.MatchString("connection refused", err.Error())
		if !matched {
			log.Error(err)
		}
		return false
	}
	c := &entities.Cluster{}
	err = json.Unmarshal(body, c)
	if err != nil {
		log.Error(err)
		return false
	}
	return c.Leader
}
