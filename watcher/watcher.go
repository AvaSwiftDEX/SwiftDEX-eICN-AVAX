package watcher

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/kimroniny/SuperRunner-eICN-eth2/SR2PC"
	"github.com/kimroniny/SuperRunner-eICN-eth2/client"
	ethclientext "github.com/kimroniny/SuperRunner-eICN-eth2/ethclientExt"
	"github.com/kimroniny/SuperRunner-eICN-eth2/logger"
	"github.com/kimroniny/SuperRunner-eICN-eth2/metrics/metrics"
	"github.com/kimroniny/SuperRunner-eICN-eth2/sdk"
	"github.com/sirupsen/logrus"
)

type SyncHeaderData struct {
	Number *big.Int
	Root   common.Hash
}

type CrossReceiveData struct {
	chainId *big.Int
	Data1   []byte
	Data2   []byte
}

type Watcher struct {
	ctx               context.Context
	httpUrl           string
	wsUrl             string
	chainId           *big.Int
	address           common.Address
	httpClient        *ethclientext.EthclientExt
	wsClient          *ethclientext.EthclientExt
	transmitterClient *client.ITransmitterClient
	contractSDK       *sdk.ContractSDK
	headerCh          chan *SyncHeaderData
	crossReceiveCh    chan *CrossReceiveData
	metricsCh         chan *metrics.MetricsData
	collectorClient   *metrics.CollectorClient
	log               *logrus.Entry
}

var MetricsCMPhase = []string{
	"MASTER_ISSUE",
	"WORKER_PREPARE_UNCONFIRM",
	"WORKER_PREPARE_CONFIRM",
	"MASTER_RECEIVE_UNCONFIRM",
	"MASTER_RECEIVE_CONFIRM",
	"MASTER_COMMIT_UNCONFIRM",
	"MASTER_COMMIT_CONFIRM",
	"MASTER_ROLLBACK_UNCONFIRM",
	"MASTER_ROLLBACK_CONFIRM",
	"WORKER_COMMIT_UNCONFIRM",
	"WORKER_COMMIT_CONFIRM",
	"WORKER_ROLLBACK_UNCONFIRM",
	"WORKER_ROLLBACK_CONFIRM",
}

func NewWatcher(
	ctx context.Context,
	httpUrl string,
	wsUrl string,
	address common.Address,
	chainId *big.Int,
	transmitter client.ITransmitterClient,
	contractSDK *sdk.ContractSDK,
	collectorURL string,
) (*Watcher, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if logger.GetLogger() == nil {
		logger.InitLogger("")
	}

	// 验证参数
	if chainId == nil {
		return nil, fmt.Errorf("chainId cannot be nil")
	}
	if chainId.Sign() < 0 {
		return nil, fmt.Errorf("chainId cannot be negative")
	}
	if transmitter == nil {
		return nil, fmt.Errorf("transmitter cannot be nil")
	}
	if collectorURL == "" {
		return nil, fmt.Errorf("collectorURL cannot be empty")
	}

	wc := &Watcher{
		ctx:               ctx,
		httpUrl:           httpUrl,
		wsUrl:             wsUrl,
		address:           address,
		chainId:           chainId,
		transmitterClient: &transmitter,
		contractSDK:       contractSDK,
		httpClient:        nil,
		wsClient:          nil,
		headerCh:          make(chan *SyncHeaderData, 1024),
		crossReceiveCh:    make(chan *CrossReceiveData, 1024),
		metricsCh:         make(chan *metrics.MetricsData, 1024),
		collectorClient:   metrics.NewCollectorClient(collectorURL),
		log:               logger.NewComponent("Watcher"),
	}

	// 如果提供了 HTTP URL，尝试连接
	if httpUrl != "" {
		client, err := ethclientext.Dial(httpUrl)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to HTTP endpoint: %v", err)
		}
		// 验证连接是否正常
		_, err = client.BlockNumber(ctx)
		if err != nil {
			client.Close()
			return nil, fmt.Errorf("failed to verify HTTP connection: %v", err)
		}
		wc.httpClient = client
	}

	// 如果提供了 WebSocket URL，尝试连接
	if wsUrl != "" {
		client, err := ethclientext.Dial(wsUrl)
		if err != nil {
			// 如果 HTTP 客户端已经创建，需要关闭它
			if wc.httpClient != nil {
				wc.httpClient.Close()
			}
			return nil, fmt.Errorf("failed to connect to WebSocket endpoint: %v", err)
		}
		// 验证连接是否正常
		_, err = client.BlockNumber(ctx)
		if err != nil {
			client.Close()
			if wc.httpClient != nil {
				wc.httpClient.Close()
			}
			return nil, fmt.Errorf("failed to verify WebSocket connection: %v", err)
		}
		wc.wsClient = client
	}

	// 验证 collector 是否有效
	if _, err := wc.collectorClient.GetMetrics(); err != nil {
		return nil, fmt.Errorf("failed to verify collector: %v", err)
	}

	return wc, nil
}

