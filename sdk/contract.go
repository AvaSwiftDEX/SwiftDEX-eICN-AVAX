package sdk

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kimroniny/SuperRunner-eICN-eth2/SR2PC"
	"github.com/kimroniny/SuperRunner-eICN-eth2/constants"
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

type ClientState struct {
	MsgHeight     *big.Int
	TrustedHeight *big.Int
}

type EventSyncHeader struct {
	chainId *big.Int
	height  *big.Int
}

type CacheData struct {
	CrossMessage *SR2PC.CrossMessage
	Proof        *[]byte
}

type RetryCacheData struct {
	Identifier   string
	CrossMessage *SR2PC.CrossMessage
	Root         *common.Hash
}

const (
	RETRYID_RetryPrepareConfirmCM   = "RetryPrepareConfirmCM"
	RETRYID_RetryPrepareUnconfirmCM = "RetryPrepareUnconfirmCM"
	RETRYID_RetryRollbackConfirmCM  = "RetryRollbackConfirmCM"
)

type UnlockShadowLockData struct {
	ChainId *big.Int
	Height  *big.Int
	Hash    common.Hash
}

// ContractSDK 结构体
type ContractSDK struct {
	ctx               context.Context
	URL               string
	HttpClient        *ethclientext.EthclientExt
	ChainNativeId     *big.Int
	ClientStates      map[string]*ClientState // key: chainID, value: height
	MsgCache          map[string]*Queue       // key: chainID + height, value: queue
	PrivateKey        *ecdsa.PrivateKey
	PublicKey         *ecdsa.PublicKey
	ChainId           *big.Int
	Address           common.Address
	InstanceCM        *SR2PC.SR2PC
	InstanceHDR       *SR2PC.SR2PC
	Serv2SDK_CM       chan *CrossData            // 容量为1024的通道
	Serv2SDK_HDR      chan *HeaderData           // 容量为1024的通道
	WaitCMHashCh      chan *CMHashData           // 容量为1024的通道
	WaitHDRHashCh     chan *HDRHashData          // 容量为1024的通道
	WaitRetryHashCh   chan *CMHashData           // 容量为1024的通道
	WatchSyncHeaderCh chan *EventSyncHeader      // 容量为1024的通道
	UnlockCh          chan *UnlockShadowLockData // 容量为1024的通道
	RetryCache        map[string]*Queue          // 存储待 retry 的跨链消息
	Stop              chan struct{}              // 停止信号通道
	mutex             sync.Mutex
	log               *logrus.Entry
	eICNAsync         bool
}

func cmPhaseStr(phase uint8) string {
	switch phase {
	case 0:
		return "PREPARE"
	case 1:
		return "RESPONSE"
	case 2:
		return "ABORT"
	case 3:
		return "COMMIT"
	case 4:
		return "ROLLBACK"
	default:
		return "UNKNOWN"
	}
}

// NewContractSDK 创建一个新的 ContractSDK 实例
func NewContractSDK(ctx context.Context, url string, chainId *big.Int, address common.Address, privateKey *ecdsa.PrivateKey, eICNAsync bool) *ContractSDK {
	httpclient, err := ethclientext.Dial(url)
	if err != nil {
		panic(err)
	}
	// get public key
	var publicKeyECDSA *ecdsa.PublicKey
	publicKeyECDSA = nil
	if privateKey != nil {
		publicKeyECDSA = privateKey.Public().(*ecdsa.PublicKey)
	}
	// get chainID
	chainNativeID, err := httpclient.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// get instance
	instanceCM, err := SR2PC.NewSR2PC(address, httpclient)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	instanceHDR, err := SR2PC.NewSR2PC(address, httpclient)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// init logger
	if logger.GetLogger() == nil {
		logger.InitLogger("")
	}
	return &ContractSDK{
		ctx:               ctx,
		URL:               url,
		Address:           address,
		ChainId:           chainId,
		ChainNativeId:     chainNativeID,
		ClientStates:      make(map[string]*ClientState),
		MsgCache:          make(map[string]*Queue),
		PrivateKey:        privateKey,
		PublicKey:         publicKeyECDSA,
		HttpClient:        httpclient,
		InstanceCM:        instanceCM,
		InstanceHDR:       instanceHDR,
		Serv2SDK_CM:       make(chan *CrossData, 1024),
		Serv2SDK_HDR:      make(chan *HeaderData, 1024),
		WaitCMHashCh:      make(chan *CMHashData, 1024),
		WaitHDRHashCh:     make(chan *HDRHashData, 1024),
		WaitRetryHashCh:   make(chan *CMHashData, 1024),
		WatchSyncHeaderCh: make(chan *EventSyncHeader, 1024),
		RetryCache:        make(map[string]*Queue),
		Stop:              make(chan struct{}),
		mutex:             sync.Mutex{},
		log:               logger.NewComponent("ContractSDK"),
		eICNAsync:         eICNAsync,
		UnlockCh:          make(chan *UnlockShadowLockData, 1024),
	}
}

