package main

import (
	"ContractPoolAPI/config"
	"ContractPoolAPI/inputs"
	"io/ioutil"

	"ContractPoolAPI/helpers"

	"github.com/gin-gonic/gin"
)

func poolList(c *gin.Context) {
	inputPools, err := inputs.PoolInputs(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Error while getting input data : " + err.Error(),
		})
		return
	}

	// loading yaml file
	var poolData = config.LoadPools().PoolData
	currPoolData := helpers.PoolData(poolData, inputPools)
	for _, pools := range currPoolData.Pools {
		// get abi data from abi file name
		stake_abi_data, err := ioutil.ReadFile("ABI/" + pools.StakeTokenAddress + ".abi")
		if err != nil {
			c.JSON(400, gin.H{
				"error": pools.StakeTokenAddress + " Error reading abi file",
			})
			return
		}

		reward_abi_data, err := ioutil.ReadFile("ABI/" + pools.RewardTokenAddress + ".abi")
		if err != nil {
			c.JSON(400, gin.H{
				"error": pools.RewardTokenAddress + " Error reading abi file",
			})
			return
		}

		result, err := helpers.GetPoolInfo(c, string(stake_abi_data), string(reward_abi_data), inputPools, pools, currPoolData)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Error while getting pool info : " + err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"data": result,
		})

	}
}
