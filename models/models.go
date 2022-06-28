package models

type InputPools struct {
	ProtocolName string `yaml:"protocol_name"`
	ChainName    string `yaml:"chain_name"`
	PoolName     string `yaml:"pool_name"`
}

type OutputPools struct {
	Stake    string `yaml:"stake"`
	Get      string `yaml:"get"`
	Duration string `yaml:"duration"`
	APR      string `yaml:"apr"`
}
