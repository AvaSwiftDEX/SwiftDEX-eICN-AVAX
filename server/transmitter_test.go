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
		URL:     "http://example.com",
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

func TestSyncHeader(t *testing.T) {
	wg := &sync.WaitGroup{}
	contractSDK := sdk.NewContractSDK()
	transmitter := NewTransmitter("8080", wg, contractSDK)

	requestBody, _ := json.Marshal(RequestHeader{
		ChainID: 1,
		Number:  uint64(123),
		Root:    [32]byte{1, 2, 3}, // 示例root值
	})

	r := httptest.NewRequest(http.MethodPost, "/SyncHeader", bytes.NewReader(requestBody))
	w := httptest.NewRecorder()
	transmitter.SyncHeader(w, r)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var response ResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Errorf("Failed to decode response: %v", err)
	}

	if !response.Success {
		t.Error("Expected success to be true")
	}
}

func TestInvalidMethod(t *testing.T) {
	wg := &sync.WaitGroup{}
	contractSDK := sdk.NewContractSDK()
	transmitter := NewTransmitter("8080", wg, contractSDK)

	// 测试 SyncHeader 的 GET 请求
	r := httptest.NewRequest(http.MethodGet, "/SyncHeader", nil)
	w := httptest.NewRecorder()
	transmitter.SyncHeader(w, r)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d for GET request, got %d", http.StatusMethodNotAllowed, w.Code)
	}

	// 测试 CrossReceive 的 GET 请求
	r = httptest.NewRequest(http.MethodGet, "/CrossReceive", nil)
	w = httptest.NewRecorder()
	transmitter.CrossReceive(w, r)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d for GET request, got %d", http.StatusMethodNotAllowed, w.Code)
	}

	// 测试 RegisterEICN 的 GET 请求
	r = httptest.NewRequest(http.MethodGet, "/registerEICN", nil)
	w = httptest.NewRecorder()
	transmitter.RegisterEICN(w, r)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d for GET request, got %d", http.StatusMethodNotAllowed, w.Code)
	}
}

func TestInvalidJSON(t *testing.T) {
	wg := &sync.WaitGroup{}
	contractSDK := sdk.NewContractSDK()
	transmitter := NewTransmitter("8080", wg, contractSDK)

	invalidJSON := []byte(`{"invalid json`)

	// 测试 SyncHeader 的无效 JSON
	r := httptest.NewRequest(http.MethodPost, "/SyncHeader", bytes.NewReader(invalidJSON))
	w := httptest.NewRecorder()
	transmitter.SyncHeader(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d for invalid JSON, got %d", http.StatusBadRequest, w.Code)
	}

	// 测试 CrossReceive 的无效 JSON
	r = httptest.NewRequest(http.MethodPost, "/CrossReceive", bytes.NewReader(invalidJSON))
	w = httptest.NewRecorder()
	transmitter.CrossReceive(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d for invalid JSON, got %d", http.StatusBadRequest, w.Code)
	}

	// 测试 RegisterEICN 的无效 JSON
	r = httptest.NewRequest(http.MethodPost, "/registerEICN", bytes.NewReader(invalidJSON))
	w = httptest.NewRecorder()
	transmitter.RegisterEICN(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d for invalid JSON, got %d", http.StatusBadRequest, w.Code)
	}
}
