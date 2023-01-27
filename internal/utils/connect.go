package utils

import (
	"meteo/internal/log"
	"net"
	"time"
)

func RawConnect(host string, ports []string) bool {
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			log.Errorf("Connect to %s:%s fail: %v", host, port, err)
			return false
		}
		if conn != nil {
			defer conn.Close()
		}
	}
	return true
}
