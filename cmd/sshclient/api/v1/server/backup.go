package server

import (
	"fmt"
	Ssh "meteo/cmd/sshclient/api/v1/internal"
	"meteo/internal/config"
	"regexp"
	"strings"
)

func (p serverAPI) KodiRestart() error {

	address := fmt.Sprintf("%s:%d", config.Default.SshClient.Backup.Host, config.Default.SshClient.Backup.Port)

	link, err := Ssh.NewSSHLink(address, config.Default.SshClient.Backup.User)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	kodiShPID, err := link.Exec("pidof kodi.sh", 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	kodiBinPID, err := link.Exec("pidof kodi.bin", 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}

	cmd := fmt.Sprintf("kill -9 %s %s",
		strings.Replace(kodiBinPID, "\n", "", -1),
		strings.Replace(kodiShPID, "\n", "", -1))
	_, err = link.Exec(cmd, 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}

	return nil
}

func (p serverAPI) KodiStorageRestart() error {

	address := fmt.Sprintf("%s:%d", config.Default.SshClient.Backup.Host, config.Default.SshClient.Backup.Port)

	link, err := Ssh.NewSSHLink(address, config.Default.SshClient.Backup.User)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	commands := []string{
		"systemctl restart storage-torrents.mount",
		"systemctl restart storage-backup.mount",
		"systemctl restart storage-media.mount",
	}

	err = link.ExecPack(commands, 20, 500)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	return nil
}

func (p serverAPI) KodiStorageStop() error {

	address := fmt.Sprintf("%s:%d", config.Default.SshClient.Backup.Host, config.Default.SshClient.Backup.Port)

	link, err := Ssh.NewSSHLink(address, config.Default.SshClient.Backup.User)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	commands := []string{
		"systemctl stop storage-torrents.mount",
		"systemctl stop storage-backup.mount",
		"systemctl stop storage-media.mount",
	}

	err = link.ExecPack(commands, 20, 500)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	return nil
}

func (p serverAPI) KodiStorageStart() error {

	address := fmt.Sprintf("%s:%d", config.Default.SshClient.Backup.Host, config.Default.SshClient.Backup.Port)

	link, err := Ssh.NewSSHLink(address, config.Default.SshClient.Backup.User)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	commands := []string{
		"systemctl start storage-torrents.mount",
		"systemctl start storage-backup.mount",
		"systemctl start storage-media.mount",
	}

	err = link.ExecPack(commands, 20, 500)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	return nil
}

func (p serverAPI) StorageHealth() error {
	address := fmt.Sprintf("%s:%d", config.Default.SshClient.Backup.Host, config.Default.SshClient.Backup.Port)

	link, err := Ssh.NewSSHLink(address, config.Default.SshClient.Backup.User)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	out, err := link.Exec("df", 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}

	matched, _ := regexp.MatchString("/storage/media", out)
	if !matched {
		return fmt.Errorf("/storage/media not available")
	}

	return nil
}
