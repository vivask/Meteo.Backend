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
		Mode       string
		Server     string
		HealthPort uint
	}
	Cluster struct {
		Title     string
		Bind      string
		Port      uint
		Ca        string
		Crt       string
		Key       string
		DbLink    string
		Interface string
		VirtualIP string
		LogLevel  string
	}
	Client struct {
		Ssl    bool
		Local  string
		Remote string
		Crt    string
		Key    string
	}
	Auth struct {
		AccessTokenPrivateKeyPath  string
		AccessTokenPublicKeyPath   string
		RefreshTokenPrivateKeyPath string
		RefreshTokenPublicKeyPath  string
		PassResetCodeExpiration    int
		JwtExpiration              uint
		JwtRefreshExpiration       uint
		JwtKey                     string
	}
	Web struct {
		Title    string
		Active   bool
		Ssl      bool
		Listen   string
		Port     uint
		Ui       string
		Ca       string
		Crt      string
		Key      string
		DbLink   string
		TlsMin   string
		LogLevel string
	}
	Database struct {
		Name     string
		User     string
		Password string
		Port     uint
		Pool     struct {
			Max uint
		}
		Sync    bool
		Exclude []string
	}
	Proxy struct {
		Title   string
		Active  bool
		Listen  string
		UdpPort uint
		TcpPort uint
		DbLink  string
		Rest    struct {
			Bind string
			Port uint
			Ca   string
			Crt  string
			Key  string
		}
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
		AdSources      []string
		LogLevel       string
	}
	Server struct {
		Title  string
		Active bool
		Bind   string
		Port   uint
		Ca     string
		Crt    string
		Key    string
		DbLink string
		Radius struct {
			HealthUser   string
			HealthPasswd string
			HealthKey    string
			HealthPort   uint
			DebugMode    bool
		}
		Storage struct {
			Device     string
			MountPoint string
		}
		Transmission struct {
			User     string
			Password string
		}
		LogLevel string
	}
	Messanger struct {
		Title    string
		Active   bool
		Bind     string
		Port     uint
		Ca       string
		Crt      string
		Key      string
		Telegram struct {
			Active bool
			ChatId int64
			Key    string
			Url    string
		}
		LogLevel string
	}
	SshClient struct {
		Title  string
		Active bool
		Bind   string
		Port   uint
		Ca     string
		Crt    string
		Key    string
		DbLink string
		Git    struct {
			Host       string
			Port       uint
			User       string
			Repository string
			Remote     string
			Commit     string
		}
		PPP struct {
			Host      string
			Port      uint
			User      string
			Interface string
			Script    string
		}
		Mikrotik struct {
			Hosts      []string
			Ports      []uint
			Users      []string
			Repository string
		}
		Vpn struct {
			Host string
			Port uint
			User string
			List string
		}
		Main struct {
			Host string
			Port uint
			User string
		}
		Backup struct {
			Host string
			Port uint
			User string
		}
		LogLevel string
	}
	Schedule struct {
		Title    string
		Active   bool
		Bind     string
		Port     uint
		Ca       string
		Crt      string
		Key      string
		DbLink   string
		LogLevel string
	}
	Esp32 struct {
		Title    string
		Active   bool
		Bind     string
		Port     uint
		Ca       string
		Crt      string
		Key      string
		DbLink   string
		Mac      string
		Check    bool
		LogLevel string
	}
}

func (d Config) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

// Parse get all config support in app
func Parse() Config {
	if err := viperInstance.Unmarshal(&Default); err != nil {
		log.Fatalf("Fail to read configuration %v", err)
	}
	return Default
}

// Viper instance
func Viper() *viper.Viper {
	return viperInstance
}
