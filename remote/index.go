package remote

import (
	"github.com/pefish/go-error"
	"github.com/pefish/go-format"
	"github.com/pefish/go-http"
	signature2 "github.com/pefish/storm-golang-sdk/signature"
	"time"
)

type RemoteOptionFunc func(options *RemoteOption)

type RemoteOption struct {
	timeout time.Duration
}

var defaultHttpRequestOption = RemoteOption{
	timeout: 10 * time.Second,
}

func WithTimeout(timeout time.Duration) RemoteOptionFunc {
	return func(option *RemoteOption) {
		option.timeout = timeout
	}
}

func NewDefaultRemoteHelper(opts ...RemoteOptionFunc) *Remote {
	option := defaultHttpRequestOption
	for _, o := range opts {
		o(&option)
	}
	go_http.Http = go_http.NewHttpRequester(go_http.WithTimeout(option.timeout))
	return &Remote{
		BaseUrl: `https://storm.zg.com`,
	}
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

func (this *Remote) postJson(path string, params interface{}) (interface{}, *go_error.ErrorInfo) {
	result := ApiResult{}
	signature := signature2.SignatureClass{
		ApiKey:    this.ApiKey,
		ApiSecret: this.ApiSecret,
	}
	sig, timestamp := signature.Sign(`POST`, path, go_format.Format.StructToMap(params))
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
	signature := signature2.SignatureClass{
		ApiKey:    this.ApiKey,
		ApiSecret: this.ApiSecret,
	}
	sig, timestamp := signature.Sign(`GET`, path, go_format.Format.StructToMap(params))
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
