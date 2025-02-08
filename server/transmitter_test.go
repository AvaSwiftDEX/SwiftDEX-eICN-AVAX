package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestCrossReceive(t *testing.T) {
	wg := &sync.WaitGroup{}
	transmitter := NewTransmitter("8080", wg)

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
	transmitter := NewTransmitter("8080", wg)

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
