package watcher

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/kimroniny/SuperRunner-eICN-eth2/client"
	"github.com/kimroniny/SuperRunner-eICN-eth2/metrics/metrics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 创建一个模拟的 TransmitterClient
type MockTransmitterClient struct {
	mock.Mock
}

func (m *MockTransmitterClient) SyncHeader(chainId *big.Int, number *big.Int, root common.Hash) error {
	args := m.Called(chainId, number, root)
	return args.Error(0)
}

func (m *MockTransmitterClient) CrossReceive(chainId *big.Int, data1 []byte, data2 []byte) error {
	args := m.Called(chainId, data1, data2)
	return args.Error(0)
}

func (m *MockTransmitterClient) RegisterEICN(url string, chainID *big.Int, key string) error {
	args := m.Called(url, chainID, key)
	return args.Error(0)
}

var _ client.ITransmitterClient = &MockTransmitterClient{}

func TestNewWatcher(t *testing.T) {
	// Start metrics collector server
	server := metrics.NewCollectorServer("localhost:8080")
	serverErrChan := make(chan error, 1)
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			serverErrChan <- err
		}
	}()
	time.Sleep(100 * time.Millisecond) // Wait for server to start

	// 准备测试数据
	ctx := context.Background()
	address := common.HexToAddress("0x1234567890")
	chainId := big.NewInt(1)
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient
	collectorURL := "http://localhost:8080"

	// 测试创建 Watcher，不使用真实的 URL
	watcher, err := NewWatcher(ctx, "", "", address, chainId, transmitterClient, nil, collectorURL)

	// 验证结果
	assert.NoError(t, err)
	assert.NotNil(t, watcher)
	assert.Equal(t, "", watcher.httpUrl)
	assert.Equal(t, "", watcher.wsUrl)
	assert.Equal(t, address, watcher.address)
	assert.Equal(t, chainId, watcher.chainId)
	assert.NotNil(t, watcher.headerCh)
	assert.NotNil(t, watcher.crossReceiveCh)
	assert.NotNil(t, watcher.metricsCh)
	assert.NotNil(t, watcher.collectorClient)
	assert.Nil(t, watcher.httpClient)
	assert.Nil(t, watcher.wsClient)

	// Stop the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Stop(ctx)

	// Check for server errors
	select {
	case err := <-serverErrChan:
		t.Errorf("Server error: %v", err)
	default:
	}
}

func TestWatcher_SendHeader(t *testing.T) {
	// Start metrics collector server
	server := metrics.NewCollectorServer("localhost:8080")
	serverErrChan := make(chan error, 1)
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			serverErrChan <- err
		}
	}()
	time.Sleep(100 * time.Millisecond) // Wait for server to start

	// 准备测试数据
	ctx, cancel := context.WithCancel(context.Background())
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient
	collectorURL := "http://localhost:8080"
	watcher, _ := NewWatcher(ctx, "", "", common.Address{}, big.NewInt(1), transmitterClient, nil, collectorURL)

	// 设置 mock 期望
	mockClient.On("SyncHeader", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// 启动 SendHeader
	go watcher.SendHeader()

	// 发送测试数据
	testHeader := &SyncHeaderData{
		Number: big.NewInt(100),
		Root:   common.HexToHash("0x1234"),
	}
	watcher.headerCh <- testHeader

	// 等待处理完成
	time.Sleep(100 * time.Millisecond)

	// 验证 mock 是否被调用
	mockClient.AssertCalled(t, "SyncHeader", mock.Anything, mock.Anything, mock.Anything)

	// 清理
	cancel()

	// Stop the server
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Stop(ctx)

	// Check for server errors
	select {
	case err := <-serverErrChan:
		t.Errorf("Server error: %v", err)
	default:
	}
}

func TestWatcher_CrossReceive(t *testing.T) {
	// Start metrics collector server
	server := metrics.NewCollectorServer("localhost:8080")
	serverErrChan := make(chan error, 1)
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			serverErrChan <- err
		}
	}()
	time.Sleep(100 * time.Millisecond) // Wait for server to start
	// 准备测试数据
	ctx, cancel := context.WithCancel(context.Background())
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient
	collectorURL := "http://localhost:8080"
	watcher, _ := NewWatcher(ctx, "", "", common.Address{}, big.NewInt(1), transmitterClient, nil, collectorURL)

	// 设置 mock 期望
	mockClient.On("CrossReceive", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// 启动 CrossReceive
	go watcher.CrossReceive()

	// 发送测试数据
	testData := &CrossReceiveData{
		chainId: big.NewInt(1),
		Data1:   []byte("test1"),
		Data2:   []byte("test2"),
	}
	watcher.crossReceiveCh <- testData

	// 等待处理完成
	time.Sleep(100 * time.Millisecond)

	// 验证 mock 是否被调用
	mockClient.AssertCalled(t, "CrossReceive", mock.Anything, mock.Anything, mock.Anything)

	// 清理
	cancel()

	// Stop the server
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Stop(ctx)

	// Check for server errors
	select {
	case err := <-serverErrChan:
		t.Errorf("Server error: %v", err)
	default:
	}
}

