package vip

import (
	"context"
	"fmt"
	"sync"
	"time"

	"meteo/internal/config"
	"meteo/internal/kit"
	"meteo/internal/log"
)

var reloadSchedulerRunning = false
var mux sync.Mutex
var stop chan struct{}

func StartupVirtualIP(ctx context.Context) error {

	stop = make(chan struct{})

	if config.Default.App.Mode == "single" {
		setLeader(true)
		return nil
	}

	leader = false

	go ListenUdpPort(ctx, config.Default.Cluster.Bind, config.Default.Cluster.Port)

	//go ListenUdptPort(ctx, config.Default.Cluster.Bind, config.Default.Cluster.Port)

	log.Infof("Network interface=%s virtual-ip=%s", config.Default.Cluster.Interface, config.Default.Cluster.VirtualIP)

	ifaces, err := getAllInterfaces()
	if err != nil {
		return fmt.Errorf("network interfaces are not available: %w", err)
	}
	log.Debugf("Available network interfaces: %v", ifaces)

	netlinkNetworkConfigurator, err := NewNetlinkNetworkConfigurator(config.Default.Cluster.VirtualIP, config.Default.Cluster.Interface)
	if err != nil {
		return fmt.Errorf("network failure: %w", err)
	}
	paramas := &Params{
		Remote:  fmt.Sprintf("%s:%d", config.Default.Client.Remote, config.Default.Cluster.Port),
		Handler: Handler,
	}

	vipManager := NewVIPManager(netlinkNetworkConfigurator, paramas)

	if err := vipManager.Start(); err != nil {
		return fmt.Errorf("start failed: %w", err)
	}

	<-ctx.Done()

	return nil
}

func Handler(alive, leader bool) {
	if !kit.IsMain() {
		go reloadScheduler()
	}
}

func setStateReloadScheduler(state bool) {
	mux.Lock()
	reloadSchedulerRunning = state
	mux.Unlock()
}

func getStateReloadScheduler() bool {
	mux.Lock()
	defer mux.Unlock()
	return reloadSchedulerRunning
}

func reloadScheduler() {
	if getStateReloadScheduler() {
		setStateReloadScheduler(false)
		stop <- struct{}{}
		return
	}
	setStateReloadScheduler(true)

	timer := time.After(5 * time.Second)
	select {
	case <-timer:
		_, err := kit.PutInt("/schedule/jobs/reload", nil)
		if err != nil {
			log.Error(err)
		} else {
			log.Info("Sheduler reloaded success")
		}
	case <-stop:
	}
	setStateReloadScheduler(false)
}
