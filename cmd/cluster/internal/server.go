package vip

import (
	"context"
	"fmt"
	"meteo/internal/log"
	"net"
	"time"
)

const maxBufferSize = 1024
const timeout = 2 * time.Second

func ListenUdpPort(ctx context.Context, address string, port uint) error {
	// listen to incoming udp packets
	dst := fmt.Sprintf("%s:%d", address, port)
	pc, err := net.ListenPacket("udp", dst)
	if err != nil {
		return fmt.Errorf("can't listen %s, error: %w", dst, err)
	}
	defer pc.Close()

	go func() {
		for {
			buf := make([]byte, maxBufferSize)
			n, addr, err := pc.ReadFrom(buf)
			if err != nil {
				continue
			}
			deadline := time.Now().Add(timeout)
			err = pc.SetWriteDeadline(deadline)
			if err != nil {
				continue
			}
			go echo(pc, addr, buf[:n])
		}
	}()

	<-ctx.Done()
	log.Warningf("Udpd shutdown")
	return ctx.Err()
}

func echo(pc net.PacketConn, addr net.Addr, buf []byte) {
	pc.WriteTo(buf, addr)
}
