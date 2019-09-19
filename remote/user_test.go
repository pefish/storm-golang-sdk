package remote

import (
	"fmt"
	"testing"
	"github.com/pefish/go-decimal"
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
