package metrics

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

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
	TxHash			common.Hash
}
