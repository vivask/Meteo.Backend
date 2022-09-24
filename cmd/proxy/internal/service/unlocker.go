package service

import (
	"fmt"
	"meteo/internal/client"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/log"
	repo "meteo/internal/repo/proxy"
	"strings"
)

type Unlocker struct {
	repo     repo.ProxyService
	cli      *client.Client
	unlocked map[string]struct{}
	ignored  map[string]struct{}
}

func NewUnlocker(r repo.ProxyService, c *client.Client) *Unlocker {

	return &Unlocker{
		repo:     r,
		cli:      c,
		unlocked: make(map[string]struct{}),
		ignored:  make(map[string]struct{}),
	}
}

func (un *Unlocker) LoadHosts() (unlocked int, ignorered int) {
	hosts, err := un.repo.GetAllAutoToVpn(dto.Pageable{})
	if err != nil {
		log.Error(err)
		return
	}
	for _, host := range *hosts {
		name := host.ID
		if !strings.HasSuffix(name, ".") {
			name += "."
		}
		un.unlocked[name] = struct{}{}
		unlocked++
	}
	ignor, err := un.repo.GetAllIgnore(dto.Pageable{})
	if err != nil {
		log.Error(err)
		return
	}
	for _, host := range *ignor {
		name := host.ID
		if !strings.HasSuffix(name, ".") {
			name += "."
		}
		un.ignored[name] = struct{}{}
		ignorered++
	}
	return
}

func (un *Unlocker) Exist(name string) bool {
	_, exist := un.unlocked[name]
	return exist
}

func (un *Unlocker) Ignore(name string) bool {
	_, exist := un.ignored[name]
	return exist
}

func (un *Unlocker) AddIgnore(name string) bool {
	if un.Exist(name) && !un.Ignore(name) {
		un.ignored[name] = struct{}{}
		return true
	}
	return false
}

func (un *Unlocker) AddHost(name string) bool {
	if !un.Exist(name) && !un.Ignore(name) {
		un.unlocked[name] = struct{}{}
		return true
	}
	return false
}

func (un *Unlocker) RemoveHost(name string) {
	if un.Exist(name) {
		delete(un.unlocked, name)
	}
}

func (un *Unlocker) RemoveIgnore(name string) {
	if un.Ignore(name) {
		delete(un.ignored, name)
	}
}

func LoadUnlocker(repo repo.ProxyService, c *client.Client) *Unlocker {

	list := NewUnlocker(repo, c)
	unlocked, ignored := list.LoadHosts()
	log.Info("Loaded ", unlocked, " unlocked hosts from database, ignore: ", ignored)
	return list
}

func (un *Unlocker) InsertManual(host entities.ToVpnManual) error {
	if !un.AddHost(host.Name) {
		return fmt.Errorf("Host [%s] exist, can't insert", host.Name)
	}
	err := un.repo.AddManualToVpn(&host)
	if err != nil {
		return fmt.Errorf("can't insert records to tovpn_manuals: %w", err)
	}
	_, err = un.cli.PutInt("/safe/sshclient/tovpn", host, client.SSHCLIENT)
	if err != nil {
		log.Errorf("error GET sshclient: %v", err)
	}
	return nil
}

func (un *Unlocker) UpdateManual(host entities.ToVpnManual) error {
	err := un.RemoveManual(host.ID)
	if err != nil {
		return fmt.Errorf("RemoveManual error: %w", err)
	}

	err = un.InsertManual(host)
	if err != nil {
		return fmt.Errorf("InsertManual error: %w", err)
	}
	return nil
}

func (un *Unlocker) RemoveManual(id uint32) error {
	host, err := un.repo.GetManualToVpn(id)
	if err != nil {
		return fmt.Errorf("GetManualToVpn error: %w", err)
	}
	un.RemoveHost(host.Name)

	err = un.repo.DelManualToVpn(host.ID)
	if err != nil {
		return fmt.Errorf("DelManualToVpn error: %w", err)
	}

	_, err = un.cli.PutInt("/safe/sshclient/rmvpn", host, client.SSHCLIENT)
	if err != nil {
		log.Errorf("error GET sshclient: %v", err)
	}
	return nil
}

func (un *Unlocker) InsertHost(name string) error {
	if !un.AddHost(name) {
		return fmt.Errorf("Host [%s] exist, can't insert", name)
	}
	err := un.repo.AddAutoToVpn(&entities.ToVpnAuto{ID: name})
	if err != nil {
		return fmt.Errorf("can't insert records to tovpn_autos: %w", err)
	}
	host := entities.ToVpnManual{Name: name}
	_, err = un.cli.PutInt("/safe/sshclient/tovpn", host, client.SSHCLIENT)
	if err != nil {
		log.Errorf("error GET sshclient: %v", err)
	}
	return nil
}

func (un *Unlocker) DeleteHost(id string) error {
	err := un.repo.DelAutoToVpn(id)
	if err != nil {
		return fmt.Errorf("DelAutoToVpn error: %w", err)
	}
	un.RemoveHost(id)
	host := entities.ToVpnManual{Name: id}
	_, err = un.cli.PutInt("/safe/sshclient/rmvpn", host, client.SSHCLIENT)
	if err != nil {
		log.Errorf("error GET sshclient: %v", err)
	}
	return nil
}

func (un *Unlocker) IgnoreHost(name string) error {
	err := un.repo.AddIgnoreToVpn(name)
	if err != nil {
		return fmt.Errorf("AddIgnoreToVpn error: %w", err)
	}
	un.AddIgnore(name)

	err = un.repo.DelAutoToVpn(name)
	if err != nil {
		return fmt.Errorf("DelAutoToVpn error: %w", err)
	}
	un.RemoveHost(name)
	host := entities.ToVpnManual{Name: name}
	_, err = un.cli.PutInt("/safe/sshclient/rmvpn", host, client.SSHCLIENT)
	if err != nil {
		log.Errorf("error GET sshclient: %v", err)
	}

	return nil
}

func (un *Unlocker) DeleteIgnore(host string) error {
	err := un.repo.DelIgnoreToVpn(host)
	if err != nil {
		return fmt.Errorf("DelIgnoreToVpn error: %w", err)
	}
	un.RemoveIgnore(host)
	return nil
}

func (un *Unlocker) RestoreHost(name string) error {
	err := un.repo.RestoreAutoToVpn(name)
	if err != nil {
		return fmt.Errorf("can't insert records to tovpn_autos: %w", err)
	}
	un.AddHost(name)

	host := entities.ToVpnManual{Name: name}
	_, err = un.cli.PutInt("/safe/sshclient/tovpn", host, client.SSHCLIENT)
	if err != nil {
		log.Errorf("error GET sshclient: %v", err)
	}

	err = un.repo.DelIgnoreToVpn(name)
	if err != nil {
		return fmt.Errorf("DelIgnoreToVpn error: %w", err)
	}
	un.RemoveIgnore(name)

	return nil
}
