package vip

import (
	"context"
	"meteo/internal/log"
	"sync"
	"time"

	"github.com/balacode/udpt"
)

var KEY = []byte{
	0xF3, 0xD7, 0xDA, 0x69, 0xF0, 0x23, 0x22, 0x29,
	0xEC, 0x83, 0x61, 0xB9, 0xCD, 0x6E, 0x0A, 0x27,
	0x72, 0x66, 0x3B, 0xC4, 0x44, 0x3E, 0x85, 0x91,
	0x4A, 0x62, 0x99, 0x5F, 0xFE, 0x6D, 0x20, 0x54,
}

var remote_alive bool
var alive sync.Mutex

func ListenUdptPort(ctx context.Context, address string, port uint) {

	udpt.Config.Address = address
	udpt.Config.Port = int(port)
	udpt.Config.ReplyTimeout = 1 * time.Second
	udpt.Config.AESKey = KEY
	udpt.Config.VerboseSender = false
	udpt.Config.VerboseReceiver = false

	log.Debug("Running the receiver %s:%d", address, port)

	write := func(name string, data []byte) error {
		alive.Lock()
		log.Debug("Udpd write name: %s, data: %s", name, string(data))
		remote_alive = (name == "request" && string(data) == PASSPHRASE)
		alive.Unlock()
		return nil
	}

	read := func(name string) ([]byte, error) {
		//log.Infof("Receiver's read(), name: %s", name)
		return nil, nil
	}

	udpt.RunReceiver(write, read)

	<-ctx.Done()

	log.Warning("UDPD Shutdown")
}

func isAlive() bool {
	alive.Lock()
	defer alive.Unlock()
	return remote_alive
}
