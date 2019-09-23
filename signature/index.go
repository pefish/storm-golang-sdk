package signature

import (
	"github.com/pefish/go-crypto"
	"github.com/pefish/go-reflect"
	"sort"
	"strings"
	"time"
)

func NewDefaultSignatureHelper() *SignatureClass {
	return &SignatureClass{}
}

type SignatureClass struct {
	ApiKey    string
	ApiSecret string
}

var Signature = SignatureClass{}

func (this *SignatureClass) SetApiKey(apiKey string) {
	this.ApiKey = apiKey
}

func (this *SignatureClass) SetApiSecret(apiSecret string) {
	this.ApiSecret = apiSecret
}

func (this *SignatureClass) VerifySignature(timestamp string, method string, apiPath string, params map[string]interface{}) string {
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
	return go_crypto.Crypto.HmacSha256ToHex(toSignStr, this.ApiSecret)
}

func (this *SignatureClass) Sign(method string, apiPath string, params map[string]interface{}) (string, string) {
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
