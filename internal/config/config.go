package config

import (
	"encoding/json"
	"meteo/internal/log"

	"github.com/spf13/viper"
)

var viperInstance = viper.New()
var Default Config

func init() {
	//log.Debug("INIT CONFIG")
}

// Config struct
type Config struct {
	// Application config
	App struct {
		Mode       string
		Server     string
		HealthPort uint
		Api        string
	}
	// Database config
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
	// Cluster config
	Cluster struct {
		Title string
		Api   struct {
			Bind string
			Port uint
			Ca   string
			Crt  string
			Key  string
		}
		Client struct {
			Crt string
			Key string
		}
		DbLink    string
		Interface string
		Vrid      uint
		VirtualIP string
		Priority  uint
		Delay     uint
		LogLevel  string
	}
	// Client config
	Client struct {
		Ssl    bool
		Local  string
		Remote string
	}
	// JWT config
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
	// WEB config
	Web struct {
		Title    string
		Active   bool
		Ssl      bool
		Ui       string
		Listen   string
		Port     uint
		Ca       string
		Crt      string
		Key      string
		DbLink   string
		TlsMin   string
		LogLevel string
	}
	// Proxy config
	Proxy struct {
		Title   string
		Active  bool
		Listen  string
		UdpPort uint
		TcpPort uint
		DbLink  string
		Api     struct {
			Bind string
			Port uint
			Ca   string
			Crt  string
			Key  string
		}
		Client struct {
			Crt string
			Key string
		}
		EvictMetrics   bool
		NsVpn          []string
		NsDirect       []string
		NsProvider     []string
		Resolvers      []string
		HealthHost     string
		ReserveTimeout uint
		LockTimeout    uint
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
	// Radius config
	Radius struct {
		Title  string
		Active bool
		Api    struct {
			Bind string
			Port uint
			Ca   string
			Crt  string
			Key  string
		}
		Client struct {
			Crt string
			Key string
		}
		DbLink       string
		HealthUser   string
		HealthPasswd string
		HealthKey    string
		HealthPort   uint
		DebugMode    bool
		LogLevel     string
	}
	// Media config
	Media struct {
		Title  string
		Active bool
		Api    struct {
			Bind string
			Port uint
			Ca   string
			Crt  string
			Key  string
		}
		Client struct {
			Crt string
			Key string
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
	// Nut config
	Nut struct {
		Title  string
		Active bool
		Api    struct {
			Bind string
			Port uint
			Ca   string
			Crt  string
			Key  string
		}
		Client struct {
			Crt string
			Key string
		}
		Driver   string
		Port     uint
		ApiUser  string
		ApiPass  string
		LogLevel string
	}
	// Messanger config
	Messanger struct {
		Title  string
		Active bool
		Api    struct {
			Bind string
			Port uint
			Ca   string
			Crt  string
			Key  string
		}
		Client struct {
			Crt string
			Key string
		}
		Telegram struct {
			Active bool
			ChatId int64
			Key    string
			Url    string
		}
		LogLevel string
	}
	// Ssh client config
	SshClient struct {
		Title  string
		Active bool
		Api    struct {
			Bind string
			Port uint
			Ca   string
			Crt  string
			Key  string
		}
		Client struct {
			Crt string
			Key string
		}
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
			Host      string
			Port      uint
			User      string
			ListVpn   string
			ListLocal string
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
	// Schedule config
	Schedule struct {
		Title  string
		Active bool
		Api    struct {
			Bind string
			Port uint
			Ca   string
			Crt  string
			Key  string
		}
		Client struct {
			Crt string
			Key string
		}
		DbLink   string
		LogLevel string
	}
	// Esp32 config
	Esp32 struct {
		Title  string
		Active bool
		Api    struct {
			Bind string
			Port uint
			Ca   string
			Crt  string
			Key  string
		}
		Client struct {
			Crt string
			Key string
		}
		Sensors struct {
			Bmp280   bool
			Aht25    bool
			Mics6814 bool
			Radsens  bool
			Ze08     bool
			Ds18b20  bool
		}
		DbLink   string
		Mac      string
		Check    bool
		LogLevel string
	}
	// Radio config
	Radio struct {
		Title  string
		Active bool
		Api    struct {
			Bind string
			Port uint
			Ca   string
			Crt  string
			Key  string
		}
		Client struct {
			Crt string
			Key string
		}
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