func (wc *Watcher) Run() {
	go wc.MonitorBlock()
	go wc.SendHeader()
	go wc.MonitorEvent()
	go wc.CrossReceive()
	go wc.MonitorMetrics()
	go wc.Metrics()
	go wc.MonitorError()
	go wc.MonitorSyncHeader()
}

func (wc *Watcher) MonitorBlock() {
	if wc.wsClient == nil {
		wc.log.Error("ws client is nil")
		return
	}
	if wc.httpClient == nil {
		wc.log.Error("http client is nil")
		return
	}

	wc.log.Info("start to monitor block")

	headers := make(chan *types.Header)
	sub, err := wc.wsClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		wc.log.WithError(err).Error("failed to subscribe to new headers")
		return
	}
	for {
		select {
		case <-wc.ctx.Done():
			return
		case err := <-sub.Err():
			wc.log.WithError(err).Error("subscription MonitorBlock error")
			return
		case header := <-headers:
			wc.log.Info(fmt.Sprintf("find a new header: block height: %d, block hash: %s", header.Number.Uint64(), header.Hash().Hex()))

			header, err := wc.httpClient.HeaderByNumber(wc.ctx, header.Number)
			if err != nil {
				wc.log.WithError(err).Error("failed to get block by hash")
				continue
			}

			// 发送 header 到 headerCh
			wc.headerCh <- &SyncHeaderData{
				Number: header.Number,
				Root:   header.Root,
			}
		}
	}
}

func (wc *Watcher) SendHeader() {
	for {
		select {
		case <-wc.ctx.Done():
			return
		case header := <-wc.headerCh:
			// TODO retrieve root according to the realistic blockchain
			// now it is a mock function
			root, err := wc.GetRoot(nil, header.Number)
			if err != nil {
				wc.log.WithError(err).Error("failed to get root")
				continue
			}
			wc.log.WithFields(logrus.Fields{
				"method":     "SendHeader",
				"chainID":    wc.chainId,
				"blockNum":   header.Number.Uint64(),
				"headerRoot": root.Hex(),
			}).Debug("call target server's SyncHeader")

			err = (*wc.transmitterClient).SyncHeader(wc.chainId, header.Number, root)
			if err != nil {
				wc.log.WithError(err).Error("failed to send header to transmitter client")
				continue
			}
		}
	}
}

