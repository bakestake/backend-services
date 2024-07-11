package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllChainsHandler() gin.HandlerFunc{

	return func(c *gin.Context) {

		var networks = GetNetworksArray();

		c.JSON(http.StatusOK, gin.H{"aprs": networks})

	}

}