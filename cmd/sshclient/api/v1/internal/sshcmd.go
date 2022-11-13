package Ssh

import (
	"bytes"
	"context"
	"meteo/internal/log"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/sync/errgroup"
)

func (p sshlink) Config() *ssh.ClientConfig {
	return p.cnf
}

func (p sshlink) Close() {
	p.connection.Close()
	p.session.Close()
}

func (p sshlink) Exec(cmd string, wait time.Duration) (req string, err error) {
	var stdout bytes.Buffer
	p.session.Stdout = &stdout

	timeout := wait * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	g, _ := errgroup.WithContext(ctx)
	defer cancel()

	g.Go(func() error {
		log.Debugf("CMD: %s", cmd)
		return p.session.Run(cmd)
	})

	hostUsed(p.host)

	return stdout.String(), g.Wait()
}

func (p sshlink) ExecPack(commands []string, wait, pause time.Duration) error {

	timeout := wait * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	g, _ := errgroup.WithContext(ctx)
	defer cancel()

	g.Go(func() error {
		for _, cmd := range commands {
			log.Debugf("CMD: %s", cmd)
			err := p.session.Run(cmd)
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
