package remote

import (
	"github.com/pefish/go-error"
	"github.com/pefish/go-format"
	"github.com/pefish/go-http"
	"github.com/pefish/go-json"
	"github.com/pefish/go-reflect"
	signature2 "github.com/pefish/storm-golang-sdk/signature"
	"net/http"
	"time"
)

type RemoteOptionFunc func(options *RemoteOption)

type Remote struct {
	baseUrl          string
	signatureManager *signature2.SignatureClass
	httpRequester    *go_http.HttpClass
}

type RemoteOption struct {
	timeout   time.Duration
	apiKey    string
	apiSecret string
}

func WithTimeout(timeout time.Duration) RemoteOptionFunc {
	return func(option *RemoteOption) {
		option.timeout = timeout
	}
}

func WithKey(apiKey string, apiSecret string) RemoteOptionFunc {
	return func(option *RemoteOption) {
		option.apiKey = apiKey
		option.apiSecret = apiSecret
	}
}

func NewRemote(baseUrl string, opts ...RemoteOptionFunc) *Remote {
	option := RemoteOption{
		timeout: 10 * time.Second,
	}
	for _, o := range opts {
		o(&option)
	}
	return &Remote{
		baseUrl: baseUrl,
		signatureManager: &signature2.SignatureClass{
			ApiKey:    option.apiKey,
			ApiSecret: option.apiSecret,
		},
		httpRequester: go_http.NewHttpRequester(go_http.WithTimeout(option.timeout)),
	}
}

type ApiResult struct {
	Msg                  string      `json:"msg"`
	Code                 uint64      `json:"code"`
	Data                 interface{} `json:"data"`
	InternalErrorMessage string      `json:"internal_msg"`
}

func (this *Remote) postJson(path string, params interface{}) (interface{}, *go_error.ErrorInfo) {
	result := ApiResult{}
	sig, timestamp := this.signatureManager.Sign(`POST`, path, go_format.Format.StructToMap(params))
	resp, body := go_http.Http.Post(go_http.RequestParam{
		Url: this.baseUrl + path,
		Headers: map[string]interface{}{
			`BIZ-API-KEY`:       this.signatureManager.ApiKey,
			`BIZ-API-SIGNATURE`: sig,
			`BIZ-API-TIMESTAMP`: timestamp,
		},
		Params: params,
	})
	isValidRequest := this.verifyReturnData(resp, body)
	if !isValidRequest {
		return nil, &go_error.ErrorInfo{
			ErrorCode:    go_error.INTERNAL_ERROR_CODE,
			ErrorMessage: `response signature verify error`,
		}
	}
	go_format.Format.MapToStruct(go_json.Json.Parse(body).(map[string]interface{}), &result)
	if result.Code != 0 {
		return nil, &go_error.ErrorInfo{
			ErrorCode:            result.Code,
			ErrorMessage:         result.Msg,
			InternalErrorMessage: result.InternalErrorMessage,
			Data:                 result.Data,
		}
	}
	return result.Data, nil
}

func (this *Remote) getJson(path string, params interface{}) (interface{}, *go_error.ErrorInfo) {
	result := ApiResult{}
	sig, timestamp := this.signatureManager.Sign(`GET`, path, go_format.Format.StructToMap(params))
	resp, body := go_http.Http.Get(go_http.RequestParam{
		Url: this.baseUrl + path,
		Headers: map[string]interface{}{
			`BIZ-API-KEY`:       this.signatureManager.ApiKey,
			`BIZ-API-SIGNATURE`: sig,
			`BIZ-API-TIMESTAMP`: timestamp,
		},
		Params: params,
	})
	isValidRequest := this.verifyReturnData(resp, body)
	if !isValidRequest {
		return nil, &go_error.ErrorInfo{
			ErrorCode:    go_error.INTERNAL_ERROR_CODE,
			ErrorMessage: `response signature verify error`,
		}
	}
	go_format.Format.MapToStruct(go_json.Json.Parse(body).(map[string]interface{}), &result)
	if result.Code != 0 {
		return nil, &go_error.ErrorInfo{
			ErrorCode:            result.Code,
			ErrorMessage:         result.Msg,
			InternalErrorMessage: result.InternalErrorMessage,
			Data:                 result.Data,
		}
	}
	return result.Data, nil
}

func (this *Remote) verifyReturnData(resp *http.Response, body string) bool {
	timeStamp := resp.Header.Get(`BIZ_TIMESTAMP`)
	signatureStr := resp.Header.Get(`BIZ_RESP_SIGNATURE`)
	if timeStamp == `` || signatureStr == `` {
		return false
	}
	nowTimestamp := time.Now().UnixNano() / 1e6
	if nowTimestamp-go_reflect.Reflect.MustToInt64(timeStamp) > 30*1000 {
		return false
	}
	return this.signatureManager.VerifyResponseSignature(signatureStr, timeStamp, body)
}