func (wc *Watcher) MonitorEvent() {
	if wc.httpClient == nil {
		wc.log.Error("http client is nil")
		return
	}
	if wc.wsClient == nil {
		wc.log.Error("ws client is nil")
		return
	}
	wc.log.Info("start to monitor event")
	instance, err := SR2PC.NewSR2PC(wc.address, wc.wsClient)
	if err != nil {
		wc.log.WithError(err).Error("failed to create SR2PC instance")
		return
	}
	logs := make(chan *SR2PC.SR2PCSendCMHash)
	sub, err := instance.WatchSendCMHash(nil, logs)
	if err != nil {
		wc.log.WithError(err).Error("failed to watch SendCMHash")
		return
	}
	for {
		select {
		case <-wc.ctx.Done():
			return
		case err := <-sub.Err():
			wc.log.WithError(err).Error("subscription MonitorEvent error")
			return
		case vLog := <-logs:
			wc.log.Info(fmt.Sprintf("find a new log(SendCMHash), cmHash: %s, status: %d", hex.EncodeToString(vLog.CmHash[:]), vLog.Status))
			cm, err := instance.GetCMByHash(nil, vLog.CmHash)
			if err != nil {
				wc.log.WithError(err).Error("failed to get CM by hash")
				continue
			}
			cmBytes, err := json.Marshal(cm)
			if err != nil {
				wc.log.WithError(err).Error("failed to marshal CM")
				continue
			}
			proof, err := wc.GetProof(instance, &cm)
			if err != nil {
				wc.log.WithError(err).Error("failed to get proof")
				continue
			}
			proofBytes, err := json.Marshal(proof)
			if err != nil {
				wc.log.WithError(err).Error("failed to marshal proof")
				continue
			}
			wc.crossReceiveCh <- &CrossReceiveData{
				chainId: cm.TargetChainId,
				Data1:   cmBytes,
				Data2:   proofBytes,
			}
		}
	}
}

func (wc *Watcher) GetRoot(instance *SR2PC.SR2PC, height *big.Int) (common.Hash, error) {
	// TODO: get root from instance
	hash := sha256.Sum256([]byte(fmt.Sprintf("root of block height on chain: %d\n chainId: %d", height, wc.chainId)))
	return hash, nil
}

func (wc *Watcher) GetProof(instance *SR2PC.SR2PC, cm *SR2PC.CrossMessage) ([]byte, error) {
	// TODO: get proof from instance
	hash := sha256.Sum256([]byte(fmt.Sprintf("root of block height on chain: %d\n chainId: %d", cm.ExpectedHeight, wc.chainId)))
	proof := hash[:]
	return proof, nil
}

func (wc *Watcher) CrossReceive() {
	for {
		select {
		case <-wc.ctx.Done():
			return
		case data := <-wc.crossReceiveCh:
			wc.log.WithFields(logrus.Fields{
				"method":  "CrossReceive",
				"chainID": data.chainId,
			}).Debug("call transmitter client's CrossReceive")
			err := (*wc.transmitterClient).CrossReceive(data.chainId, data.Data1, data.Data2)
			if err != nil {
				wc.log.WithError(err).Error("failed to send crossReceive to transmitter client")
				continue
			}
		}
	}
}

func (wc *Watcher) MonitorMetrics() {
	if wc.httpClient == nil {
		wc.log.Error("http client is nil")
		return
	}
	if wc.wsClient == nil {
		wc.log.Error("ws client is nil")
		return
	}
	wc.log.Info("start to monitor metrics")
	instance, err := SR2PC.NewSR2PC(wc.address, wc.wsClient)
	if err != nil {
		wc.log.Error("new SR2PC instance, while error: ", err)
		return
	}
	logs := make(chan *SR2PC.SR2PCMetrics)
	sub, err := instance.WatchMetrics(nil, logs)
	if err != nil {
		wc.log.Error("watchMetrics, while error: ", err)
		return
	}
	for {
		select {
		case <-wc.ctx.Done():
			return
		case err := <-sub.Err():
			wc.log.Error("subscribeMetrics, the sub error: ", err)
			return
		case vLog := <-logs:
			wc.log.Info(
				fmt.Sprintf(
					"find a new log(Metrics), transactionHash: %s, phase: %s, height: %d, isConfirmed: %t, byHeader: %t",
					hex.EncodeToString(vLog.TransactionHash[:]),
					MetricsCMPhase[vLog.Phase],
					vLog.Height,
					vLog.IsConfirmed,
					vLog.ByHeader,
				),
			)
			// TODO: add gas used
			// receipt, err := wc.contractSDK.HttpClient.TransactionReceipt(wc.ctx, vLog.Raw.TxHash)
			// if err != nil {
			// 	wc.log.WithError(err).Error("failed to get transaction receipt")
			// 	continue
			// }
			// gasUsed := receipt.GasUsed
			wc.metricsCh <- &metrics.MetricsData{
				TransactionHash: vLog.TransactionHash,
				CmHash:          vLog.CmHash,
				ChainId:         vLog.ChainId,
				Height:          vLog.Height,
				Phase:           vLog.Phase,
				IsConfirmed:     vLog.IsConfirmed,
				ByHeader:        vLog.ByHeader,
				Timestamp:       uint64(time.Now().UnixMilli()),
				TxHash:          vLog.Raw.TxHash,
				Root:            vLog.Root,
				FromChainId:     vLog.FromChainId,
				FromHeight:      vLog.FromHeight,
			}
		}
	}
}

