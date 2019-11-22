package remote

import (
	"github.com/pefish/go-error"
	"github.com/pefish/go-format"
)

type ListDepositTransactionParam struct {
	Chain string `json:"chain,omitempty"`
	TxId  string `json:"tx_id"`
}
type ListDepositTransactionReturn struct {
	UserId        uint64  `db:"user_id" json:"user_id"`
	Currency      string  `db:"currency" json:"currency"`
	Chain         string  `db:"chain" json:"chain"`
	Amount        float64 `db:"amount" json:"amount"`
	Address       string  `db:"address" json:"address"`
	Status        int64   `db:"status" json:"status"`
	Height        string  `db:"height" json:"height"`
	BlockId       string  `db:"block_id" json:"block_id"`
	TxId          string  `db:"tx_id" json:"tx_id"`
	Confirmations int64   `db:"confirmations" json:"confirmations"`
	OutputIndex   int64   `db:"output_index" json:"output_index"`
	Tag           string  `db:"tag" json:"tag"`
	CreatedAt     string  `db:"created_at" json:"created_at"`
}

func (this *Remote) ListDepositTransaction(param ListDepositTransactionParam) ([]ListDepositTransactionReturn, *go_error.ErrorInfo) {
	path := `/api/storm/v1/deposit/transactions`
	data, err := this.getJson(path, param)
	if err != nil {
		return nil, err
	}
	results := []ListDepositTransactionReturn{}
	go_format.Format.SliceToStruct(data.([]interface{}), &results)
	return results, nil
}
