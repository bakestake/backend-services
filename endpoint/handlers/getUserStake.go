package handlers

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	Getter "endpoints/getter_artifact"
)

func GetUserStake() gin.HandlerFunc {
	return func(c *gin.Context) {

		userAddressString := c.Param("address")
		network := c.Param("network")


		userAddressbytes := []byte(userAddressString)

		userAddress := common.BytesToAddress(userAddressbytes)

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

		contractAddress := common.HexToAddress("0x26705aD938791e61Aa64a2a9D808378805aec819")
		instance, err2 := Getter.NewArtifacts(contractAddress, client)

		if err2 != nil {
			fmt.Println("error creating contract instance:", err2.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating contract instance"})
			return
		}

		stake, err3 := instance.GetUserStakes(&bind.CallOpts{}, userAddress)

		if err3 != nil {
			fmt.Println("error getting response from contract:", err3.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting response from contract"})
			return
		}

		var budsAmount = big.NewInt(0);

		for i := 0; i < len(stake); i++{
			budsAmount.Add(budsAmount,stake[i].BudsAmount)
		}

		c.JSON(http.StatusOK, gin.H{"userStake" : budsAmount})
	}
}
