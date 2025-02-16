package metrics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kimroniny/SuperRunner-eICN-eth2/logger"
	"github.com/sirupsen/logrus"
)

// CollectorClient is used to send metrics to the collector server
type CollectorClient struct {
	serverURL string
	client    *http.Client
	log       *logrus.Entry
}

// NewCollectorClient creates a new instance of CollectorClient
func NewCollectorClient(serverURL string) *CollectorClient {
	if logger.GetLogger() == nil {
		logger.InitLogger()
	}
	return &CollectorClient{
		serverURL: serverURL,
		client:    &http.Client{},
		log:       logger.NewComponent("CollectorClient"),
	}
}

// CollectMetricsEvent sends a metrics event to the collector server
func (cc *CollectorClient) CollectMetricsEvent(data MetricsData) error {
	return sendJSONRequest(cc.client, fmt.Sprintf("%s/metrics/collect", cc.serverURL), data)
}

// GetMetrics retrieves all metrics data from the collector server
func (cc *CollectorClient) GetMetrics() ([]MetricsData, error) {
	url := fmt.Sprintf("%s/metrics/get", cc.serverURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := cc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("server returned error: %s (status code: %d)", string(body), resp.StatusCode)
	}

	var metrics []MetricsData
	if err := json.NewDecoder(resp.Body).Decode(&metrics); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	cc.log.WithField("count", len(metrics)).Info("Retrieved metrics data from server")
	return metrics, nil
}

// Helper function for sending JSON requests
func sendJSONRequest(client *http.Client, url string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal request data: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server returned error: %s (status code: %d)", string(body), resp.StatusCode)
	}

	return nil
}
