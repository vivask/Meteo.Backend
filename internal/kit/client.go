package kit

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"meteo/internal/config"
	"meteo/internal/log"
	"net"
	"net/http"
	URL "net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type ServicePort int

//var remoteDead bool = false

type confService struct {
	port ServicePort
	ca   string
	crt  string
	key  string
}

type Service struct {
	port   ServicePort
	client *http.Client
}

var (
	WEB       ServicePort = 443
	CLUSTER   ServicePort = 10000
	PROXY     ServicePort = 11000
	SCHEDULE  ServicePort = 12000
	SSHCLIENT ServicePort = 13000
	MESSANGER ServicePort = 14000
	RADIUS    ServicePort = 15000
	NUT       ServicePort = 16000
	ESP32     ServicePort = 17000
	MEDIA     ServicePort = 18000
)

const (
	INTERNAL = false
	EXTERNAL = true
)

type Client struct {
	local   string
	remote  string
	clients map[string]*Service
}

func NewClient() (*Client, error) {
	readConfig()
	return &Client{
		local:  config.Default.Client.Local,
		remote: config.Default.Client.Remote,
		clients: map[string]*Service{
			"cluster": GetService(confService{
				port: CLUSTER,
				ca:   config.Default.Cluster.Api.Ca,
				crt:  config.Default.Cluster.Client.Crt,
				key:  config.Default.Cluster.Client.Key,
			}),
			"web": GetService(confService{
				port: WEB,
				ca:   config.Default.Web.Ca,
				crt:  config.Default.Cluster.Client.Crt,
				key:  config.Default.Cluster.Client.Key,
			}),
			"proxy": GetService(confService{
				port: PROXY,
				ca:   config.Default.Proxy.Api.Ca,
				crt:  config.Default.Proxy.Client.Crt,
				key:  config.Default.Proxy.Client.Key,
			}),
			"sshclient": GetService(confService{
				port: SSHCLIENT,
				ca:   config.Default.SshClient.Api.Ca,
				crt:  config.Default.SshClient.Client.Crt,
				key:  config.Default.SshClient.Client.Key,
			}),
			"messanger": GetService(confService{
				port: MESSANGER,
				ca:   config.Default.Messanger.Api.Ca,
				crt:  config.Default.Messanger.Client.Crt,
				key:  config.Default.Messanger.Client.Key,
			}),
			"schedule": GetService(confService{
				port: SCHEDULE,
				ca:   config.Default.Schedule.Api.Ca,
				crt:  config.Default.Schedule.Client.Crt,
				key:  config.Default.Schedule.Client.Key,
			}),
			"radius": GetService(confService{
				port: RADIUS,
				ca:   config.Default.Radius.Api.Ca,
				crt:  config.Default.Radius.Client.Crt,
				key:  config.Default.Radius.Client.Key,
			}),
			"esp32": GetService(confService{
				port: ESP32,
				ca:   config.Default.Esp32.Api.Ca,
				crt:  config.Default.Esp32.Client.Crt,
				key:  config.Default.Esp32.Client.Key,
			}),
			"media": GetService(confService{
				port: MEDIA,
				ca:   config.Default.Media.Api.Ca,
				crt:  config.Default.Media.Client.Crt,
				key:  config.Default.Media.Client.Key,
			}),
			"nut": GetService(confService{
				port: NUT,
				ca:   config.Default.Nut.Api.Ca,
				crt:  config.Default.Nut.Client.Crt,
				key:  config.Default.Nut.Client.Key,
			}),
		},
	}, nil
}

func readConfig() {
	WEB = ServicePort(config.Default.Web.Port)
	CLUSTER = ServicePort(config.Default.Cluster.Api.Port)
	PROXY = ServicePort(config.Default.Proxy.Api.Port)
	SCHEDULE = ServicePort(config.Default.Schedule.Api.Port)
	SSHCLIENT = ServicePort(config.Default.SshClient.Api.Port)
	MESSANGER = ServicePort(config.Default.Messanger.Api.Port)
	RADIUS = ServicePort(config.Default.Radius.Api.Port)
	ESP32 = ServicePort(config.Default.Esp32.Api.Port)
	MEDIA = ServicePort(config.Default.Media.Api.Port)
	NUT = ServicePort(config.Default.Nut.Api.Port)
}

