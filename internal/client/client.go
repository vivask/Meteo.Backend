package client

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"meteo/internal/config"
	"meteo/internal/leader"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net"
	"net/http"
	URL "net/url"
	"time"
)

type ServicePort int

var (
	HTTP      ServicePort = 80
	HTTPS     ServicePort = 443
	CLUSTER   ServicePort = ServicePort(config.Default.Cluster.Port)
	PROXY     ServicePort = ServicePort(config.Default.Proxy.RestPort)
	SCHEDULE  ServicePort = 12000
	SSHCLIENT ServicePort = 13000
	TELEGRAM  ServicePort = 14000
	XU4       ServicePort = 15000
)

const (
	INTERNAL = false
	EXTERNAL = true
)

const (
	POST   = true
	PUT    = false
	GET    = true
	DELETE = false
)

type Client struct {
	lead   *leader.Leader
	local  string
	remote string
	cli    *http.Client
}

func New(lead *leader.Leader) (*Client, error) {
	var client *http.Client
	if !config.Default.Client.Ssl {
		client = http.DefaultClient
	} else {
		caCert, err := ioutil.ReadFile(config.Default.Client.Ca)
		if err != nil {
			return nil, fmt.Errorf("error read CA: %w", err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		cert, err := tls.LoadX509KeyPair(config.Default.Client.Crt, config.Default.Client.Key)
		if err != nil {
			log.Fatalf("server: loadkeys: %s", err)
		}

		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}).DialContext,
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
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
	}

	return &Client{
		lead:   lead,
		local:  config.Default.Client.Local,
		remote: config.Default.Client.Remote,
		cli:    client,
	}, nil
}

func (c *Client) internal(port ServicePort) string {
	return fmt.Sprintf("https://%s:%d%s", c.local, port, utils.GetCurrentApi())

}

func (c *Client) external(port ServicePort) string {
	return fmt.Sprintf("https://%s:%d%s", c.remote, port, utils.GetCurrentApi())

}

func (c *Client) post(path string, r interface{}, post, ext bool, port ServicePort) ([]byte, error) {
	jsonStr, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("error JSON Marshal: %w", err)
	}
	var url string
	if ext {
		url = fmt.Sprintf("%s%s", c.external(port), path)
	} else {
		url = fmt.Sprintf("%s%s", c.internal(port), path)
	}

	method := http.MethodPost
	if !post {
		method = http.MethodPut
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("unable to create http request due to error %w", err)
	}

	resp, err := c.cli.Do(req)
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
		log.Errorf("Client response error: %v", err)
		return nil, err
	}

	return body, nil
}

func (c *Client) get(path string, get, ext bool, port ServicePort) ([]byte, error) {
	var url string
	if ext {
		url = fmt.Sprintf("%s%s", c.external(port), path)
	} else {
		url = fmt.Sprintf("%s%s", c.internal(port), path)
	}

	method := http.MethodGet
	if !get {
		method = http.MethodDelete
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, fmt.Errorf("unable to create http request due to error %w", err)
	}

	resp, err := c.cli.Do(req)
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
		log.Errorf("Client response error: %v", err)
		return nil, err
	}

	return body, nil
}

func (c *Client) PostInt(url string, r interface{}, port ServicePort) (body []byte, err error) {
	return c.post(url, r, POST, INTERNAL, port)
}

func (c *Client) PostExt(url string, r interface{}, port ServicePort) (body []byte, err error) {
	if c.lead.IsAliveRemote() {
		return c.post(url, r, POST, EXTERNAL, port)
	} else {
		return nil, fmt.Errorf("remote server is dead")
	}
}

func (c *Client) PostMaster(url string, r interface{}, port ServicePort) (body []byte, err error) {
	if c.lead.IsMaster() {
		return c.PostInt(url, r, port)
	} else {
		return c.PostExt(url, r, port)
	}
}

func (c *Client) PostSlave(url string, r interface{}, port ServicePort) (body []byte, err error) {
	if c.lead.IsMaster() {
		return c.PostExt(url, r, port)
	} else {
		return c.PostInt(url, r, port)
	}
}

func (c *Client) PutInt(url string, r interface{}, port ServicePort) (body []byte, err error) {
	return c.post(url, r, PUT, INTERNAL, port)
}

func (c *Client) PutExt(url string, r interface{}, port ServicePort) (body []byte, err error) {
	if c.lead.IsAliveRemote() {
		return c.post(url, r, PUT, EXTERNAL, port)
	} else {
		return nil, fmt.Errorf("remote server is dead")
	}
}

func (c *Client) PutMaster(url string, r interface{}, port ServicePort) (body []byte, err error) {
	if c.lead.IsMaster() {
		return c.PutInt(url, r, port)
	} else {
		return c.PutExt(url, r, port)
	}
}

func (c *Client) PutSlave(url string, r interface{}, port ServicePort) (body []byte, err error) {
	if c.lead.IsMaster() {
		return c.PutExt(url, r, port)
	} else {
		return c.PutInt(url, r, port)
	}
}

func (c *Client) GetInt(url string, port ServicePort) (body []byte, err error) {
	return c.get(url, GET, INTERNAL, port)
}

func (c *Client) GetExt(url string, port ServicePort) (body []byte, err error) {
	if c.lead.IsAliveRemote() {
		return c.get(url, GET, EXTERNAL, port)
	} else {
		return nil, fmt.Errorf("remote server is dead")
	}
}

func (c *Client) GetMaster(url string, port ServicePort) (body []byte, err error) {
	if c.lead.IsMaster() {
		return c.GetInt(url, port)
	} else {
		return c.GetExt(url, port)
	}
}

func (c *Client) GetSlave(url string, port ServicePort) (body []byte, err error) {
	if c.lead.IsMaster() {
		return c.GetExt(url, port)
	} else {
		return c.GetInt(url, port)
	}
}

func (c *Client) DeleteInt(url string, port ServicePort) (body []byte, err error) {
	return c.get(url, DELETE, INTERNAL, port)
}

func (c *Client) DeleteExt(url string, port ServicePort) (body []byte, err error) {
	if c.lead.IsAliveRemote() {
		return c.get(url, DELETE, EXTERNAL, port)
	} else {
		return nil, fmt.Errorf("remote server is dead")
	}
}

func (c *Client) DeleteMaster(url string, port ServicePort) (body []byte, err error) {
	if c.lead.IsMaster() {
		return c.DeleteInt(url, port)
	} else {
		return c.DeleteExt(url, port)
	}
}

func (c *Client) DeleteSlave(url string, port ServicePort) (body []byte, err error) {
	if c.lead.IsMaster() {
		return c.DeleteExt(url, port)
	} else {
		return c.DeleteInt(url, port)
	}
}
