package runtime

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type MysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
	MaxOpen  int
	MaxIdle  int
	MaxLife  int
	Migrate  bool
}

type Infura struct {
	Https string
	Wss   string
}

type Contract struct {
	IndexPath      string
	MarketAddress  string
	PurchasedEvent string
	FallbackIndex  int64
	IndexFile      string
	DialInterval   int64
	PullInterval   int64
	RequestTimeout int64
	Step           int64
}

type Config struct {
	Mysql    MysqlConfig
	Contract Contract
	Infura   Infura
}

func loadConfig(configFile string) (*Config, error) {
	config := &Config{}

	content, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
