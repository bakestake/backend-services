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

func GetCurrentAprHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var networks = GetNetworksArray()

		// Use a map to store chain names as keys and APRs as values
		aprs := make(map[string]string)

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

			apr, err3 := instance.GetCurrentApr(&bind.CallOpts{})

			if err3 != nil {
				fmt.Println("error getting response from contract:", err3.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting response from contract"})
				return
			}

			aprs[network] = apr.String()
		}

		c.JSON(http.StatusOK, aprs)
	}
}
