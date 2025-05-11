package metrics

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/kimroniny/SuperRunner-eICN-eth2/logger"
	"github.com/sirupsen/logrus"
)

// Client represents a connected websocket client
type Client struct {
	conn      *websocket.Conn
	send      chan MetricsData
	closeChan chan struct{}
}

// CollectorServer handles the collection and storage of metrics data
type CollectorServer struct {
	storage    *Storage
	log        *logrus.Entry
	server     *http.Server
	address    string
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan MetricsData
	clientsMu  sync.Mutex
	upgrader   websocket.Upgrader
}

// NewCollectorServer creates a new instance of CollectorServer
func NewCollectorServer(address string) *CollectorServer {
	if logger.GetLogger() == nil {
		logger.InitLogger("")
	}
	cs := &CollectorServer{
		storage:    NewStorage(),
		log:        logger.NewComponent("CollectorServer"),
		address:    address,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan MetricsData, 256),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // Allow all origins
			},
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/metrics/collect", cs.handleCollectMetrics)
	mux.HandleFunc("/metrics/get", cs.handleGetMetrics)
	mux.HandleFunc("/metrics/subscribe", cs.handleWebSocketSubscribe)

	cs.server = &http.Server{
		Addr:    address,
		Handler: mux,
	}

	return cs
}

// Start starts the collector server
func (cs *CollectorServer) Start() error {
	go cs.runClientManager()

	listener, err := net.Listen("tcp", cs.address)
	if err != nil {
		return fmt.Errorf("failed to start listener: %v", err)
	}

	cs.log.Infof("CollectorServer started on %s", cs.address)
	return cs.server.Serve(listener)
}

// Stop stops the collector server
func (cs *CollectorServer) Stop(ctx context.Context) error {
	cs.clientsMu.Lock()
	for client := range cs.clients {
		close(client.send)
		client.conn.Close()
	}
	cs.clientsMu.Unlock()

	return cs.server.Shutdown(ctx)
}

// runClientManager manages websocket client connections
func (cs *CollectorServer) runClientManager() {
	for {
		select {
		case client := <-cs.register:
			cs.clientsMu.Lock()
			cs.clients[client] = true
			cs.clientsMu.Unlock()
			cs.log.Infof("Client connected. Total connected clients: %d", len(cs.clients))

		case client := <-cs.unregister:
			cs.clientsMu.Lock()
			if _, ok := cs.clients[client]; ok {
				delete(cs.clients, client)
				close(client.send)
			}
			cs.clientsMu.Unlock()
			cs.log.Infof("Client disconnected. Total connected clients: %d", len(cs.clients))

		case data := <-cs.broadcast:
			cs.broadcastToClients(data)
		}
	}
}

// broadcastToClients sends metrics data to all connected clients
func (cs *CollectorServer) broadcastToClients(data MetricsData) {
	cs.clientsMu.Lock()
	defer cs.clientsMu.Unlock()

	for client := range cs.clients {
		select {
		case client.send <- data:
		default:
			close(client.send)
			delete(cs.clients, client)
		}
	}
}

// handleCollectMetrics handles POST requests for collecting metrics data
func (cs *CollectorServer) handleCollectMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data MetricsData
	if err := decodeJSONBody(w, r, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cs.storage.Store(data)

	// Broadcast the new metrics data to all connected clients
	cs.broadcast <- data

	cs.log.WithFields(logrus.Fields{
		"transactionHash": hex.EncodeToString(data.TransactionHash[:]),
		"cmHash":          hex.EncodeToString(data.CmHash[:]),
		"chainId":         data.ChainId.String(),
		"height":          data.Height.String(),
		"phase":           data.Phase,
		"isConfirmed":     data.IsConfirmed,
		"byHeader":        data.ByHeader,
		"timestamp":       data.Timestamp,
		"txHash":          data.TxHash.String(),
	}).Info("Collected new metrics event")

	w.WriteHeader(http.StatusOK)
}

// handleGetMetrics handles GET requests for retrieving metrics data
func (cs *CollectorServer) handleGetMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	metrics := cs.GetMetrics()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		cs.log.Errorf("Failed to encode metrics data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	cs.log.WithField("count", len(metrics)).Info("Received metrics data")
}

// handleWebSocketSubscribe handles websocket connections for real-time metrics updates
func (cs *CollectorServer) handleWebSocketSubscribe(w http.ResponseWriter, r *http.Request) {
	conn, err := cs.upgrader.Upgrade(w, r, nil)
	if err != nil {
		cs.log.Errorf("Failed to upgrade connection to WebSocket: %v", err)
		return
	}

	client := &Client{
		conn:      conn,
		send:      make(chan MetricsData, 256),
		closeChan: make(chan struct{}),
	}

	// Register the new client
	cs.register <- client

	// Start the client handlers
	go cs.handleClientWrite(client)
	go cs.handleClientRead(client)
}

// handleClientWrite writes messages to the websocket connection
func (cs *CollectorServer) handleClientWrite(client *Client) {
	defer func() {
		client.conn.Close()
	}()

	for {
		select {
		case data, ok := <-client.send:
			if !ok {
				// Channel was closed
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			// Marshal the metrics data to JSON
			jsonData, err := json.Marshal(data)
			if err != nil {
				cs.log.Errorf("Failed to marshal metrics data: %v", err)
				return
			}

			// Write the JSON data to the websocket
			err = client.conn.WriteMessage(websocket.TextMessage, jsonData)
			if err != nil {
				cs.log.Errorf("Failed to write to websocket: %v", err)
				return
			}

		case <-client.closeChan:
			return
		}
	}
}

// handleClientRead reads messages from the websocket connection
func (cs *CollectorServer) handleClientRead(client *Client) {
	defer func() {
		cs.unregister <- client
		client.conn.Close()
	}()

	for {
		_, _, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				cs.log.Errorf("WebSocket read error: %v", err)
			}
			close(client.closeChan)
			break
		}
		// We don't expect any messages from clients, just handle the connection
	}
}

// GetMetrics returns all collected metrics
func (cs *CollectorServer) GetMetrics() []MetricsData {
	return cs.storage.GetAll()
}

// Helper function for JSON decoding
func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("Content-Type header is not application/json")
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&dst); err != nil {
		return fmt.Errorf("failed to decode request body: %v", err)
	}

	return nil
}
