package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
	ethclientext "github.com/kimroniny/SuperRunner-eICN-eth2/ethclientExt"
	"github.com/spf13/cobra"
)

func sendBackgroundWorkload(cfg *config.Config, stopCh chan struct{}) error {
	ctx := context.Background()
	senderPrivateKey, err := cfg.ReadPrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	senderPublicKey := senderPrivateKey.Public()
	senderAddress := crypto.PubkeyToAddress(*senderPublicKey.(*ecdsa.PublicKey))

	recipientPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	recipientAddress := crypto.PubkeyToAddress(*recipientPrivateKey.Public().(*ecdsa.PublicKey))

	httpclient, err := ethclientext.Dial(cfg.Chain.HTTPURL)
	if err != nil {
		log.Fatal(err)
	}
	ChainID, err := httpclient.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	amount := big.NewInt(10000)
	gasLimit := uint64(100000)

	gasPrice, err := httpclient.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-stopCh:
			return nil
		default:
		}
		nonce, err := httpclient.PendingNonceAt(ctx, senderAddress)
		if err != nil {
			log.Println(err)
			return err
		}
		transaction := types.NewTransaction(nonce, recipientAddress, amount, uint64(gasLimit), gasPrice, nil)
		signedTx, err := types.SignTx(transaction, types.NewEIP155Signer(ChainID), senderPrivateKey)
		if err != nil {
			log.Println(err)
			return err
		}
		err = httpclient.SendTransaction(ctx, signedTx)
		if err != nil {
			log.Println(err)
			return err
		}
		receipt, err := httpclient.WaitTransactionReceipt(ctx, signedTx.Hash(), 10000*time.Millisecond)
		if err != nil {
			log.Println(err)
			return err
		}
		if receipt.Status == types.ReceiptStatusFailed {
			log.Println("transaction failed")
			return errors.New("transaction failed")
		}
		log.Printf("transaction sent: %s\n", signedTx.Hash().Hex())
		time.Sleep(1000 * time.Millisecond)
	}

}

func send(cfgs []string) {
	wg := sync.WaitGroup{}
	stopCh := make(chan struct{})
	for _, configPath := range cfgs {
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := sendBackgroundWorkload(cfg, stopCh)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}()
	}

	// Set up signal handling
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Wait for interrupt signal
	<-sigCh
	fmt.Println("\nReceived interrupt signal. Shutting down...")

	// Signal all goroutines to stop
	close(stopCh)

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Background workload shutdown complete")
}

func main() {
	var configPath string

	var rootCmd = &cobra.Command{
		Use:   "background",
		Short: "Send background workload to keep mining",
		Long:  "A CLI tool to send background workload to keep mining",
		RunE: func(cmd *cobra.Command, args []string) error {
			configPaths := strings.Split(configPath, ";")
			send(configPaths)
			return nil
		},
	}

	rootCmd.Flags().StringVar(&configPath, "config_paths", "", "Config file path, separated by semicolon")

	// Mark required flags
	rootCmd.MarkFlagRequired("config_paths")

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
