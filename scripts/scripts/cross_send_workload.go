package scripts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kimroniny/SuperRunner-eICN-eth2/SR2PC"
	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
	ethclientext "github.com/kimroniny/SuperRunner-eICN-eth2/ethclientExt"
)

type CrossSendWorkloadArg struct {
	WriteConflictRate uint64 // 0-100
	// ForgedProbability uint64 // 0-100 // in relayer
	// FinalizedLatency  uint64 // in deployment
	TransactionNumber uint64
	ChainIDs          []*big.Int
	AppIdentifier     string
}

type IssueArgs struct {
	ChainIDs      []*big.Int
	Value         *big.Int
	AppIdentifier string
	AppValueId    *big.Int
}

func generateIssueArgs(args CrossSendWorkloadArg) ([]IssueArgs, uint64, []*big.Int) {
	issueArgs := make([]IssueArgs, args.TransactionNumber)
	groupNumber := args.TransactionNumber * (100 - args.WriteConflictRate) / 100
	groupCapacity := args.TransactionNumber / groupNumber
	rest := args.TransactionNumber - groupCapacity*groupNumber

	// TODO:
	// init contract value by groupNumber

	initialValues := make([]*big.Int, groupNumber)
	for i := range int(groupNumber) {
		initialValues[i] = big.NewInt(10000)
	}

	var cnt uint64 = 0
	for range int(groupCapacity) {
		for j := range int(groupNumber) {
			issueArgs[cnt] = IssueArgs{ // TODO:
				ChainIDs:      args.ChainIDs,
				Value:         big.NewInt(10),
				AppIdentifier: args.AppIdentifier,
				AppValueId:    big.NewInt(int64(j)),
			}
			cnt++
		}
	}
	for i := range int(rest) {
		issueArgs[cnt] = IssueArgs{ // TODO:
			ChainIDs:      args.ChainIDs,
			Value:         big.NewInt(10),
			AppIdentifier: args.AppIdentifier,
			AppValueId:    big.NewInt(int64(i)),
		}
		cnt++
	}
	if cnt != args.TransactionNumber {
		log.Fatal("cnt != args.TransactionNumber", cnt, args.TransactionNumber)
	}
	fmt.Println("group number: ", groupNumber)
	fmt.Println("group capacity: ", groupCapacity)
	fmt.Println("rest: ", rest)

	return issueArgs, groupNumber, initialValues
}

func initContractValue(ctx context.Context, config *config.Config, workerCfgs []*config.Config, groupNumber uint64, initialValues []*big.Int) error {

	initAppStateValuesArgs := InitAppStateValuesArg{
		Values: initialValues,
	}
	for _, workerCfg := range workerCfgs {
		if err := InitAppStateValues(ctx, workerCfg, initAppStateValuesArgs); err != nil {
			return err
		}
	}
	return nil
}

func CrossSendWorkload(ctx context.Context, config *config.Config, workerCfgs []*config.Config, args CrossSendWorkloadArg) error {
	fmt.Println("cross send workload start")

	issueArgs, groupNumber, initialValues := generateIssueArgs(args)

	// init contract value by groupNumber
	if err := initContractValue(ctx, config, workerCfgs, groupNumber, initialValues); err != nil {
		return err
	}

	// get private key
	privateKey, err := config.ReadPrivateKey()
	if err != nil {
		return err
	}

	// get public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// get client
	client, err := ethclientext.Dial(config.Chain.HTTPURL)
	if err != nil {
		fmt.Println("dial client, while error: ", err)
		return err
	}

	// get instance
	instance, err := SR2PC.NewSR2PC(config.Chain.Address, client)
	if err != nil {
		return err
	}

	// get chainID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return err
	}

	// get auth
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}

	// set gas limit
	auth.GasLimit = uint64(4000000)
	fmt.Println("gas tip: ", auth.GasLimit)

	// set gas price
	gasPrice := 10
	auth.GasPrice = big.NewInt(int64(gasPrice))

	// set value
	auth.Value = big.NewInt(0)

	// send tx
	var txHashes []common.Hash
	for _, arg := range issueArgs {
		// get nonce
		nonce, err := client.PendingNonceAt(ctx, fromAddress)
		if err != nil {
			return err
		}
		auth.Nonce = big.NewInt(int64(nonce))

		// send tx
		tx, err := instance.CrossSend(auth, arg.ChainIDs, arg.Value, arg.AppIdentifier, arg.AppValueId)
		if err != nil {
			return err
		}

		// append tx hash
		txHashes = append(txHashes, tx.Hash())
	}
	

	blockHeights := make(map[uint64]struct{})
	// wait tx
	for _, txHash := range txHashes {
		receipt, err := client.WaitTransactionReceipt(ctx, txHash, 10000*time.Millisecond)
		if err != nil {
			return err
		}
		if receipt.Status == types.ReceiptStatusFailed {
			return errors.New("receipt failed, txhash: " + txHash.Hex())
		}
		fmt.Println("tx succeed: ", txHash.Hex())
		blockHeights[receipt.BlockNumber.Uint64()] = struct{}{}
	}
	for blockHeight := range blockHeights {
		fmt.Println("block height: ", blockHeight)
	}

	return nil
}
