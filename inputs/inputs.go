package inputs

import (
	"contract_pool/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PoolInputs(c *gin.Context) (models.InputPools, error) {
	var inputPools models.InputPools
	var err error

	protocol := c.PostForm("protocol")
	if protocol == "" {
		err = fmt.Errorf("protocol is required")
	}

	chain := c.PostForm("chain")
	if chain == "" {
		err = fmt.Errorf("chain is required")
	}

	inputPools = models.InputPools{
		ProtocolName: protocol,
		ChainName:    chain,
	}

	if err != nil {
		return inputPools, err
	} else {
		return inputPools, nil
	}

}
