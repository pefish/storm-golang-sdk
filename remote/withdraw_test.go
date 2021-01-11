package remote

import (
	"fmt"
	"testing"
)

func TestRemote_Withdraw(t *testing.T) {
	result := remote.Withdraw(WithdrawParam{
		Currency:  `BTC`,
		Chain:     `Btc`,
		RequestId: `185-aa`,
		Address:   `3HqH1qGAqNWPpbrvyGjnRxNEjcUKD4e6ea`,
		Amount:    `0.01`,
	})
	fmt.Printf(`%#v`, result)
}

func TestRemote_ListWithdrawTransaction(t *testing.T) {
	results, err := remote.ListWithdrawTransaction(ListWithdrawTransactionParam{
		TxId: `67e885c11f0dacf982dd2d1e10a7a62c37e454d9ead827eab7a96124fc628629`,
	})
	if err != nil {
		panic(err)
	}
	if results[0].UserId <= 0 {
		t.Error()
	}
}
