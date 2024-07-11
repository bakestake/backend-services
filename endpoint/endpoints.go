package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"endpoints/handlers"
)

func main() {
	// cwd, err := os.Getwd()
    // if err != nil {
    //     log.Fatal(err)
    // }
    // log.Printf("Current working directory: %s", cwd)

	// envPath := filepath.Join(cwd, "/.env")

    err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}	

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbInfo := fmt.Sprintf("host=postgres port=5432 user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/getApr", handlers.GetCurrentAprHandler())
	router.GET("/totalStakedBudsAcrossAllChains", handlers.GetGlobalStakedBudsHandler())
	router.GET("/getLocalBudsCount", handlers.GetLocalStakedBuds())
	router.GET("/getStakersCount", handlers.GetNumberOfStakers())
	router.GET("/getRewards/:address", handlers.GetStakingRewardsForUserHandler())
	router.GET("/getFees/:networkName", handlers.GetCctxFees())
	router.GET("/getBudsBalance/:networkName/:address")
	router.GET("/getCurrentBlockNumber/:networkName", handlers.GetBlockNumber())
	router.GET("/getEvents/:networkName/:event", handlers.GetEventConfirmation())
	router.GET("/mostBaked", handlers.GetMostStaked(db))
	router.GET("/mostRekt", handlers.GetMostRaided(db))


	// Run the http server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start the server", err)
	}

}