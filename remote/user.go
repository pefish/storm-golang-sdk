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
	path := `/api/storm/v1/balance`
	data, err := this.getJson(path, nil)
	if err != nil {
		return nil, err
	}
	results := []ListBalanceReturn{}
	go_format.Format.SliceToStruct(data.([]interface{}), &results)
	return results, nil
}

type ListUserCurrencyReturn struct {
	WithdrawLimitDaily            float64 `json:"withdraw_limit_daily"`
	MaxWithdrawAmount             float64 `json:"max_withdraw_amount"`
	WithdrawCheckLimit            float64 `json:"withdraw_check_limit"`
	Currency                      string  `json:"currency"`
	Chain                         string  `json:"chain"`
	ContractAddress               string  `json:"contract_address"`
	Decimals                      uint64  `json:"decimals"`
	DepositConfirmationThreshold  uint64  `json:"deposit_confirmation_threshold"`
	WithdrawConfirmationThreshold uint64  `json:"withdraw_confirmation_threshold"`
	NetworkFeeCurrency            string  `json:"network_fee_currency"`
	NetworkFeeDecimal             uint64  `json:"network_fee_decimal"`
	HasTag                        uint64  `json:"has_tag"`
	MaxTagLength                  uint64  `json:"max_tag_length"`
	IsWithdrawEnable              uint64  `json:"is_withdraw_enable"`
	IsDepositEnable               uint64  `json:"is_deposit_enable"`
}

func (this *Remote) ListUserCurrencies() ([]ListUserCurrencyReturn, *go_error.ErrorInfo) {
	path := `/api/storm/v1/user-currencies`
	data, err := this.getJson(path, nil)
	if err != nil {
		return nil, err
	}
	results := []ListUserCurrencyReturn{}
	go_format.Format.SliceToStruct(data.([]interface{}), &results)
	return results, nil
}

type GetUserCurrencyParam struct {
	Currency string `json:"currency" validate:"required"`
	Chain    string `json:"chain" validate:"required"`
}

func (this *Remote) GetUserCurrency(param GetUserCurrencyParam) (*ListUserCurrencyReturn, *go_error.ErrorInfo) {
	path := `/api/storm/v1/user-currency`
	data, err := this.getJson(path, param)
	if err != nil {
		return nil, err
	}
	result := ListUserCurrencyReturn{}
	go_format.Format.MapToStruct(data.(map[string]interface{}), &result)
	return &result, nil
}
