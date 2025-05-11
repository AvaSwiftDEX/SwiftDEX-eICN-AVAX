
package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/kimroniny/SuperRunner-eICN-eth2/metrics/metrics"
)

// RunExample starts the metrics client example
func main() {
	// Parse command line arguments
	serverURL := flag.String("server", "http://localhost:8080", "Server URL to connect to")
	flag.Parse()

	// Create a new collector client
	client := metrics.NewCollectorClient(*serverURL)

	// Set up a channel to handle clean shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Subscribe to metrics updates
	err := client.SubscribeToMetrics(func(data metrics.MetricsData) {
		fmt.Printf("Received metrics event:\n")
		fmt.Printf("  Transaction Hash: %s\n", hex.EncodeToString(data.TransactionHash[:]))
		fmt.Printf("  CM Hash: %s\n", hex.EncodeToString(data.CmHash[:]))
		fmt.Printf("  Chain ID: %s\n", data.ChainId.String())
		fmt.Printf("  Height: %s\n", data.Height.String())
		fmt.Printf("  Phase: %d\n", data.Phase)
		fmt.Printf("  Is Confirmed: %t\n", data.IsConfirmed)
		fmt.Printf("  By Header: %t\n", data.ByHeader)
		fmt.Printf("  Timestamp: %d\n", data.Timestamp)
		fmt.Println("----------------------------")
	})

	if err != nil {
		fmt.Printf("Error subscribing to metrics: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Subscribed to metrics from %s\n", *serverURL)
	fmt.Println("Press Ctrl+C to exit")

	// Wait for interrupt signal
	<-interrupt
	fmt.Println("Shutting down...")

	// Close the websocket connection
	client.UnsubscribeFromMetrics()
	time.Sleep(time.Second) // Give it time to close properly
}
