package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/kimroniny/SuperRunner-eICN-eth2/metrics/metrics"
)

// enum MetricsCMPhase {
// 	0 MASTER_ISSUE,
// 	1 WORKER_PREPARE_UNCONFIRM,
// 	2 WORKER_PREPARE_CONFIRM,
// 	3 MASTER_RECEIVE_UNCONFIRM,
// 	4 MASTER_RECEIVE_CONFIRM,
// 	5 MASTER_COMMIT_UNCONFIRM,
// 	6 MASTER_COMMIT_CONFIRM,
// 	7 MASTER_ROLLBACK_UNCONFIRM,
// 	8 MASTER_ROLLBACK_CONFIRM,
// 	9 WORKER_COMMIT_UNCONFIRM,
// 	10 WORKER_COMMIT_CONFIRM,
// 	11 WORKER_ROLLBACK_UNCONFIRM,
// 	12 WORKER_ROLLBACK_CONFIRM
// }

const (
	MetricsCMPhaseMasterIssue uint8 = iota
	MetricsCMPhaseWorkerPrepareUnconfirm
	MetricsCMPhaseWorkerPrepareConfirm
	MetricsCMPhaseMasterReceiveUnconfirm
	MetricsCMPhaseMasterReceiveConfirm
	MetricsCMPhaseMasterCommitUnconfirm
	MetricsCMPhaseMasterCommitConfirm
	MetricsCMPhaseMasterRollbackUnconfirm
	MetricsCMPhaseMasterRollbackConfirm
	MetricsCMPhaseWorkerCommitUnconfirm
	MetricsCMPhaseWorkerCommitConfirm
	MetricsCMPhaseWorkerRollbackUnconfirm
	MetricsCMPhaseWorkerRollbackConfirm
)

type TransactionStorage struct {
	Metrics          []metrics.MetricsData `json:"metrics"`
	CoordinatorChain *big.Int              `json:"coordinatorChain"` // if Phase == 0, CoordinatorChain is the chain id
	WorkerChains     []*big.Int            `json:"workerChains"`
	WorkerChainSet   map[string]bool       `json:"workerChainSet"`
	WorkerChainSize  int                   `json:"workerChainSize"`
	Finished         bool                  `json:"finished"`
}

func (ts *TransactionStorage) MarshalJSON() ([]byte, error) {
	type Alias struct {
		Metrics          []metrics.MetricsData `json:"metrics"`
		CoordinatorChain string                `json:"coordinatorChain"`
		WorkerChains     []string              `json:"workerChains"`
		WorkerChainSet   map[string]bool       `json:"workerChainSet"`
		WorkerChainSize  int                   `json:"workerChainSize"`
		Finished         bool                  `json:"finished"`
	}

	workerChains := make([]string, len(ts.WorkerChains))
	for i, chain := range ts.WorkerChains {
		workerChains[i] = chain.String()
	}

	alias := Alias{
		Metrics:          ts.Metrics,
		CoordinatorChain: ts.CoordinatorChain.String(),
		WorkerChains:     workerChains,
		WorkerChainSet:   ts.WorkerChainSet,
		WorkerChainSize:  ts.WorkerChainSize,
		Finished:         ts.Finished,
	}

	return json.Marshal(alias)
}

func (ts *TransactionStorage) UnmarshalJSON(data []byte) error {
	type Alias struct {
		Metrics          []metrics.MetricsData `json:"metrics"`
		CoordinatorChain string                `json:"coordinatorChain"`
		WorkerChains     []string              `json:"workerChains"`
		WorkerChainSet   map[string]bool       `json:"workerChainSet"`
		WorkerChainSize  int                   `json:"workerChainSize"`
		Finished         bool                  `json:"finished"`
	}

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	ts.Metrics = alias.Metrics
	ts.WorkerChainSet = alias.WorkerChainSet
	ts.WorkerChainSize = alias.WorkerChainSize
	ts.Finished = alias.Finished

	// Convert CoordinatorChain from string to *big.Int
	coordinatorChain := new(big.Int)
	if _, ok := coordinatorChain.SetString(alias.CoordinatorChain, 10); !ok {
		return fmt.Errorf("invalid coordinator chain value: %s", alias.CoordinatorChain)
	}
	ts.CoordinatorChain = coordinatorChain

	// Convert WorkerChains from []string to []*big.Int
	ts.WorkerChains = make([]*big.Int, len(alias.WorkerChains))
	for i, chainStr := range alias.WorkerChains {
		chain := new(big.Int)
		if _, ok := chain.SetString(chainStr, 10); !ok {
			return fmt.Errorf("invalid worker chain value at index %d: %s", i, chainStr)
		}
		ts.WorkerChains[i] = chain
	}

	return nil
}

type Analyzer struct {
	transHashStorage map[string]*TransactionStorage
	totalNumber      int
	finishedNumber   int
}

