package Ssh

import (
	"errors"
	"fmt"
	"meteo/internal/log"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

var finger string

func (p sshclient) GetFinger(username, host string) (string, error) {

	var keyErr *knownhosts.KeyError
	kh, err := checkKnownHosts()
	if err != nil {
		return "", fmt.Errorf("check knownhosts error: %w", err)
	}

	cfg := &ssh.ClientConfig{
		User: username,
		Auth: makeKeyring(host),
		HostKeyCallback: ssh.HostKeyCallback(func(host string, remote net.Addr, pubKey ssh.PublicKey) error {
			hErr := kh(host, remote, pubKey)
			if errors.As(hErr, &keyErr) && len(keyErr.Want) > 0 {
				log.Infof("WARNING: %v is not a key of %s, either a MiTM attack or %s has reconfigured the host pub key.", hostKeyString(pubKey), host, host)
				return keyErr
			} else if errors.As(hErr, &keyErr) && len(keyErr.Want) == 0 {
				//log.Infof("WARNING: %s is not trusted, adding this key: %q to known_hosts file.", host, hostKeyString(pubKey))
				touch(host, remote, pubKey)
			}
			return nil
		}),
	}

	link := fmt.Sprintf("%s:22", host)
	conn, err := ssh.Dial("tcp", link, cfg)
	if err == nil {
		conn.Close()
	}

	return finger, nil
}

func touch(host string, remote net.Addr, pubKey ssh.PublicKey) {
	knownHost := knownhosts.Normalize(remote.String())
	finger = knownhosts.Line([]string{knownHost}, pubKey)
}
