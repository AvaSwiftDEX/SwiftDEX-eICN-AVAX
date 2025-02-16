package metrics

import (
	"context"
	"math/big"
	"net/http"
	"testing"
	"time"
)

func TestCollectorServer(t *testing.T) {
	// 创建服务器实例
	server := NewCollectorServer("localhost:8080")
	errChan := make(chan error, 1)

	// 在后台启动服务器
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			errChan <- err
			t.Errorf("Failed to start server: %v", err)
		}
	}()

	// 等待服务器启动
	time.Sleep(100 * time.Millisecond)

	// 创建测试数据
	testData := MetricsData{
		TransactionHash: [32]byte{1, 2, 3},
		CmHash:          [32]byte{4, 5, 6},
		ChainId:         big.NewInt(1),
		Height:          big.NewInt(100),
		Phase:           1,
		IsConfirmed:     true,
		ByHeader:        false,
	}

	// 存储测试数据
	server.storage.Store(testData)

	// 测试获取数据
	metrics := server.GetMetrics()
	if len(metrics) != 1 {
		t.Errorf("Expected 1 metric, got %d", len(metrics))
	}

	// 验证获取的数据
	if metrics[0].ChainId.Cmp(testData.ChainId) != 0 {
		t.Errorf("Expected ChainId %v, got %v", testData.ChainId, metrics[0].ChainId)
	}
	if metrics[0].Height.Cmp(testData.Height) != 0 {
		t.Errorf("Expected Height %v, got %v", testData.Height, metrics[0].Height)
	}
	if metrics[0].Phase != testData.Phase {
		t.Errorf("Expected Phase %v, got %v", testData.Phase, metrics[0].Phase)
	}

	// 关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Stop(ctx); err != nil {
		t.Errorf("Failed to stop server: %v", err)
	}

	// 检查服务器是否有错误
	select {
	case err := <-errChan:
		t.Errorf("Server error: %v", err)
	default:
	}
}

func TestCollectorServerConcurrency(t *testing.T) {
	server := NewCollectorServer("localhost:8081")
	errChan := make(chan error, 1)

	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			errChan <- err
			t.Errorf("Failed to start server: %v", err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	// 并发测试
	const numGoroutines = 10
	done := make(chan bool)

	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			testData := MetricsData{
				TransactionHash: [32]byte{byte(index)},
				CmHash:          [32]byte{byte(index)},
				ChainId:         big.NewInt(int64(index)),
				Height:          big.NewInt(100),
				Phase:           uint8(index),
				IsConfirmed:     true,
				ByHeader:        false,
			}
			server.storage.Store(testData)
			done <- true
		}(i)
	}

	// 等待所有goroutine完成
	for i := 0; i < numGoroutines; i++ {
		<-done
	}

	// 验证数据
	metrics := server.GetMetrics()
	if len(metrics) != numGoroutines {
		t.Errorf("Expected %d metrics, got %d", numGoroutines, len(metrics))
	}

	// 关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Stop(ctx); err != nil {
		t.Errorf("Failed to stop server: %v", err)
	}

	// 检查服务器是否有错误
	select {
	case err := <-errChan:
		t.Errorf("Server error: %v", err)
	default:
	}
}
