package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
	"github.com/kimroniny/SuperRunner-eICN-eth2/metrics/metrics"
	"github.com/kimroniny/SuperRunner-eICN-eth2/sdk"
)

// transaction hash -> chainId -> phase -> []MetricsData

func getGas(ctx context.Context, csdk *sdk.ContractSDK, tranHash string) (uint64, error) {
	receipt, err := csdk.HttpClient.TransactionReceipt(ctx, common.HexToHash(tranHash))
	if err != nil {
		return 0, err
	}
	return receipt.GasUsed, nil
}

func getRoot(ctx context.Context, csdk *sdk.ContractSDK, chainId *big.Int, height *big.Int) ([32]byte, error) {
	root, err := csdk.InstanceCM.GetRoot(&bind.CallOpts{Context: ctx}, chainId, height)
	if err != nil {
		return [32]byte{}, err
	}
	return root, nil
}

func ExtractObservationFromFile(filename string, configs []*config.Config) ([]byte, error) {
	// Read and parse the JSON file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tranStorage map[string]*TransactionStorage
	err = json.Unmarshal(data, &tranStorage)
	if err != nil {
		return nil, err
	}

	return ExtractObservation(tranStorage, configs)
}

func ExtractObservation(tranStorage map[string]*TransactionStorage, configs []*config.Config) ([]byte, error) {

	ctx := context.Background()
	sdks := make(map[string]*sdk.ContractSDK)
	for _, cfg := range configs {
		// run contractSDK
		privateKey, err := cfg.ReadPrivateKey()
		if err != nil {
			return nil, err
		}
		contractSDK := sdk.NewContractSDK(ctx, cfg.Chain.HTTPURL, cfg.Chain.ID, cfg.Chain.Address, privateKey, cfg.EICN.Async)
		sdks[cfg.Chain.ID.String()] = contractSDK
	}
	observation := make(map[string]map[string]map[string][]metrics.MetricsData)
	fmt.Println("tranStorage: ", len(tranStorage))
	for tranHash, storage := range tranStorage {
		for _, data := range storage.Metrics {
			if _, ok := observation[tranHash]; !ok {
				observation[tranHash] = make(map[string]map[string][]metrics.MetricsData)
			}
			if _, ok := observation[tranHash][data.ChainId.String()]; !ok {
				observation[tranHash][data.ChainId.String()] = make(map[string][]metrics.MetricsData)
			}
			if _, ok := observation[tranHash][data.ChainId.String()][data.PhaseStr()]; !ok {
				observation[tranHash][data.ChainId.String()][data.PhaseStr()] = make([]metrics.MetricsData, 0)
			}
			// get sdk
			sdk := sdks[data.ChainId.String()]
			if sdk == nil {
				fmt.Println("sdk is nil, chainId: ", data.ChainId.String())
				continue
			}
			// get gas
			gas, err := getGas(ctx, sdk, data.TxHash.String())
			fmt.Println("get gas: ", gas)
			if err == nil {

				data.Gas = gas
			}
			if err != nil {
				fmt.Println("get gas error: ", err)
				fmt.Println("chainId: ", data.ChainId.String(), "tranHash: ", tranHash)
			}
			// get real root
			root, err := getRoot(ctx, sdk, data.FromChainId, data.FromHeight)
			if err == nil {
				if data.PhaseStr() == "WorkerPrepareUnconfirm" {
					if data.Root != root {
						data.Forged = true
					}
				}
				data.RealRoot = root
			}
			observation[tranHash][data.ChainId.String()][data.PhaseStr()] = append(
				observation[tranHash][data.ChainId.String()][data.PhaseStr()],
				data,
			)
		}
	}
	return json.Marshal(observation)
}
