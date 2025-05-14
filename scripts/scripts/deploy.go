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
	"github.com/kimroniny/SuperRunner-eICN-eth2/SR2PC/AppState"
	"github.com/kimroniny/SuperRunner-eICN-eth2/SR2PC/Filter"
	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
	ethclientext "github.com/kimroniny/SuperRunner-eICN-eth2/ethclientExt"
)

func Deploy(ctx context.Context, config *config.Config) (common.Address, error) {
	privateKey, err := config.ReadPrivateKey()
	if err != nil {
		return common.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	client, err := ethclientext.Dial(config.Chain.HTTPURL)
	if err != nil {
		log.Fatal(err)
	}
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
	auth.GasLimit = uint64(6000000)
	fmt.Println("gas tip: ", auth.GasLimit)
	gasPrice := 10
	auth.GasPrice = big.NewInt(int64(gasPrice))

	_chainId := config.Chain.ID
	address, tx, instance, err := SR2PC.DeploySR2PC(auth, client, _chainId, big.NewInt(int64(config.Chain.ExpectedTrustDelta)))
	if err != nil {
		return common.Address{}, err
	}

	receipt, err := client.WaitTransactionReceipt(ctx, tx.Hash(), 10000*time.Millisecond)
	if err != nil {
		return common.Address{}, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return common.Address{}, errors.New("receipt failed")
	}
	fmt.Println("gas used: ", receipt.GasUsed)
	fmt.Println("address: ", receipt.ContractAddress.Hex())
	fmt.Println("tx: ", tx.Hash().Hex())
	_ = instance
	return address, nil
}

func DeployLibraries(ctx context.Context, config *config.Config) (common.Address, error) {
	privateKey, err := config.ReadPrivateKey()
	if err != nil {
		return common.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	client, err := ethclientext.Dial(config.Chain.HTTPURL)
	if err != nil {
		log.Fatal(err)
	}
	balance, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		return common.Address{}, err
	}
	fmt.Println("balance of fromAddress: ", balance, fromAddress.Hex())
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
	auth.GasLimit = uint64(6000000)
	fmt.Println("gas tip: ", auth.GasLimit)
	gasPrice := 10
	auth.GasPrice = big.NewInt(int64(gasPrice))

	address, tx, instance, err := Filter.DeployFilter(auth, client)
	if err != nil {
		return common.Address{}, err
	}

	receipt, err := client.WaitTransactionReceipt(ctx, tx.Hash(), 10000*time.Millisecond)
	if err != nil {
		return common.Address{}, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return common.Address{}, errors.New("receipt failed")
	}
	fmt.Println("gas used: ", receipt.GasUsed)
	fmt.Println("address: ", receipt.ContractAddress.Hex())
	fmt.Println("tx: ", tx.Hash().Hex())
	_ = instance
	return address, nil
}

func DeployAppState(ctx context.Context, config *config.Config) (common.Address, error) {
	privateKey, err := config.ReadPrivateKey()
	if err != nil {
		return common.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	client, err := ethclientext.Dial(config.Chain.HTTPURL)
	if err != nil {
		log.Fatal(err)
	}
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
	auth.GasLimit = uint64(6000000)
	fmt.Println("gas tip: ", auth.GasLimit)
	gasPrice := 10
	auth.GasPrice = big.NewInt(int64(gasPrice))

	address, tx, instance, err := AppState.DeployAppState(auth, client)
	if err != nil {
		return common.Address{}, err
	}

	receipt, err := client.WaitTransactionReceipt(ctx, tx.Hash(), 10000*time.Millisecond)
	if err != nil {
		return common.Address{}, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return common.Address{}, errors.New("receipt failed")
	}
	fmt.Println("gas used: ", receipt.GasUsed)
	fmt.Println("address: ", receipt.ContractAddress.Hex())
	fmt.Println("tx: ", tx.Hash().Hex())
	_ = instance
	return address, nil
}

func RegisterAppState(ctx context.Context, config *config.Config) error {
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
	instance, err := SR2PC.NewSR2PC(config.Chain.Address, client)
	if err != nil {
		return err
	}

	if config.Chain.AppStateName == "" {
		return errors.New("app state name is empty")
	}
	tx, err := instance.RegisterApp(auth, config.Chain.AppStateName, config.Chain.AppStateAddress)
	if err != nil {
		return err
	}

	receipt, err := client.WaitTransactionReceipt(ctx, tx.Hash(), 10000*time.Millisecond)
	if err != nil {
		return err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return errors.New("receipt failed")
	}
	return nil
}