func TestWatcher_Metrics(t *testing.T) {
	// 启动 metrics collector server
	server := metrics.NewCollectorServer("localhost:8084")
	serverErrChan := make(chan error, 1)
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			serverErrChan <- err
		}
	}()
	time.Sleep(100 * time.Millisecond)

	// 准备测试数据
	ctx, cancel := context.WithCancel(context.Background())
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient
	collectorURL := "http://localhost:8084"
	watcher, _ := NewWatcher(ctx, "", "", common.Address{}, big.NewInt(1), transmitterClient, nil, collectorURL)

	// 启动 Metrics 处理
	go watcher.Metrics()

	// 发送测试数据
	testData := &metrics.MetricsData{
		TransactionHash: [32]byte{1, 2, 3},
		CmHash:          [32]byte{4, 5, 6},
		ChainId:         big.NewInt(1),
		Height:          big.NewInt(100),
		Phase:           1,
		IsConfirmed:     true,
		ByHeader:        false,
	}
	watcher.metricsCh <- testData

	// 等待处理完成
	time.Sleep(200 * time.Millisecond)

	// 验证数据是否被正确存储
	storedMetrics := server.GetMetrics()
	assert.Equal(t, 1, len(storedMetrics))
	assert.Equal(t, testData.TransactionHash, storedMetrics[0].TransactionHash)
	assert.Equal(t, testData.CmHash, storedMetrics[0].CmHash)
	assert.Equal(t, 0, testData.ChainId.Cmp(storedMetrics[0].ChainId))
	assert.Equal(t, 0, testData.Height.Cmp(storedMetrics[0].Height))
	assert.Equal(t, testData.Phase, storedMetrics[0].Phase)
	assert.Equal(t, testData.IsConfirmed, storedMetrics[0].IsConfirmed)
	assert.Equal(t, testData.ByHeader, storedMetrics[0].ByHeader)

	// 清理
	cancel()
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Stop(ctx)

	// 检查服务器是否有错误
	select {
	case err := <-serverErrChan:
		t.Errorf("Server error: %v", err)
	default:
	}
}

func TestWatcher_ContextCancellation(t *testing.T) {
	// Start metrics collector server
	server := metrics.NewCollectorServer("localhost:8080")
	serverErrChan := make(chan error, 1)
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			serverErrChan <- err
		}
	}()
	time.Sleep(100 * time.Millisecond) // Wait for server to start

	// 准备测试数据
	ctx, cancel := context.WithCancel(context.Background())
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient
	collectorURL := "http://localhost:8080"
	watcher, _ := NewWatcher(ctx, "", "", common.Address{}, big.NewInt(1), transmitterClient, nil, collectorURL)

	// 启动所有监听器
	done := make(chan bool)
	go func() {
		watcher.SendHeader()
		done <- true
	}()

	go func() {
		watcher.CrossReceive()
		done <- true
	}()

	go func() {
		watcher.Metrics()
		done <- true
	}()

	// 取消 context
	cancel()

	// 等待所有 goroutine 退出
	timeout := time.After(1 * time.Second)
	count := 0
	for count < 3 {
		select {
		case <-done:
			count++
		case <-timeout:
			t.Fatal("Timeout waiting for goroutines to exit")
		}
	}

	// Stop the server
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Stop(ctx)

	// Check for server errors
	select {
	case err := <-serverErrChan:
		t.Errorf("Server error: %v", err)
	default:
	}
}

