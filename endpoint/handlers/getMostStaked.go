package handlers

import (
	"database/sql"
	Getter "endpoints/getter_artifact"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)



func GetMostStaked(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var networks = GetNetworksArray()
		maxChain := networks[0]
		var maxDif int64

		for i := 0; i < len(networks); i++ {
			url, err := GetNetworkRpc(networks[i])
			if err != nil {
				fmt.Println("error getting RPC for chain:", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting RPC for chain"})
				return
			}

			var lastStakedBuds int64
			query := "SELECT staked_buds FROM daily_stats WHERE chain_id = $1"
			err = db.QueryRow(query, i+1).Scan(&lastStakedBuds)
			if err != nil {
				fmt.Println("error fetching last staked buds from db:", err.Error())
				log.Print(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching last staked buds from db"})
				return
			}

			client, err := ethclient.Dial(url)
			if err != nil {
				fmt.Println("error creating client for chain:", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating client for chain"})
				return
			}

			contractAddress := common.HexToAddress("0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe")
			instance, err := Getter.NewArtifacts(contractAddress, client)
			if err != nil {
				fmt.Println("error creating contract instance:", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating contract instance"})
				return
			}

			buds, err := instance.GetlocalStakedBuds(&bind.CallOpts{})
			if err != nil {
				fmt.Println("error getting response from contract:", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting response from contract"})
				return
			}

			if buds.Cmp(big.NewInt(lastStakedBuds)) > 0 {
				diff := new(big.Int).Sub(buds, big.NewInt(lastStakedBuds))
				if diff.Cmp(big.NewInt(maxDif)) > 0 {
					maxDif = diff.Int64()
					maxChain = networks[i]
				}
			}

			updateQuery := `UPDATE daily_stats SET staked_buds = $1 WHERE chain_id = $2`
			_, err = db.Exec(updateQuery, buds.Int64(), i+1)
			if err != nil {
				fmt.Println("Failed to update staked buds record:", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update staked buds record"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"most_staked": maxChain, "stakedAmount":maxDif})
	}
}
