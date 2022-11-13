package vip

import (
	"meteo/internal/config"
	"time"

	"github.com/balacode/udpt"
)

func UdptSend(pack string) error {
	udpt.Config.Address = config.Default.Client.Remote
	udpt.Config.Port = int(config.Default.Cluster.Port)
	udpt.Config.ReplyTimeout = 1 * time.Second
	udpt.Config.AESKey = KEY
	udpt.Config.VerboseSender = false
	udpt.Config.VerboseReceiver = false

	_ = udpt.Send("request", []byte(pack))
	return nil
}
