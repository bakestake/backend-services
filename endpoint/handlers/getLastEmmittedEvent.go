package handlers

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	LibGlobal "endpoints/lib_artifact"
)

type StakeConfirmationPayload struct {
	User      common.Address `json:"user"`
	StartBlock big.Int       `json:"startblock"`
}

type LogStake struct {
	User           common.Address
	TokenId        *big.Int
	BudsAmount     *big.Int
	TimeStamp      *big.Int
	LocalStakedBuds *big.Int
	LatestAPR      *big.Int
}

type LogRaids struct {
	User                    common.Address
	IsSuccess               bool
	IsBoosted               bool
	RewardTaken             *big.Int
	BoostsUsedInLastSevenDays *big.Int
}

func GetEventConfirmation() gin.HandlerFunc {
	return func(c *gin.Context) {

		var payload StakeConfirmationPayload

		network := c.Param("networkName")
		event := c.Param("event")

		if err := c.ShouldBindJSON(&payload); err != nil {
			fmt.Println("Error: invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		url, err := GetNetworkRpc(network)
		if err != nil {
			fmt.Println("error getting RPC for chain: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting RPC for chain"})
			return
		}

		client, err := ethclient.Dial(url)
		if err != nil {
			fmt.Println("error creating client for chain: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating client for chain"})
			return
		}

		contractAddress := common.HexToAddress("0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe")

		currentBlock, err := client.BlockNumber(context.Background())
		if err != nil {
			fmt.Println("error getting block number: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting block number"})
			return
		}

		query := ethereum.FilterQuery{
			FromBlock: &payload.StartBlock,
			ToBlock:   big.NewInt(int64(currentBlock)),
			Addresses: []common.Address{
				contractAddress,
			},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}

		contractAbi, err := abi.JSON(strings.NewReader(string(LibGlobal.ArtifactsABI)))
		if err != nil {
			log.Fatal(err)
		}

		logStakeSig := []byte("Staked(address,uint256,uint256,uint256,uint256,uint256)")
		logStakeSigHash := crypto.Keccak256Hash(logStakeSig)

		logRaidSig := []byte("Raided(address,bool,bool,uint256,uint256)")
		logRaidSigHash := crypto.Keccak256Hash(logRaidSig)

		for _, vLog := range logs {
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())

			switch vLog.Topics[0].Hex() {
			case logStakeSigHash.Hex():
				if event == "1" {
					fmt.Printf("Log Name: Stake\n")

					var stakeEvent LogStake

					err := contractAbi.UnpackIntoInterface(&stakeEvent, "Staked", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}

					stakeEvent.User = common.HexToAddress(vLog.Topics[1].Hex())
					stakeEvent.BudsAmount = new(big.Int).SetBytes(vLog.Data)

					if stakeEvent.User == payload.User {
						c.JSON(http.StatusOK, gin.H{"stake_found": true})
						return
					}
				}

			case logRaidSigHash.Hex():
				if event == "2" {
					fmt.Printf("Log Name: Raided\n")

					var raidEvent LogRaids

					err := contractAbi.UnpackIntoInterface(&raidEvent, "Raided", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}

					raidEvent.User = common.HexToAddress(vLog.Topics[1].Hex())
					raidEvent.IsSuccess = vLog.Topics[2].Big().Cmp(big.NewInt(1)) == 0

					if raidEvent.User == payload.User {
						c.JSON(http.StatusOK, gin.H{
							"raid_status": gin.H{
								"isSuccess": raidEvent.IsSuccess,
							},
						})
						return
					}
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{"stake_found": false})
	}
}
