package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"github.com/kimroniny/SuperRunner-eICN-eth2/sdk"
)

func TestCrossReceive(t *testing.T) {
	wg := &sync.WaitGroup{}
	contractSDK := sdk.NewContractSDK()
	transmitter := NewTransmitter("8080", wg, contractSDK)

	requestBody, _ := json.Marshal(RequestBody{
		Data1: []byte("test1"),
		Data2: []byte("test2"),
	})

	r := httptest.NewRequest(http.MethodPost, "/CrossReceive", bytes.NewReader(requestBody))
	w := httptest.NewRecorder()
	transmitter.CrossReceive(w, r)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestRegisterEICN(t *testing.T) {
	wg := &sync.WaitGroup{}
	contractSDK := sdk.NewContractSDK()
	transmitter := NewTransmitter("8080", wg, contractSDK)

	requestBody, _ := json.Marshal(RegisterRequest{
		URL:    "http://example.com",
		ChainID: 123,
	})

	r := httptest.NewRequest(http.MethodPost, "/registerEICN", bytes.NewReader(requestBody))
	w := httptest.NewRecorder()
	transmitter.RegisterEICN(w, r)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestTransmitterCrossReceiveIntegration(t *testing.T) {
	wg := &sync.WaitGroup{}
	contractSDK := sdk.NewContractSDK()
	transmitter := NewTransmitter("8080", wg, contractSDK)

	go contractSDK.Run()

	requestBody, _ := json.Marshal(RequestBody{
		Data1: []byte("test_cm"),
		Data2: []byte("test_proof"),
	})

	r := httptest.NewRequest(http.MethodPost, "/CrossReceive", bytes.NewReader(requestBody))
	w := httptest.NewRecorder()
	transmitter.CrossReceive(w, r)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	contractSDK.Stop <- struct{}{}
}
