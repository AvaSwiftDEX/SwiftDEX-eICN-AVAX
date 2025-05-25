package scripts

import (
	"fmt"
	"math/big"
	"testing"
)

func TestGenerateIssueArgs(t *testing.T) {
	args := CrossSendWorkloadArg{
		WriteConflictRate: 10,
		TransactionNumber: 100,
		ChainIDs:          []*big.Int{big.NewInt(1), big.NewInt(2)},
		AppIdentifier:     "test",
	}
	issueArgs, _, _ := generateIssueArgs(args)
	fmt.Println(len(issueArgs))
}
