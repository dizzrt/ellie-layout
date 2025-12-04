package conf

import (
	"os"
	"sync"

	"github.com/dizzrt/ellie-layout/internal/infra/utils"
	"github.com/dizzrt/ellie/config"
	"github.com/google/uuid"
)

// inject by ldflags
// e.g. go build -ldflags "-X github.com/dizzrt/ellie-layout/internal/conf.Version=x.y.z"
var (
	Service     string = "ellie-layout"
	Version     string = "dev"
	Hostname, _        = os.Hostname()
	ServiceID   string = uuid.NewString()
)

var (
	ac   *AppConfig
	once sync.Once
)

type AppConfig struct {
	ENV     string
	Address string

	Log    Log    `mapstructure:"log"`
	Server Server `mapstructure:"server"`
}

type Server struct {
	GRPC GRPCServer `mapstructure:"grpc"`
	HTTP HTTPServer `mapstructure:"http"`
}

type GRPCServer struct {
	Addr string `mapstructure:"addr"`
}

type HTTPServer struct {
	Addr string `mapstructure:"addr"`
}

type Log struct {
	File       string `mapstructure:"file"`
	Symlink    string `mapstructure:"symlink"`
	Level      string `mapstructure:"level"`
	MaxAge     string `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	OutputType string `mapstructure:"output_type"`
}

func GetAppConfig() *AppConfig {
	once.Do(func() {
		c := config.NewStdViperConfig()
		if err := c.Load(); err != nil {
			panic(err)
		}

		var tmp AppConfig
		if err := c.Unmarshal(&tmp); err != nil {
			panic(err)
		}

		tmp.ENV = c.V().GetString("ENV")
		addr, err := utils.GetLocalAddress()
		if err != nil {
			panic(err)
		}
		tmp.Address = addr

		ac = &tmp
	})

	return ac
}
