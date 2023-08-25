package vip

import (
	"context"
	"encoding/json"
	"fmt"
	"meteo/cmd/cluster/internal/VRRP"
	"meteo/internal/config"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/log"
	"net"
	"regexp"
	"sync"
	"time"
)

const (
	TIMEOUT = 800 * time.Millisecond
)

var (
	leader         bool = false
	aliveRemote    bool = false
	healthy        bool = true
	managerStarted bool = false
	m              sync.Mutex
)

type VIPHandler func(alive, leader bool)

type Params struct {
	Remote  string
	Handler VIPHandler
}

type VIPManager struct {
	stop                chan bool
	finished            chan bool
	started             chan bool
	networkConfigurator NetworkConfigurator
	port                string
	remoteAlive         bool
	handler             VIPHandler
	vr                  *VRRP.VirtualRouter
}

func NewVIPManager(networkConfigurator NetworkConfigurator, p *Params) *VIPManager {
	return &VIPManager{
		networkConfigurator: networkConfigurator,
		port:                p.Remote,
		handler:             p.Handler,
		vr:                  VRRP.NewVirtualRouter(byte(config.Default.Cluster.Vrid), config.Default.Cluster.Interface, false, VRRP.IPv4),
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

func (manager *VIPManager) Start(ctx context.Context, started chan bool) error {
	//Enable VRRP
	virtualIP, _, _ := net.ParseCIDR(fmt.Sprintf("%s/24", config.Default.Cluster.VirtualIP))
	manager.vr.AddIPvXAddr(virtualIP)
	manager.vr.Enroll(VRRP.Init2Master, func() {
		log.Info("init to MASTER")
	})
	manager.vr.Enroll(VRRP.Backup2Master, func() {
		log.Info("backup to MASTER")
	})
	manager.vr.Enroll(VRRP.Init2Master, func() {
		log.Info("init to BACKUP")
	})
	manager.vr.Enroll(VRRP.Master2Backup, func() {
		log.Info("master to BACKUP")
	})
	go manager.vr.StartWithEventSelector()

	manager.started = started
	manager.stop = make(chan bool, 1)
	manager.finished = make(chan bool, 1)
	ticker := time.NewTicker(TIMEOUT * 2)

	if kit.IsAliveRemote() {
		setLeader(kit.IsBackup())
	}

	manager.deleteIP(false)

	go func() {
		for {
			select {
			case <-ticker.C:
				remoteAlive := kit.IsAliveRemote()
				if manager.remoteAlive != remoteAlive {
					if remoteAlive {
						log.Info("Remote server is alive")
					} else {
						log.Info("Remote server is dead")
					}
					manager.remoteAlive = remoteAlive
				}

				if kit.IsHealthyInt("/proxy/health") {
					manager.vr.SetPriorityAndMasterAdvInterval(byte(config.Default.Cluster.Priority), time.Millisecond*800)
				} else {
					manager.vr.SetPriorityAndMasterAdvInterval(byte(1), time.Millisecond*800)
				}

				if manager.isRemoteDead(remoteAlive) {
					continue
				}
				if manager.isSelfMain() {
					continue
				}
				if manager.isRemoteNotLeader() {
					continue
				}
				manager.setFollowing()

			case <-manager.stop:
			case <-ctx.Done():
				log.Debugf("Virtual IP Manager done: %v", ctx.Err())

				if IsLeader() {
					manager.deleteIP(true)
				}

				close(manager.finished)
				managerStarted = false

				return
			}
		}
	}()

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
	healthy = isServicesHealthy()
	if kit.IsMain() && (healthy || !managerStarted) {
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
		log.Infof("LEADING")
		manager.addIP(true)
		ManagerStarted(manager.started)
		manager.handler(manager.remoteAlive, true)
	}
}

func (manager *VIPManager) setFollowing() {
	if IsLeader() {
		setLeader(false)
		log.Info("FOLLOWING")
		manager.deleteIP(true)
		ManagerStarted(manager.started)
		manager.handler(manager.remoteAlive, false)
	}
}

func (manager *VIPManager) Stop() {
	close(manager.stop)

	<-manager.finished
	manager.vr.Stop()
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

func IsAliveRemote() bool {
	m.Lock()
	defer m.Unlock()
	return aliveRemote
}

func isServicesHealthy() bool {
	if !kit.IsHealthyInt("/esp32/health") {
		if healthy {
			log.Warning("Service ESP32 not available")
		}
		return false
	}

	if !kit.IsHealthyInt("/radius/health") {
		if healthy {
			log.Warning("Service Radius not available")
		}
		return false
	}

	if !kit.IsHealthyInt("/proxy/health") {
		if healthy {
			log.Warning("Service Proxy not available")
		}
		return false
	}
	return true
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

func ManagerStarted(started chan bool) {
	if !managerStarted {
		managerStarted = true
		started <- true
	}
}