// Run 持续监听通道，同时监听停止信号
func (sdk *ContractSDK) Run() {
	go sdk.ListenDataFromServer()
	go sdk.WaitCMHashData()
	go sdk.WaitHDRHashData()
	go sdk.WatchSyncHeaderEvent()
	go sdk.WaitRetryHashData()
	go sdk.RetryUponUnlock()
}

func (sdk *ContractSDK) ListenDataFromServer() {
	for {
		select {
		case data := <-sdk.Serv2SDK_CM:
			sdk.DealCounterpartCM(data)
		case data := <-sdk.Serv2SDK_HDR:
			sdk.SyncHeader(data)
		case <-sdk.ctx.Done():
			sdk.log.Info("Stopping ContractSDK...")
			return
		}
	}
}

func (sdk *ContractSDK) WatchSyncHeaderEvent() {
	for {
		select {
		case <-sdk.ctx.Done():
			return
		case syncHeader := <-sdk.WatchSyncHeaderCh:
			sdk.UpdateTrustedHeight(syncHeader.chainId, syncHeader.height)
		}
	}
}

func (sdk *ContractSDK) FindSyncHeader(chainId *big.Int, height *big.Int) {
	sdk.WatchSyncHeaderCh <- &EventSyncHeader{chainId: chainId, height: height}
}

func (sdk *ContractSDK) UpdateTrustedHeight(chainId *big.Int, height *big.Int) {
	chainIdStr := chainId.String()
	if _, ok := sdk.ClientStates[chainIdStr]; !ok {
		sdk.ClientStates[chainIdStr] = &ClientState{
			MsgHeight:     big.NewInt(0),
			TrustedHeight: height,
		}
	} else {
		if height.Cmp(sdk.ClientStates[chainIdStr].TrustedHeight) > 0 {
			sdk.ClientStates[chainIdStr].TrustedHeight = height
			sdk.log.WithFields(logrus.Fields{
				"method": "UpdateTrustedHeight",
			}).Info(fmt.Sprintf("Update TrustedHeight to %d for Chain#%d", height, chainId))
			sdk.NotifyCM(chainId, height)
		}
	}
}

func (sdk *ContractSDK) NotifyCM(chainId *big.Int, height *big.Int) {
	chainIdStr := chainId.String()
	heightStr := height.String()
	q, ok := sdk.MsgCache[chainIdStr+heightStr]
	if !ok {
		return
	}
	sdk.log.WithFields(logrus.Fields{
		"method": "NotifyCM",
	}).Info(fmt.Sprintf("NotifyCM for Chain#%d, height: %d, cm size: %d", chainId, height, q.Size()))
	for {
		item, ok := q.Dequeue()
		if !ok {
			sdk.log.WithFields(logrus.Fields{
				"method": "NotifyCM",
			}).Debug(fmt.Sprintf("get false from q.Dequeue(), chainId: %d, height: %d, q size: %d", chainId, height, q.Size()))
			break
		}
		cmData := item.(*CacheData)
		sdk.CrossReceive(cmData.CrossMessage, cmData.Proof)
	}
}

