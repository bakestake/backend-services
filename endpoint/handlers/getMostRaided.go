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


func GetMostRaided(db *sql.DB) gin.HandlerFunc {
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

			var lastLostBuds int64
			query := "SELECT lost_buds FROM daily_stats WHERE chain_id = $1"
			err = db.QueryRow(query, i+1).Scan(&lastLostBuds)
			if err != nil {
				log.Print(err.Error())
				fmt.Println("error fetching last lost buds from db:", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching last lost buds from db"})
				return
			}

			client, err := ethclient.Dial(url)
			if err != nil {
				fmt.Println("error creating client for chain:", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating client for chain"})
				return
			}

			contractAddress := common.HexToAddress("0x26705aD938791e61Aa64a2a9D808378805aec819")
			instance, err := Getter.NewArtifacts(contractAddress, client)
			if err != nil {
				fmt.Println("error creating contract instance:", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating contract instance"})
				return
			}

			budsLost, err := instance.GetBudsLostToRaids(&bind.CallOpts{})
			if err != nil {
				fmt.Println("error getting response from contract:", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting response from contract"})
				return
			}

			if budsLost.Cmp(big.NewInt(lastLostBuds)) > 0 {
				diff := new(big.Int).Sub(budsLost, big.NewInt(lastLostBuds))
				if diff.Cmp(big.NewInt(maxDif)) > 0 {
					maxDif = diff.Int64()
					maxChain = networks[i]
				}
			}

			updateQuery := `UPDATE daily_stats SET lost_buds = $1 WHERE chain_id = $2`
			_, err = db.Exec(updateQuery, budsLost, i+1)
			if err != nil {
				fmt.Println("Failed to update lost_buds record:", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lost_buds record"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"most_raided": maxChain})
	}
}
