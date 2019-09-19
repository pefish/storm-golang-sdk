package remote

import (
	"github.com/pefish/go-error"
	"github.com/pefish/go-format"
)

type ListBalanceReturn struct {
	Currency string `json:"currency"`
	Chain    string `json:"chain"`
	Avail    string `json:"avail"`
	Freeze   string `json:"freeze"`
}

func (this *Remote) ListBalance() ([]ListBalanceReturn, *go_error.ErrorInfo) {
	path := `/api/storm-wallet/v1/balance`
	data, err := this.getJson(path, nil)
	if err != nil {
		return nil, err
	}
	results := []ListBalanceReturn{}
	go_format.Format.SliceToStruct(data.([]interface{}), &results)
	return results, nil
}
