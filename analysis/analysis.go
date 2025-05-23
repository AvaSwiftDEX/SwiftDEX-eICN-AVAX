package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
	"github.com/kimroniny/SuperRunner-eICN-eth2/logger"
	"github.com/kimroniny/SuperRunner-eICN-eth2/metrics/metrics"
)

func main() {
	// 定义命令行参数
	serverURL := flag.String("metrics-server-url", "127.0.0.1:8090", "metrics collector server websocket url")
	logfile := flag.String("logfile", "logs/analysis.log", "log file path")
	totalNumber := flag.Int("total-number", 0, "total number of transactions")
	identifier := flag.String("identifier", "", "identifier of the analysis")
	configfiles := flag.String("config-files", "config.yaml", "config files, separated by comma")
	extract := flag.Bool("extract", false, "only extract observation")
	flag.Parse()

	// 初始化日志
	if logger.GetLogger() == nil {
		logger.InitLogger(*logfile)
	}
	log := logger.GetLogger()

	// Set up a channel to handle clean shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Create a new collector client
	client := metrics.NewCollectorClient(*serverURL)

	// Create a new analyzer
	analyzer := NewAnalyzer(*totalNumber)

	// all config files
	configs := make([]*config.Config, 0)
	configs_filename := strings.Split(*configfiles, ",")
	for _, config_filename := range configs_filename {
		cfg, err := config.LoadConfig(config_filename)
		if err != nil {
			log.Errorf("failed to load config: %v", err)
			os.Exit(1)
		}
		configs = append(configs, cfg)
	}

	temp_filename := fmt.Sprintf("temp/%s.json", *identifier)
	if *extract {
		fmt.Println("extract observation from file: ", temp_filename)
		observation, err := ExtractObservationFromFile(temp_filename, configs)
		if err != nil {
			log.Errorf("Error extracting observation: %v\n", err)
			return
		}
		os.WriteFile(fmt.Sprintf("observation/%s.json", *identifier), observation, 0644)
		return
	}

	// Create a new routine to check if all transactions are finished
	stopped := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-stopped:
				return
			default:
				time.Sleep(time.Second * 3)
				log.Infof("Finished number: %d, Total number: %d\n", analyzer.finishedNumber, *totalNumber)
				if analyzer.finishedNumber == *totalNumber {
					log.Infof("All transactions are finished\n")
					stopped <- true
					return
				}
			}
		}
	}()

	// Subscribe to metrics updates
	err := client.SubscribeToMetrics(analyzer.AnalysisMetrics)

	if err != nil {
		log.Errorf("Error subscribing to metrics: %v\n", err)
		os.Exit(1)
	}

	log.Infof("Subscribed to metrics from %s\n", *serverURL)
	log.Infof("Press Ctrl+C to exit")

	// Wait for interrupt signal
	select {
	case <-interrupt:
		stopped <- true
		log.Infof("Interrupted! Shutting down...")
	case <-stopped:
		log.Infof("Finished! Shutting down...")
	}

	// Close the websocket connection
	client.UnsubscribeFromMetrics()
	time.Sleep(time.Second) // Give it time to close properly

	// write analyzer.transHashStorage to temp file
	temp_data, err := json.Marshal(analyzer.transHashStorage)
	if err != nil {
		log.Errorf("Error marshalling transHashStorage: %v\n", err)
		return
	}
	os.WriteFile(temp_filename, temp_data, 0644)

	observation, err := ExtractObservation(analyzer.transHashStorage, configs)
	if err != nil {
		log.Errorf("Error extracting observation: %v\n", err)
		return
	}
	os.WriteFile(fmt.Sprintf("observation/%s.json", *identifier), observation, 0644)
}
