package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"ethereum_blockchain/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"fmt"
	"math/big"
	"os"
	"gopkg.in/yaml.v2"
)


type Configuration struct {
	NodeAddress 	string
	Key		string
	Paraphrase	string
}

func main(){
	// connect to an ethereum node  hosted by infura
	file, _ := os.Open("conf.yaml")
	defer file.Close()
	config := Configuration{}
	err :=  yaml.Unmarshal([]byte(file), &Configuration{})
	if err != nil {
		fmt.Println("error:", err)
	}

	blockchain, err := ethclient.Dial(config.NodeAddress)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(config.Key), config.Paraphrase)

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	address, _, _, _:= contracts.DeployRating(auth, blockchain, big.NewInt(5), "Some Provider", "Some Consumer")

	fmt.Printf("Contract pending deploy: 0x%x\n", address)}