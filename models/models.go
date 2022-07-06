package models

type InputPools struct {
	ProtocolName string `yaml:"protocol_name"`
	ChainName    string `yaml:"chain_name"`
	PoolName     string `yaml:"pool_name"`
}

type OutputPools struct {
	Stake          string `yaml:"stake"`
	Get            string `yaml:"get"`
	BlockCountDown string `yaml:"block_count_down"`
	APR            string `yaml:"apr"`
	TotalStaked    string `yaml:"total_staked"`
}
