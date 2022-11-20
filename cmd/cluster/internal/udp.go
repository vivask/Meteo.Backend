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
		for {
			buf := make([]byte, maxBufferSize)
			n, addr, err := conn.ReadFrom(buf)
			if err != nil {
				continue
			}
			//log.Debugf("udp server packet-received: bytes=%d from=%s", n, addr.String())

			go echo(conn, addr, buf[:n])
		}
	}()

	<-ctx.Done()
	log.Debugf("UDP server done: %v", ctx.Err())

	return nil
}

func echo(conn net.PacketConn, addr net.Addr, buf []byte) {
	_, err := conn.WriteTo(buf, addr)
	if err != nil && aliveRemote {
		log.Debugf("UDP write error: %v", err)
	}
}
