package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kimroniny/SuperRunner-eICN-eth2/SR2PC"
)

const PRIVATE_KEY = "c45ba5d6de0e502aefd23c98b40a2c9018e2e0286dde4fdb542ded619cefc8bd"

func deploy(ctx context.Context, client *ethclient.Client) {
	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice := 1

	// get chainID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(4000000)
	
	auth.GasPrice = big.NewInt(int64(gasPrice))

	_chainId := big.NewInt(1)
	address, tx, instance, err := SR2PC.DeploySR2PC(auth, client, _chainId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("address: ", address.Hex())
	fmt.Println("tx: ", tx.Hash().Hex())
	_ = instance
}

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

	deploy(ctx, client)
}
