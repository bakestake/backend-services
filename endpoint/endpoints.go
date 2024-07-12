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

func InitiateDBWithData(db *sql.DB) error {
	// SQL statements to insert initial data
	statements := []string{
		`INSERT INTO daily_stats (chain_id, name, staked_buds, lost_buds) VALUES
			(1, 'amoy', 1000, 200),
			(2, 'bscTestnet', 1500, 300),
			(3, 'fuji', 450, 150),
			(4, 'arbSepolia', 200, 350),
			(5, 'baseSepolia', 750, 100);`,
		`INSERT INTO user_data (user_address, user_name, buds_won, games_won) VALUES
			('0x123abc...', 'Alice', 500, 10),
			('0x456def...', 'Bob', 700, 15);`,
	}

	// Execute each SQL statement
	for _, statement := range statements {
		_, err := db.Exec(statement)
		if err != nil {
			return fmt.Errorf("failed to execute SQL statement: %v", err)
		}
	}

	fmt.Println("Database initiated successfully with initial data.")
	return nil
}

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

	err = InitiateDBWithData(db)
	if err != nil {
		log.Fatal("Failed to initiate database:", err)
	}

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