func (sdk *ContractSDK) ShouldReceiveCM(cm *SR2PC.CrossMessage) bool {
	chainIdStr := cm.SourceChainId.String()
	height := cm.ExpectedHeight
	clientState, ok := sdk.ClientStates[chainIdStr]
	if !ok || height.Cmp(clientState.TrustedHeight) > 0 {
		return false
	}
	return true
}

func (sdk *ContractSDK) CacheCM(cm *SR2PC.CrossMessage, proof *[]byte) {
	chainIdStr := cm.SourceChainId.String()
	height := cm.ExpectedHeight
	cacheKey := chainIdStr + height.String()
	if _, ok := sdk.MsgCache[cacheKey]; !ok {
		sdk.MsgCache[cacheKey] = NewQueue()
	}
	sdk.MsgCache[cacheKey].Enqueue(
		&CacheData{
			CrossMessage: cm,
			Proof:        proof,
		},
	)
}

func (sdk *ContractSDK) DealCounterpartCM(data *CrossData) {
	cm := SR2PC.CrossMessage{}
	err := json.Unmarshal(data.Cm, &cm)
	if err != nil {
		return
	}
	// store the CM to cache
	var proof []byte
	err = json.Unmarshal(data.Proof, &proof)
	if err != nil {
		return
	}
	if !sdk.eICNAsync && !sdk.ShouldReceiveCM(&cm) {
		sdk.log.WithFields(logrus.Fields{
			"method": "DealCounterpartCM",
		}).Info(
			fmt.Sprintf("CM(chainId: %d, height: %d, expectedHeight: %d, nonce: %d)'s header has not been trusted yet",
				cm.SourceChainId,
				cm.SourceHeight,
				cm.ExpectedHeight,
				cm.Nonce,
			),
		)
		sdk.CacheCM(&cm, &proof)
	} else {
		sdk.log.WithFields(logrus.Fields{
			"method": "DealCounterpartCM",
		}).Info(
			fmt.Sprintf("CM(chainId: %d, height: %d, expectedHeight: %d, nonce: %d)'s will be transmitted asynchronously",
				cm.SourceChainId,
				cm.SourceHeight,
				cm.ExpectedHeight,
				cm.Nonce,
			),
		)
		sdk.CrossReceive(&cm, &proof)
	}
}

// TransmitterCrossReceive is called by server.Transmitter, send data to ContractSDK
func (sdk *ContractSDK) TransmitterCrossReceive(cm []byte, proof []byte) {
	sdk.Serv2SDK_CM <- &CrossData{Cm: cm, Proof: proof}
}

// CrossReceive handles the data received from the channel (specific logic to be determined later)
func (sdk *ContractSDK) CrossReceive(cm *SR2PC.CrossMessage, proof *[]byte) {
	sdk.mutex.Lock()
	defer sdk.mutex.Unlock()

	if cm.TargetChainId.Cmp(sdk.ChainId) != 0 {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Error("chain id not match")
		return
	}

	sdk.log.WithFields(logrus.Fields{
		"method": "CrossReceive",
	}).Info(
		"Received CM and proof data, source chainID: ", cm.SourceChainId,
		", target chainID: ", cm.TargetChainId,
		", phase: ", cmPhaseStr(cm.Phase),
		", nonce: ", cm.Nonce,
		", source height: ", cm.SourceHeight,
		", cm input height: ", cm.CmInputHeight,
	)

	// get nonce
	fromAddress := crypto.PubkeyToAddress(*sdk.PublicKey)
	nonce, err := sdk.HttpClient.PendingNonceAt(sdk.ctx, fromAddress)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Error(err)
		return
	}

	// get auth
	auth, err := bind.NewKeyedTransactorWithChainID(sdk.PrivateKey, sdk.ChainNativeId)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Error(err)
		return
	}

	// set auth
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(4000000)
	gasPrice := constants.GasFee
	auth.GasPrice = big.NewInt(int64(gasPrice))

	// send transaction
	tx, err := sdk.InstanceCM.CrossReceive(auth, *cm, *proof)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Error(err)
		return
	}
	sdk.log.WithFields(logrus.Fields{
		"method": "CrossReceive",
	}).Debug("txHash: ", tx.Hash().Hex())

	// send hash to channel
	hash := tx.Hash()
	sdk.WaitCMHashCh <- &CMHashData{Hash: hash, CrossData: nil}
}

