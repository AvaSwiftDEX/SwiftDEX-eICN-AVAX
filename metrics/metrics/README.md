# Metrics Collector Server

The Metrics Collector Server is a service for collecting, storing, and distributing metrics data for the SuperRunner system.

## Features

- HTTP API for collecting and retrieving metrics data
- Real-time metrics updates via WebSocket
- Client library for easy integration

## HTTP API

### Collect Metrics

```
POST /metrics/collect
Content-Type: application/json

{
  "TransactionHash": "0x...",
  "CmHash": "0x...",
  "ChainId": "1",
  "Height": "100",
  "Phase": 1,
  "IsConfirmed": true,
  "ByHeader": false
}
```

### Get Metrics

```
GET /metrics/get
```

Returns all collected metrics data as a JSON array.

## WebSocket API

### Subscribe to Metrics

Connect to `/metrics/subscribe` using WebSocket to receive real-time metrics updates.

Example:
```
ws://localhost:8080/metrics/subscribe
```

Each time a new metrics event is collected, it will be sent to all connected websocket clients.

## Using the Client Library

The `CollectorClient` provides a convenient way to interact with the server.

### Basic Usage

```go
// Create a client
client := metrics.NewCollectorClient("http://localhost:8080")

// Send metrics
data := metrics.MetricsData{...}
err := client.CollectMetricsEvent(data)

// Get all metrics
metrics, err := client.GetMetrics()
```

### Subscribe to Real-time Updates

```go
// Create a client
client := metrics.NewCollectorClient("http://localhost:8080")

// Subscribe to metrics updates
err := client.SubscribeToMetrics(func(data metrics.MetricsData) {
    // Handle received metrics data
    fmt.Printf("Received metrics: %+v\n", data)
})

// Unsubscribe when done
client.UnsubscribeFromMetrics()
```

## Examples

See the `examples` directory for complete working examples:

- `client_wrapper` - Example of using the client library to subscribe to metrics
- `websocket_client` - Direct WebSocket client example without using the client library 