package handlers

import (
	"database/sql"
	Getter "endpoints/getter_artifact"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)



func GetMostRaided(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {

		var networks = GetNetworksArray();		

		maxChain := networks[0];

		var maxDif int64

		for i := 0; i < len(networks); i++{

			url, err := GetNetworkRpc(networks[i]);
			
			if err != nil {
				fmt.Println("error getting RPC for chain : ", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting RPC for chain"})
				return
			}

			var last_staked_buds int64
			query := "SELECT lost_buds FROM daily_stats WHERE chain_id = $1"
			err = db.QueryRow(query, i+1).Scan(&last_staked_buds)

			if err != nil {
				fmt.Println("error fetching last staked buds from db: ", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching last staked buds from db"})
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

			buds, err3 := instance.GetBudsLostToRaids(&bind.CallOpts{})

			if err3 != nil {
				fmt.Println("error getting response from contract : ", err3.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error  getting response from contract"})
				return
			}

			if buds.Cmp(big.NewInt(last_staked_buds)) > 0 {
				diff := buds.Sub(buds, big.NewInt(last_staked_buds))
				if diff.Cmp(big.NewInt(maxDif)) > 0 {
					maxChain = networks[i]
				}
			}

			updateQuery := `UPDATE daily_stats SET lost_buds = $1, WHERE chain_id = $2`
			_, err = db.Exec(updateQuery, buds, i+1)

			if err != nil {
			fmt.Println("Failed to update lost_buds buds record")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ailed to update staked buds record"})
			return
		}
		}

		c.JSON(http.StatusOK, gin.H{"most_stakes": maxChain})

	}

}