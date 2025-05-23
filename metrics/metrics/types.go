package metrics

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// enum MetricsCMPhase {
// 	MASTER_ISSUE,
// 	WORKER_PREPARE_UNCONFIRM,
// 	WORKER_PREPARE_CONFIRM,
// 	MASTER_RECEIVE_UNCONFIRM,
// 	MASTER_RECEIVE_CONFIRM,
// 	MASTER_COMMIT_UNCONFIRM,
// 	MASTER_COMMIT_CONFIRM,
// 	MASTER_ROLLBACK_UNCONFIRM,
// 	MASTER_ROLLBACK_CONFIRM,
// 	WORKER_COMMIT_UNCONFIRM,
// 	WORKER_COMMIT_CONFIRM,
// 	WORKER_ROLLBACK_UNCONFIRM,
// 	WORKER_ROLLBACK_CONFIRM
// }

var phaseToString = map[uint8]string{
	0:  "MasterIssue",
	1:  "WorkerPrepareUnconfirm",
	2:  "WorkerPrepareConfirm",
	3:  "MasterReceiveUnconfirm",
	4:  "MasterReceiveConfirm",
	5:  "MasterCommitUnconfirm",
	6:  "MasterCommitConfirm",
	7:  "MasterRollbackUnconfirm",
	8:  "MasterRollbackConfirm",
	9:  "WorkerCommitUnconfirm",
	10: "WorkerCommitConfirm",
	11: "WorkerRollbackUnconfirm",
	12: "WorkerRollbackConfirm",
}

// MetricsData represents the metrics data structure
type MetricsData struct {
	TransactionHash [32]byte    `json:"transactionHash"`
	CmHash          [32]byte    `json:"cmHash"`
	ChainId         *big.Int    `json:"chainId"`
	Height          *big.Int    `json:"height"`
	Phase           uint8       `json:"phase"`
	IsConfirmed     bool        `json:"isConfirmed"`
	ByHeader        bool        `json:"byHeader"`
	Timestamp       uint64      `json:"timestamp"`
	TxHash          common.Hash `json:"txHash"`
	Forged          bool        `json:"forged"`
	Retry           bool        `json:"retry"`
	Gas             uint64      `json:"gas"`
	Root            [32]byte    `json:"root"`
	RealRoot        [32]byte    `json:"realRoot"`
	FromChainId     *big.Int    `json:"fromChainId"`
	FromHeight      *big.Int    `json:"fromHeight"`
}

func (m *MetricsData) PhaseStr() string {
	if str, ok := phaseToString[m.Phase]; ok {
		return str
	}
	return "Unknown"
}

func (m MetricsData) MarshalJSON() ([]byte, error) {
	type Alias MetricsData
	return json.Marshal(&struct {
		TransactionHash string `json:"transactionHash"`
		CmHash          string `json:"cmHash"`
		ChainId         string `json:"chainId"`
		Height          string `json:"height"`
		Phase           string `json:"phase"`
		TxHash          string `json:"txHash"`
		Root            string `json:"root"`
		RealRoot        string `json:"realRoot"`
		FromChainId     string `json:"fromChainId"`
		FromHeight      string `json:"fromHeight"`
		Alias
	}{
		TransactionHash: hex.EncodeToString(m.TransactionHash[:]),
		CmHash:          hex.EncodeToString(m.CmHash[:]),
		ChainId:         m.ChainId.String(),
		Height:          m.Height.String(),
		Phase:           m.PhaseStr(),
		TxHash:          m.TxHash.String(),
		Root:            hex.EncodeToString(m.Root[:]),
		RealRoot:        hex.EncodeToString(m.RealRoot[:]),
		FromChainId:     m.FromChainId.String(),
		FromHeight:      m.FromHeight.String(),
		Alias:           Alias(m),
	})
}

func (m *MetricsData) UnmarshalJSON(data []byte) error {
	type Alias MetricsData
	aux := &struct {
		TransactionHash string `json:"transactionHash"`
		CmHash          string `json:"cmHash"`
		ChainId         string `json:"chainId"`
		Height          string `json:"height"`
		Phase           string `json:"phase"`
		TxHash          string `json:"txHash"`
		Root            string `json:"root"`
		RealRoot        string `json:"realRoot"`
		FromChainId     string `json:"fromChainId"`
		FromHeight      string `json:"fromHeight"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Convert TransactionHash
	transactionHash, err := hex.DecodeString(aux.TransactionHash)
	if err != nil {
		return fmt.Errorf("invalid transaction hash: %s", err)
	}
	copy(m.TransactionHash[:], transactionHash)

	// Convert CmHash
	cmHash, err := hex.DecodeString(aux.CmHash)
	if err != nil {
		return fmt.Errorf("invalid cm hash: %s", err)
	}
	copy(m.CmHash[:], cmHash)

	// Convert ChainId
	chainId := new(big.Int)
	if _, ok := chainId.SetString(aux.ChainId, 10); !ok {
		return fmt.Errorf("invalid chain id: %s", aux.ChainId)
	}
	m.ChainId = chainId

	// Convert Height
	height := new(big.Int)
	if _, ok := height.SetString(aux.Height, 10); !ok {
		return fmt.Errorf("invalid height: %s", aux.Height)
	}
	m.Height = height

	// Convert Phase
	var phase uint8 = 255
	for k, v := range phaseToString {
		if v == aux.Phase {
			phase = k
			break
		}
	}
	if phase == 255 {
		return fmt.Errorf("invalid phase: %s", aux.Phase)
	}
	m.Phase = phase

	// Convert TxHash
	txHash := common.HexToHash(aux.TxHash)
	m.TxHash = txHash

	// Convert Root
	root, err := hex.DecodeString(aux.Root)
	if err != nil {
		return fmt.Errorf("invalid root: %s", err)
	}
	copy(m.Root[:], root)

	// Convert RealRoot
	realRoot, err := hex.DecodeString(aux.RealRoot)
	if err != nil {
		return fmt.Errorf("invalid real root: %s", err)
	}
	copy(m.RealRoot[:], realRoot)

	// Convert FromChainId
	fromChainId := new(big.Int)
	if _, ok := fromChainId.SetString(aux.FromChainId, 10); !ok {
		return fmt.Errorf("invalid from chain id: %s", aux.FromChainId)
	}
	m.FromChainId = fromChainId

	// Convert FromHeight
	fromHeight := new(big.Int)
	if _, ok := fromHeight.SetString(aux.FromHeight, 10); !ok {
		return fmt.Errorf("invalid from height: %s", aux.FromHeight)
	}
	m.FromHeight = fromHeight

	return nil
}
