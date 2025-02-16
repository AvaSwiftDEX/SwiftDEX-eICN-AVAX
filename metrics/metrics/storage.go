package metrics

import (
	"bytes"
	"math/big"
	"sync"
)

// Storage handles the storage and retrieval of MetricsData
type Storage struct {
	mu      sync.RWMutex
	metrics []MetricsData
}

// NewStorage creates a new instance of Storage
func NewStorage() *Storage {
	return &Storage{
		metrics: make([]MetricsData, 0),
	}
}

// Store adds a new MetricsData to the storage
func (s *Storage) Store(data MetricsData) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.metrics = append(s.metrics, data)
}

// GetAll returns all stored MetricsData
func (s *Storage) GetAll() []MetricsData {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]MetricsData, len(s.metrics))
	copy(result, s.metrics)
	return result
}

// GetByTransactionHash returns MetricsData for a specific transaction hash
func (s *Storage) GetByTransactionHash(hash [32]byte) []MetricsData {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []MetricsData
	for _, m := range s.metrics {
		if bytes.Equal(m.TransactionHash[:], hash[:]) {
			result = append(result, m)
		}
	}
	return result
}

// GetByChainId returns MetricsData for a specific chain ID
func (s *Storage) GetByChainId(chainId *big.Int) []MetricsData {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []MetricsData
	for _, m := range s.metrics {
		if m.ChainId.Cmp(chainId) == 0 {
			result = append(result, m)
		}
	}
	return result
}

// GetByPhase returns MetricsData for a specific phase
func (s *Storage) GetByPhase(phase uint8) []MetricsData {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []MetricsData
	for _, m := range s.metrics {
		if m.Phase == phase {
			result = append(result, m)
		}
	}
	return result
}

// Clear removes all stored metrics data
func (s *Storage) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.metrics = make([]MetricsData, 0)
}
