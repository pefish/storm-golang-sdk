package signature

import (
	"github.com/pefish/go-crypto"
	"github.com/pefish/go-reflect"
	"sort"
	"strings"
	"time"
)

func NewDefaultSignatureHelper() *Signature {
	return &Signature{}
}

type Signature struct {
	ApiKey    string
	ApiSecret string
}

func (this *Signature) VerifySignature(secret string, timestamp string, method string, apiPath string, params map[string]interface{}) string {
	sortedStr := ``
	var keys []string
	for k, v := range params {
		if v != nil { // nil参数不参与签名
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		sortedStr += k + `=` + go_reflect.Reflect.ToString(params[k]) + `&`
	}
	sortedStr = strings.TrimSuffix(sortedStr, `&`)
	toSignStr := method + `|` + apiPath + `|` + timestamp + `|` + sortedStr
	return go_crypto.Crypto.HmacSha256ToHex(toSignStr, secret)
}

func (this *Signature) Sign(method string, apiPath string, params map[string]interface{}) (string, string) {
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
	return go_crypto.Crypto.HmacSha256ToHex(toSignStr, this.ApiSecret), timestamp
}
