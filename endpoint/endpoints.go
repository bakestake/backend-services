package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"endpoints/handlers"
)

func main() {
    err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}	

	router := gin.Default()

	router.GET("/getAPR", handlers.GetCurrentAprHandler())
	router.GET("/getGlobalLiq", handlers.GetGlobalStakedBudsHandler())
	router.GET("/getLocalLiq", handlers.GetLocalStakedBuds())
	router.GET("/getStakersCount", handlers.GetNumberOfStakers())

	// Run the http server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start the server", err)
	}

}