package handlers

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	CC "endpoints/CC_artifact"
)


func GetCctxFees() gin.HandlerFunc {
	return func(c *gin.Context) {

		network := c.Param("networkName")
		budsAmount := c.Param("budsAmount")
		tokenID := c.Param("tokenId")
		userAddress := c.Param("userAddress")
		destEid := c.Param("destEid")

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

		buds := new(big.Int)
		token := new(big.Int)

		buds.SetString(budsAmount, 10)
		token.SetString(tokenID, 10)

		destinationEid , err := strconv.ParseInt(destEid,10, 32)

		if err !=  nil{
			fmt.Println("error converting destId to uint32")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error converting destId to uint32"})
			return
		}

		fmt.Println("Calling GetCctxFees with parameters:")
		fmt.Printf("DestEid: %d, BudsAmount: %s, TokenID: %s, Address: %s\n", uint32(destinationEid), buds, token, common.HexToAddress(userAddress))

		opts := &bind.CallOpts{
			Pending:     true,
		}

		res, err := instance.GetCctxFees(opts, uint32(destinationEid) , buds, token, common.HexToAddress(userAddress))

		if err != nil {
			fmt.Println("error getting response from contract:", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting response from contract", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"cctx_fees": res})
	}
}
