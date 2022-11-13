package Ssh

import (
	repo "meteo/internal/repo/sshclient"

	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"
)

// SshClient api interface
type SshClient interface {
	GitPush(repo string) error
	SftpDownload(remoteFile, localFile, hostname string, port uint, config *ssh.ClientConfig) error
	SftpUpload(remoteFile, localFile, hostname string, port uint, config *ssh.ClientConfig) error
	GetFinger(username, host string) (string, error)
}

type sshclient struct {
	repo repo.SshClientService
}

// NewSshClient get sshclient service instance
func NewSshClient(db *gorm.DB) SshClient {
	r := repo.NewSshClientService(db)
	SetRepozitory(r)
	return &sshclient{
		repo: r,
	}
}
