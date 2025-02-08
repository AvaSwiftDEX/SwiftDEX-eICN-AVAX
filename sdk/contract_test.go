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
