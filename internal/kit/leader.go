package kit

import "sync"

type Leader struct {
	server      string
	aliveRemote bool
	mu          sync.Mutex
	localIP     string
	remoteIP    string
}

func NewLeader(m bool, l_IP, r_IP, server string) *Leader {
	return &Leader{
		server:      server,
		aliveRemote: false,
		localIP:     l_IP,
		remoteIP:    r_IP,
	}
}

func (l *Leader) IsAliveRemote() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	ret := l.aliveRemote
	return ret
}

func (l *Leader) SetAliveRemote(alive bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.aliveRemote = alive
}

func (l *Leader) LocalIP() string {
	return l.localIP
}

func (l *Leader) RemoreIP() string {
	return l.remoteIP
}

func (l *Leader) Server() string {
	return l.server
}
