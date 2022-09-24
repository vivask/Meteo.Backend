package service

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"meteo/cmd/proxy/internal/specialized"
	"meteo/internal/config"
	"meteo/internal/log"

	"github.com/miekg/dns"
	"golang.org/x/sync/errgroup"
)

type ServerState int

const (
	STATE_VPN ServerState = iota
	STATE_DIRECT
	STATE_PROV
)

const (
	defaultCacheSize       = 65536
	connectionTimeout      = 10 * time.Second
	connectionsPerUpstream = 2
	refreshQueueSize       = 2048
	timerResolution        = 1 * time.Second
	verifyPerUpstream      = 3
	timeReachable          = 300 * time.Millisecond
	timeVerifyRetry        = 5 * time.Second
)

// Server is a caching DNS proxy that upgrades DNS to DNS over TLS.
type Server struct {
	cache           *cache
	cacheOn         bool
	pools           []*pool
	rq              chan *dns.Msg
	dial            func(addr string, cfg *tls.Config) (net.Conn, error)
	mu              sync.RWMutex
	currentTime     time.Time
	startTime       time.Time
	vpnServers      []string
	directServers   []string
	providerServers []string
	servers         []*dns.Server
	state           ServerState
	zo              *Zones
	bl              *BlackList
	blkListOn       bool
	TTL             uint32
	un              *Unlocker
	unlockerOn      bool
	mutex           sync.Mutex
	queu            map[string]struct{}
}

// NewServer constructs a new server but does not start it, use Run to start it afterwards.
// Calling New(0) is valid and comes with working defaults:
// * If cacheSize is 0 a default value will be used. to disable caches use a negative value.
// * If no upstream servers are specified default ones will be used.
func NewServer() *Server {
	cacheSize := config.Default.Proxy.CacheSize
	switch {
	case cacheSize == 0:
		cacheSize = defaultCacheSize
	case cacheSize < 0:
		cacheSize = 0
	}
	cache, err := newCache(cacheSize, config.Default.Proxy.EvictMetrics)
	if err != nil {
		log.Fatal("Unable to initialize the cache")
	}

	ttl, err := time.ParseDuration(config.Default.Proxy.UpdateInterval)
	if err != nil {
		log.Fatal(err)
	}

	s := &Server{
		cache:   cache,
		cacheOn: config.Default.Proxy.Cached,
		rq:      make(chan *dns.Msg, refreshQueueSize),
		dial: func(addr string, cfg *tls.Config) (net.Conn, error) {
			return tls.Dial("tcp", addr, cfg)
		},
		vpnServers:      config.Default.Proxy.NsVpn,
		directServers:   config.Default.Proxy.NsDirect,
		providerServers: config.Default.Proxy.NsProvider,
		state:           STATE_VPN,
		TTL:             uint32(ttl),
		blkListOn:       config.Default.Proxy.AdBlock,
		unlockerOn:      config.Default.Proxy.Unlocker,
		queu:            map[string]struct{}{},
	}
	return s
}

func (s *Server) CacheClear() {
	cacheSize := config.Default.Proxy.CacheSize
	switch {
	case cacheSize == 0:
		cacheSize = defaultCacheSize
	case cacheSize < 0:
		cacheSize = 0
	}
	cache, err := newCache(cacheSize, config.Default.Proxy.EvictMetrics)
	if err != nil {
		log.Fatal("Unable to initialize the cache")
	}
	s.cache = cache
}

func (s *Server) createConnectors() {
	f := func(servers []string, tlsOn bool) {
		s.pools = []*pool{}
		if tlsOn {
			for _, addr := range servers {
				s.pools = append(s.pools, newPool(connectionsPerUpstream, s.tlsconnector(addr)))
			}
		} else {
			for _, addr := range servers {
				s.pools = append(s.pools, newPool(connectionsPerUpstream, s.connector(addr)))
			}
		}
	}
	switch s.state {
	case STATE_VPN:
		f(s.vpnServers, true)
	case STATE_DIRECT:
		f(s.directServers, true)
	case STATE_PROV:
		f(s.providerServers, false)
	default:
		log.Fatalf("Unknown server state: %v", s.state)
	}
}

func (s *Server) connector(upstreamServer string) func() (*dns.Conn, error) {
	return func() (*dns.Conn, error) {
		c := new(dns.Client)
		conn, err := c.Dial(upstreamServer)
		if err != nil {
			return nil, fmt.Errorf("failed connect to server: %s, error: %w", upstreamServer, err)
		}
		return &dns.Conn{Conn: conn}, nil
	}
}