func (wc *Watcher) Metrics() {
	for {
		select {
		case <-wc.ctx.Done():
			return
		case data := <-wc.metricsCh:
			wc.log.WithFields(logrus.Fields{
				"method":          "Metrics",
				"transactionHash": hex.EncodeToString(data.TransactionHash[:]),
			}).Debug("Sending metrics event to collector")

			if err := wc.collectorClient.CollectMetricsEvent(*data); err != nil {
				wc.log.WithError(err).Error("Failed to send metrics event to collector")
			}
		}
	}
}

func (wc *Watcher) MonitorError() {
	if wc.httpClient == nil {
		wc.log.Error("http client is nil")
		return
	}
	if wc.wsClient == nil {
		wc.log.Error("ws client is nil")
		return
	}
	wc.log.Info("start to monitor errors")
	instance, err := SR2PC.NewSR2PC(wc.address, wc.wsClient)
	if err != nil {
		wc.log.Error("new SR2PC instance, while error: ", err)
		return
	}
	logs := make(chan *SR2PC.SR2PCError)
	sub, err := instance.WatchError(nil, logs)
	if err != nil {
		wc.log.Error("watchError, while error: ", err)
		return
	}
	logsWarning := make(chan *SR2PC.SR2PCWarning)
	subWarning, err := instance.WatchWarning(nil, logsWarning)
	if err != nil {
		wc.log.Error("watchWarning, while error: ", err)
		return
	}
	for {
		select {
		case <-wc.ctx.Done():
			return
		case err := <-sub.Err():
			wc.log.Error("subscribeErrorEvent, the sub error: ", err)
			return
		case err := <-subWarning.Err():
			wc.log.Error("subscribeWarningEvent, the sub error: ", err)
			return
		case vLog := <-logs:
			wc.log.Warning(fmt.Sprintf("find a new log(Error), transactionHash: %s, sourceChainId: %d, targetChainId: %d, phase: %d, reason: %s, others: %s", hex.EncodeToString(vLog.Cm.Payload[1]), vLog.Cm.SourceChainId, vLog.Cm.TargetChainId, vLog.Cm.Phase, vLog.Reason, hex.EncodeToString(vLog.Others)))
		case vLog := <-logsWarning:
			wc.log.Warning(fmt.Sprintf("find a new log(Warning), reason: %s, others: %s", vLog.Reason, hex.EncodeToString(vLog.Others)))
		}
	}
}

func (wc *Watcher) MonitorSyncHeader() {
	if wc.httpClient == nil {
		wc.log.Error("http client is nil")
		return
	}
	if wc.wsClient == nil {
		wc.log.Error("ws client is nil")
		return
	}
	wc.log.Info("start to monitor sync header")
	instance, err := SR2PC.NewSR2PC(wc.address, wc.wsClient)
	if err != nil {
		wc.log.Error("new SR2PC instance, while error: ", err)
		return
	}
	logs := make(chan *SR2PC.SR2PCSyncHeader)
	sub, err := instance.WatchSyncHeader(nil, logs)
	if err != nil {
		wc.log.Error("watchSyncHeader, while error: ", err)
		return
	}
	for {
		select {
		case <-wc.ctx.Done():
			return
		case err := <-sub.Err():
			wc.log.Error("subscribeSyncHeader, the sub error: ", err)
			return
		case vLog := <-logs:
			wc.log.Info(fmt.Sprintf("find a new log(SyncHeader), chainId: %d, height: %d, root: %s", vLog.ChainId, vLog.Height, hex.EncodeToString(vLog.Root[:])))
			wc.contractSDK.FindSyncHeader(vLog.ChainId, vLog.Height)
		}
	}
}
