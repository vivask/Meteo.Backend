package Ssh

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func (p sshclient) SftpDownload(remoteFile, localFile, hostname string, port uint, config *ssh.ClientConfig) error {
	link := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := ssh.Dial("tcp", link, config)
	if err != nil {
		return fmt.Errorf("ssh connect error: %w", err)
	}
	defer conn.Close()

	sc, err := sftp.NewClient(conn)
	if err != nil {
		return fmt.Errorf("unable to start SFTP subsystem: %w", err)
	}
	defer sc.Close()

	srcFile, err := sc.OpenFile(remoteFile, (os.O_RDONLY))
	if err != nil {
		return fmt.Errorf("unable to open remote file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(localFile)
	if err != nil {
		return fmt.Errorf("unable to open local file: %w", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("unable to download remote file: %w", err)
	}

	return nil
}

func (p sshclient) SftpUpload(remoteFile, localFile, hostname string, port uint, config *ssh.ClientConfig) error {
	link := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := ssh.Dial("tcp", link, config)
	if err != nil {
		return fmt.Errorf("ssh connect error: %w", err)
	}
	defer conn.Close()

	sc, err := sftp.NewClient(conn)
	if err != nil {
		return fmt.Errorf("unable to start SFTP subsystem: %w", err)
	}
	defer sc.Close()

	srcFile, err := os.Open(localFile)
	if err != nil {
		return fmt.Errorf("unable to open local file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := sc.Create(remoteFile)
	if err != nil {
		return fmt.Errorf("unable to open remote file: %w", err)
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("unable copy file: %w", err)
	}

	return nil
}
