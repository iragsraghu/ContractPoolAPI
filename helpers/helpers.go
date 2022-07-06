package helpers

import (
	"ContractPoolAPI/config"
	"ContractPoolAPI/models"
	"fmt"

	poolpkg "ContractPoolAPI/pools"

	"github.com/chenzhijie/go-web3"
	"github.com/gin-gonic/gin"
)

func PoolData(configPoolData []config.PoolData, inputPools models.InputPools) config.PoolData {
	var pool_data config.PoolData
	for _, pools := range configPoolData {
		if pools.ProtocolName == inputPools.ProtocolName && pools.ChainName == inputPools.ChainName {
			pool_data = pools
		}
	}
	return pool_data
}

// get pool info from web3
func GetPoolInfo(c *gin.Context, stake_abi_data string, reward_abi_data string, main_contract_abi_data string, inputPools models.InputPools, pools config.Pools, pool_data config.PoolData) (models.OutputPools, error) {
	var result models.OutputPools
	var error1 error
	web3, err := web3.NewWeb3(pool_data.RPC)
	if err != nil {
		error1 = fmt.Errorf("Error while getting web3 : " + err.Error())
	}
	// creating stake token contract
	stakeSymbol, totalStaked, errors := poolpkg.GetStakedTokenDetails(web3, stake_abi_data, pools)
	if errors != nil {
		error1 = fmt.Errorf("Error while getting stake token details : " + errors[0].Error())
	}
	// end

	// creating reward token contract
	rewardSymbol, errors := poolpkg.GetRewardTokenDetails(web3, reward_abi_data, pools)
	if errors != nil {
		error1 = fmt.Errorf("Error while getting reward token details : " + errors[0].Error())
	}
	// end

	// get Current block
	blockCount, errors := poolpkg.GetMainContractDetails(web3, main_contract_abi_data, pools)
	if errors != nil {
		error1 = fmt.Errorf("Error while getting main contract details : " + errors[0].Error())
	}
	// end

	// assigning values to output pools
	result = models.OutputPools{
		Stake:          stakeSymbol,
		Get:            rewardSymbol,
		BlockCountDown: blockCount,
		TotalStaked:    totalStaked,
	}
	// end

	// returning result
	if error1 != nil {
		return result, error1
	} else {
		return result, nil
	}
	// end
}

// convert seconds to days, hours, minutes, seconds
func GetCountDown(remainingBlockInSec uint64) string {
	var days, hours, minutes, seconds uint64
	days = remainingBlockInSec / 86400
	remainingBlockInSec = remainingBlockInSec % 86400
	hours = remainingBlockInSec / 3600
	remainingBlockInSec = remainingBlockInSec % 3600
	minutes = remainingBlockInSec / 60
	remainingBlockInSec = remainingBlockInSec % 60
	seconds = remainingBlockInSec
	return fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds", days, hours, minutes, seconds)
}
