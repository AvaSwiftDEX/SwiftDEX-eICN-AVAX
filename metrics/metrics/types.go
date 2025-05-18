package metrics

import (
	"encoding/hex"
	"encoding/json"
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
	FromChainId     *big.Int    `json:"fromChainId"`
	FromHeight      *big.Int    `json:"fromHeight"`
}

func (m *MetricsData) PhaseStr() string {
	if str, ok := phaseToString[m.Phase]; ok {
		return str
	}
	return "Unknown"
}

func (m *MetricsData) MarshalJSON() ([]byte, error) {
	type Alias MetricsData
	return json.Marshal(&struct {
		Phase           string `json:"phase"`
		TxHash          string `json:"txHash"`
		ChainId         string `json:"chainId"`
		Height          string `json:"height"`
		TransactionHash string `json:"transactionHash"`
		CmHash          string `json:"cmHash"`
		Root            string `json:"root"`
		FromChainId     string `json:"fromChainId"`
		FromHeight      string `json:"fromHeight"`
		*Alias
	}{
		Phase:           m.PhaseStr(),
		TxHash:          m.TxHash.String(),
		ChainId:         m.ChainId.String(),
		Height:          m.Height.String(),
		TransactionHash: hex.EncodeToString(m.TransactionHash[:]),
		CmHash:          hex.EncodeToString(m.CmHash[:]),
		Root:            hex.EncodeToString(m.Root[:]),
		FromChainId:     m.FromChainId.String(),
		FromHeight:      m.FromHeight.String(),
		Alias:           (*Alias)(m),
	})
}
