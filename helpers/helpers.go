package helpers

import (
	"ContractPoolAPI/config"
	"ContractPoolAPI/models"
	"fmt"

	"github.com/chenzhijie/go-web3"
	"github.com/gin-gonic/gin"
)

func PoolData(configPoolData []config.PoolData, inputPools models.InputPools) config.PoolData {
	var pool_data config.PoolData
	for _, pools := range configPoolData {
		fmt.Println("pools", pools.ProtocolName, pools.ChainName)
		fmt.Println("pools", inputPools.ProtocolName, inputPools.ChainName)
		if pools.ProtocolName == inputPools.ProtocolName && pools.ChainName == inputPools.ChainName {
			pool_data = pools
		}
	}
	fmt.Println("pool_data", pool_data)
	return pool_data
}

func GetPoolInfo(c *gin.Context, stake_abi_data string, reward_abi_data string, inputPools models.InputPools, pools config.Pools, pool_data config.PoolData) (models.OutputPools, error) {
	var result models.OutputPools
	var error1 error
	web3, err := web3.NewWeb3(pool_data.RPC)
	if err != nil {
		error1 = fmt.Errorf("Error while getting web3 : " + err.Error())
	}
	// creating stake token contract
	stakeTokenContract, err := web3.Eth.NewContract(stake_abi_data, pools.StakeTokenAddress)
	if err != nil {
		error1 = fmt.Errorf("Error while getting stake token contract : " + err.Error())
	}

	// creating reward token contract
	rewardTokenContract, err := web3.Eth.NewContract(reward_abi_data, pools.RewardTokenAddress)
	if err != nil {
		error1 = fmt.Errorf("Error while getting reward token contract : " + err.Error())
	}

	stakeSymbol, err := stakeTokenContract.Call("symbol")
	if err != nil {
		error1 = fmt.Errorf("Error while getting stake symbol : " + err.Error())
	}
	rewardSymbol, err := rewardTokenContract.Call("symbol")
	if err != nil {
		error1 = fmt.Errorf("Error while getting reward symbol : " + err.Error())
	}
	stakeSym := fmt.Sprintf("%v", stakeSymbol)
	rewardSym := fmt.Sprintf("%v", rewardSymbol)
	result = models.OutputPools{
		Stake: stakeSym,
		Get:   rewardSym,
	}

	if error1 != nil {
		return result, error1
	} else {
		return result, nil
	}
}
