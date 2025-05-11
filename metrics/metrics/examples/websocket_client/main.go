package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kimroniny/SuperRunner-eICN-eth2/metrics/metrics"
)

func main() {
	// Parse command line arguments
	serverAddr := flag.String("server", "localhost:8080", "Server address to connect to")
	flag.Parse()

	// Set up a channel to handle clean shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Create the websocket URL
	u := url.URL{Scheme: "ws", Host: *serverAddr, Path: "/metrics/subscribe"}
	log.Printf("Connecting to %s", u.String())

	// Connect to the server
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer c.Close()

	// Channel to receive metrics
	metricsChannel := make(chan metrics.MetricsData)

	// Start a goroutine to read messages from the server
	go func() {
		defer close(metricsChannel)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}

			// Parse the message into a MetricsData struct
			var data metrics.MetricsData
			if err := json.Unmarshal(message, &data); err != nil {
				log.Println("Failed to parse metrics data:", err)
				continue
			}

			metricsChannel <- data
		}
	}()

	// Main loop to process received metrics and handle termination
	for {
		select {
		case data, ok := <-metricsChannel:
			if !ok {
				return
			}
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

		case <-interrupt:
			log.Println("Interrupted, closing connection...")

			// Cleanly close the connection by sending a close message
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Write close error:", err)
				return
			}

			// Wait for the server to close the connection
			select {
			case <-metricsChannel:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
