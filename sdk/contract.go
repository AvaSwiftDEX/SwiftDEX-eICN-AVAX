package sdk

import (
	"fmt"
)

// CrossData 结构体表示通道传输的数据
type CrossData struct {
	Cm    []byte
	Proof []byte
}

// ContractSDK 结构体
type ContractSDK struct {
	Serv2SDK chan CrossData // 容量为1024的通道
	Stop     chan struct{}  // 停止信号通道
}

// NewContractSDK 创建一个新的 ContractSDK 实例
func NewContractSDK() *ContractSDK {
	return &ContractSDK{
		Serv2SDK: make(chan CrossData, 1024),
		Stop:     make(chan struct{}),
	}
}

// Run 持续监听通道，同时监听停止信号
func (sdk *ContractSDK) Run() {
	for {
		select {
		case data := <-sdk.Serv2SDK:
			sdk.CrossReceive(data.Cm, data.Proof)
		case <-sdk.Stop:
			fmt.Println("Stopping ContractSDK...")
			return
		}
	}
}

// CrossReceive 处理通道收到的数据（具体逻辑稍后确定）
func (sdk *ContractSDK) CrossReceive(cm []byte, proof []byte) {
	fmt.Println("Received data:", string(cm), string(proof))
}

// TransmitterCrossReceive 用于 server.Transmitter 调用，将数据发送到 ContractSDK
func (sdk *ContractSDK) TransmitterCrossReceive(cm []byte, proof []byte) {
	sdk.Serv2SDK <- CrossData{Cm: cm, Proof: proof}
}
