package config

import (
	"encoding/json"
	"meteo/internal/log"

	"github.com/spf13/viper"
)

var viperInstance = viper.New()
var Default Config

func init() {
	log.Debug("INIT CONFIG")
}

// Config struct
type Config struct {
	App struct {
		Version string `yaml:"version"`
		Master  bool   `yaml:"master"`
		Dir     string `yaml:"dir"`
	}
	Cluster struct {
		Bind      string
		Port      uint
		Interface string
		VirtualIP string
	}
	Client struct {
		Ssl    bool
		Local  string
		Remote string
		Ca     string
		Crt    string
		Key    string
	}
	Web struct {
		Title  string
		Active bool
		Ssl    bool
		Listen string
		Port   uint
		Ui     string
		Ca     string
		Crt    string
		Key    string
		TlsMin string
	}
	Database struct {
		URL  string
		Pool struct {
			Max uint
		}
		Sync    bool
		Exclude []string
	}
	Proxy struct {
		Title          string
		Active         bool
		Listen         string
		UdpPort        uint
		TcpPort        uint
		RestBind       string
		RestPort       uint
		EvictMetrics   bool
		NsVpn          []string
		NsDirect       []string
		NsProvider     []string
		Resolvers      []string
		Cached         bool
		CacheSize      int
		Unlocker       bool
		UpdateInterval string
		BlockIPv4      string
		BlockIPv6      string
		AdBlock        bool
		AdResource     []string
	}
}

func (d Config) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

// Parse get all config support in app
func Parse() Config {
	if err := viperInstance.Unmarshal(&Default); err != nil {
		log.Fatal("Fail to read configuration", err)
	}
	return Default
}

// Viper instance
func Viper() *viper.Viper {
	return viperInstance
}
