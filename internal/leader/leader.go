package leader

import "sync"

type Leader struct {
	leader      bool
	master      bool
	aliveRemote bool
	mu          sync.Mutex
	localIP     string
	remoteIP    string
}

func New(m bool, l_IP, r_IP string) *Leader {
	return &Leader{
		master:      m,
		leader:      false,
		aliveRemote: false,
		localIP:     l_IP,
		remoteIP:    r_IP,
	}
}

func (l *Leader) IsLeader() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	ret := l.leader
	return ret
}

func (l *Leader) SetLeader(lead bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.leader = lead
}

func (l *Leader) IsMaster() bool {
	return l.master
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

func (l *Leader) Self() string {
	if l.master {
		return "MASTER"
	}
	return "SLAVE"
}

func (l *Leader) Other() string {
	if l.master {
		return "SLAVE"
	}
	return "MASTER"
}

func (l *Leader) LocalIP() string {
	return l.localIP
}

func (l *Leader) RemoreIP() string {
	return l.remoteIP
}
