package repo

import (
	"bytes"
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/utils"
	"net"
	"sort"
	"time"

	"github.com/tatsushid/go-fastping"
)

func (p proxyService) AddHomeZoneHost(host entities.Homezone) error {
	host.ID = utils.HashString32(fmt.Sprintf("%s%s", host.DomainName, host.IPv4))
	err := p.db.Create(&host).Error
	if err != nil {
		return fmt.Errorf("error insert homezones: %w", err)
	}
	return nil
}

func (p proxyService) GetAllHomeZoneHosts() (*[]entities.Homezone, error) {
	var hosts []entities.Homezone
	err := p.db.Order("ip").Find(&hosts).Error
	if err != nil {
		return nil, fmt.Errorf("error read homezones: %w", err)
	}
	sort.Slice(hosts, func(i, j int) bool {
		iIP := net.ParseIP(hosts[i].IPv4)
		jIP := net.ParseIP(hosts[j].IPv4)
		return bytes.Compare(iIP, jIP) < 0
	})
	err = pinger(hosts)
	if err != nil {
		return nil, fmt.Errorf("pinger error: %w", err)
	}
	return &hosts, nil
}

func pinger(hosts []entities.Homezone) error {
	ips := map[string]bool{}
	for _, host := range hosts {
		ips[host.IPv4] = false
	}

	p := fastping.NewPinger()
	for k := range ips {
		ra, err := net.ResolveIPAddr("ip4:icmp", k)
		if err != nil {
			return fmt.Errorf("ip parse error: %w", err)
		}
		p.AddIPAddr(ra)
	}
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		ips[addr.String()] = true
	}
	p.OnIdle = func() {
		for i, host := range hosts {
			hosts[i].Active = ips[host.IPv4]
		}
	}
	err := p.Run()
	if err != nil {
		return fmt.Errorf("ip ping error: %w", err)
	}
	return nil
}
