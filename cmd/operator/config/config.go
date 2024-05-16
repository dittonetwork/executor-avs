package config

type Config struct {
	Executor        Executor        `mapstructure:"executor"`
	Ethereum        Ethereum        `mapstructure:"ethereum"`
	DittoEntryPoint DittoEntryPoint `mapstructure:"ditto_entrypoint"`
}

type Executor struct {
	PrivateKey string `mapstructure:"private_key"`
}

type Ethereum struct {
	NodeURL string `mapstructure:"node_url"`
}

type DittoEntryPoint struct {
	ContractAddress string `mapstructure:"contract_address"`
}
