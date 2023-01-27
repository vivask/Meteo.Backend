package mikrotik

import (
	"fmt"
	SSH "meteo/cmd/sshclient/api/v1/internal"
	"meteo/internal/config"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const REQUEST_WAIT = 5 // in seconds

func (p mikrotikAPI) AddManualHostToVpn(c *gin.Context) {

	var host entities.ToVpnManual
	if err := c.ShouldBind(&host); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.PutInVpn(host)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mikrotikAPI) RemoveManualHostFromVpn(c *gin.Context) {

	var host entities.ToVpnManual
	if err := c.ShouldBind(&host); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.RemoveFromVpn(host)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mikrotikAPI) AddAutoHostToVpn(c *gin.Context) {

	var auto entities.ToVpnAuto
	if err := c.ShouldBind(&auto); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}
	err := p.PutInVpn(entities.ToVpnManual{Name: auto.ID, AccesList: entities.AccesList{ID: "tovpn"}})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mikrotikAPI) RemoveAutoHostFromVpn(c *gin.Context) {

	var auto entities.ToVpnAuto
	if err := c.ShouldBind(&auto); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.RemoveFromVpn(entities.ToVpnManual{Name: auto.ID, AccesList: entities.AccesList{ID: "tovpn"}})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mikrotikAPI) PutInVpn(host entities.ToVpnManual) error {

	bind := fmt.Sprintf("%s:%d", config.Default.SshClient.Vpn.Host, config.Default.SshClient.Vpn.Port)

	link, err := SSH.NewSSHLink(bind, config.Default.SshClient.Vpn.User)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	name := strings.TrimSuffix(host.Name, ".")
	if len(host.ListID) == 0 {
		host.ListID = config.Default.SshClient.Vpn.ListVpn
	}

	req, err := link.Exec("/ip firewall address-list print", REQUEST_WAIT)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	//log.Debugf("REQ: %s", req)
	split := strings.Split(req, "\n")
	found := "-1"
	for _, s := range split {
		words := strings.Fields(s)
		if len(words) > 2 && words[1] == host.AccesList.ID && words[2] == host.Name {
			log.Debugf("ID: %s, List: %s, Name: %s", words[0], words[1], words[2])
			found = words[0]
		}
	}

	if found == "-1" {
		cmd := fmt.Sprintf("/ip firewall address-list add list=%s address=%s", host.AccesList.ID, name)
		_, err = link.Exec(cmd, 1)
		if err != nil {
			return fmt.Errorf("ssh exec error: %w", err)
		}
	}

	return nil
}

func (p mikrotikAPI) RemoveFromVpn(host entities.ToVpnManual) error {

	bind := fmt.Sprintf("%s:%d", config.Default.SshClient.Vpn.Host, config.Default.SshClient.Vpn.Port)

	link, err := SSH.NewSSHLink(bind, config.Default.SshClient.Vpn.User)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	if len(host.ListID) == 0 {
		host.ListID = config.Default.SshClient.Vpn.ListVpn
	}

	req, err := link.Exec("/ip firewall address-list print", REQUEST_WAIT)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	//log.Debugf("REQ: %s", req)
	split := strings.Split(req, "\n")
	found := "-1"
	for _, s := range split {
		words := strings.Fields(s)
		if len(words) > 2 && words[1] == host.AccesList.ID && words[2] == host.Name {
			log.Debugf("ID: %s, List: %s, Name: %s", words[0], words[1], words[2])
			found = words[0]
		}
	}
	if found == "-1" {
		log.Errorf("not found host [%s]", host.Name)
		return fmt.Errorf("not found host in access list")
	}

	cmd := fmt.Sprintf("/ip firewall address-list remove %s", found)
	_, err = link.Exec(cmd, 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	return nil
}

func (p mikrotikAPI) VpnListsSync(c *gin.Context) {

	err := p.SyncVpnLists()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mikrotikAPI) SyncVpnLists() error {
	bind := fmt.Sprintf("%s:%d", config.Default.SshClient.Vpn.Host, config.Default.SshClient.Vpn.Port)

	link, err := SSH.NewSSHLink(bind, config.Default.SshClient.Vpn.User)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	err = ClearAddressLists(link)
	if err != nil {
		return fmt.Errorf("address lists clear error: %w", err)
	}

	err = p.CreateAdressLists(link)
	if err != nil {
		return fmt.Errorf("create address lists error: %w", err)
	}

	return nil
}

func ClearAddressLists(link SSH.SshLink) error {

	req, err := link.Exec("/ip firewall address-list print", REQUEST_WAIT)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}

	count := 0
	split := strings.Split(req, "\n")
	for _, s := range split {
		words := strings.Fields(s)
		if len(words) > 0 {
			if _, err := strconv.Atoi(words[0]); err == nil {
				count++
			}
		}
	}

	for i := count - 1; i >= 0; i-- {
		cmd := fmt.Sprintf("/ip firewall address-list remove %d", i)
		_, err = link.Exec(cmd, 1)
		if err != nil {
			return fmt.Errorf("ssh exec error: %w", err)
		}
	}

	return nil
}

func (p mikrotikAPI) CreateAdressLists(link SSH.SshLink) error {
	manual, err := p.repo.GetAllManualToVpn(dto.Pageable{})
	if err != nil {
		return err
	}

	for _, host := range *manual {
		cmd := fmt.Sprintf("/ip firewall address-list add list=%s address=%s", host.AccesList.ID, host.Name)
		_, err = link.Exec(cmd, 1)
		if err != nil {
			return fmt.Errorf("ssh exec error: %w", err)
		}
	}

	auto, err := p.repo.GetAllAutoToVpn(dto.Pageable{})
	if err != nil {
		return err
	}
	for _, host := range *auto {
		cmd := fmt.Sprintf("/ip firewall address-list add list=%s address=%s", config.Default.SshClient.Vpn.ListVpn, host.ID)
		_, err = link.Exec(cmd, 1)
		if err != nil {
			return fmt.Errorf("ssh exec error: %w", err)
		}
	}
	return nil
}
