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

func GetStakingRewardsForUserHandler() gin.HandlerFunc{

	return func(c *gin.Context) {

		userAddressString := c.Param("address")

		userAddress := common.HexToAddress(userAddressString)

		var networks = GetNetworksArray();

		var stakers [len(networks)]string;

		for i := 0; i < len(networks); i++{

			url, err := GetNetworkRpc(networks[i]);

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

			contractAddress := common.HexToAddress("0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe")
			instance, err2 := Getter.NewArtifacts(contractAddress, client)

			if err2 != nil {
				fmt.Println("error creating contract instance : ", err2.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating contract instance"})
				return
			}
			fmt.Println("response for address : ", userAddress)

			res, err3 := instance.GetRewardsForUser(&bind.CallOpts{},userAddress)

			fmt.Println("response from contract : ", res)

			if err3 != nil {
				fmt.Println("error getting response from contract : ", err3.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error  getting response from contract"})
				return
			}else{
				stakers[i]  = res.Div(res,big.NewInt(1e18)).String()
			}
		}

		c.JSON(http.StatusOK, gin.H{"staking_rewards": stakers})

	}

}