func (sdk *ContractSDK) WaitCMHashData() {
	for {
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
				}).Error(err)
				return
			}
			if receipt.Status == types.ReceiptStatusFailed {
				sdk.log.WithFields(logrus.Fields{
					"method": "WaitCMHashData",
				}).Error("CrossReceive transaction failed: ", cmHash.Hash.Hex())
				// sdk.ctx.Done()
				return
			}
			sdk.log.WithFields(logrus.Fields{
				"method": "WaitCMHashData",
			}).Debug("CrossReceive transaction success: ", cmHash.Hash.Hex())
			sdk.ParseRetryEvent(receipt)
			// TODO: check whether the tx needs to be resend
			// sdk.InstanceCM.ParseSendCMHash()
			// boundContract := bind.NewBoundContract(sdk.Address, SR2PC.SR2PCMetaData.GetAbi(), sdk.HttpClient, auth, sdk.InstanceCM.SR2PCCaller)
			// for _, log := range receipt.Logs {
			// 	sdk.InstanceCM.C
			// }
		case <-sdk.ctx.Done():
			return
		}
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
	}).Info(fmt.Sprintf("Received Header#%d from Chain#%d ", data.number, data.chainId))

	// get nonce
	fromAddress := crypto.PubkeyToAddress(*sdk.PublicKey)
	nonce, err := sdk.HttpClient.PendingNonceAt(sdk.ctx, fromAddress)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "SyncHeader",
		}).Error(err)
		return
	}

	// get auth
	auth, err := bind.NewKeyedTransactorWithChainID(sdk.PrivateKey, sdk.ChainNativeId)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "SyncHeader",
		}).Error(err)
		return
	}

	// set auth
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(10000000)
	gasPrice := constants.GasFee
	auth.GasPrice = big.NewInt(int64(gasPrice))

	// sdk.HttpClient.EstimateGas()

	// send transaction
	header := SR2PC.SR2PCBlockHeader{
		Height: data.number,
		Root:   data.root,
	}
	tx, err := sdk.InstanceHDR.SyncHeader(auth, data.chainId, header)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "SyncHeader",
		}).Error(err)
		return
	}
	sdk.log.WithFields(logrus.Fields{
		"method": "SyncHeader",
	}).Debug("txHash: ", tx.Hash().Hex())

	// send hash to channel
	hash := tx.Hash()
	sdk.WaitHDRHashCh <- &HDRHashData{Hash: hash, HdrData: data}
}

func (sdk *ContractSDK) WaitHDRHashData() {
	for {
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
				}).Error(err)
				return
			}
			if receipt.Status == types.ReceiptStatusFailed {
				sdk.log.WithFields(logrus.Fields{
					"method": "WaitHDRHashData",
				}).Error("SyncHeader transaction failed: ", hdrHash.Hash.Hex())
				// sdk.ctx.Done()
				return
			}
			sdk.log.WithFields(logrus.Fields{
				"method": "WaitHDRHashData",
			}).Debug("SyncHeader transaction success: ", hdrHash.Hash.Hex(), " chainId: ", hdrHash.HdrData.chainId, " height: ", hdrHash.HdrData.number)
			sdk.ParseRetryEvent(receipt)
			sdk.ParseUnlockEvent(receipt)
			sdk.ParseDebugEvent(receipt)
		case <-sdk.ctx.Done():
			return
		}
	}
}

