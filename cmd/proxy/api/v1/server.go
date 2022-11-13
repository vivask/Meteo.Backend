package v1

import (
	"context"
	"fmt"
	"io"
	"time"

	"meteo/cmd/proxy/api/v1/tools"
	"meteo/internal/config"
	"meteo/internal/entities"
	"meteo/internal/log"
	"os"
)

var started bool
var ctx context.Context
var cancel context.CancelFunc

func (p proxyAPI) Start() error {
	err := p.configLocalResolver()
	if err != nil {
		return fmt.Errorf("can't configure resolver: %w", err)
	}

	p.dns.SetZones(tools.LoadZones(p.repo))
	p.dns.SetBlackList(tools.LoadAdBlock())
	p.dns.SetUnlocker(tools.LoadUnlocker(p.repo))

	ctx, cancel = context.WithCancel(context.Background())
	timeout := time.After(1 * time.Second)
	fErr := make(chan error, 1)
	go func(fErr chan error) {
		fErr <- p.dns.Run(ctx)
	}(fErr)

	select {
	case e := <-fErr:
		if e != nil {
			return fmt.Errorf("proxy server failed to start, error: %w", e)
		}
	case <-timeout:
	}

	started = true

	go p.verifyServerState()
	log.Debugf("Proxy server success started")
	return nil
}

func (p proxyAPI) stop() {
	cancel()
	started = false
	log.Infof("DNS Server success stoped")
}

func (p proxyAPI) configLocalResolver() error {
	if len(config.Default.Proxy.Resolvers) == 0 {
		return fmt.Errorf("undefined local resolver")
	}

	if config.Default.Proxy.Resolvers[0] == "no" {
		return nil
	}

	const fName = "/etc/resolv.conf"
	file, err := os.OpenFile(fName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("can't open file: %s, error: %w", fName, err)
	}
	defer file.Close()

	io.WriteString(file, "#Generated by meteo daemon\n")

	for _, resolver := range config.Default.Proxy.Resolvers {
		nameserver := fmt.Sprintf("nameserver %s\n", resolver)
		io.WriteString(file, nameserver)
	}

	return nil
}

func (p proxyAPI) verifyServerState() {
	for {
		timer := time.After(p.dns.GetTimerVerifyRetry())
		select {
		case <-timer:
			p.dns.VerifyState()
		case <-ctx.Done():
			return
		}
	}
}

func (p proxyAPI) SyncZones(host entities.Homezone) {
	if len(host.Address) == 0 {
		p.dns.GetZones().RemoveHost(host.Name)
	} else {
		p.dns.GetZones().UpdateHost(host.Name, host.Address)
	}
}

func (p proxyAPI) IsStarted() bool {
	return started
}
