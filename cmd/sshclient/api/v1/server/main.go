package server

import (
	"fmt"
	SSH "meteo/cmd/sshclient/api/v1/internal"
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

func (p serverAPI) ServerShutdown(adderess string, port uint, username, cmd string) error {

	bind := fmt.Sprintf("%s:%d", adderess, port)

	link, err := SSH.NewSSHLink(bind, username)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	_, err = link.Exec(cmd, 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}

	return nil
}
