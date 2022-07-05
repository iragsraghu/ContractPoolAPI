package helpers

import (
	"ContractPoolAPI/config"
	"ContractPoolAPI/models"
	"fmt"
	"strconv"

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

func GetPoolInfo(c *gin.Context, stake_abi_data string, reward_abi_data string, main_contract_abi_data string, inputPools models.InputPools, pools config.Pools, pool_data config.PoolData) (models.OutputPools, error) {
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
	stakeSymbol, err := stakeTokenContract.Call("symbol")
	if err != nil {
		error1 = fmt.Errorf("Error while getting stake symbol : " + err.Error())
	}
	stakeSym := fmt.Sprintf("%v", stakeSymbol)
	// end

	// creating reward token contract
	rewardTokenContract, err := web3.Eth.NewContract(reward_abi_data, pools.RewardTokenAddress)
	if err != nil {
		error1 = fmt.Errorf("Error while getting reward token contract : " + err.Error())
	}
	rewardSymbol, err := rewardTokenContract.Call("symbol")
	if err != nil {
		error1 = fmt.Errorf("Error while getting reward symbol : " + err.Error())
	}
	rewardSym := fmt.Sprintf("%v", rewardSymbol)
	// end

	// get Current block
	var countDown string
	currentBlock, err := web3.Eth.GetBlockNumber()
	if err != nil {
		error1 = fmt.Errorf("Error while getting current block : " + err.Error())
	}
	contract, err := web3.Eth.NewContract(main_contract_abi_data, pools.StakeContractAddress)
	if err != nil {
		error1 = fmt.Errorf("Error while getting main contract : " + err.Error())
	}
	if pools.EndBlock == "nil" {
		countDown = "Nil"
	} else {
		end_block, err := contract.Call(pools.EndBlock)
		if err != nil {
			error1 = fmt.Errorf("Error while getting end block : " + err.Error())
		}
		_endBlock := fmt.Sprintf("%v", end_block)                   // convert to string
		endBlock, _ := strconv.ParseUint(string(_endBlock), 10, 64) // convert to uint64

		remainingBlock := endBlock - currentBlock     // get remaining block
		remainingBlockInSec := remainingBlock * 3     // convert to seconds
		countDown = GetCountDown(remainingBlockInSec) // get count down time into readable format
	}

	// assigning values to output pools
	result = models.OutputPools{
		Stake:          stakeSym,
		Get:            rewardSym,
		BlockCountDown: countDown,
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
