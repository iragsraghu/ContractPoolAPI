package config

import (
	"log"
	"runtime"

	"github.com/spf13/viper"
)

type PoolsConfig struct {
	PoolData []PoolData `yaml:"pools_data"`
}

type PoolData struct {
	ProtocolName string  `yaml:"protocol_name"`
	ChainName    string  `yaml:"chain_name"`
	RPC          string  `yaml:"rpc"`
	Pools        []Pools `yaml:"pools"`
}

type Pools struct {
	PoolName             string `yaml:"pool_name"`
	StakeContractAddress string `yaml:"stake_contract_address"`
	StakeTokenAddress    string `yaml:"stake_token_address"`
	RewardTokenAddress   string `yaml:"reward_token_address"`
}

func LoadPools() PoolsConfig {
	var poolsConfig PoolsConfig
	//For Local Config
	if runtime.GOOS == "darwin" || runtime.GOOS == "windows" {
		viper.SetConfigName("default_dev")
		// Set the path to look for the configurations file
		viper.AddConfigPath("./config")
		// Enable VIPER to read Environment Variables
		viper.AutomaticEnv()
		viper.SetConfigType("yml")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err.Error())
		}
		err := viper.Unmarshal(&poolsConfig)
		if err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}
		return poolsConfig
	}
	return poolsConfig
}
