package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kimroniny/SuperRunner-eICN-eth2/SR2PC"
	ethclientext "github.com/kimroniny/SuperRunner-eICN-eth2/ethclientExt"
)

const READ_KEY_FROM_FILE = true
const PRIVATE_KEY = "c45ba5d6de0e502aefd23c98b40a2c9018e2e0286dde4fdb542ded619cefc8bd"

// const HTTP_URL = "http://10.210.143.97:7545"
// const WS_URL = "ws://10.210.143.97:7545"
const HTTP_URL = "http://127.0.0.1:8545"
const WS_URL = "ws://127.0.0.1:8546"

func readPrivateKeyFromHex() (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		return nil, err
	}
	// fmt.Println("private key: ", privateKey)
	return privateKey, nil
}

func readPrivateKeyFromFile() (*ecdsa.PrivateKey, error) {
	const keystoreFilename = "node/keystore/UTC--2025-02-08T06-12-23.376721660Z--a8a410a56f93e14fb5a71f5968958851915b6909"
	keyjson, err := os.ReadFile(keystoreFilename)
	if err != nil {
		return nil, err
	}
	auth := ""
	key, err := keystore.DecryptKey(keyjson, auth)
	if err != nil {
		return nil, err
	}
	// fmt.Println("key address: ", key.Address.Hex())
	return key.PrivateKey, nil
}

func readPrivateKey(readFromFile bool) (*ecdsa.PrivateKey, error) {
	if readFromFile {
		return readPrivateKeyFromFile()
	} else {
		return readPrivateKeyFromHex()
	}
}

func deploy(ctx context.Context, client *ethclientext.EthclientExt) (common.Address, error) {
	privateKey, err := readPrivateKey(READ_KEY_FROM_FILE)
	if err != nil {
		return common.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return common.Address{}, err
	}

	// get chainID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return common.Address{}, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return common.Address{}, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(4000000)
	fmt.Println("gas tip: ", auth.GasLimit)
	gasPrice := 10
	auth.GasPrice = big.NewInt(int64(gasPrice))

	_chainId := big.NewInt(1)
	address, tx, instance, err := SR2PC.DeploySR2PC(auth, client, _chainId)
	if err != nil {
		return common.Address{}, err
	}

	receipt, err := client.WaitTransactionReceipt(ctx, tx.Hash(), 3000*time.Millisecond)
	if err != nil {
		return common.Address{}, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return common.Address{}, errors.New("receipt failed")
	}

	fmt.Println("address: ", receipt.ContractAddress.Hex())
	fmt.Println("tx: ", tx.Hash().Hex())
	_ = instance
	return address, nil
}

func SendTransaction(ctx context.Context, client *ethclientext.EthclientExt, address common.Address) {
	privateKey, err := readPrivateKey(READ_KEY_FROM_FILE)
	if err != nil {
		log.Fatal(err)
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
		return
	}
	gasPrice := 10000

	// get chainID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(4000000)
	fmt.Println("gas tip: ", auth.GasLimit)
	auth.GasPrice = big.NewInt(int64(gasPrice))

	instance, err := SR2PC.NewSR2PC(address, client)
	if err != nil {
		log.Fatal(err)
		return
	}

	tx, err := instance.CrossSend(auth, []*big.Int{big.NewInt(2), big.NewInt(3)}, big.NewInt(100))
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("tx: ", tx.Hash().Hex())

	receipt, err := client.WaitTransactionReceipt(ctx, tx.Hash(), 3000*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatal("transaction failed")
		return
	}

	status, err := instance.GetTransStatus(nil, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("trans status: ", status)
}

func MonitorEvent(ctx context.Context, address common.Address, ws_client *ethclientext.EthclientExt, client *ethclientext.EthclientExt) {
	fmt.Println(">>> MonitorEvent")
	instance, err := SR2PC.NewSR2PC(address, ws_client)
	if err != nil {
		log.Fatal(err)
		return
	}
	logs := make(chan *SR2PC.SR2PCSendCMHash)
	sub, err := instance.WatchSendCMHash(nil, logs)
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(">>>>>>>>>>>>>>>>>>>> find new log <<<<<<<<<<<<<<<<<<<<")
			fmt.Println("cmHash: ", hex.EncodeToString(vLog.CmHash[:]))
			fmt.Println("status: ", vLog.Status)
		}
	}
}

func MonitorBlock(ctx context.Context, ws_client *ethclientext.EthclientExt, client *ethclientext.EthclientExt) {
	headers := make(chan *types.Header)
	sub, err := ws_client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Number.Uint64(), header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByNumber(ctx, header.Number)
			if err != nil {
				log.Fatal("get block by hash, while error: ", err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}

func work() {
	ctx := context.Background()

	client, err := ethclientext.Dial(HTTP_URL)
	if err != nil {
		log.Fatal(err)
	}

	ws_client, err := ethclientext.Dial(WS_URL)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("we have a connection")

	latestHeight, err := client.BlockNumber(ctx)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("latest block height: ", latestHeight)

	networkID, _ := client.NetworkID(ctx)
	fmt.Println("network ID: ", networkID)

	address, err := deploy(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("address: ", address.Hex())

	_ = ws_client
	go MonitorEvent(ctx, address, ws_client, client)
	// go MonitorBlock(ctx, ws_client, client)
	// go MonitorEventNative(ctx, address, ws_client, client)

	for i := 1; i <= 3; i++ {
		fmt.Println("Iteration:", i)
		time.Sleep(1 * time.Second)
		SendTransaction(ctx, client, address)
	}

	time.Sleep(2 * time.Second)
}

func main() {
	work()
}
