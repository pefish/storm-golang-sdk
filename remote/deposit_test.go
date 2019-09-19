package remote

import (
	"fmt"
	"testing"
)

func TestRemote_ListDepositTransaction(t *testing.T) {
	results, err := remote.ListDepositTransaction(ListDepositTransactionParam{
		TxId: `67e885c11f0dacf982dd2d1e10a7a62c37e454d9ead827eab7a96124fc628628`,
	})
	if err != nil {
		panic(err)
	}
	if results[0].UserId <= 0 {
		t.Error()
	}
}

func TestRemote_GetDepositTransaction(t *testing.T) {
	result, err := remote.GetDepositTransaction(GetDepositTransactionParam{
		Uuid: `hsjghjsfghafjagfha1`,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf(`%#v`, result)
}