func (sdk *ContractSDK) CrossRetry(identifier string, cm *SR2PC.CrossMessage, root *common.Hash, unlockHashStr string) {
	sdk.mutex.Lock()
	defer sdk.mutex.Unlock()

	if cm.TargetChainId.Cmp(sdk.ChainId) != 0 {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Error("chain id not match")
		return
	}

	sdk.log.WithFields(logrus.Fields{
		"method": "CrossRetry",
	}).Info(fmt.Sprintf("CrossRetry for CM(chainId: %d, targetChainId: %d, height: %d, expectedHeight: %d, nonce: %d), lockHash: %s",
		cm.SourceChainId,
		cm.TargetChainId,
		cm.SourceHeight,
		cm.ExpectedHeight,
		cm.Nonce,
		unlockHashStr,
	))

	// get nonce
	fromAddress := crypto.PubkeyToAddress(*sdk.PublicKey)
	nonce, err := sdk.HttpClient.PendingNonceAt(sdk.ctx, fromAddress)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossReceive",
		}).Error(err)
		return
	}

	// get auth
	auth, err := bind.NewKeyedTransactorWithChainID(sdk.PrivateKey, sdk.ChainNativeId)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossRetry",
		}).Error(err)
		return
	}

	// set auth
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(4000000)
	gasPrice := constants.GasFee
	auth.GasPrice = big.NewInt(int64(gasPrice))

	// send transaction
	tx, err := sdk.InstanceCM.CrossRetry(auth, identifier, *cm, *root)
	if err != nil {
		sdk.log.WithFields(logrus.Fields{
			"method": "CrossRetry",
		}).Error(err)
		return
	}
	sdk.log.WithFields(logrus.Fields{
		"method": "CrossRetry",
	}).Debug("txHash: ", tx.Hash().Hex())

	// send hash to channel
	hash := tx.Hash()
	sdk.WaitRetryHashCh <- &CMHashData{Hash: hash, CrossData: nil}
}

func (sdk *ContractSDK) WaitRetryHashData() {
	for {
		select {
		case cmHash := <-sdk.WaitRetryHashCh:
			receipt, err := sdk.HttpClient.WaitTransactionReceipt(
				sdk.ctx,
				cmHash.Hash,
				10000*time.Millisecond,
			)
			if err != nil {
				sdk.log.WithFields(logrus.Fields{
					"method": "WaitRetryHashData",
				}).Error(err)
				return
			}
			if receipt.Status == types.ReceiptStatusFailed {
				sdk.log.WithFields(logrus.Fields{
					"method": "WaitRetryHashData",
				}).Error("CrossRetry transaction failed: ", cmHash.Hash.Hex())
				// sdk.ctx.Done()
				return
			}
			sdk.log.WithFields(logrus.Fields{
				"method": "WaitRetryHashData",
			}).Debug("CrossRetry transaction success: ", cmHash.Hash.Hex())
			sdk.ParseRetryEvent(receipt)
		case <-sdk.ctx.Done():
			return
		}
	}
}

func (sdk *ContractSDK) RetryUponUnlock() {
	for {
		select {
		case <-sdk.ctx.Done():
			return
		case unlock := <-sdk.UnlockCh:
			unlockHashStr := hex.EncodeToString(unlock.Hash[:])
			if _, ok := sdk.RetryCache[unlockHashStr]; !ok {
				continue
			} else {
				for {
					rawData, ok := sdk.RetryCache[unlockHashStr].Dequeue()
					if !ok {
						break
					}
					retryData := rawData.(*RetryCacheData)
					sdk.CrossRetry(retryData.Identifier, retryData.CrossMessage, retryData.Root, unlockHashStr)
				}
			}

		}
	}
}

