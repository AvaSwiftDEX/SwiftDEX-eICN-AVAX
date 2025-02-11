package watcher

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/kimroniny/SuperRunner-eICN-eth2/client"
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
	// 准备测试数据
	ctx := context.Background()
	httpUrl := "http://localhost:8545"
	wsUrl := "ws://localhost:8546"
	address := common.HexToAddress("0x1234567890")
	chainId := big.NewInt(1)
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient
	// 测试创建 Watcher
	watcher, err := NewWatcher(ctx, httpUrl, wsUrl, address, chainId, transmitterClient)

	// 验证结果
	assert.NoError(t, err)
	assert.NotNil(t, watcher)
	assert.Equal(t, httpUrl, watcher.httpUrl)
	assert.Equal(t, wsUrl, watcher.wsUrl)
	assert.Equal(t, address, watcher.address)
	assert.Equal(t, chainId, watcher.chainId)
	assert.NotNil(t, watcher.headerCh)
	assert.NotNil(t, watcher.crossReceiveCh)
}

func TestWatcher_SendHeader(t *testing.T) {
	// 准备测试数据
	ctx, cancel := context.WithCancel(context.Background())
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient
	watcher, _ := NewWatcher(ctx, "", "", common.Address{}, big.NewInt(1), transmitterClient)

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
}

func TestWatcher_CrossReceive(t *testing.T) {
	// 准备测试数据
	ctx, cancel := context.WithCancel(context.Background())
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient
	watcher, _ := NewWatcher(ctx, "", "", common.Address{}, big.NewInt(1), transmitterClient)

	// 设置 mock 期望
	mockClient.On("CrossReceive", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// 启动 CrossReceive
	go watcher.CrossReceive()

	// 发送测试数据
	testData := &CrossReceiveData{
		Data1: []byte("test1"),
		Data2: []byte("test2"),
	}
	watcher.crossReceiveCh <- testData

	// 等待处理完成
	time.Sleep(100 * time.Millisecond)

	// 验证 mock 是否被调用
	mockClient.AssertCalled(t, "CrossReceive", mock.Anything, mock.Anything, mock.Anything)

	// 清理
	cancel()
}

func TestWatcher_ContextCancellation(t *testing.T) {
	// 准备测试数据
	ctx, cancel := context.WithCancel(context.Background())
	mockClient := new(MockTransmitterClient)
	var transmitterClient client.ITransmitterClient = mockClient
	watcher, _ := NewWatcher(ctx, "", "", common.Address{}, big.NewInt(1), transmitterClient)

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

	// 取消 context
	cancel()

	// 等待所有 goroutine 退出
	timeout := time.After(1 * time.Second)
	count := 0
	for count < 2 {
		select {
		case <-done:
			count++
		case <-timeout:
			t.Fatal("Timeout waiting for goroutines to exit")
		}
	}
}
