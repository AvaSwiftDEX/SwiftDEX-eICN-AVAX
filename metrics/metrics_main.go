package main

import (
	"flag"
	"sync"

	"github.com/kimroniny/SuperRunner-eICN-eth2/logger"
	"github.com/kimroniny/SuperRunner-eICN-eth2/metrics/metrics"
)

func main() {
	// 定义命令行参数
	url := flag.String("url", "127.0.0.1:8090", "metrics collector server url")
	flag.Parse()

	// 初始化日志
	if logger.GetLogger() == nil {
		logger.InitLogger("")
	}
	log := logger.GetLogger()
	wg := sync.WaitGroup{}
	// create metrics collector server
	metricsCollectorServer := metrics.NewCollectorServer(*url)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := metricsCollectorServer.Start(); err != nil {
			log.Error(err)
			return
		}
	}()

	wg.Wait()
}
