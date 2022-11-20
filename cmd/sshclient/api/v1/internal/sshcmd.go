package Ssh

import (
	"bytes"
	"context"
	"fmt"
	"meteo/internal/log"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/sync/errgroup"
)

func (p sshlink) Config() *ssh.ClientConfig {
	return p.cnf
}

func (p sshlink) Close() {
	p.conn.Close()
}

func (p sshlink) Exec(cmd string, wait time.Duration) (req string, err error) {

	session, err := p.conn.NewSession()
	if err != nil {
		return req, fmt.Errorf("ssh session error: %w", err)
	}
	defer session.Close()

	var stdout bytes.Buffer
	session.Stdout = &stdout

	timeout := wait * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	g, _ := errgroup.WithContext(ctx)
	defer cancel()

	g.Go(func() error {
		log.Debugf("CMD: %s", cmd)
		return session.Run(cmd)
	})

	hostUsed(p.host)

	err = g.Wait()

	return stdout.String(), err
}

func (p sshlink) ExecPack(commands []string, wait, pause time.Duration) error {

	timeout := wait * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	g, _ := errgroup.WithContext(ctx)
	defer cancel()

	g.Go(func() error {
		for _, cmd := range commands {
			session, err := p.conn.NewSession()
			if err != nil {
				return fmt.Errorf("ssh session error: %w", err)
			}
			defer session.Close()

			log.Debugf("CMD: %s", cmd)
			err = session.Run(cmd)
			if err != nil {
				return err
			}
			time.Sleep(pause * time.Millisecond)
		}
		return nil
	})

	hostUsed(p.host)

	return g.Wait()
}
