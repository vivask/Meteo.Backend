package vip

import (
	"context"
	"fmt"
	"meteo/internal/log"
	"net"
)

const maxBufferSize = 512

func ListenUdpPort(ctx context.Context, address string, port uint) error {
	// listen to incoming udp packets
	dst := fmt.Sprintf("%s:%d", address, port)
	conn, err := net.ListenPacket("udp", dst)
	if err != nil {
		return fmt.Errorf("can't listen %s, error: %w", dst, err)
	}
	defer conn.Close()

	go func() {
		buf := make([]byte, maxBufferSize)
		for {
			n, addr, err := conn.ReadFrom(buf)
			if err != nil {
				continue
			}
			go echo(conn, addr, buf[:n])
		}
	}()

	<-ctx.Done()
	log.Debugf("UDP server done: %v", ctx.Err())

	return nil
}

func echo(conn net.PacketConn, addr net.Addr, buf []byte) {
	conn.WriteTo(buf, addr)
}
