package config

import (
	"github.com/dittonetwork/executor-avs/pkg/config/viper"
)

const configPath = "./cmd/operator/config/config.yml"

type Config struct {
	Ethereum        Ethereum        `mapstructure:"ethereum"`
	DittoEntryPoint DittoEntryPoint `mapstructure:"ditto_entrypoint"`
}

type Ethereum struct {
	NodeURL string `mapstructure:"node_url"`
}

type DittoEntryPoint struct {
	ContractAddress string `mapstructure:"contract_address"`
}

func New() Config {
	var cfg Config
	viper.Load(configPath, &cfg)

	return cfg
}
