package ethclientext

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthclientExt struct {
	*ethclient.Client
}

func Dial(url string) (*EthclientExt, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &EthclientExt{client}, nil
}

func (ecext *EthclientExt) Accounts(ctx context.Context) ([]common.Address, error) {
	var result []common.Address
	err := ecext.Client.Client().CallContext(ctx, &result, "eth_accounts")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ec *EthclientExt) SendWrapTransaction(ctx context.Context, tx *types.Transaction) error {
	// TODO: send Transaction without signature
	data, err := tx.MarshalBinary()
	if err != nil {
		return err
	}
	return ec.Client.Client().CallContext(ctx, nil, "eth_sendRawTransaction", hexutil.Encode(data))
}

func (ec *EthclientExt) WaitTransactionReceipt(ctx context.Context, txhash common.Hash, timeout time.Duration) (*types.Receipt, error) {
	starttime := time.Now()
	for {
		txreceipt, err := ec.TransactionReceipt(ctx, txhash)
		if txreceipt != nil {
			return txreceipt, nil
		}
		if err != nil {
			time.Sleep(500 * time.Millisecond)
		}
		if time.Since(starttime) > timeout {
			return nil, fmt.Errorf("timeout waiting for transaction receipt")
		}
	}
}
