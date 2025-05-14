package metrics

import (
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
	TransactionHash [32]byte
	CmHash          [32]byte
	ChainId         *big.Int
	Height          *big.Int
	Phase           uint8
	IsConfirmed     bool
	ByHeader        bool
	Timestamp       uint64
	TxHash          common.Hash
}

func (m *MetricsData) PhaseStr() string {
	if str, ok := phaseToString[m.Phase]; ok {
		return str
	}
	return "Unknown"
}
