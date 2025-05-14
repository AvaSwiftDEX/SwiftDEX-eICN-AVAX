package scripts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
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

type CrossSendArg struct {
	ChainIDs      []*big.Int
	Value         *big.Int
	AppIdentifier string
	AppValueId    *big.Int
}

func CrossSend(ctx context.Context, config *config.Config, args []CrossSendArg) error {
	fmt.Println("cross send start")

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
	for _, arg := range args {
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
	}

	return nil
}
