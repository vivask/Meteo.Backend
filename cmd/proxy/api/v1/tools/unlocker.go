package tools

import (
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/log"
	repo "meteo/internal/repo/proxy"
	"strings"
)

type Unlocker struct {
	repo     repo.ProxyService
	unlocked map[string]struct{}
	ignored  map[string]struct{}
}

func NewUnlocker(r repo.ProxyService) *Unlocker {

	return &Unlocker{
		repo:     r,
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
	ignor, err := un.repo.GetAllIgnoreAutoToVpn(dto.Pageable{})
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

func LoadUnlocker(repo repo.ProxyService) *Unlocker {

	list := NewUnlocker(repo)
	unlocked, ignored := list.LoadHosts()
	log.Info("Loaded ", unlocked, " unlocked hosts from database, ignore: ", ignored)
	return list
}

func (un *Unlocker) AddAutoHostToVpn(name string) error {
	if !un.AddHost(name) {
		return fmt.Errorf("Host [%s] exist, can't insert", name)
	}
	err := un.repo.AddAutoToVpn(&entities.ToVpnAuto{ID: name})
	if err != nil {
		return fmt.Errorf("can't insert records to tovpn_autos: %w", err)
	}
	return nil
}
