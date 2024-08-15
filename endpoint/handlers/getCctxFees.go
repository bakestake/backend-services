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
	BudsAmount string       `json:"budsAmount"`
	TokenID    string       `json:"tokenId"`
	Address    string       `json:"userAddress"`
	DestEid    uint32       `json:"destEid"`
}

func GetCctxFees() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload CCTX_Payload

		network := c.Param("networkName")

		if err := c.ShouldBindJSON(&payload); err != nil {
			fmt.Println("Error: invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		url, err := GetNetworkRpc(network)
		if err != nil {
			fmt.Println("error getting RPC for chain:", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting RPC for chain"})
			return
		}

		client, err := ethclient.Dial(url)
		if err != nil {
			fmt.Println("error creating client for chain:", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating client for chain"})
			return
		}

		contractAddress := common.HexToAddress("0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe")
		instance, err := CC.NewArtifacts(contractAddress, client)
		if err != nil {
			fmt.Println("error creating contract instance:", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating contract instance"})
			return
		}

		budsAmount := new(big.Int)
		tokenID := new(big.Int)

		budsAmount.SetString(payload.BudsAmount, 10)
		tokenID.SetString(payload.TokenID, 10)
		address := common.HexToAddress(payload.Address)

		fmt.Println("Calling GetCctxFees with parameters:")
		fmt.Printf("DestEid: %d, BudsAmount: %s, TokenID: %s, Address: %s\n", payload.DestEid, budsAmount.String(), tokenID.String(), address.Hex())

		opts := &bind.CallOpts{
			Pending:     true,
		}

		res, err := instance.GetCctxFees(opts, payload.DestEid, budsAmount, tokenID, address)
		if err != nil {
			fmt.Println("error getting response from contract:", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting response from contract", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"cctx_fees": res})
	}
}