func (s *Server) tlsconnector(upstreamServer string) func() (*dns.Conn, error) {
	return func() (*dns.Conn, error) {
		tlsConf := &tls.Config{
			// Force TLS 1.2 as minimum version.
			MinVersion: tls.VersionTLS12,
		}
		dialableAddress := upstreamServer
		serverComponents := strings.Split(upstreamServer, "@")
		if len(serverComponents) == 2 {
			servername, port, err := net.SplitHostPort(serverComponents[0])
			if err != nil {
				log.Warningf("Failed to parse DNS-over-TLS upstream address: %v", err)
				return nil, err
			}
			tlsConf.ServerName = servername
			dialableAddress = serverComponents[1] + ":" + port
		}
		conn, err := s.dial(dialableAddress, tlsConf)
		if err != nil {
			log.Warningf("Failed to connect to DNS-over-TLS upstream: %v", err)
			return nil, err
		}
		return &dns.Conn{Conn: conn}, nil
	}
}

// Run runs the server. The server will gracefully shutdown when context is canceled.
func (s *Server) Run(ctx context.Context) error {

	s.createConnectors()

	mux := dns.NewServeMux()
	mux.Handle(".", s)

	tcpAddr := fmt.Sprintf("%s:%d", config.Default.Proxy.Listen, config.Default.Proxy.TcpPort)
	udpAddr := fmt.Sprintf("%s:%d", config.Default.Proxy.Listen, config.Default.Proxy.UdpPort)

	//log.Infof("tcpAddr: %s", tcpAddr)
	//log.Infof("udpAddr: %s", udpAddr)

	tcp := &dns.Server{
		Addr:         tcpAddr,
		Net:          "tcp",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	udp := &dns.Server{
		Addr:         udpAddr,
		Net:          "udp",
		Handler:      mux,
		UDPSize:      65535,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	s.servers = []*dns.Server{tcp, udp}

	g, ctx := errgroup.WithContext(ctx)

	go func() {
		<-ctx.Done()
		for _, s := range s.servers {
			_ = s.Shutdown()
		}
		for _, p := range s.pools {
			p.shutdown()
		}
	}()

	go s.refresher(ctx)
	go s.timer(ctx)

	for _, s := range s.servers {
		s := s
		g.Go(func() error { return s.ListenAndServe() })
	}

	s.startTime = time.Now()
	log.Infof("DNS over TLS forwarder listening on %v", udpAddr)
	return g.Wait()
}

// ServeDNS implements miekg/dns.Handler for Server.
func (s *Server) ServeDNS(w dns.ResponseWriter, q *dns.Msg) {
	inboundIP, _, _ := net.SplitHostPort(w.RemoteAddr().String())
	log.Debugf("Question from %s: %q", inboundIP, q.Question[0])
	m := s.getAnswer(q)
	if m == nil {
		dns.HandleFailed(w, q)
		return
	}
	if err := w.WriteMsg(m); err != nil {
		log.Warningf("Write message failed, message: %v, error: %v", m, err)
	}
}

type debugStats struct {
	CacheMetrics       specialized.CacheMetrics
	CacheLen, CacheCap int
	Uptime             string
}

// DebugHandler returns an http.Handler that serves debug stats.
func (s *Server) DebugHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		buf, err := json.MarshalIndent(debugStats{
			s.cache.c.Metrics(),
			s.cache.c.Len(),
			s.cache.c.Cap(),
			time.Since(s.startTime).String(),
		}, "", " ")
		if err != nil {
			http.Error(w, "Unable to retrieve debug info", http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(buf)
	})
}

func (s *Server) extractNS(server string) (dialableAddress string, err error) {
	serverComponents := strings.Split(server, "@")
	if len(serverComponents) == 2 {
		_, port, err := net.SplitHostPort(serverComponents[0])
		if err != nil {
			log.Warningf("Failed to parse DNS-over-TLS upstream address: %v", err)
			return "", err
		}
		dialableAddress = serverComponents[1] + ":" + port
	}
	return
}

func (s *Server) isAvailableServers(servers []string) bool {
	for i := 0; i < verifyPerUpstream; i++ {
		for _, server := range servers {
			ns, err := s.extractNS(server)
			if err == nil {
				_, err := net.DialTimeout("tcp", ns, timeReachable)
				if err == nil {
					return true
				}
			}
			log.Debugf("Retrying %q [%d/%d]...", ns, i+1, verifyPerUpstream)
		}
	}
	return false
}

func (s *Server) setServerState(ss ServerState) {
	s.state = ss
	s.createConnectors()
}

func (s *Server) VerifyState() {
	switch s.state {
	case STATE_VPN:
		if s.isAvailableServers(s.vpnServers) {
			return
		}
		s.setServerState(STATE_DIRECT)
		log.Info("Proxy down to DIRECT")
	case STATE_DIRECT:
		if s.isAvailableServers(s.vpnServers) {
			s.setServerState(STATE_VPN)
			log.Info("Proxy up to VPN")
			return
		} else if s.isAvailableServers(s.directServers) {
			return
		}
		s.setServerState(STATE_PROV)
		log.Info("Proxy down to RESERVE")
	case STATE_PROV:
		if s.isAvailableServers(s.vpnServers) {
			s.setServerState(STATE_VPN)
			log.Info("Proxy up to VPN")
		} else if s.isAvailableServers(s.directServers) {
			s.setServerState(STATE_DIRECT)
			log.Info("Proxy up to DIRECT")
		}
	default:
		log.Errorf("Unknown server state: %v", s.state)
	}
}

func (s *Server) getAnswer(q *dns.Msg) *dns.Msg {

	question := q.Question[0]

	lover := strings.ToLower(question.Name)
	if (question.Qtype == dns.TypeA || question.Qtype == dns.TypeAAAA) && s.zo.Contains(lover) {
		m := &dns.Msg{}
		m.SetReply(q)

		head := dns.RR_Header{
			Name:   question.Name,
			Rrtype: question.Qtype,
			Class:  dns.ClassINET,
			Ttl:    s.TTL,
		}

		line := &dns.A{
			Hdr: head,
			A:   net.ParseIP(s.zo.Address(lover)),
		}

		m.Answer = append(m.Answer, line)

		log.Debugf("LOCAL QN: %s", question.Name)

		return m
	}

	if (question.Qtype == dns.TypeA || question.Qtype == dns.TypeAAAA) && s.blkListOn && s.bl.Contains(question.Name) {
		m := &dns.Msg{}
		m.SetReply(q)

		head := dns.RR_Header{
			Name:   question.Name,
			Rrtype: question.Qtype,
			Class:  dns.ClassINET,
			Ttl:    s.TTL,
		}

		var line dns.RR
		if question.Qtype == dns.TypeA {
			line = &dns.A{
				Hdr: head,
				A:   net.ParseIP(config.Default.Proxy.BlockIPv4),
			}
		} else {
			line = &dns.AAAA{
				Hdr:  head,
				AAAA: net.ParseIP(config.Default.Proxy.BlockIPv6),
			}
		}
		m.Answer = append(m.Answer, line)

		log.Debugf("BLOCKED QN: %s", question.Name)

		return m
	}

	if s.cacheOn {
		m, ok := s.cache.get(q)
		// Cache HIT.
		if ok {
			return m
		}
		// If there is a cache HIT with an expired TTL, speculatively return the cache entry anyway with a short TTL, and refresh it.
		if !ok && m != nil {
			s.refresh(q)
			return m
		}
	}

	// If there is a cache MISS, forward the message upstream and return the answer.
	// miek/dns does not pass a context so we fallback to Background.
	return s.forwardMessageAndCacheResponse(q)
}

func (s *Server) refresh(q *dns.Msg) {
	select {
	case s.rq <- q:
	default:
	}
}

func (s *Server) refresher(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case q := <-s.rq:
			s.forwardMessageAndCacheResponse(q)
		}
	}
}

