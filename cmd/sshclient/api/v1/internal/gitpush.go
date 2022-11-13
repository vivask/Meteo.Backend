package Ssh

import (
	"fmt"
	"meteo/internal/config"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func (p sshclient) GitPush(repo string) error {
	r, err := git.PlainOpen(repo)
	if err != nil {
		return fmt.Errorf("unable open git repo: %w", err)
	}

	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("unable read git worktree: %w", err)
	}

	status, _ := w.Status()
	if status.IsClean() {
		return fmt.Errorf("unable clean git status: %w", err)
	}

	_, err = w.Add(".")
	if err != nil {
		return fmt.Errorf("unable add git: %w", err)
	}

	commit, err := w.Commit(config.Default.SshClient.Git.Commit, &git.CommitOptions{
		Author: &object.Signature{
			Name: config.Default.SshClient.Git.User,
			When: time.Now(),
		},
	})
	if err != nil {
		return err
	}
	_, err = r.CommitObject(commit)
	if err != nil {
		return fmt.Errorf("unable commit git: %w", err)
	}

	/*host := fmt.Sprintf("%s:%d", config.Default.SshClient.Git.Host, config.Default.SshClient.Git.Port)
	auth, err := getSSHAuth(host)
	if err != nil {
		return fmt.Errorf("git authentication fail: %w", err)
	}*/

	auth, err := p.getHTTPAuth(config.Default.SshClient.Git.Repository)
	if err != nil {
		return fmt.Errorf("git authentication fail: %w", err)
	}

	err = r.Push(&git.PushOptions{RemoteName: config.Default.SshClient.Git.Remote, Auth: auth})
	if err != nil {
		return fmt.Errorf("unable push git: %w", err)
	}

	//return p.usedKey(host)
	return p.usedHTTPAccount(config.Default.SshClient.Git.Repository)
}

/*func (p sshclient) makeKey(host string) (signer ssh.Signer, err error) {
	if row, err := p.repo.GetKeyGitByOwner(host); err != nil {
		return nil, fmt.Errorf("get key from repo: %w", err)
	} else {
		signer, err = ssh.ParsePrivateKey([]byte(row.Finger))
		if err != nil {
			return nil, fmt.Errorf("parse key error: %w", err)
		}
	}
	return
}*/

/*func (p sshclient) usedKey(host string) error {
	err := p.repo.UpTimeSshKey(host)
	if err != nil {
		return fmt.Errorf("update timestamp gitKeys: %w", err)
	}
	return nil
}*/

func (p sshclient) usedHTTPAccount(service string) error {
	err := p.repo.UpTimeGitUsers(service)
	if err != nil {
		return fmt.Errorf("update timestamp gitUsers: %w", err)
	}
	return nil
}

/*
func (p sshclient) getSSHAuth(host string) (*gitssh.PublicKeys, error) {
	singer, err := p.makeKey(host)
	if err != nil {
		return nil, fmt.Errorf("make key error: %w", err)
	}

	var keyErr *knownhosts.KeyError
	kh, err := p.checkKnownHosts()
	if err != nil {
		return nil, fmt.Errorf("check knownhosts erro:r %w", err)
	}

	auth := &gitssh.PublicKeys{
		Signer: singer,
		/*HostKeyCallbackHelper: gitssh.HostKeyCallbackHelper{HostKeyCallback: ssh.InsecureIgnoreHostKey()},*/
/*HostKeyCallbackHelper: gitssh.HostKeyCallbackHelper{
			HostKeyCallback: ssh.HostKeyCallback(func(host string, remote net.Addr, pubKey ssh.PublicKey) error {
				hErr := kh(host, remote, pubKey)
				if errors.As(hErr, &keyErr) && len(keyErr.Want) > 0 {
					log.Infof("WARNING: %v is not a key of %s, either a MiTM attack or %s has reconfigured the host pub key.", hostKeyString(pubKey), host, host)
					return keyErr
				} else if errors.As(hErr, &keyErr) && len(keyErr.Want) == 0 {
					log.Infof("WARNING: %s is not trusted, adding this key: %q to known_hosts file.", host, hostKeyString(pubKey))
					return p.addHostKey(host, remote, pubKey)
				}
				return nil
			}),
		},
	}

	return auth, nil
}*/

func (p sshclient) getHTTPAuth(service string) (*http.BasicAuth, error) {

	row, err := p.repo.GetUserKeyByService(service)
	if err != nil {
		return nil, fmt.Errorf("not found account for service: %s, error: %w", service, err)
	}
	auth := &http.BasicAuth{
		Username: row.Username,
		Password: row.Password,
	}
	return auth, nil
}
