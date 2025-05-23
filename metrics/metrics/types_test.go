package metrics

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestMetricsDataJSON(t *testing.T) {
	fmt.Println("TestMetricsDataJSON")
	// Create test data
	testData := MetricsData{
		TransactionHash: [32]byte{1, 2, 3},
		CmHash:          [32]byte{4, 5, 6},
		ChainId:         big.NewInt(1),
		Height:          big.NewInt(100),
		Phase:           0, // MasterIssue
		IsConfirmed:     true,
		ByHeader:        false,
		Timestamp:       1234567890,
		TxHash:          common.HexToHash("0x123"),
		Forged:          true,
		Retry:           false,
		Gas:             21000,
		Root:            [32]byte{7, 8, 9},
		RealRoot:        [32]byte{10, 11, 12},
		FromChainId:     big.NewInt(2),
		FromHeight:      big.NewInt(200),
	}

	// Test marshaling
	jsonData, err := json.Marshal(testData)
	assert.NoError(t, err)
	fmt.Println(string(jsonData))

	// Test unmarshaling
	var decoded MetricsData
	err = json.Unmarshal(jsonData, &decoded)
	assert.NoError(t, err)

	// Verify fields
	assert.Equal(t, testData.TransactionHash, decoded.TransactionHash)
	assert.Equal(t, testData.CmHash, decoded.CmHash)
	assert.Equal(t, testData.ChainId.String(), decoded.ChainId.String())
	assert.Equal(t, testData.Height.String(), decoded.Height.String())
	assert.Equal(t, testData.Phase, decoded.Phase)
	assert.Equal(t, testData.IsConfirmed, decoded.IsConfirmed)
	assert.Equal(t, testData.ByHeader, decoded.ByHeader)
	assert.Equal(t, testData.Timestamp, decoded.Timestamp)
	assert.Equal(t, testData.TxHash, decoded.TxHash)
	assert.Equal(t, testData.Forged, decoded.Forged)
	assert.Equal(t, testData.Retry, decoded.Retry)
	assert.Equal(t, testData.Gas, decoded.Gas)
	assert.Equal(t, testData.Root, decoded.Root)
	assert.Equal(t, testData.RealRoot, decoded.RealRoot)
	assert.Equal(t, testData.FromChainId.String(), decoded.FromChainId.String())
	assert.Equal(t, testData.FromHeight.String(), decoded.FromHeight.String())

	// Test PhaseStr
	assert.Equal(t, "MasterIssue", testData.PhaseStr())
}

func TestMetricsDataJSONInvalid(t *testing.T) {
	invalidJSON := `{
		"transactionHash": "invalid",
		"cmHash": "invalid",
		"chainId": "invalid",
		"height": "invalid",
		"phase": "InvalidPhase",
		"txHash": "invalid"
	}`

	var decoded MetricsData
	err := json.Unmarshal([]byte(invalidJSON), &decoded)
	assert.Error(t, err)
}