func TestWatcher_SendHeader_Error(t *testing.T) {
	// Start metrics collector server
	server := metrics.NewCollectorServer("localhost:8080")
	serverErrChan := make(chan error, 1)
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			serverErrChan <- err
		}
	}()
	time.Sleep(100 * time.Millisecond) // Wait for server to start

	// 准备测试数据
	ctx, cancel := context.WithCancel(context.Background())
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient
	collectorURL := "http://localhost:8080"
	watcher, _ := NewWatcher(ctx, "", "", common.Address{}, big.NewInt(1), transmitterClient, nil, collectorURL)

	// 设置 mock 期望返回错误
	expectedErr := fmt.Errorf("sync header error")
	mockClient.On("SyncHeader", mock.Anything, mock.Anything, mock.Anything).Return(expectedErr)

	// 启动 SendHeader
	done := make(chan bool)
	go func() {
		watcher.SendHeader()
		done <- true
	}()

	// 发送测试数据
	testHeader := &SyncHeaderData{
		Number: big.NewInt(100),
		Root:   common.HexToHash("0x1234"),
	}
	watcher.headerCh <- testHeader

	// 等待处理完成
	time.Sleep(100 * time.Millisecond)

	// 验证 mock 是否被调用
	mockClient.AssertCalled(t, "SyncHeader", mock.Anything, mock.Anything, mock.Anything)

	// 清理
	cancel()
	<-done

	// Stop the server
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Stop(ctx)

	// Check for server errors
	select {
	case err := <-serverErrChan:
		t.Errorf("Server error: %v", err)
	default:
	}
}

func TestWatcher_CrossReceive_Error(t *testing.T) {
	// Start metrics collector server
	server := metrics.NewCollectorServer("localhost:8080")
	serverErrChan := make(chan error, 1)
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			serverErrChan <- err
		}
	}()
	time.Sleep(100 * time.Millisecond) // Wait for server to start

	// 准备测试数据
	ctx, cancel := context.WithCancel(context.Background())
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient
	collectorURL := "http://localhost:8080"
	watcher, _ := NewWatcher(ctx, "", "", common.Address{}, big.NewInt(1), transmitterClient, nil, collectorURL)

	// 设置 mock 期望返回错误
	expectedErr := fmt.Errorf("cross receive error")
	mockClient.On("CrossReceive", mock.Anything, mock.Anything, mock.Anything).Return(expectedErr)

	// 启动 CrossReceive
	done := make(chan bool)
	go func() {
		watcher.CrossReceive()
		done <- true
	}()

	// 发送测试数据
	testData := &CrossReceiveData{
		chainId: big.NewInt(1),
		Data1:   []byte("test1"),
		Data2:   []byte("test2"),
	}
	watcher.crossReceiveCh <- testData

	// 等待处理完成
	time.Sleep(100 * time.Millisecond)

	// 验证 mock 是否被调用
	mockClient.AssertCalled(t, "CrossReceive", mock.Anything, mock.Anything, mock.Anything)

	// 清理
	cancel()
	<-done

	// Stop the server
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Stop(ctx)

	// Check for server errors
	select {
	case err := <-serverErrChan:
		t.Errorf("Server error: %v", err)
	default:
	}
}

func TestWatcher_InvalidChainId(t *testing.T) {
	// Start metrics collector server
	server := metrics.NewCollectorServer("localhost:8080")
	serverErrChan := make(chan error, 1)
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			serverErrChan <- err
		}
	}()
	time.Sleep(100 * time.Millisecond) // Wait for server to start
	// 准备测试数据
	ctx := context.Background()
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient
	collectorURL := "http://localhost:8080"

	// 测试使用无效的 chainId
	watcher, err := NewWatcher(ctx, "", "", common.Address{}, nil, transmitterClient, nil, collectorURL)
	assert.Error(t, err)
	assert.Nil(t, watcher)

	// 测试使用负数 chainId
	watcher, err = NewWatcher(ctx, "", "", common.Address{}, big.NewInt(-1), transmitterClient, nil, collectorURL)
	assert.Error(t, err)
	assert.Nil(t, watcher)

	// Stop the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Stop(ctx)

	// Check for server errors
	select {
	case err := <-serverErrChan:
		t.Errorf("Server error: %v", err)
	default:
	}
}

func TestWatcher_InvalidCollectorURL(t *testing.T) {
	// 准备测试数据
	ctx := context.Background()
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient

	// 测试使用无效的 collector URL
	watcher, err := NewWatcher(ctx, "", "", common.Address{}, big.NewInt(1), transmitterClient, nil, "invalid-url")
	assert.Error(t, err)
	assert.Nil(t, watcher)

	// 测试使用空的 collector URL
	watcher, err = NewWatcher(ctx, "", "", common.Address{}, big.NewInt(1), transmitterClient, nil, "")
	assert.Error(t, err)
	assert.Nil(t, watcher)
}
