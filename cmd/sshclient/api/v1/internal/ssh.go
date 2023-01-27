package Ssh

import (
	"encoding/base64"
	"errors"
	"fmt"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/log"
	"net"
	"os"
	"strings"
	"time"

	repo "meteo/internal/repo/sshclient"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

type SshLink interface {
	Config() *ssh.ClientConfig
	Close()
	Exec(cmd string, wait time.Duration) (req string, err error)
	ExecPack(commands []string, wait, pause time.Duration) error
}

type sshlink struct {
	link     string
	username string
	host     string
	port     string
	cnf      *ssh.ClientConfig
	conn     *ssh.Client
}

var repozitory repo.SshClientService = nil

func SetRepozitory(r repo.SshClientService) {
	repozitory = r
}

func NewSSHLink(address, username string) (SshLink, error) {
	split := strings.Split(address, ":")
	if len(split) != 2 || repozitory == nil {
		return nil, fmt.Errorf("invalid address")
	}
	c, err := GetConfig(username, split[0])
	if err != nil {
		return nil, fmt.Errorf("ssh config error: %w", err)
	}
	conn, err := ssh.Dial("tcp", address, c)
	if err != nil {
		return nil, fmt.Errorf("ssh dial error: %w", err)
	}

	link := &sshlink{
		link:     address,
		username: username,
		host:     split[0],
		port:     split[1],
		cnf:      c,
		conn:     conn,
	}

	return link, nil
}

func GetConfig(username, host string) (*ssh.ClientConfig, error) {
	var keyErr *knownhosts.KeyError
	kh, err := checkKnownHosts()
	if err != nil {
		return nil, fmt.Errorf("check knownhosts erro:r %w", err)
	}

	return &ssh.ClientConfig{
		User: username,
		Auth: makeKeyring(host),
		HostKeyCallback: ssh.HostKeyCallback(func(host string, remote net.Addr, pubKey ssh.PublicKey) error {
			hErr := kh(host, remote, pubKey)
			if errors.As(hErr, &keyErr) && len(keyErr.Want) > 0 {
				log.Infof("WARNING: %v is not a key of %s, either a MiTM attack or %s has reconfigured the host pub key.", hostKeyString(pubKey), host, host)
				return keyErr
			} else if errors.As(hErr, &keyErr) && len(keyErr.Want) == 0 {
				log.Infof("WARNING: %s is not trusted, adding this key: %q to known_hosts file.", host, hostKeyString(pubKey))
				return addHostKey(host, remote, pubKey)
			}
			return nil
		}),
	}, nil
}

func checkKnownHosts() (ssh.HostKeyCallback, error) {
	fName, err := createKnownHosts()
	if err != nil {
		return nil, fmt.Errorf("create knownhost file error: %w", err)
	}
	return knownhosts.New(fName)
}

func createKnownHosts() (fName string, err error) {

	rows, err := repozitory.GetAllSshHosts(dto.Pageable{})
	if err != nil {
		return fName, fmt.Errorf("get knownhosts error:% w", err)
	}
	f, err := os.CreateTemp("/tmp", "knownhosts")
	if err != nil {
		return fName, fmt.Errorf("can't open file %s error: %w", f.Name(), err)
	}
	defer f.Close()
	for _, row := range rows {
		_, err = f.WriteString(fmt.Sprintf("%s\n", row.Finger))
		if err != nil {
			return fName, fmt.Errorf("can't write file %s error: %w", f.Name(), err)
		}
	}
	return f.Name(), nil
}

func makeKeyring(host string) (signers []ssh.AuthMethod) {
	if keys, err := repozitory.GetSshKeysByHost(host); err != nil {
		log.Errorf("Get keys from repo fail: %v", err)
	} else {
		for _, key := range keys {
			signer, err := func(finger string) (signer ssh.Signer, err error) {
				signer, err = ssh.ParsePrivateKey([]byte(finger))
				return signer, err
			}(key.Finger)
			if err != nil {
				log.Errorf("Key make error from signature: %s, %v", key.Owner, err)
			} else {
				signers = append(signers, ssh.PublicKeys(signer))
				err := repozitory.UpTimeSshKey(key.Owner)
				if err != nil {
					log.Errorf("Up time key error: %v", err)
				}
			}
		}
	}
	return
}

func addHostKey(host string, remote net.Addr, pubKey ssh.PublicKey) error {

	knownHost := knownhosts.Normalize(remote.String())
	finger := knownhosts.Line([]string{knownHost}, pubKey)

	_, err := repozitory.AddSshHost(entities.SshHosts{Host: knownHost, Finger: finger})
	if err != nil {
		return fmt.Errorf("add knownhost fail: %w", err)
	}
	return nil
}

func SaveConfig(fName string, body string) error {
	f, err := os.OpenFile(fName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("unable open file [%s], error: %w", fName, err)
	}
	_, err = f.WriteString(body)
	if err != nil {
		return fmt.Errorf("unable write file [%s], error: %w", fName, err)
	}
	f.Close()
	return nil
}

func hostUsed(hostname string) error {
	err := repozitory.UpTimeSshHosts(hostname)
	if err != nil {
		return fmt.Errorf("update timestamp sshhost fail: %w", err)
	}
	return nil
}

func hostKeyString(k ssh.PublicKey) string {
	return k.Type() + " " + base64.StdEncoding.EncodeToString(k.Marshal())
}
