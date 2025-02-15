package watcher

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/kimroniny/SuperRunner-eICN-eth2/SR2PC"
	"github.com/kimroniny/SuperRunner-eICN-eth2/client"
	ethclientext "github.com/kimroniny/SuperRunner-eICN-eth2/ethclientExt"
	"github.com/kimroniny/SuperRunner-eICN-eth2/logger"
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

type MetricsData struct {
	TransactionHash [32]byte
	CmHash          [32]byte
	ChainId         *big.Int
	Height          *big.Int
	Phase           uint8
	IsConfirmed     bool
	ByHeader        bool
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
	headerCh          chan *SyncHeaderData
	crossReceiveCh    chan *CrossReceiveData
	metricsCh         chan *MetricsData
	log               *logrus.Entry
}

func NewWatcher(
	ctx context.Context,
	httpUrl string,
	wsUrl string,
	address common.Address,
	chainId *big.Int,
	transmitter client.ITransmitterClient,
) (*Watcher, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	wc := &Watcher{
		ctx:               ctx,
		httpUrl:           httpUrl,
		wsUrl:             wsUrl,
		address:           address,
		chainId:           chainId,
		transmitterClient: &transmitter,
		httpClient:        nil,
		wsClient:          nil,
		headerCh:          make(chan *SyncHeaderData, 1024),
		crossReceiveCh:    make(chan *CrossReceiveData, 1024),
		log:               logger.NewComponent("Watcher"),
	}
	if httpUrl != "" {
		client, err := ethclientext.Dial(httpUrl)
		if err != nil {
			return nil, err
		}
		wc.httpClient = client
	}
	if wsUrl != "" {
		client, err := ethclientext.Dial(wsUrl)
		if err != nil {
			return nil, err
		}
		wc.wsClient = client
	}
	return wc, nil
}

func (wc *Watcher) Run() {
	go wc.MonitorBlock()
	go wc.SendHeader()
	go wc.MonitorEvent()
	go wc.CrossReceive()
	go wc.MonitorMetrics()
}

func (wc *Watcher) MonitorBlock() {
	if wc.wsClient == nil {
		wc.log.Fatal("ws client is nil")
	}
	if wc.httpClient == nil {
		wc.log.Fatal("http client is nil")
	}

	wc.log.Info("start to monitor block")

	headers := make(chan *types.Header)
	sub, err := wc.wsClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		wc.log.Fatal("subscribeHeader, the sub error: ", err)
	}
	for {
		select {
		case <-wc.ctx.Done():
			return
		case err := <-sub.Err():
			wc.log.Fatal("subscribeHeader, the sub error: ", err)
		case header := <-headers:
			wc.log.Info(fmt.Sprintf("find a new header: block height: %d, block hash: %s", header.Number.Uint64(), header.Hash().Hex()))

			header, err := wc.httpClient.HeaderByNumber(wc.ctx, header.Number)
			if err != nil {
				wc.log.Fatal("get block by hash, while error: ", err)
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
			// 处理接收到的 header
			wc.log.WithFields(logrus.Fields{
				"method":     "SendHeader",
				"chainID":    wc.chainId,
				"blockNum":   header.Number.Uint64(),
				"headerRoot": header.Root.Hex(),
			}).Info("call target server's SyncHeader")
			err := (*wc.transmitterClient).SyncHeader(wc.chainId, header.Number, header.Root)
			if err != nil {
				wc.log.Fatal("send header to transmitter client, while error: ", err)
			}
		}
	}
}

func (wc *Watcher) MonitorEvent() {
	if wc.httpClient == nil {
		wc.log.Fatal("http client is nil")
	}
	if wc.wsClient == nil {
		wc.log.Fatal("ws client is nil")
	}
	wc.log.Info("start to monitor event")
	instance, err := SR2PC.NewSR2PC(wc.address, wc.wsClient)
	if err != nil {
		wc.log.Fatal("new SR2PC instance, while error: ", err)
	}
	logs := make(chan *SR2PC.SR2PCSendCMHash)
	sub, err := instance.WatchSendCMHash(nil, logs)
	if err != nil {
		wc.log.Fatal("watchSendCMHash, while error: ", err)
	}
	for {
		select {
		case <-wc.ctx.Done():
			return
		case err := <-sub.Err():
			wc.log.Fatal("subscribeSendCMHash, the sub error: ", err)
		case vLog := <-logs:
			wc.log.Info(fmt.Sprintf("find a new log, cmHash: %s, status: %d", hex.EncodeToString(vLog.CmHash[:]), vLog.Status))
			cm, err := instance.GetCMByHash(nil, vLog.CmHash)
			if err != nil {
				wc.log.Fatal("get cm by hash, while error: ", err)
				continue
			}
			cmBytes, err := json.Marshal(cm)
			if err != nil {
				wc.log.Fatal("marshal cm, while error: ", err)
				continue
			}
			proof, err := wc.GetProof(instance, &cm)
			if err != nil {
				wc.log.Fatal("get proof, while error: ", err)
				continue
			}
			proofBytes, err := json.Marshal(proof)
			if err != nil {
				wc.log.Fatal("marshal proof, while error: ", err)
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

func (wc *Watcher) GetProof(instance *SR2PC.SR2PC, cm *SR2PC.SR2PCCrossMessage) ([]byte, error) {
	// TODO: get proof from instance
	proof := []byte(fmt.Sprintf("root of block height: %d\n", cm.SourceHeight))
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
			}).Info("call transmitter client's CrossReceive")
			err := (*wc.transmitterClient).CrossReceive(data.chainId, data.Data1, data.Data2)
			if err != nil {
				wc.log.Fatal("send crossReceive to transmitter client, while error: ", err)
			}
		}
	}
}

func (wc *Watcher) MonitorMetrics() {
	if wc.httpClient == nil {
		wc.log.Fatal("http client is nil")
	}
	if wc.wsClient == nil {
		wc.log.Fatal("ws client is nil")
	}
	wc.log.Info("start to monitor metrics")
	instance, err := SR2PC.NewSR2PC(wc.address, wc.wsClient)
	if err != nil {
		wc.log.Fatal("new SR2PC instance, while error: ", err)
	}
	logs := make(chan *SR2PC.SR2PCMetrics)
	sub, err := instance.WatchMetrics(nil, logs)
	if err != nil {
		wc.log.Fatal("watchMetrics, while error: ", err)
	}
	for {
		select {
		case <-wc.ctx.Done():
			return
		case err := <-sub.Err():
			wc.log.Fatal("subscribeMetrics, the sub error: ", err)
		case vLog := <-logs:
			wc.log.Info(fmt.Sprintf("find a new log, metrics: %v", vLog))
			wc.metricsCh <- &MetricsData{
				TransactionHash: vLog.TransactionHash,
				CmHash:          vLog.CmHash,
				ChainId:         vLog.ChainId,
				Height:          vLog.Height,
				Phase:           vLog.Phase,
				IsConfirmed:     vLog.IsConfirmed,
				ByHeader:        vLog.ByHeader,
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
			}).Info("call transmitter client's Metrics")
			// TODO: call collector client
		}
	}
}
