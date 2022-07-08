package pools

import (
	"contract_pool/config"
	"fmt"
	"strconv"

	"github.com/chenzhijie/go-web3"
	"github.com/ethereum/go-ethereum/common"
)

func GetStakedTokenDetails(web3 *web3.Web3, stake_abi_data string, pools config.Pools) (string, string, []error) {
	var errors []error
	stakeTokenContract, err := web3.Eth.NewContract(stake_abi_data, pools.StakeTokenAddress)
	if err != nil {
		errors = append(errors, fmt.Errorf("Error while getting stake token contract : "+err.Error()))
	}
	stakeSymbol, err := stakeTokenContract.Call("symbol")
	if err != nil {
		errors = append(errors, fmt.Errorf("Error while getting stake symbol : "+err.Error()))
	}
	stakeSym := fmt.Sprintf("%v", stakeSymbol)

	address := common.HexToAddress(pools.StakeContractAddress)
	args := []interface{}{
		address,
	}
	stakedTokens, err := stakeTokenContract.Call("balanceOf", args...)
	if err != nil {
		errors = append(errors, fmt.Errorf("Error while getting staked tokens : "+err.Error()))
	}
	totalStaked := fmt.Sprintf("%v", stakedTokens)
	return stakeSym, totalStaked, errors
}

func GetRewardTokenDetails(web3 *web3.Web3, reward_abi_data string, pools config.Pools) (string, []error) {
	var errors []error
	rewardTokenContract, err := web3.Eth.NewContract(reward_abi_data, pools.RewardTokenAddress)
	if err != nil {
		errors = append(errors, fmt.Errorf("Error while getting reward token contract : "+err.Error()))
	}
	rewardSymbol, err := rewardTokenContract.Call("symbol")
	if err != nil {
		errors = append(errors, fmt.Errorf("Error while getting reward symbol : "+err.Error()))
	}
	rewardSym := fmt.Sprintf("%v", rewardSymbol)
	return rewardSym, errors
}

func GetMainContractDetails(web3 *web3.Web3, main_abi_data string, pools config.Pools) (string, []error) {
	var countDown string
	var errors []error

	currentBlock, err := web3.Eth.GetBlockNumber()
	if err != nil {
		errors = append(errors, fmt.Errorf("Error while getting current block : "+err.Error()))
	}
	contract, err := web3.Eth.NewContract(main_abi_data, pools.StakeContractAddress)
	if err != nil {
		errors = append(errors, fmt.Errorf("Error while getting main contract : "+err.Error()))
	}
	if pools.EndBlock == "nil" {
		countDown = "Nil"
	} else {
		end_block, err := contract.Call(pools.EndBlock)
		if err != nil {
			errors = append(errors, fmt.Errorf("Error while getting end block : "+err.Error()))
		}
		_endBlock := fmt.Sprintf("%v", end_block)                   // convert to string
		endBlock, _ := strconv.ParseUint(string(_endBlock), 10, 64) // convert to uint64

		remainingBlock := endBlock - currentBlock     // get remaining block
		remainingBlockInSec := remainingBlock * 3     // convert to seconds
		countDown = GetCountDown(remainingBlockInSec) // get count down time into readable format
	}
	return countDown, errors
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
