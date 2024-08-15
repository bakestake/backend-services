package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRaidFees() gin.HandlerFunc {
	return func(c *gin.Context) {

		var chain = c.Param("networkName")

		var fees = 0.0

		switch{
			case chain == "amoy":
				fees = 0.5
			case chain == "bscTestnet":
				fees = 0.0041
			case chain == "fuji":
				fees = 0.05
			case chain == "arbSepolia":
				fees = 0.5
			case chain == "baseSepolia":
				fees = 0.005
			case chain == "beraTestnet":
				fees = 0.005
			default:
				fmt.Println("error : No such chain configured")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "No such chain"})
		}

		c.JSON(http.StatusOK, gin.H{"raid_fees": fees})

	}
}