func NewAnalyzer(totalNumber int) *Analyzer {
	return &Analyzer{
		transHashStorage: make(map[string]*TransactionStorage),
		totalNumber:      totalNumber,
	}
}

func (aly *Analyzer) AnalysisMetrics(data metrics.MetricsData) {
	transactionHash := hex.EncodeToString(data.TransactionHash[:])
	if _, ok := aly.transHashStorage[transactionHash]; !ok {
		aly.transHashStorage[transactionHash] = &TransactionStorage{
			Metrics:          make([]metrics.MetricsData, 0),
			CoordinatorChain: big.NewInt(0),
			WorkerChains:     make([]*big.Int, 0),
			WorkerChainSet:   make(map[string]bool),
			WorkerChainSize:  0,
			Finished:         false,
		}
	}

	cmHash := hex.EncodeToString(data.CmHash[:])

	fmt.Printf("Received metrics event:\n")
	fmt.Printf("  Transaction Hash: %s\n", transactionHash)
	fmt.Printf("  CM Hash: %s\n", cmHash)
	fmt.Printf("  Chain ID: %s\n", data.ChainId.String())
	fmt.Printf("  Height: %s\n", data.Height.String())
	fmt.Printf("  Phase: %s\n", data.PhaseStr())
	fmt.Printf("  Is Confirmed: %t\n", data.IsConfirmed)
	fmt.Printf("  By Header: %t\n", data.ByHeader)
	fmt.Printf("  Timestamp: %d\n", data.Timestamp)
	fmt.Printf("  Tx Hash: %s\n", data.TxHash.String())
	fmt.Printf("----------------------------\n")

	aly.transHashStorage[transactionHash].Metrics = append(aly.transHashStorage[transactionHash].Metrics, data)

	// if phase == 0, set coordinator chain id
	switch data.Phase {
	case MetricsCMPhaseMasterIssue:
		aly.transHashStorage[transactionHash].CoordinatorChain = data.ChainId
	case MetricsCMPhaseMasterCommitUnconfirm,
		MetricsCMPhaseMasterCommitConfirm,
		MetricsCMPhaseMasterRollbackUnconfirm,
		MetricsCMPhaseMasterRollbackConfirm:
		// if master has committed or rollbacked, then all worker chains should be added to the worker chains list
		for _, metric := range aly.transHashStorage[transactionHash].Metrics {
			if metric.Phase == MetricsCMPhaseWorkerPrepareUnconfirm ||
				metric.Phase == MetricsCMPhaseWorkerPrepareConfirm ||
				metric.Phase == MetricsCMPhaseWorkerRollbackUnconfirm ||
				metric.Phase == MetricsCMPhaseWorkerRollbackConfirm ||
				metric.Phase == MetricsCMPhaseWorkerCommitUnconfirm ||
				metric.Phase == MetricsCMPhaseWorkerCommitConfirm {
				if _, ok := aly.transHashStorage[transactionHash].WorkerChainSet[metric.ChainId.String()]; !ok {
					aly.transHashStorage[transactionHash].WorkerChains = append(
						aly.transHashStorage[transactionHash].WorkerChains,
						metric.ChainId,
					)
					// false means the worker chain is not confirmed
					aly.transHashStorage[transactionHash].WorkerChainSet[metric.ChainId.String()] = false
				}
			}
		}
		if aly.transHashStorage[transactionHash].WorkerChainSize == 0 {
			for _, metric := range aly.transHashStorage[transactionHash].Metrics {
				if metric.Phase == MetricsCMPhaseMasterIssue {
					// if master has issued, then we could calculate the worker chain size
					aly.transHashStorage[transactionHash].WorkerChainSize++
				}
			}
		}
	case MetricsCMPhaseWorkerCommitConfirm,
		MetricsCMPhaseWorkerRollbackConfirm:
		aly.transHashStorage[transactionHash].WorkerChainSet[data.ChainId.String()] = true
		fmt.Printf("Worker chain(#%s) is confirmed\n", data.ChainId.String())
		finished := 0
		if data.Phase == MetricsCMPhaseWorkerRollbackConfirm {
			finished = 1 // rollback means one worker chain has been abort in prepare phase
		}
		for _, metric := range aly.transHashStorage[transactionHash].WorkerChainSet {
			if metric {
				finished++
				fmt.Printf("Confirmed worker chain number: %d / %d\n", finished, aly.transHashStorage[transactionHash].WorkerChainSize)
			}
		}
		if finished == aly.transHashStorage[transactionHash].WorkerChainSize {
			aly.transHashStorage[transactionHash].Finished = true
			aly.finishedNumber++
			fmt.Printf("Transaction finished (number: %d / %d): %s \n", aly.finishedNumber, aly.totalNumber, transactionHash)
		}
	}
}
