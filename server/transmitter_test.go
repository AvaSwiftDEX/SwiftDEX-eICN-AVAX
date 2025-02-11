package server

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kimroniny/SuperRunner-eICN-eth2/sdk"
)

// 创建测试所需的通用参数
func setupTestSDK(t *testing.T) *sdk.ContractSDK {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("无法生成私钥: %v", err)
	}

	ctx := context.Background()
	url := "http://localhost:8545"                          // 测试用URL
	chainId := big.NewInt(1)                                // 测试用chainId
	address := "0x742d35Cc6634C0532925a3b844Bc454e4438f44e" // 测试用合约地址

	return sdk.NewContractSDK(ctx, url, chainId, common.HexToAddress(address), privateKey)
}

func TestCrossReceive(t *testing.T) {
	wg := &sync.WaitGroup{}
	contractSDK := setupTestSDK(t)
	storage := make(map[string]string)
	transmitter := NewTransmitter("127.0.0.1", 8080, wg, contractSDK, storage)

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
	contractSDK := setupTestSDK(t)
	storage := make(map[string]string)
	transmitter := NewTransmitter("127.0.0.1", 8080, wg, contractSDK, storage)

	requestBody, _ := json.Marshal(RegisterRequest{
		URL:     "http://example.com",
		ChainID: big.NewInt(123),
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
	contractSDK := setupTestSDK(t)
	storage := make(map[string]string)
	transmitter := NewTransmitter("127.0.0.1", 8080, wg, contractSDK, storage)

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
	contractSDK := setupTestSDK(t)
	storage := make(map[string]string)
	transmitter := NewTransmitter("127.0.0.1", 8080, wg, contractSDK, storage)

	requestBody, _ := json.Marshal(RequestHeader{
		ChainID: big.NewInt(1),
		Number:  big.NewInt(123),
		Root:    common.HexToHash("0x1234567890123456789012345678901234567890123456789012345678901234"),
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
	contractSDK := setupTestSDK(t)
	storage := make(map[string]string)
	transmitter := NewTransmitter("127.0.0.1", 8080, wg, contractSDK, storage)

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
	contractSDK := setupTestSDK(t)
	storage := make(map[string]string)
	transmitter := NewTransmitter("127.0.0.1", 8080, wg, contractSDK, storage)

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
