package remote

import (
	"errors"
	"github.com/pefish/go-error"
	"github.com/pefish/go-format"
	"github.com/pefish/go-http"
	"github.com/pefish/go-json"
	go_logger "github.com/pefish/go-logger"
	"github.com/pefish/go-reflect"
	signature2 "github.com/pefish/storm-golang-sdk/signature"
	"net/http"
	"time"
)

type RemoteOptionFunc func(options *RemoteOption)

type Remote struct {
	baseUrl          string
	signatureManager *signature2.SignatureClass
	timeout          time.Duration

	logger go_logger.InterfaceLogger
}

type RemoteOption struct {
	timeout    time.Duration
	reqPubKey  string
	reqPrivKey string
	resPubKey  string
	logger go_logger.InterfaceLogger
}

func WithTimeout(timeout time.Duration) RemoteOptionFunc {
	return func(option *RemoteOption) {
		option.timeout = timeout
	}
}

func WithLogger(logger go_logger.InterfaceLogger) RemoteOptionFunc {
	return func(option *RemoteOption) {
		option.logger = logger
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
		timeout: option.timeout,
		logger: option.logger,
	}
}

type ApiResult struct {
	Msg                  string      `json:"msg"`
	Code                 uint64      `json:"code"`
	Data                 interface{} `json:"data"`
	InternalErrorMessage string      `json:"internal_msg"`
}

func (r *Remote) postJson(path string, params interface{}) (interface{}, *go_error.ErrorInfo) {
	result := ApiResult{}
	sig, timestamp := r.signatureManager.SignRequest(`POST`, path, go_format.Format.StructToMap(params))
	resp, body, err := go_http.NewHttpRequester(go_http.WithTimeout(r.timeout), go_http.WithLogger(r.logger)).Post(go_http.RequestParam{
		Url: r.baseUrl + path,
		Headers: map[string]interface{}{
			`STM-REQ-KEY`:       r.signatureManager.ReqPubKey,
			`STM-REQ-SIGNATURE`: sig,
			`STM-REQ-TIMESTAMP`: timestamp,
		},
		Params: params,
	})
	if err != nil {
		return nil, go_error.Wrap(err)
	}
	isValidRequest := r.verifyReturnData(resp, body)
	if !isValidRequest {
		return nil, go_error.Wrap(errors.New(`response signature verify error`))
	}
	bodyJson, err := go_json.Json.Parse(body)
	if err != nil {
		return nil, go_error.Wrap(err)
	}
	go_format.Format.MapToStruct(bodyJson.(map[string]interface{}), &result)
	if result.Code != 0 {
		return nil, go_error.WrapWithAll(errors.New(result.Msg), result.Code, result.Data)
	}
	return result.Data, nil
}

func (r *Remote) getJson(path string, params interface{}) (interface{}, *go_error.ErrorInfo) {
	result := ApiResult{}
	sig, timestamp := r.signatureManager.SignRequest(`GET`, path, go_format.Format.StructToMap(params))
	resp, body, err := go_http.NewHttpRequester(go_http.WithTimeout(r.timeout), go_http.WithLogger(r.logger)).Get(go_http.RequestParam{
		Url: r.baseUrl + path,
		Headers: map[string]interface{}{
			`STM-REQ-KEY`:       r.signatureManager.ReqPubKey,
			`STM-REQ-SIGNATURE`: sig,
			`STM-REQ-TIMESTAMP`: timestamp,
		},
		Params: params,
	})
	if err != nil {
		return nil, go_error.Wrap(err)
	}
	isValidRequest := r.verifyReturnData(resp, body)
	if !isValidRequest {
		return nil, go_error.Wrap(errors.New(`response signature verify error`))
	}
	bodyJson, err := go_json.Json.Parse(body)
	if err != nil {
		return nil, go_error.Wrap(err)
	}
	go_format.Format.MapToStruct(bodyJson.(map[string]interface{}), &result)
	if result.Code != 0 {
		return nil, go_error.WrapWithAll(errors.New(result.Msg), result.Code, result.Data)
	}
	return result.Data, nil
}

func (r *Remote) verifyReturnData(resp *http.Response, body string) bool {
	timeStamp := resp.Header.Get(`STM-RES-TIMESTAMP`)
	signatureStr := resp.Header.Get(`STM-RES-SIGNATURE`)
	if timeStamp == `` || signatureStr == `` {
		return false
	}
	nowTimestamp := time.Now().UnixNano() / 1e6
	if nowTimestamp-go_reflect.Reflect.MustToInt64(timeStamp) > 30*1000 {
		return false
	}
	return r.signatureManager.VerifyResponseSignature(signatureStr, timeStamp, body)
}