func (sdk *ContractSDK) ParseUnlockEvent(receipt *types.Receipt) {
	sdk.log.WithFields(logrus.Fields{
		"method": "ParseUnlockEvent",
	}).Debug("CacheRetryCM: ", func() string {
		keys := make([]string, 0, len(sdk.RetryCache))
		for k := range sdk.RetryCache {
			keys = append(keys, fmt.Sprintf("lockHash: %s, size: %d", k, sdk.RetryCache[k].Size()))
		}
		return strings.Join(keys, ", ")
	}())
	for _, log := range receipt.Logs {
		eventRPC, err := sdk.InstanceCM.ParseUnlockShadowLock(*log)
		if err == nil {
			sdk.log.WithFields(logrus.Fields{
				"method": "ParseUnlockEvent",
			}).Info(fmt.Sprintf("UnlockShadowLock event SDL.Hash: %s, SDL.ChainId: %d, SDL.Height: %d", hex.EncodeToString(eventRPC.LockHash[:]), eventRPC.ChainId, eventRPC.Height))
			sdk.UnlockCh <- &UnlockShadowLockData{
				ChainId: eventRPC.ChainId,
				Height:  eventRPC.Height,
				Hash:    eventRPC.LockHash,
			}
		}
	}
}

func (sdk *ContractSDK) ParseDebugEvent(receipt *types.Receipt) {
	for _, log := range receipt.Logs {
		eventDebug, err := sdk.InstanceCM.ParseDebug(*log)
		if err == nil {
			sdk.log.WithFields(logrus.Fields{
				"method": "ParseDebugEvent",
			}).Debug("Debug event: ", eventDebug.Msg)
		}
	}
}

func (sdk *ContractSDK) cacheRetryCM(lockHashStr string, identifier string, cm *SR2PC.CrossMessage, root *common.Hash) {
	if _, ok := sdk.RetryCache[lockHashStr]; !ok {
		sdk.RetryCache[lockHashStr] = NewQueue()
	}
	sdk.RetryCache[lockHashStr].Enqueue(
		&RetryCacheData{
			Identifier:   identifier,
			CrossMessage: cm,
			Root:         root,
		},
	)
}

func (sdk *ContractSDK) ParseRetryEvent(receipt *types.Receipt) {
	for _, log := range receipt.Logs {
		eventRPC, err := sdk.InstanceCM.ParseRetryPrepareConfirmCM(*log)
		if err == nil {
			lockHashStr := hex.EncodeToString(eventRPC.LockHash[:])
			sdk.log.WithFields(logrus.Fields{
				"method": "ParseRetryEvent",
			}).Info(fmt.Sprintf("RetryPrepareConfirmCM event: %s", lockHashStr))
			sdk.cacheRetryCM(lockHashStr, RETRYID_RetryPrepareConfirmCM, &eventRPC.Cm, &common.Hash{})
			break
		}
		eventRPU, err := sdk.InstanceCM.ParseRetryPrepareUnconfirmCM(*log)
		if err == nil {
			lockHashStr := hex.EncodeToString(eventRPU.LockHash[:])
			sdk.log.WithFields(logrus.Fields{
				"method": "ParseRetryEvent",
			}).Info("RetryPrepareUnconfirmCM event: ", lockHashStr)
			tmpRoot := common.Hash{}
			tmpRoot.SetBytes(eventRPU.Root[:])
			sdk.cacheRetryCM(lockHashStr, RETRYID_RetryPrepareUnconfirmCM, &eventRPU.Cm, &tmpRoot)
			break
		}
		eventRRC, err := sdk.InstanceCM.ParseRetryRollbackConfirmCM(*log)
		if err == nil {
			lockHashStr := hex.EncodeToString(eventRRC.LockHash[:])
			sdk.log.WithFields(logrus.Fields{
				"method": "ParseRetryEvent",
			}).Info("RetryRollbackConfirmCM event: ", lockHashStr)
			sdk.cacheRetryCM(lockHashStr, RETRYID_RetryRollbackConfirmCM, &eventRRC.Cm, &common.Hash{})
			break
		}
		// eventError, err := sdk.InstanceCM.ParseError(*log)
		// if err == nil {
		// 	sdk.log.WithFields(logrus.Fields{
		// 		"method": "ParseRetryEvent",
		// 	}).Info("Error event from ParseError: ", eventError.Reason)
		// } else {
		// 	sdk.log.WithFields(logrus.Fields{
		// 		"method": "ParseRetryEvent",
		// 	}).Debug("Error event from ParseError while error: ", err)
		// 	break
		// }
	}
}
