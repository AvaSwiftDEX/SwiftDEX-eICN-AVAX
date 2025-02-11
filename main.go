package main

import (
	"context"
	"flag"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kimroniny/SuperRunner-eICN-eth2/client"
	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
	"github.com/kimroniny/SuperRunner-eICN-eth2/scripts"
	sdk "github.com/kimroniny/SuperRunner-eICN-eth2/sdk"
	"github.com/kimroniny/SuperRunner-eICN-eth2/server"
	"github.com/kimroniny/SuperRunner-eICN-eth2/watcher"
)

func main() {
	// 定义命令行参数
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	ctx := context.Background()

	// 使用命令行参数中指定的配置文件路径
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	address, err := scripts.Deploy(ctx, cfg)
	if err != nil {
		panic(err)
	}
	cfg.Chain.Address = common.HexToAddress(address.Hex())
	log.Printf("Deployed contract at address: %s", address.Hex())

	wg := sync.WaitGroup{}

	// run contractSDK
	privateKey, err := crypto.HexToECDSA(cfg.Chain.KeyHex)
	if err != nil {
		panic(err)
	}
	contractSDK := sdk.NewContractSDK(ctx, cfg.Chain.HTTPURL, cfg.Chain.ID, cfg.Chain.Address, privateKey)
	wg.Add(1)
	go contractSDK.Run()

	// create storage
	storage := make(map[*big.Int]string)

	// run server
	transmitterServer := server.NewTransmitter(cfg.HTTP.Host, cfg.HTTP.Port, &wg, contractSDK, storage)
	wg.Add(1)
	go transmitterServer.StartServer()

	// create transmitter client
	transmitterClient := client.NewTransmitterClient(storage)

	// run watcher
	watcher, err := watcher.NewWatcher(ctx, cfg.Chain.HTTPURL, cfg.Chain.WSURL, cfg.Chain.Address, cfg.Chain.ID, transmitterClient)
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go watcher.Run()

	wg.Wait()
}
