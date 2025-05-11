package metrics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/kimroniny/SuperRunner-eICN-eth2/logger"
	"github.com/sirupsen/logrus"
)

// CollectorClient is used to send metrics to the collector server
type CollectorClient struct {
	serverURL string
	client    *http.Client
	log       *logrus.Entry
	wsConn    *websocket.Conn
	wsMutex   sync.Mutex
}

// NewCollectorClient creates a new instance of CollectorClient
func NewCollectorClient(serverURL string) *CollectorClient {
	if logger.GetLogger() == nil {
		logger.InitLogger("")
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

// SubscribeToMetrics subscribes to real-time metrics updates via websocket
// The provided callback function will be called whenever new metrics data is received
func (cc *CollectorClient) SubscribeToMetrics(callback func(MetricsData)) error {
	cc.wsMutex.Lock()
	defer cc.wsMutex.Unlock()

	// Close existing connection if any
	if cc.wsConn != nil {
		cc.wsConn.Close()
		cc.wsConn = nil
	}

	// Parse the server URL to create websocket URL
	serverURL := cc.serverURL
	if serverURL[:7] == "http://" {
		serverURL = "ws://" + serverURL[7:]
	} else if serverURL[:8] == "https://" {
		serverURL = "wss://" + serverURL[8:]
	} else {
		serverURL = "ws://" + serverURL
	}

	// Create websocket URL
	u := url.URL{Scheme: "", Host: "", Path: "/metrics/subscribe"}
	wsURL := fmt.Sprintf("%s%s", serverURL, u.Path)

	// Connect to the websocket server
	cc.log.Infof("Connecting to websocket: %s", wsURL)
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to websocket: %v", err)
	}
	cc.wsConn = conn

	// Start reading messages in a goroutine
	go func() {
		defer cc.closeWsConn()
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				cc.log.Errorf("Websocket read error: %v", err)
				return
			}

			// Parse the message
			var data MetricsData
			if err := json.Unmarshal(message, &data); err != nil {
				cc.log.Errorf("Failed to parse metrics data: %v", err)
				continue
			}

			// Call the callback function with the received data
			callback(data)
		}
	}()

	return nil
}

// UnsubscribeFromMetrics closes the websocket connection
func (cc *CollectorClient) UnsubscribeFromMetrics() {
	cc.closeWsConn()
}

// closeWsConn safely closes the websocket connection
func (cc *CollectorClient) closeWsConn() {
	cc.wsMutex.Lock()
	defer cc.wsMutex.Unlock()

	if cc.wsConn != nil {
		cc.wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		cc.wsConn.Close()
		cc.wsConn = nil
	}
}

// Helper function for sending JSON requests
func sendJSONRequest(client *http.Client, url string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
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
