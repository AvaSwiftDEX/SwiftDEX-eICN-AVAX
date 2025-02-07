package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	url := "http://10.210.143.97:7545"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")

	latestHeight, err := client.BlockNumber(ctx)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("latest block height: ", latestHeight)

	networkID, _ := client.NetworkID(ctx)
	fmt.Println("network ID: ", networkID)
}
