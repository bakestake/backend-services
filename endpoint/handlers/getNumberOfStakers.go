package handlers

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	Getter "endpoints/getter_artifact"
)

func GetNumberOfStakers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var networks = GetNetworksArray()

		// Use a map to store chain names as keys and number of stakers as values
		numberOfStakers := make(map[string]string)

		for _, network := range networks {
			url, err := GetNetworkRpc(network)

			if err != nil {
				fmt.Println("error getting RPC for chain:", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting RPC for chain"})
				return
			}

			client, err1 := ethclient.Dial(url)

			if err1 != nil {
				fmt.Println("error creating client for chain:", err1.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating client for chain"})
				return
			}

			contractAddress := common.HexToAddress("0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe")
			instance, err2 := Getter.NewArtifacts(contractAddress, client)

			if err2 != nil {
				fmt.Println("error creating contract instance:", err2.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating contract instance"})
				return
			}

			res, err3 := instance.GetNumberOfStakers(&bind.CallOpts{})

			if err3 != nil {
				fmt.Println("error getting response from contract:", err3.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting response from contract"})
				return
			}

			numberOfStakers[network] = res.String()
		}

		c.JSON(http.StatusOK, numberOfStakers)
	}
}
