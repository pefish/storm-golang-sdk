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
	timeout    time.Duration
	reqPubKey  string
	reqPrivKey string
	resPubKey  string
}

func WithTimeout(timeout time.Duration) RemoteOptionFunc {
	return func(option *RemoteOption) {
		option.timeout = timeout
	}
}

func WithKey(reqPubKey string, reqPrivKey string, resPubKey string) RemoteOptionFunc {
	return func(option *RemoteOption) {
		option.reqPubKey = reqPubKey
		option.reqPrivKey = reqPrivKey
		option.resPubKey = resPubKey
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
			ReqPubKey:  option.reqPubKey,
			ReqPrivKey: option.reqPrivKey,
			ResPubKey:  option.resPubKey,
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
	sig, timestamp := this.signatureManager.SignRequest(`POST`, path, go_format.Format.StructToMap(params))
	resp, body := go_http.Http.Post(go_http.RequestParam{
		Url: this.baseUrl + path,
		Headers: map[string]interface{}{
			`STM-REQ-KEY`:       this.signatureManager.ReqPubKey,
			`STM-REQ-SIGNATURE`: sig,
			`STM-REQ-TIMESTAMP`: timestamp,
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
	bodyJson, err := go_json.Json.Parse(body)
	if err != nil {
		return nil, &go_error.ErrorInfo{
			ErrorCode:    go_error.INTERNAL_ERROR_CODE,
			ErrorMessage: `parse body error`,
			Err: err,
		}
	}
	go_format.Format.MapToStruct(bodyJson.(map[string]interface{}), &result)
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
	sig, timestamp := this.signatureManager.SignRequest(`GET`, path, go_format.Format.StructToMap(params))
	resp, body := go_http.Http.Get(go_http.RequestParam{
		Url: this.baseUrl + path,
		Headers: map[string]interface{}{
			`STM-REQ-KEY`:       this.signatureManager.ReqPubKey,
			`STM-REQ-SIGNATURE`: sig,
			`STM-REQ-TIMESTAMP`: timestamp,
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
	bodyJson, err := go_json.Json.Parse(body)
	if err != nil {
		return nil, &go_error.ErrorInfo{
			ErrorCode:    go_error.INTERNAL_ERROR_CODE,
			ErrorMessage: `parse body error`,
			Err: err,
		}
	}
	go_format.Format.MapToStruct(bodyJson.(map[string]interface{}), &result)
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
	timeStamp := resp.Header.Get(`STM-RES-TIMESTAMP`)
	signatureStr := resp.Header.Get(`STM-RES-SIGNATURE`)
	if timeStamp == `` || signatureStr == `` {
		return false
	}
	nowTimestamp := time.Now().UnixNano() / 1e6
	if nowTimestamp-go_reflect.Reflect.MustToInt64(timeStamp) > 30*1000 {
		return false
	}
	return this.signatureManager.VerifyResponseSignature(signatureStr, timeStamp, body)
}
