package sdk

import (
	"encoding/hex"
	"fmt"
)

// CrossData 结构体表示通道传输的数据
type CrossData struct {
	Cm    []byte
	Proof []byte
}

// block header
type HeaderData struct {
	chainId int
	number  uint64
	root    [32]byte
}

// ContractSDK 结构体
type ContractSDK struct {
	Serv2SDK_CM  chan *CrossData  // 容量为1024的通道
	Serv2SDK_HDR chan *HeaderData // 容量为1025的通道
	Stop         chan struct{}   // 停止信号通道
}

// NewContractSDK 创建一个新的 ContractSDK 实例
func NewContractSDK() *ContractSDK {
	return &ContractSDK{
		Serv2SDK_CM:  make(chan *CrossData, 1024),
		Serv2SDK_HDR: make(chan *HeaderData, 1024),
		Stop:         make(chan struct{}),
	}
}

// Run 持续监听通道，同时监听停止信号
func (sdk *ContractSDK) Run() {
	for {
		select {
		case data := <-sdk.Serv2SDK_CM:
			sdk.CrossReceive(data)
		case data := <-sdk.Serv2SDK_HDR:
			sdk.SyncHeader(data)
		case <-sdk.Stop:
			fmt.Println("Stopping ContractSDK...")
			return
		}
	}
}

// CrossReceive 处理通道收到的数据（具体逻辑稍后确定）
func (sdk *ContractSDK) CrossReceive(data *CrossData) {
	fmt.Println("Received CM and proof data:", string(data.Cm), string(data.Proof))
}

// TransmitterCrossReceive 用于 server.Transmitter 调用，将数据发送到 ContractSDK
func (sdk *ContractSDK) TransmitterCrossReceive(cm []byte, proof []byte) {
	sdk.Serv2SDK_CM <- &CrossData{Cm: cm, Proof: proof}
}

func (sdk *ContractSDK) SyncHeader(data *HeaderData) {
	fmt.Printf("Received Header data: %s\n", hex.EncodeToString(data.root[:]))
}

func (sdk *ContractSDK) TransmitterSyncHeader(chainId int, number uint64, root [32]byte) {
	sdk.Serv2SDK_HDR <- &HeaderData{chainId: chainId, number: number, root: root}
}
