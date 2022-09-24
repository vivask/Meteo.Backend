package service

import (
	"bytes"
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/log"
	repo "meteo/internal/repo/proxy"
	"net"
	"sort"
	s "strings"

	"github.com/mostlygeek/arp"
)

type Host struct {
	Name string
	IP   net.IP
}

type HostList []Host

type Zones struct {
	repo repo.ProxyService
	data map[string]string
}

func NewZone(r repo.ProxyService) *Zones {
	return &Zones{
		repo: r,
		data: make(map[string]string),
	}
}

func (z *Zones) AddToTable(name, ip string) error {
	name = s.Trim(name, " ")
	name = s.ToLower(name)
	ip = s.Trim(ip, " ")
	if len(name) == 0 || len(ip) == 0 {
		return fmt.Errorf("invalid values [%s] or [%s]", name, ip)
	}
	host := entities.Homezone{
		DomainName: name,
		IPv4:       ip,
	}
	mac := arp.Search(ip)
	if len(mac) > 0 {
		host.Mac = mac
	}
	err := z.repo.AddHomeZoneHost(host)
	if err != nil {
		return fmt.Errorf("create host error: %w", err)
	}

	return nil
}

func (z *Zones) Parse(address string) (name, ip string, err error) {

	split := s.Split(address, "/")
	if split[0] == "address=" {
		server := s.Trim(split[1], " ")
		ip := s.Trim(split[2], " ")
		if len(server) == 0 || len(ip) == 0 {
			return "", "", fmt.Errorf("invalid values [%s] or [%s]", server, ip)
		}
		return server, ip, nil
	}

	return "", "", fmt.Errorf("invalid input data %v", split)
}

func (z *Zones) RemoveHost(host string) bool {
	if !s.HasSuffix(host, ".") {
		host += "."
	}
	host = s.ToLower(host)
	if _, exist := z.data[host]; exist {
		delete(z.data, host)
		return true
	}
	return false
}

func (z *Zones) AddHost(host, ip string) bool {
	if !s.HasSuffix(host, ".") {
		host += "."
	}
	host = s.ToLower(host)
	if _, exist := z.data[host]; !exist {
		z.data[host] = ip
		return true
	}
	return false
}

func (z *Zones) UpdateHost(host, ip string) bool {
	if !s.HasSuffix(host, ".") {
		host += "."
	}
	host = s.ToLower(host)
	z.data[host] = ip
	return true
}

func (z *Zones) Address(host string) string {
	ip, ok := z.data[host]
	if ok {
		return ip
	}
	return ""
}

func LoadZones(repo repo.ProxyService) (list *Zones) {
	list = NewZone(repo)
	hosts, err := repo.GetAllHomeZoneHosts()
	if err != nil {
		log.Errorf("get host from repo error: %v", err)
		return
	}
	count := 0
	for _, host := range *hosts {
		if list.AddHost(host.DomainName, host.IPv4) {
			count++
		} else {
			log.Warningf("Bad host name: %s, or ip: %s", host.DomainName, host.IPv4)
		}
	}
	log.Infof("Loaded %d hosts for home zone", count)
	return
}

func (z *Zones) Contains(host string) bool {
	_, ok := z.data[host]
	return ok
}

func rankHosts(hosts map[string]string) HostList {
	pl := make(HostList, len(hosts))
	i := 0
	for host, ip := range hosts {
		ip := net.ParseIP(ip)
		pl[i] = Host{host, ip}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		return bytes.Compare(pl[i].IP, pl[j].IP) < 0
	})
	return pl
}
