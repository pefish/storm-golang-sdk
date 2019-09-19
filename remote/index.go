package remote

import (
	"fmt"
	"github.com/pefish/go-crypto"
	"github.com/pefish/go-error"
	"github.com/pefish/go-format"
	"github.com/pefish/go-http"
	"github.com/pefish/go-reflect"
	"sort"
	"strings"
	"time"
)

func NewDefaultRemoteHelper() *Remote {
	return &Remote{
		BaseUrl: `https://storm.zg.com`,
	}
}

func (this *Remote) SetTimeout(timeout time.Duration) {
	go_http.Http.SetTimeout(timeout)
}

type ApiResult struct {
	Msg  string      `json:"msg"`
	Code uint64      `json:"code"`
	Data interface{} `json:"data"`
}

type Remote struct {
	BaseUrl   string
	ApiKey    string
	ApiSecret string
}

func (this *Remote) sign(method string, apiPath string, params map[string]interface{}) (string, string) {
	sortedStr := ``
	var keys []string
	fmt.Println(params)
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

func (this *Remote) postJson(path string, params interface{}) (interface{}, *go_error.ErrorInfo) {
	result := ApiResult{}
	sig, timestamp := this.sign(`POST`, path, go_format.Format.StructToMap(params))
	go_http.Http.PostJsonForStruct(go_http.RequestParam{
		Url: this.BaseUrl + path,
		Headers: map[string]interface{}{
			`BIZ-API-KEY`:       this.ApiKey,
			`BIZ-API-SIGNATURE`: sig,
			`BIZ-API-TIMESTAMP`: timestamp,
		},
		Params: params,
	}, &result)
	if result.Code != 0 {
		return nil, &go_error.ErrorInfo{
			ErrorCode:    result.Code,
			ErrorMessage: result.Msg,
			Data:         result.Data,
		}
	}
	return result.Data, nil
}

func (this *Remote) getJson(path string, params interface{}) (interface{}, *go_error.ErrorInfo) {
	result := ApiResult{}
	sig, timestamp := this.sign(`GET`, path, go_format.Format.StructToMap(params))
	go_http.Http.GetForStruct(go_http.RequestParam{
		Url: this.BaseUrl + path,
		Headers: map[string]interface{}{
			`BIZ-API-KEY`:       this.ApiKey,
			`BIZ-API-SIGNATURE`: sig,
			`BIZ-API-TIMESTAMP`: timestamp,
		},
		Params: params,
	}, &result)
	if result.Code != 0 {
		return nil, &go_error.ErrorInfo{
			ErrorCode:    result.Code,
			ErrorMessage: result.Msg,
			Data:         result.Data,
		}
	}
	return result.Data, nil
}
