package remote

import (
	"github.com/pefish/go-error"
	"github.com/pefish/go-format"
)

type GetNewDepositAddressParam struct {
	Currency string `json:"currency"`
	Chain    string `json:"chain"`
	Index    uint64 `json:"index"`
}
type GetNewDepositAddressResult struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
}

func (r *Remote) GetNewDepositAddress(param GetNewDepositAddressParam) (*GetNewDepositAddressResult, *go_error.ErrorInfo) {
	path := `/api/storm/v1/new-address`
	data, err := r.postJson(path, param)
	if err != nil {
		return nil, err
	}
	result := GetNewDepositAddressResult{}
	go_format.Format.MapToStruct(data.(map[string]interface{}), &result)
	return &result, nil
}

type ValidateAddressParam struct {
	Currency string `json:"currency"`
	Chain    string `json:"chain"`
	Address  string `json:"address"`
	Tag      string `json:"tag"`
}

func (r *Remote) ValidateAddress(param ValidateAddressParam) (bool, *go_error.ErrorInfo) {
	path := `/api/storm/v1/validate-address`
	data, err := r.getJson(path, param)
	if err != nil {
		return false, err
	}
	return data.(bool), nil
}

type IsPlatformParam struct {
	Currency string `json:"currency"`
	Chain    string `json:"chain"`
	Address  string `json:"address"`
	Tag      string `json:"tag"`
}

func (r *Remote) IsPlatformAddress(param IsPlatformParam) (bool, *go_error.ErrorInfo) {
	path := `/api/storm/v1/is-platform-address`
	data, err := r.getJson(path, param)
	if err != nil {
		return false, err
	}
	return data.(bool), nil
}
