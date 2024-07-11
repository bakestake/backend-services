package handlers

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func GetNetworkRpc(networkRPC string) (string, error){

	switch  {
		case networkRPC == "polygon":
			str := os.Getenv("RPC_URL_POLYGON");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "amoy":
			str := os.Getenv("RPC_URL_AMOY");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "avax":
			str := os.Getenv("RPC_URL_AVAX");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "fuji":
			str := os.Getenv("RPC_URL_FUJI");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "bsc":
			str := os.Getenv("RPC_URL_BSC");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "bscTestnet":
			str := os.Getenv("RPC_URL_BSCTESTNET");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "bera":
			str := os.Getenv("RPC_URL_BERA");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "beraTestnet":
			str := os.Getenv("RPC_URL_BERATESTNET");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "arb":
			str := os.Getenv("RPC_URL_ARB");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "arbSepolia":
			str := os.Getenv("RPC_URL_ARBSEPOLIA");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "base":
			str := os.Getenv("RPC_URL_BASE");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "baseSepolia":
			str := os.Getenv("RPC_URL_BASESEPOLIA");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "core":
			str := os.Getenv("RPC_URL_CORE");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;

		case networkRPC == "coreTestnet":
			str := os.Getenv("RPC_URL_CORETESTNET");
			if str == ""{
				return "", fmt.Errorf("RPC Url not configured for given chain", networkRPC);
			}
			return str, nil;
	}

	return "", fmt.Errorf("network is not configured")
}

// Settin up the client
func InitNetwork(networkRPC string) {
	var err error
	client, err = ethclient.Dial(networkRPC)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
}

// Import wallet
func ImportWallet(privateKey string) (common.Address, *ecdsa.PrivateKey) {
	importedPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := importedPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, importedPrivateKey
}


func PrepareTransaction(publicKey common.Address, privateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	nonce, err := client.PendingNonceAt(context.Background(), publicKey)
	if err != nil {
	log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
	log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
	log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
	panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(15000000) // It can change
	auth.GasPrice = gasPrice

	return auth
}

func GetNetworksArray() [5]string {
    return [5]string{"amoy", "bscTestnet", "fuji", "arbSepolia", "baseSepolia"/*, "beraTestnet","coreTestnet"*/}
}

func GetChainIds() [5] string{
	return [5]string{"1","2","3","4","5"}
}