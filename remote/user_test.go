package remote

import (
	"fmt"
	"github.com/pefish/go-decimal"
	"testing"
)

func TestRemote_ListBalance(t *testing.T) {
	results, err := remote.ListBalance()
	if err != nil {
		panic(err)
	}
	fmt.Printf(`%#v`, results)
	if go_decimal.Decimal.Start(results[0].Avail).Lte(0) {
		t.Error()
	}
}

func TestRemote_ListUserCurrencies(t *testing.T) {
	results, err := remote.ListUserCurrencies()
	if err != nil {
		panic(err)
	}
	fmt.Printf(`%#v`, results)
}

func TestRemote_GetUserCurrency(t *testing.T) {
	result, err := remote.GetUserCurrency(GetUserCurrencyParam{
		Currency: `ETH`,
		Chain: `Eth`,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf(`%#v`, result)
}
