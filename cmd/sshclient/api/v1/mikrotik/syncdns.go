package mikrotik

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"net/http"
	"os"
	"regexp"

	SSH "meteo/cmd/sshclient/api/v1/internal"

	"github.com/gin-gonic/gin"
)

func (p mikrotikAPI) RouterSyncDNS(c *gin.Context) {

	var hosts []entities.Homezone
	if err := c.ShouldBind(&hosts); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.SyncDNS(&hosts)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mikrotikAPI) SyncDNS(hosts *[]entities.Homezone) error {

	bind := fmt.Sprintf("%s:%d", config.Default.SshClient.PPP.Host, config.Default.SshClient.PPP.Port)

	link, err := SSH.NewSSHLink(bind, config.Default.SshClient.PPP.User)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	_, err = link.Exec("/system script run remove-all-static-dns", 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	fName, err := prepareFile(hosts)
	if err != nil {
		return fmt.Errorf("prepare file error: %w", err)
	}
	err = p.ssh.SftpUpload("static-dns", fName, config.Default.SshClient.PPP.Host, config.Default.SshClient.PPP.Port, link.Config())
	os.Remove(fName)
	if err != nil {
		return fmt.Errorf("file upload on: %s, error: %w", config.Default.SshClient.PPP.Host, err)
	}

	req, err := link.Exec("/import file-name=static-dns", 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	matched, _ := regexp.MatchString("Script file loaded and executed successfully", req)
	if !matched {
		return fmt.Errorf("ssh response error: %s", req)
	}

	_, err = link.Exec("/file remove static-dns", 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	return nil
}

func prepareFile(hosts *[]entities.Homezone) (string, error) {

	file, err := os.CreateTemp("/tmp", "zones")
	if err != nil {
		return "", fmt.Errorf("can't create tmp file: %w", err)
	}
	defer file.Close()

	for _, host := range *hosts {
		str := fmt.Sprintf("/ip dns static add name=%s address=%s\n", host.Name, host.Address)
		_, err := file.WriteString(str)
		if err != nil {
			os.Remove(file.Name())
			return "", fmt.Errorf("can't write to file [%s] error: %w", file.Name(), err)
		}
	}
	return file.Name(), nil
}
