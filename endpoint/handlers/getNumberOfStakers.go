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

func GetNumberOfStakers() gin.HandlerFunc{

	return func(c *gin.Context) {

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

			contractAddress := common.HexToAddress("0x26705aD938791e61Aa64a2a9D808378805aec819")
			instance, err2 := Getter.NewArtifacts(contractAddress, client)

			if err2 != nil {
				fmt.Println("error creating contract instance : ", err2.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating contract instance"})
				return
			}

			res, err3 := instance.GetNumberOfStakers(&bind.CallOpts{})

			if err3 != nil {
				fmt.Println("error getting response from contract : ", err3.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error  getting response from contract"})
				return
			}else{
				stakers[i] = res.String();
			}
		}

		c.JSON(http.StatusOK, gin.H{"no_of_stakers": stakers})

	}

}