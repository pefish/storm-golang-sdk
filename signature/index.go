package signature

import (
	"github.com/pefish/go-reflect"
	"sort"
	"strings"
	"time"
)

type SignatureClass struct {
	ReqPubKey  string
	ReqPrivKey string
	ResPubKey  string
}

func (this *SignatureClass) VerifyResponseSignature(signature string, timestamp string, body string) bool {
	return VerifySignature(body+`|`+timestamp, signature, this.ResPubKey)
}

func (this *SignatureClass) SignRequest(method string, apiPath string, params map[string]interface{}) (string, string) {
	sortedStr := ``
	var keys []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		sortedStr += k + `=` + go_reflect.Reflect.ToString(params[k]) + `&`
	}
	sortedStr = strings.TrimSuffix(sortedStr, `&`)
	timestamp := go_reflect.Reflect.ToString(time.Now().UnixNano() / 1e6)
	toSignStr := method + `|` + apiPath + `|` + timestamp + `|` + sortedStr
	return SignMessage(toSignStr, this.ReqPrivKey), timestamp
}
