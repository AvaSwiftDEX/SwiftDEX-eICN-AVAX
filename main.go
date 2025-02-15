package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

	// 创建一个带取消功能的上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 设置信号处理
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

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
	go func() {
		defer wg.Done()
		contractSDK.Run()
	}()

	// create storage
	// use string as key, *big.Int with the same value is different key
	storage := make(map[string]string)

	// run server
	transmitterServer := server.NewTransmitter(cfg.HTTP.Host, cfg.HTTP.Port, &wg, contractSDK, storage)
	wg.Add(1)
	go func() {
		defer wg.Done()
		transmitterServer.StartServer()
	}()

	// create transmitter client
	transmitterClient := client.NewTransmitterClient(storage)

	// run watcher
	watcher, err := watcher.NewWatcher(ctx, cfg.Chain.HTTPURL, cfg.Chain.WSURL, cfg.Chain.Address, cfg.Chain.ID, transmitterClient)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		watcher.Run()
	}()

	// 等待中断信号
	sig := <-sigChan
	log.WithField("signal", sig.String()).Info("收到退出信号，开始优雅关闭...")

	// 取消上下文
	cancel()

	// 等待所有goroutine完成
	wg.Wait()
	log.Info("所有服务已停止，程序退出")
}
