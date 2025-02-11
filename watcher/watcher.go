package watcher

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/kimroniny/SuperRunner-eICN-eth2/SR2PC"
	"github.com/kimroniny/SuperRunner-eICN-eth2/client"
	ethclientext "github.com/kimroniny/SuperRunner-eICN-eth2/ethclientExt"
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
	headerCh          chan *SyncHeaderData
	crossReceiveCh    chan *CrossReceiveData
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
}

func (wc *Watcher) MonitorBlock() {
	if wc.wsClient == nil {
		panic("ws client is nil")
	}
	if wc.httpClient == nil {
		panic("http client is nil")
	}
	headers := make(chan *types.Header)
	sub, err := wc.wsClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-wc.ctx.Done():
			return
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Number.Uint64(), header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			header, err := wc.httpClient.HeaderByNumber(wc.ctx, header.Number)
			if err != nil {
				log.Fatal("get block by hash, while error: ", err)
			}

			fmt.Println(header.Hash().Hex())    // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(header.Number.Uint64()) // 3477413
			fmt.Println(header.Time)            // 1529525947
			fmt.Println(header.Nonce.Uint64())  // 130524141876765836

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
			fmt.Printf("Received new header from headerCh of MonitorBlock, number: %d, hash: %s\n",
				header.Number.Uint64(),
				header.Root.Hex())
			err := (*wc.transmitterClient).SyncHeader(wc.chainId, header.Number, header.Root)
			if err != nil {
				log.Fatal("send header to transmitter client, while error: ", err)
			}
		}
	}
}

func (wc *Watcher) MonitorEvent() {
	if wc.httpClient == nil {
		panic("http client is nil")
	}
	if wc.wsClient == nil {
		panic("ws client is nil")
	}
	fmt.Println(">>> MonitorEvent")
	instance, err := SR2PC.NewSR2PC(wc.address, wc.wsClient)
	if err != nil {
		panic(err)
	}
	logs := make(chan *SR2PC.SR2PCSendCMHash)
	sub, err := instance.WatchSendCMHash(nil, logs)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case <-wc.ctx.Done():
			return
		case err := <-sub.Err():
			panic(err)
		case vLog := <-logs:
			fmt.Println(">>>>>>>>>>>>>>>>>>>> find new log <<<<<<<<<<<<<<<<<<<<")
			fmt.Println("cmHash: ", hex.EncodeToString(vLog.CmHash[:]))
			fmt.Println("status: ", vLog.Status)
			cm, err := instance.GetCMByHash(nil, vLog.CmHash)
			if err != nil {
				log.Fatal("get cm by hash, while error: ", err)
				continue
			}
			cmBytes, err := json.Marshal(cm)
			if err != nil {
				log.Fatal("marshal cm, while error: ", err)
				continue
			}
			proof, err := wc.GetProof(instance, &cm)
			if err != nil {
				log.Fatal("get proof, while error: ", err)
				continue
			}
			wc.crossReceiveCh <- &CrossReceiveData{
				chainId: cm.TargetChainId,
				Data1:   cmBytes,
				Data2:   proof,
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
			fmt.Println("crossReceive: ", data)
			err := (*wc.transmitterClient).CrossReceive(data.chainId, data.Data1, data.Data2)
			if err != nil {
				log.Fatal("send crossReceive to transmitter client, while error: ", err)
			}
		}
	}
}
