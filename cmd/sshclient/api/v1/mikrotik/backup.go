package mikrotik

import (
	"context"
	"fmt"
	"meteo/internal/config"
	"meteo/internal/log"
	"meteo/internal/utils"
	"os"
	"path/filepath"
	"regexp"
	s "strings"
	"time"

	SSH "meteo/cmd/sshclient/api/v1/internal"

	"golang.org/x/sync/errgroup"
)

const DOWNLOAD_TIMEOUT = 20

func (p mikrotikAPI) MikrotikBackup(adderess string, port uint, username string) error {

	bind := fmt.Sprintf("%s:%d", adderess, port)

	link, err := SSH.NewSSHLink(bind, username)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	verbose, err := link.Exec("/export verbose", 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}

	if len(verbose) == 0 {
		return fmt.Errorf("received a zero length response from the host: %s", adderess)
	}
	idx := s.Index(verbose, "\n")
	verbose = utils.LeftTrunc(verbose, idx+1)
	configFile := config.Default.SshClient.Mikrotik.Repository + "/" + adderess
	_, err = os.Stat(configFile)
	if err == nil {
		savedMD5, err := utils.GetMD5FileSum(configFile)
		if err != nil {
			return fmt.Errorf("GetMD5FileSum error: %w", err)
		}
		receivedMD5 := utils.GetMD5StringlnSum(verbose)
		if savedMD5 != receivedMD5 {
			log.Infof("Backuped: %s", configFile)
			err = SSH.SaveConfig(configFile, verbose)
			if err != nil {
				return err
			}
			err = p.downloadBackupFile(adderess, port)
			if err != nil {
				return err
			}
		}
	} else {
		log.Infof("Backuped: %s", configFile)
		err = SSH.SaveConfig(configFile, verbose)
		if err != nil {
			return err
		}
		err = p.downloadBackupFile(adderess, port)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p mikrotikAPI) Backup() (err error) {

	g, _ := errgroup.WithContext(context.Background())
	for i := range config.Default.SshClient.Mikrotik.Hosts {
		i := i
		g.Go(func() error {
			return p.MikrotikBackup(config.Default.SshClient.Mikrotik.Hosts[i],
				config.Default.SshClient.Mikrotik.Ports[i],
				config.Default.SshClient.Mikrotik.Users[i])
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("MikrotikBackup routine error: %w", err)
	}

	err = p.ssh.GitPush(config.Default.SshClient.Mikrotik.Repository)
	if err != nil {
		return fmt.Errorf("gitPush error: %w", err)
	}
	files, err := filepath.Glob(config.Default.SshClient.Mikrotik.Repository + "/*.backup")
	if err != nil {
		return fmt.Errorf("glob error: %w", err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return fmt.Errorf("remove file [%s] error: %w", f, err)
		}
	}
	return nil
}

func (p mikrotikAPI) downloadBackupFile(host string, port uint) error {

	address := fmt.Sprintf("%s:%d", host, port)

	link, err := SSH.NewSSHLink(address, config.Default.SshClient.Backup.User)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()
	fileName := fmt.Sprintf("%s-%v", host, time.Now().Format(time.RFC3339))
	cmd := fmt.Sprintf("/system backup save name=%s", fileName)
	req, err := link.Exec(cmd, 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	matched, _ := regexp.MatchString("Configuration backup saved", req)
	if !matched {
		return fmt.Errorf("ssh response error: %s", req)
	}

	backupFile := fileName + ".backup"
	sftpReq := make(chan error)
	repository := fmt.Sprintf("%s/%s", config.Default.SshClient.Mikrotik.Repository, backupFile)
	go func(sftpReq chan error) {
		sftpReq <- p.ssh.SftpDownload(backupFile, repository, host, port, link.Config())
	}(sftpReq)
	timeout := time.After(DOWNLOAD_TIMEOUT * time.Second)
	select {
	case sftpErr := <-sftpReq:
		if sftpErr != nil {
			return fmt.Errorf("unable to download remote file: %s", backupFile)
		}
	case <-timeout:
		return fmt.Errorf("iimed out file download from: %s", host)
	}
	cmd = fmt.Sprintf("/file remove \"%s\"", backupFile)
	_, err = link.Exec(cmd, 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	log.Debugf("backup file download success")
	return nil
}
