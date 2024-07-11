package handlers

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func GetBlockNumber() gin.HandlerFunc{

	return func(c *gin.Context){

		network := c.Param("networkName")

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

			blockNo, err2 := client.BlockNumber(c)

			if err != nil {
				fmt.Println("error getting block number for chain : ", err2.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting block number for chain"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"block_number": blockNo})

	}
}