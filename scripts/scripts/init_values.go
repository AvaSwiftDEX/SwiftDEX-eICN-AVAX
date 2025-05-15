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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kimroniny/SuperRunner-eICN-eth2/SR2PC/AppState"
	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
	ethclientext "github.com/kimroniny/SuperRunner-eICN-eth2/ethclientExt"
)

type InitAppStateValuesArg struct {
	Values []*big.Int
}

func InitAppStateValues(ctx context.Context, config *config.Config, args InitAppStateValuesArg) error {
	privateKey, err := config.ReadPrivateKey()
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	client, err := ethclientext.Dial(config.Chain.HTTPURL)
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return err
	}

	// get chainID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(6000000)
	fmt.Println("gas tip: ", auth.GasLimit)
	gasPrice := 10
	auth.GasPrice = big.NewInt(int64(gasPrice))

	// get instance
	instance, err := AppState.NewAppState(config.Chain.AppStateAddress, client)
	if err != nil {
		return err
	}

	tx, err := instance.InitValue(auth, args.Values)
	if err != nil {
		return err
	}

	receipt, err := client.WaitTransactionReceipt(ctx, tx.Hash(), 10000*time.Millisecond)
	if err != nil {
		return err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return errors.New("init app state values receipt failed")
	}
	return nil
}