func GetService(c confService) *Service {
	if len(c.ca) == 0 {
		return nil
	}

	caCert, err := os.ReadFile(c.ca)
	if err != nil {
		log.Fatalf("can't read cert %s for port %d: %v", c.ca, c.port, err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cert, err := tls.LoadX509KeyPair(c.crt, c.key)
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			MaxConnsPerHost:       100,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
			},
		},
	}
	return &Service{port: c.port, client: client}
}

func GetServiceName(url string) string {
	split := strings.Split(url, "/")
	if len(split) <= 1 {
		return url
	}
	return split[1]
}

func urlPrepare(port ServicePort, address string) string {
	var sb strings.Builder
	sb.WriteString("https://")
	sb.WriteString(address)
	sb.WriteString(":")
	sb.WriteString(strconv.Itoa(int(port)))
	sb.WriteString(config.Default.App.Api)
	return sb.String()
}

func (c *Client) internal(port ServicePort) string {
	return urlPrepare(port, c.local)
}

func (c *Client) external(port ServicePort) string {
	return urlPrepare(port, c.remote)
}

type params struct {
	service *Service
	url     string
	method  string
}

func (c *Client) prepare(path string, method string, ext bool) (*params, error) {

	var serviceName string = GetServiceName(path)
	var service *Service
	var ok bool

	if service, ok = c.clients[serviceName]; !ok {
		return nil, fmt.Errorf("unknown service [%s]", serviceName)
	}

	if service == nil {
		return nil, fmt.Errorf("not implemented service [%s]", serviceName)
	}

	var sb strings.Builder
	if ext {
		sb.WriteString(c.external(service.port))
	} else {
		sb.WriteString(c.internal(service.port))
	}
	sb.WriteString(path)

	return &params{
		method:  method,
		service: service,
		url:     sb.String(),
	}, nil
}

func getJsonRequest(url, method string, r interface{}) (*http.Request, error) {
	jsonStr, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("error JSON Marshal: %w", err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("unable to create http request due to error %w", err)
	}
	return req, err
}

func getMultipartRequest(url, method string, r interface{}) (req *http.Request, err error) {
	if body, ok := r.(*bytes.Buffer); ok {
		req, err = http.NewRequest(method, url, bytes.NewReader(body.Bytes()))
		if err != nil {
			return nil, fmt.Errorf("unable to create http request due to error %w", err)
		}
	} else {
		return nil, fmt.Errorf("bad input data: %w", err)
	}
	return
}

func Do(service *Service, req *http.Request) ([]byte, error) {
	resp, err := service.client.Do(req)
	if err != nil {
		switch e := err.(type) {
		case *URL.Error:
			return nil, fmt.Errorf("url.Error received on http request: %w", e)
		default:
			return nil, fmt.Errorf("unexpected error received: %w", err)
		}
	}

	body, err := FromJSON(resp)
	if err != nil {
		return nil, fmt.Errorf("serialization error: %w", err)
	}

	if string(body) == "404 page not found" {
		return nil, fmt.Errorf("404 page [%s] not found ", resp.Request.URL)
	}

	return body, nil
}

func (c *Client) post(path, content string, r interface{}, method string, ext bool) ([]byte, error) {

	p, err := c.prepare(path, method, ext)
	if err != nil {
		return nil, fmt.Errorf("prepare params error: %w", err)
	}

	var req *http.Request
	if content == "application/json" {
		req, err = getJsonRequest(p.url, p.method, r)
	} else {
		req, err = getMultipartRequest(p.url, p.method, r)
	}
	if err != nil {
		return nil, fmt.Errorf("prepare request error: %w", err)
	}
	req.Header.Set("Content-Type", content)

	body, err := Do(p.service, req)
	if err != nil {
		return nil, fmt.Errorf("can't doing a request: %w", err)
	}

	//log.Debugf("POST: %v", p.url)

	return body, nil
}