func (s *Server) timer(ctx context.Context) {
	t := time.NewTicker(timerResolution)
	for {
		select {
		case <-ctx.Done():
			t.Stop()
			return
		case t := <-t.C:
			s.mu.Lock()
			s.currentTime = t
			s.mu.Unlock()
		}
	}
}

func (s *Server) now() time.Time {
	s.mu.RLock()
	t := s.currentTime
	s.mu.RUnlock()
	return t
}

func (s *Server) Lookup(host string) *dns.Msg {
	if !strings.HasSuffix(host, ".") {
		host += "."
	}
	q := &dns.Msg{}
	q.SetQuestion(host, dns.TypeA)
	q.RecursionDesired = true

	return s.forwardMessageAndCacheResponse(q)
}

func (s *Server) forwardMessageAndCacheResponse(q *dns.Msg) (m *dns.Msg) {
	m = s.forwardMessageAndGetResponse(q)
	// Let's retry a few times if we can't resolve it at the first try.
	for c := 0; m == nil && c < connectionsPerUpstream; c++ {
		log.Debugf("Retrying %q [%d/%d]...", q.Question, c+1, connectionsPerUpstream)
		m = s.forwardMessageAndGetResponse(q)
	}
	if m == nil {
		log.Infof("Giving up on %q after %d connection retries.", q.Question, connectionsPerUpstream)
		return nil
	}

	if s.unlockerOn {
		name := q.Question[0].Name
		if !(strings.HasSuffix(name, ".tv.") || strings.HasSuffix(name, ".tv") || len(name) > 25) {
			go s.UnlockIfLocked(name, m, s.un)
		}
	}

	if s.cacheOn {
		s.cache.put(q, m)
	}

	return m
}

