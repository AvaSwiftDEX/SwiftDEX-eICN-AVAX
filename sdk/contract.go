package sdk

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kimroniny/SuperRunner-eICN-eth2/SR2PC"
	ethclientext "github.com/kimroniny/SuperRunner-eICN-eth2/ethclientExt"
	"github.com/kimroniny/SuperRunner-eICN-eth2/logger"
	"github.com/sirupsen/logrus"
)

// CrossData 结构体表示通道传输的数据
type CrossData struct {
	Cm    []byte
	Proof []byte
}

type CMHashData struct {
	Hash      common.Hash
	CrossData *CrossData
}

// block header
type HeaderData struct {
	chainId *big.Int
	number  *big.Int
	root    common.Hash
}

type HDRHashData struct {
	Hash    common.Hash
	HdrData *HeaderData
}

// ContractSDK 结构体
type ContractSDK struct {
	ctx           context.Context
	URL           string
	HttpClient    *ethclientext.EthclientExt
	ChainNativeId *big.Int
	PrivateKey    *ecdsa.PrivateKey
	PublicKey     *ecdsa.PublicKey
	ChainId       *big.Int
	Address       common.Address
	Serv2SDK_CM   chan *CrossData   // 容量为1024的通道
	Serv2SDK_HDR  chan *HeaderData  // 容量为1025的通道
	WaitCMHashCh  chan *CMHashData  // 容量为1024的通道
	WaitHDRHashCh chan *HDRHashData // 容量为1024的通道
	Stop          chan struct{}     // 停止信号通道
	mutex         sync.Mutex
	log           *logrus.Entry
}

// NewContractSDK 创建一个新的 ContractSDK 实例
func NewContractSDK(ctx context.Context, url string, chainId *big.Int, address common.Address, privateKey *ecdsa.PrivateKey) *ContractSDK {
	httpclient, err := ethclientext.Dial(url)
	if err != nil {
		panic(err)
	}
	// get public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil
	}
	// get chainID
	chainNativeID, err := httpclient.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &ContractSDK{
		ctx:           ctx,
		URL:           url,
		Address:       address,
		ChainId:       chainId,
		ChainNativeId: chainNativeID,
		PrivateKey:    privateKey,
		PublicKey:     publicKeyECDSA,
		HttpClient:    httpclient,
		Serv2SDK_CM:   make(chan *CrossData, 1024),
		Serv2SDK_HDR:  make(chan *HeaderData, 1024),
		WaitCMHashCh:  make(chan *CMHashData, 1024),
		WaitHDRHashCh: make(chan *HDRHashData, 1024),
		Stop:          make(chan struct{}),
		mutex:         sync.Mutex{},
		log:           logger.NewComponent("ContractSDK"),
	}
}

// Run 持续监听通道，同时监听停止信号
func (sdk *ContractSDK) Run() {
	go sdk.ListenDataFromServer()
	go sdk.WaitCMHashData()
	go sdk.WaitHDRHashData()
}

func (sdk *ContractSDK) ListenDataFromServer() {
	for {
		select {
		case data := <-sdk.Serv2SDK_CM:
			sdk.CrossReceive(data)
		case data := <-sdk.Serv2SDK_HDR:
			sdk.SyncHeader(data)
		case <-sdk.ctx.Done():
			sdk.log.Info("Stopping ContractSDK...")
			return
		}
	}
}

// TransmitterCrossReceive 用于 server.Transmitter 调用，将数据发送到 ContractSDK
func (sdk *ContractSDK) TransmitterCrossReceive(cm []byte, proof []byte) {
	sdk.Serv2SDK_CM <- &CrossData{Cm: cm, Proof: proof}
}