func (c *Client) get(path string, method string, ext bool) ([]byte, error) {

	p, err := c.prepare(path, method, ext)
	if err != nil {
		return nil, fmt.Errorf("prepare params error: %w", err)
	}

	req, err := http.NewRequest(p.method, p.url, bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, fmt.Errorf("unable to create http request due to error %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	body, err := Do(p.service, req)
	if err != nil {
		return nil, fmt.Errorf("can't doing a request: %w", err)
	}

	//log.Debugf("GET: %v", p.url)

	return body, nil
}

func (c *Client) PostInt(url string, r interface{}) (body []byte, err error) {
	return c.post(url, "application/json", r, http.MethodPost, INTERNAL)
}

func (c *Client) PostFormInt(url, content string, r interface{}) (body []byte, err error) {
	return c.post(url, content, r, http.MethodPost, INTERNAL)
}

func (c *Client) PostExt(url string, r interface{}) (body []byte, err error) {
	if config.Default.App.Mode == "single" {
		log.Debug("Server is in single mode")
		return nil, errors.New("single mode on, can't external send")
	}
	return c.post(url, "application/json", r, http.MethodPost, EXTERNAL)
	/*if IsAliveRemote() {
		remoteDead = false
		return c.post(url, "application/json", r, http.MethodPost, EXTERNAL)
	} else {
		if !remoteDead {
			remoteDead = true
			return nil, fmt.Errorf("remote server is dead")
		} else {
			return nil, nil
		}
	}*/
}

func (c *Client) PutInt(url string, r interface{}) (body []byte, err error) {
	return c.post(url, "application/json", r, http.MethodPut, INTERNAL)
}

func (c *Client) PutExt(url string, r interface{}) (body []byte, err error) {
	if config.Default.App.Mode == "single" {
		log.Debug("Server is in single mode")
		return nil, errors.New("single mode on, can't external send")
	}
	return c.post(url, "application/json", r, http.MethodPut, EXTERNAL)
	/*if IsAliveRemote() {
		remoteDead = false
		return c.post(url, "application/json", r, http.MethodPut, EXTERNAL)
	} else {
		if !remoteDead {
			remoteDead = true
			return nil, fmt.Errorf("remote server is dead")
		} else {
			return nil, nil
		}
	}*/
}

func (c *Client) PutMain(url string, r interface{}) (body []byte, err error) {
	if config.Default.App.Server == "main" {
		return c.PutInt(url, r)
	} else {
		return c.PutExt(url, r)
	}
}

func (c *Client) PutBackup(url string, r interface{}) (body []byte, err error) {
	if config.Default.App.Server == "backup" {
		return c.PutInt(url, r)
	} else {
		return c.PutExt(url, r)
	}
}

func (c *Client) GetInt(url string) (body []byte, err error) {
	return c.get(url, http.MethodGet, INTERNAL)
}

func (c *Client) GetExt(url string) (body []byte, err error) {
	if config.Default.App.Mode == "single" {
		log.Debug("Server is in single mode")
		return nil, errors.New("single mode on, can't external send")
	}
	return c.get(url, http.MethodGet, EXTERNAL)
	/*if IsAliveRemote() {
		remoteDead = false
		return c.get(url, http.MethodGet, EXTERNAL)
	} else {
		if !remoteDead {
			remoteDead = true
			return nil, fmt.Errorf("remote server is dead")
		} else {
			return nil, nil
		}
	}*/
}

func (c *Client) GetMain(url string) (body []byte, err error) {
	if config.Default.App.Server == "main" {
		return c.GetInt(url)
	} else {
		return c.GetExt(url)
	}
}

func (c *Client) GetBackup(url string) (body []byte, err error) {
	if config.Default.App.Server == "backup" {
		return c.GetInt(url)
	} else {
		return c.GetExt(url)
	}
}

func (c *Client) DeleteInt(url string) (body []byte, err error) {
	return c.get(url, http.MethodDelete, INTERNAL)
}

func (c *Client) DeleteExt(url string) (body []byte, err error) {
	if config.Default.App.Mode == "single" {
		log.Debug("Server is in single mode")
		return nil, errors.New("single mode on, can't external send")
	}
	return c.get(url, http.MethodDelete, EXTERNAL)
	/*if IsAliveRemote() {
		remoteDead = false
		return c.get(url, http.MethodDelete, EXTERNAL)
	} else {
		if !remoteDead {
			remoteDead = true
			return nil, fmt.Errorf("remote server is dead")
		} else {
			return nil, nil
		}
	}*/
}
