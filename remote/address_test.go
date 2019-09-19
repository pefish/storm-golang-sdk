package remote

import (
	"fmt"
	"testing"
)

func TestRemote_GetNewDepositAddress(t *testing.T) {
	result, err := remote.GetNewDepositAddress(GetNewDepositAddressParam{
		Currency: `BTC`,
		Chain:    `Btc`,
		Index:    1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf(`address: %s`, result.Address)
}

func TestRemote_ValidateAddress(t *testing.T) {
	result, err := remote.ValidateAddress(ValidateAddressParam{
		Currency: `BTC`,
		Chain:    `Btc`,
		Address:  `3HqH1qGAqNWPpbrvyGjnRxNEjcUKD4e6ea`,
	})
	if err != nil {
		panic(err)
	}
	if result != true {
		t.Error()
	}
}

func TestRemote_IsPlatformAddress(t *testing.T) {
	result, err := remote.IsPlatformAddress(IsPlatformParam{
		Currency: `BTC`,
		Chain:    `Btc`,
		Address:  `3HqH1qGAqNWPpbrvyGjnRxNEjcUKD4e6ea`,
	})
	if err != nil {
		panic(err)
	}
	if result != false {
		t.Error()
	}
}
