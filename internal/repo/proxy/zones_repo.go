package repo

import (
	"bytes"
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/utils"
	"net"
	"sort"
	"time"

	"github.com/tatsushid/go-fastping"
)

func (p proxyService) GetAllHomeZoneHosts() (*[]entities.Homezone, error) {
	var hosts []entities.Homezone
	err := p.db.Order("address").Find(&hosts).Error
	if err != nil {
		return nil, fmt.Errorf("error read homezones: %w", err)
	}
	sort.Slice(hosts, func(i, j int) bool {
		iIP := net.ParseIP(hosts[i].Address)
		jIP := net.ParseIP(hosts[j].Address)
		return bytes.Compare(iIP, jIP) < 0
	})
	err = pinger(hosts)
	if err != nil {
		return nil, fmt.Errorf("pinger error: %w", err)
	}
	return &hosts, nil
}

func (p proxyService) AddHomeZoneHost(host entities.Homezone) (uint32, error) {
	tx := p.db.Begin()
	host.ID = utils.HashString32(fmt.Sprintf("%s%s", host.Name, host.Address))
	err := tx.Create(&host).Error
	if err != nil {
		return 0, fmt.Errorf("error insert homezones: %w", err)
	}
	_, err = kit.PutInt("/proxy/zones/update", nil)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("proxy internal error: %w", err)
	}
	tx.Commit()
	err = p.routerZonesUpdate()
	if err != nil {
		return 0, fmt.Errorf("router zones update: %w", err)
	}
	return host.ID, nil
}

func (p proxyService) EditHomeZoneHost(host entities.Homezone) error {
	tx := p.db.Begin()
	err := p.db.Save(&host).Error
	if err != nil {
		return fmt.Errorf("error update homezones: %w", err)
	}
	_, err = kit.PutInt("/proxy/zones/update", nil)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("proxy internal error: %w", err)
	}
	tx.Commit()
	err = p.routerZonesUpdate()
	if err != nil {
		return fmt.Errorf("router zones update: %w", err)
	}
	return nil
}

func (p proxyService) DelHomeZoneHost(id uint32) error {
	tx := p.db.Begin()
	err := tx.Delete(&entities.Homezone{ID: id}).Where("id=?", id).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete homezones: %w", err)
	}
	_, err = kit.PutInt("/proxy/zones/update", nil)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("proxy internal error: %w", err)
	}
	tx.Commit()
	err = p.routerZonesUpdate()
	if err != nil {
		return fmt.Errorf("router zones update: %w", err)
	}
	return nil
}

func (p proxyService) routerZonesUpdate() error {
	var hosts []entities.Homezone
	err := p.db.Order("address").Find(&hosts).Error
	if err != nil {
		return fmt.Errorf("error read homezones: %w", err)
	}
	_, err = kit.PostInt("/sshclient/mikrotik/router/zones/update", hosts)
	if err != nil {
		return fmt.Errorf("sshclient internal error: %w", err)
	}
	return nil
}

func pinger(hosts []entities.Homezone) error {
	ips := map[string]bool{}
	for _, host := range hosts {
		ips[host.Address] = false
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
			hosts[i].Active = ips[host.Address]
		}
	}
	err := p.Run()
	if err != nil {
		return fmt.Errorf("ip ping error: %w", err)
	}
	return nil
}
