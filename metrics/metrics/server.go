package metrics

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/kimroniny/SuperRunner-eICN-eth2/logger"
	"github.com/sirupsen/logrus"
)

// CollectorServer handles the collection and storage of metrics data
type CollectorServer struct {
	storage *Storage
	log     *logrus.Entry
	server  *http.Server
	address string
}

// NewCollectorServer creates a new instance of CollectorServer
func NewCollectorServer(address string) *CollectorServer {
	if logger.GetLogger() == nil {
		logger.InitLogger()
	}
	cs := &CollectorServer{
		storage: NewStorage(),
		log:     logger.NewComponent("CollectorServer"),
		address: address,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/metrics/collect", cs.handleCollectMetrics)
	mux.HandleFunc("/metrics/get", cs.handleGetMetrics)

	cs.server = &http.Server{
		Addr:    address,
		Handler: mux,
	}

	return cs
}

// Start starts the collector server
func (cs *CollectorServer) Start() error {
	listener, err := net.Listen("tcp", cs.address)
	if err != nil {
		return fmt.Errorf("failed to start listener: %v", err)
	}

	cs.log.Infof("CollectorServer started on %s", cs.address)
	return cs.server.Serve(listener)
}

// Stop stops the collector server
func (cs *CollectorServer) Stop(ctx context.Context) error {
	return cs.server.Shutdown(ctx)
}

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

	cs.log.WithFields(logrus.Fields{
		"transactionHash": hex.EncodeToString(data.TransactionHash[:]),
		"cmHash":          hex.EncodeToString(data.CmHash[:]),
		"chainId":         data.ChainId.String(),
		"height":          data.Height.String(),
		"phase":           data.Phase,
		"isConfirmed":     data.IsConfirmed,
		"byHeader":        data.ByHeader,
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

	cs.log.WithField("count", len(metrics)).Info("Retrieved metrics data")
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
