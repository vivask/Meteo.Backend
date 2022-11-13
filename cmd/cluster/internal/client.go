package vip

import (
	"fmt"
	"io"
	"meteo/internal/config"
	"net"
	"strings"
	"time"
)

func UdpSend(msg string) (string, error) {
	dst := fmt.Sprintf("%s:%d", config.Default.Client.Remote, config.Default.Cluster.Port)
	raddr, err := net.ResolveUDPAddr("udp", dst)
	if err != nil {
		return "", fmt.Errorf("can't resolve address %s, error: %w", dst, err)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return "", fmt.Errorf("can't udp send to %s, error: %w", dst, err)
	}
	defer conn.Close()
	doneChan := make(chan error, 1)
	reader := strings.NewReader(msg)
	r := make(chan string)
	go func() {
		_, err := io.Copy(conn, reader)
		if err != nil {
			doneChan <- err
			return
		}

		//log.Debugf("udp client packet-written: bytes=%d", n)

		buffer := make([]byte, maxBufferSize)

		deadline := time.Now().Add(timeout)
		err = conn.SetReadDeadline(deadline)
		if err != nil {
			doneChan <- err
			return
		}

		nRead, _, err := conn.ReadFrom(buffer)
		if err != nil {
			doneChan <- err
			return
		}

		//log.Debugf("udp client packet-received: bytes=%d from=%s", nRead, addr.String())

		r <- string(buffer[:nRead])
	}()

	timer := time.After(timeout)
	var response string
	select {
	case <-timer:
		err = fmt.Errorf("udp dial time out")
	case err = <-doneChan:
	case response = <-r:
	}

	return response, err
}
