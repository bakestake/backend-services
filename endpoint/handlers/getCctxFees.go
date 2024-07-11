package handlers

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	CC "endpoints/CC_artifact"
)

type CCTX_Payload struct {
	budsAmount *big.Int `json:"budsAmount"`
	tokenID *big.Int `json:"tokenId"`
	address common.Address `json:user`
	destEid uint32 `json:eid`
}

func GetCctxFees() gin.HandlerFunc{

	return func(c *gin.Context) {

		var payload CCTX_Payload;

		network := c.Param("networkName")

		if err := c.ShouldBindJSON(&payload); err != nil {
			fmt.Println("Error: invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		url, err := GetNetworkRpc(network);

		if err != nil {
			fmt.Println("error getting RPC for chain : ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting RPC for chain"})
			return
		}

		client, err1 := ethclient.Dial(url);

		if err1 != nil {
			fmt.Println("error creating client for chain : ", err1.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating client for chain"})
			return
		}

		contractAddress := common.HexToAddress("0x26705aD938791e61Aa64a2a9D808378805aec819")
		instance, err2 := CC.NewArtifacts(contractAddress, client)

		if err2 != nil {
			fmt.Println("error creating contract instance : ", err2.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating contract instance"})
			return
		}


		res, err3 := instance.GetCctxFees(&bind.CallOpts{}, payload.destEid, payload.budsAmount, payload.tokenID, payload.address)

		if err3 != nil {
			fmt.Println("error getting response from contract : ", err3.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error  getting response from contract"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"cctx_fees": res})

	}

}