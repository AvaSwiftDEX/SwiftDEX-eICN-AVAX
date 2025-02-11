package main

import (
	"context"
	"flag"
	"sync"

	"github.com/kimroniny/SuperRunner-eICN-eth2/client"
	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
	"github.com/kimroniny/SuperRunner-eICN-eth2/logger"
	sdk "github.com/kimroniny/SuperRunner-eICN-eth2/sdk"
	"github.com/kimroniny/SuperRunner-eICN-eth2/server"
	"github.com/kimroniny/SuperRunner-eICN-eth2/watcher"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化日志
	logger.InitLogger()
	log := logger.GetLogger()
	// log.SetReportCaller(true)

	// 定义命令行参数
	configPath := flag.String("config", "config.yaml", "path to config file")
	debugLevel := flag.Bool("debug", true, "debug level")
	flag.Parse()

	if *debugLevel {
		log.SetLevel(logrus.DebugLevel)
	}

	ctx := context.Background()

	// 使用命令行参数中指定的配置文件路径
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	// address, err := scripts.Deploy(ctx, cfg)
	// if err != nil {
	// 	panic(err)
	// }
	// cfg.Chain.Address = common.HexToAddress(address.Hex())
	// log.Printf("Deployed contract at address: %s", address.Hex())

	wg := sync.WaitGroup{}

	// run contractSDK
	privateKey, err := cfg.ReadPrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	contractSDK := sdk.NewContractSDK(ctx, cfg.Chain.HTTPURL, cfg.Chain.ID, cfg.Chain.Address, privateKey)
	wg.Add(1)
	go contractSDK.Run()

	// create storage
	// use string as key, *big.Int with the same value is different key
	storage := make(map[string]string)

	// run server
	transmitterServer := server.NewTransmitter(cfg.HTTP.Host, cfg.HTTP.Port, &wg, contractSDK, storage)
	wg.Add(1)
	go transmitterServer.StartServer()

	// create transmitter client
	transmitterClient := client.NewTransmitterClient(storage)

	// run watcher
	watcher, err := watcher.NewWatcher(ctx, cfg.Chain.HTTPURL, cfg.Chain.WSURL, cfg.Chain.Address, cfg.Chain.ID, transmitterClient)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go watcher.Run()

	wg.Wait()
}
