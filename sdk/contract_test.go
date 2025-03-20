package sdk

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// 创建测试所需的通用参数
func setupTestSDK(t *testing.T) *ContractSDK {
	// 创建私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("无法生成私钥: %v", err)
	}

	ctx := context.Background()
	url := "http://localhost:8545"                                               // 测试用URL
	chainId := big.NewInt(1)                                                     // 测试用chainId
	address := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e") // 测试用合约地址

	return NewContractSDK(ctx, url, chainId, address, privateKey, false)
}

func TestContractSDK_Run(t *testing.T) {
	sdk := setupTestSDK(t)

	// 启动 ContractSDK 监听协程
	go sdk.Run()

	// 发送测试数据
	cm := []byte(`{"targetChainId": 1, "sourceHeight": 100}`)
	proof := []byte("test_proof")
	sdk.TransmitterCrossReceive(cm, proof)

	// 等待数据处理
	time.Sleep(100 * time.Millisecond)

	// 发送停止信号
	sdk.ctx.Done()
}

func TestContractSDK_TransmitterCrossReceive(t *testing.T) {
	sdk := setupTestSDK(t)

	// 启动监听协程
	go sdk.Run()

	// 发送测试数据
	cm := []byte(`{"targetChainId": 1, "sourceHeight": 100}`)
	proof := []byte("test_proof")
	sdk.TransmitterCrossReceive(cm, proof)

	// 等待数据处理
	time.Sleep(100 * time.Millisecond)

	// 发送停止信号
	sdk.ctx.Done()
}

// 添加新的测试函数
func TestContractSDK_TransmitterSyncHeader(t *testing.T) {
	sdk := setupTestSDK(t)

	// 启动监听协程
	go sdk.Run()

	// 发送测试数据
	chainId := big.NewInt(1)
	number := big.NewInt(100)
	root := common.HexToHash("0x1234567890123456789012345678901234567890123456789012345678901234")
	sdk.TransmitterSyncHeader(chainId, number, root)

	// 等待数据处理
	time.Sleep(100 * time.Millisecond)

	// 发送停止信号
	sdk.ctx.Done()
}

// 添加组合测试函数
func TestContractSDK_MultipleOperations(t *testing.T) {
	sdk := setupTestSDK(t)

	// 启动监听协程
	go sdk.Run()

	// 发送跨链数据
	cm := []byte(`{"targetChainId": 1, "sourceHeight": 100}`)
	proof := []byte("test_proof")
	sdk.TransmitterCrossReceive(cm, proof)

	// 发送区块头数据
	chainId := big.NewInt(1)
	number := big.NewInt(100)
	root := common.HexToHash("0x1234567890123456789012345678901234567890123456789012345678901234")
	sdk.TransmitterSyncHeader(chainId, number, root)

	// 等待数据处理
	time.Sleep(100 * time.Millisecond)

	// 发送停止信号
	sdk.ctx.Done()
}

func TestContractSDK_ChannelOperations(t *testing.T) {
	sdk := setupTestSDK(t)

	// 测试通道容量
	if cap(sdk.Serv2SDK_CM) != 1024 {
		t.Errorf("期望 Serv2SDK_CM 通道容量为 1024，得到 %d", cap(sdk.Serv2SDK_CM))
	}
	if cap(sdk.Serv2SDK_HDR) != 1024 {
		t.Errorf("期望 Serv2SDK_HDR 通道容量为 1024，得到 %d", cap(sdk.Serv2SDK_HDR))
	}
	if cap(sdk.WaitCMHashCh) != 1024 {
		t.Errorf("期望 WaitCMHashCh 通道容量为 1024，得到 %d", cap(sdk.WaitCMHashCh))
	}
	if cap(sdk.WaitHDRHashCh) != 1024 {
		t.Errorf("期望 WaitHDRHashCh 通道容量为 1024，得到 %d", cap(sdk.WaitHDRHashCh))
	}
}
