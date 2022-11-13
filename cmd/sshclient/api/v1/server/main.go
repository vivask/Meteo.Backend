package server

import (
	"fmt"
	SSH "meteo/cmd/sshclient/api/v1/internal"
	"meteo/internal/config"
	"strings"
)

func (p serverAPI) ContainerReboot(adderess string, port uint, username, name string) error {

	bind := fmt.Sprintf("%s:%d", adderess, port)

	link, err := SSH.NewSSHLink(bind, username)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	cmd := fmt.Sprintf("docker restart %s", name)
	_, err = link.Exec(cmd, 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	return nil
}

func (p serverAPI) ContainerShutdown(adderess string, port uint, username, name string) error {

	bind := fmt.Sprintf("%s:%d", adderess, port)

	link, err := SSH.NewSSHLink(bind, username)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	cmd := fmt.Sprintf("docker stop %s", name)
	_, err = link.Exec(cmd, 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	return nil
}

func (p serverAPI) ContainerStart(adderess string, port uint, username, name string) error {
	bind := fmt.Sprintf("%s:%d", adderess, port)

	link, err := SSH.NewSSHLink(bind, username)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	cmd := fmt.Sprintf("docker start %s", name)
	_, err = link.Exec(cmd, 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	return nil
}

func (p serverAPI) ServerReboot(adderess string, port uint, username string) error {

	bind := fmt.Sprintf("%s:%d", adderess, port)

	link, err := SSH.NewSSHLink(bind, username)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	cmd := "sudo systemctl reboot"
	if config.Default.App.Server == "backup" {
		cmd = strings.Replace(cmd, "sudo ", "", -1)
	}

	_, err = link.Exec(cmd, 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}

	return nil
}

func (p serverAPI) ServerShutdown(adderess string, port uint, username string) error {

	bind := fmt.Sprintf("%s:%d", adderess, port)

	link, err := SSH.NewSSHLink(bind, username)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	cmd := "sudo systemctl shutdown"
	if config.Default.App.Server == "backup" {
		cmd = strings.Replace(cmd, "sudo ", "", -1)
	}

	_, err = link.Exec(cmd, 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}

	return nil
}
