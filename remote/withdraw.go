package remote

import (
	"github.com/pefish/go-error"
	"github.com/pefish/go-format"
)

type WithdrawParam struct {
	Currency  string  `json:"currency"`
	Chain     string  `json:"chain"`
	RequestId string  `json:"request_id"`
	Address   string  `json:"address"`
	Amount    string  `json:"amount"`
	Memo      *string `json:"memo,omitempty"`
}

func (this *Remote) Withdraw(param WithdrawParam) *go_error.ErrorInfo {
	path := `/api/storm-wallet/v1/withdraw`
	_, err := this.postJson(path, param)
	if err != nil {
		return err
	}
	return nil
}


type ListWithdrawTransactionParam struct {
	Chain *string `json:"chain,omitempty"`
	TxId  string  `json:"tx_id"`
}

type ListWithdrawTransactionReturn struct {
	UserId        uint64  `db:"user_id" json:"user_id"`
	Currency      string  `db:"currency" json:"currency"`
	Chain         string  `db:"chain" json:"chain"`
	Amount        float64 `db:"amount" json:"amount"`
	FromAddress   string  `db:"from_address" json:"from_address"`
	ToAddress     string  `db:"to_address" json:"to_address"`
	Status        int64   `db:"status" json:"status"`
	Height        string  `db:"height" json:"height"`
	BlockId       string  `db:"block_id" json:"block_id"`
	TxId          string  `db:"tx_id" json:"tx_id"`
	Confirmations int64   `db:"confirmations" json:"confirmations"`
	OutputIndex   int64   `db:"output_index" json:"output_index"`
	NetworkFee    float64 `db:"network_fee" json:"network_fee"`
	Tag           string  `db:"tag" json:"tag"`
	CreatedAt     string  `db:"created_at" json:"created_at"`
}

func (this *Remote) ListWithdrawTransaction(param ListWithdrawTransactionParam) ([]ListWithdrawTransactionReturn, *go_error.ErrorInfo) {
	path := `/api/storm-wallet/v1/withdraw/transactions`
	data, err := this.getJson(path, param)
	if err != nil {
		return nil, err
	}
	results := []ListWithdrawTransactionReturn{}
	go_format.Format.SliceToStruct(data.([]interface{}), &results)
	return results, nil
}