// CrossReceive handles the data received from the channel (specific logic to be determined later)
func (sdk *ContractSDK) CrossReceive(data *CrossData) {
	sdk.mutex.Lock()
	defer sdk.mutex.Unlock()

	// parse CM 和 proof
	var cm SR2PC.SR2PCCrossMessage
	err := json.Unmarshal(data.Cm, &cm)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Fatal(err)
		return
	}
	if cm.TargetChainId.Cmp(sdk.ChainId) != 0 {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Fatal("chain id not match")
		return
	}

	var proof []byte
	err = json.Unmarshal(data.Proof, &proof)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Fatal(err)
		return
	}
	sdk.log.WithFields(logrus.Fields{
		"method": "CrossReceive",
	}).Info("Received CM and proof data: ", cm, proof)

	// get nonce
	fromAddress := crypto.PubkeyToAddress(*sdk.PublicKey)
	nonce, err := sdk.HttpClient.PendingNonceAt(sdk.ctx, fromAddress)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Fatal(err)
		return
	}

	// get auth
	auth, err := bind.NewKeyedTransactorWithChainID(sdk.PrivateKey, sdk.ChainNativeId)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Fatal(err)
		return
	}

	// set auth
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(4000000)
	gasPrice := 100
	auth.GasPrice = big.NewInt(int64(gasPrice))

	// get instance
	instance, err := SR2PC.NewSR2PC(sdk.Address, sdk.HttpClient)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Fatal(err)
		return
	}

	// send transaction
	tx, err := instance.CrossReceive(auth, cm, proof)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Fatal(err)
		return
	}
	sdk.log.WithFields(logrus.Fields{
		"method": "CrossReceive",
	}).Info("txHash: ", tx.Hash().Hex())

	// send hash to channel
	hash := tx.Hash()
	sdk.WaitCMHashCh <- &CMHashData{Hash: hash, CrossData: data}
}

func (sdk *ContractSDK) WaitCMHashData() {
	select {
	case cmHash := <-sdk.WaitCMHashCh:
		receipt, err := sdk.HttpClient.WaitTransactionReceipt(
			sdk.ctx,
			cmHash.Hash,
			10000*time.Millisecond,
		)
		if err != nil {
			sdk.log.WithFields(logrus.Fields{
				"method": "WaitCMHashData",
			}).Fatal(err)
			return
		}
		if receipt.Status == types.ReceiptStatusFailed {
			sdk.log.WithFields(logrus.Fields{
				"method": "WaitCMHashData",
			}).Fatal("CrossReceive transaction failed: ", cmHash.Hash.Hex())
			return
		}
	case <-sdk.ctx.Done():
		return
	}
}

func (sdk *ContractSDK) TransmitterSyncHeader(chainId *big.Int, number *big.Int, root common.Hash) {
	sdk.Serv2SDK_HDR <- &HeaderData{chainId: chainId, number: number, root: root}
}

func (sdk *ContractSDK) SyncHeader(data *HeaderData) {
	sdk.mutex.Lock()
	defer sdk.mutex.Unlock()

	sdk.log.WithFields(logrus.Fields{
		"method": "SyncHeader",
	}).Info("Received Header data: ", hex.EncodeToString(data.root[:]))

	// get nonce
	fromAddress := crypto.PubkeyToAddress(*sdk.PublicKey)
	nonce, err := sdk.HttpClient.PendingNonceAt(sdk.ctx, fromAddress)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "SyncHeader",
		}).Fatal(err)
		return
	}

	// get auth
	auth, err := bind.NewKeyedTransactorWithChainID(sdk.PrivateKey, sdk.ChainNativeId)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "SyncHeader",
		}).Fatal(err)
		return
	}

	// set auth
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(4000000)
	gasPrice := 100
	auth.GasPrice = big.NewInt(int64(gasPrice))

	// get instance
	instance, err := SR2PC.NewSR2PC(sdk.Address, sdk.HttpClient)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "SyncHeader",
		}).Fatal(err)
		return
	}

	// send transaction
	header := SR2PC.SR2PCBlockHeader{
		Height: data.number,
		Root:   data.root,
	}
	tx, err := instance.SyncHeader(auth, data.chainId, header)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "SyncHeader",
		}).Fatal(err)
		return
	}
	sdk.log.WithFields(logrus.Fields{
		"method": "SyncHeader",
	}).Info("txHash: ", tx.Hash().Hex())

	// send hash to channel
	hash := tx.Hash()
	sdk.WaitHDRHashCh <- &HDRHashData{Hash: hash, HdrData: data}
}

func (sdk *ContractSDK) WaitHDRHashData() {
	select {
	case hdrHash := <-sdk.WaitHDRHashCh:
		receipt, err := sdk.HttpClient.WaitTransactionReceipt(
			sdk.ctx,
			hdrHash.Hash,
			10000*time.Millisecond,
		)
		if err != nil {
			sdk.log.WithFields(logrus.Fields{
				"method": "WaitHDRHashData",
			}).Fatal(err)
			return
		}
		if receipt.Status == types.ReceiptStatusFailed {
			sdk.log.WithFields(logrus.Fields{
				"method": "WaitHDRHashData",
			}).Fatal("SyncHeader transaction failed: ", hdrHash.Hash.Hex())
			return
		}
	case <-sdk.ctx.Done():
		return
	}
}