func (s *Server) forwardMessageAndGetResponse(q *dns.Msg) (m *dns.Msg) {
	resps := make(chan *dns.Msg, len(s.pools))
	for _, p := range s.pools {
		go func(p *pool) {
			r, err := s.exchangeMessages(p, q)
			if err != nil || r == nil {
				resps <- nil
			}
			resps <- r
		}(p)
	}
	for c := 0; c < len(s.pools); c++ {
		// Return the response only if it has Rcode NoError or NXDomain, otherwise try another pool.
		if r := <-resps; r != nil && (r.Rcode == dns.RcodeSuccess || r.Rcode == dns.RcodeNameError) {
			return r
		}
	}
	return nil
}

var errNilResponse = errors.New("nil response from upstream")

func (s *Server) exchangeMessages(p *pool, q *dns.Msg) (resp *dns.Msg, err error) {
	c, err := p.get()
	if err != nil {
		return nil, err
	}
	_ = c.SetDeadline(s.now().Add(connectionTimeout))
	defer func() {
		if err == nil {
			p.put(c)
		}
	}()
	if err := c.WriteMsg(q); err != nil {
		log.Debugf("Send question message failed: %v", err)
		c.Close()
		return nil, err
	}
	resp, err = c.ReadMsg()
	if err != nil {
		log.Debugf("Error while reading message: %v", err)
		c.Close()
		return nil, err
	}
	if resp == nil {
		log.Debug(errNilResponse)
		c.Close()
		return nil, errNilResponse
	}
	return resp, err
}

func (s *Server) zones() *Zones {
	return s.zo
}

func (s *Server) SetZones(z *Zones) {
	s.zo = z
}

func (s *Server) blacklist() *BlackList {
	return s.bl
}

func (s *Server) SetBlackList(b *BlackList) {
	s.bl = b
}

func (s *Server) unlocker() *Unlocker {
	return s.un
}

func (s *Server) SetUnlocker(u *Unlocker) {
	s.un = u
}

func (s *Server) isBlocked(host string) bool {
	ports := []string{"443", "80"}
	for _, port := range ports {
		var url string
		if port == "443" {
			url = "https://" + host
		} else {
			url = "http://" + host
		}
		timeout := time.Duration(1 * time.Second)
		c := http.Client{
			Timeout: timeout,
		}
		_, err := c.Get(url)
		if err != nil && strings.Contains(err.Error(), "connection reset by peer") {
			return true
		}
	}
	return false
}

func (s *Server) UnlockIfLocked(host string, r *dns.Msg, un *Unlocker) {
	if s.state == STATE_VPN && len(r.Answer) > 0 {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		if _, ok := s.queu[host]; ok {
			return
		}
		if !un.Exist(host) && !un.Ignore(host) {
			s.queu[host] = struct{}{}
			if s.isBlocked(host) {
				if strings.HasSuffix(host, ".") {
					host = host[:len(host)-1]
				}
				err := s.un.InsertHost(host)
				if err != nil {
					log.Errorf("Failed to add host [%s] to vpn, error: %v", host, err)
				}
				log.Infof("Found and unlocked blocked host [%v] ", host)
			}
			delete(s.queu, host)
		}
	}
}

func (s *Server) GetTimerVerifyRetry() time.Duration {
	return timeVerifyRetry
}
