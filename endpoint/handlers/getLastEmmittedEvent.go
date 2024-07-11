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
	user common.Address `json:user`
	startBlock big.Int `json:eid`
}
type LogStake struct {
	user   common.Address
	tokenId big.Int
	budsAmount big.Int
	timeStamp big.Int
	localStakedBuds big.Int
	latestAPR big.Int
}

type LogRaids struct {
	user   common.Address
	isSuccess bool
	isBoosted bool
	rewardTaken big.Int
	boostsUsedInLastSevenDays big.Int
}

func GetEventConfirmation() gin.HandlerFunc{
	return func(c *gin.Context) {

		var payload StakeConfirmationPayload;

		network := c.Param("networkName")
		event := c.Param("event")

		if err := c.ShouldBindJSON(&payload); err != nil {
			fmt.Println("Error: invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		url, err := GetNetworkRpc(network);

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

		currentBlock, err := client.BlockNumber(c)

		if err != nil{
			fmt.Println("error getting block number : ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error  getting block number"})
			return
		}

		query := ethereum.FilterQuery{
			FromBlock: &payload.startBlock,
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

		logStakeSig := []byte("event Staked(address,uint256,,uint256,uint256,uint256,uint256);")
		logStakeSigHash := crypto.Keccak256Hash(logStakeSig)
		logRaidSig := []byte("event Raided(address,bool,bool,uint256,uint256)")
		logRaidSigHash := crypto.Keccak256Hash(logRaidSig)


		for _, vLog := range logs {
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)    
			fmt.Println(vLog.TxHash.Hex())   

			if event == "1"{
				switch vLog.Topics[0].Hex() {
					case logStakeSigHash.Hex():
						fmt.Printf("Log Name: Stake\n")

						var stakeEvent LogStake

						err := contractAbi.UnpackIntoInterface(&stakeEvent, "Staked", vLog.Data)
						if err != nil {
							log.Fatal(err)
						}

						stakeEvent.user = common.HexToAddress(vLog.Topics[1].Hex())
						stakeEvent.budsAmount = (*vLog.Topics[3].Big())

						if(stakeEvent.user == payload.user){
							c.JSON(http.StatusOK, gin.H{"stake_found": true})
						}
				
				}
			} else if event == "2"{
				switch vLog.Topics[0].Hex() {
					case logRaidSigHash.Hex():
						fmt.Printf("Log Name: Raided\n")

						var raidEvent LogRaids

						err := contractAbi.UnpackIntoInterface(&raidEvent, "Raided", vLog.Data)
						if err != nil {
							log.Fatal(err)
						}

						raidEvent.user = common.HexToAddress(vLog.Topics[1].Hex())
						raidEvent.isSuccess = vLog.Topics[2].Big().Cmp(big.NewInt(1)) == 0

						if raidEvent.user == payload.user {
							c.JSON(http.StatusOK, gin.H{"raid_status": gin.H{
								"isSuccess": raidEvent.isSuccess,
							}})
							return
						}

				
				}
			}

			

		}

		c.JSON(http.StatusOK, gin.H{"stake_found": false})
	}
}