package sdk

import (
	"testing"
	"time"
)

func TestContractSDK_Run(t *testing.T) {
	sdk := NewContractSDK()

	// 启动 ContractSDK 监听协程
	go sdk.Run()

	// 发送测试数据
	cm := []byte("test_cm")
	proof := []byte("test_proof")
	sdk.TransmitterCrossReceive(cm, proof)

	// 等待数据处理
	time.Sleep(100 * time.Millisecond)

	// 发送停止信号
	sdk.Stop <- struct{}{}
}

func TestContractSDK_TransmitterCrossReceive(t *testing.T) {
	sdk := NewContractSDK()

	// 启动监听协程
	go sdk.Run()

	// 发送测试数据
	cm := []byte("test_cm")
	proof := []byte("test_proof")
	sdk.TransmitterCrossReceive(cm, proof)

	// 等待数据处理
	time.Sleep(100 * time.Millisecond)

	// 发送停止信号
	sdk.Stop <- struct{}{}
}

// 添加新的测试函数
func TestContractSDK_TransmitterSyncHeader(t *testing.T) {
	sdk := NewContractSDK()

	// 启动监听协程
	go sdk.Run()

	// 发送测试数据
	chainId := 1
	number := uint64(100)
	var root [32]byte
	copy(root[:], []byte("test_root_hash_bytes_32_length__"))
	sdk.TransmitterSyncHeader(chainId, number, root)

	// 等待数据处理
	time.Sleep(100 * time.Millisecond)

	// 发送停止信号
	sdk.Stop <- struct{}{}
}

// 添加组合测试函数
func TestContractSDK_MultipleOperations(t *testing.T) {
	sdk := NewContractSDK()

	// 启动监听协程
	go sdk.Run()

	// 发送跨链数据
	cm := []byte("test_cm")
	proof := []byte("test_proof")
	sdk.TransmitterCrossReceive(cm, proof)

	// 发送区块头数据
	chainId := 1
	number := uint64(100)
	var root [32]byte
	copy(root[:], []byte("test_root_hash_bytes_32_length__"))
	sdk.TransmitterSyncHeader(chainId, number, root)

	// 等待数据处理
	time.Sleep(100 * time.Millisecond)

	// 发送停止信号
	sdk.Stop <- struct{}{}
}
