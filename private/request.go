package private

import (
	"fmt"
	"github.com/naoki-maeda/zaif-go"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type ApiClient struct {
	Key      string
	Secret   string
	Endpoint string
}

func NewApiClient(key string, secret string, endpoint string) *ApiClient {
	return &ApiClient{
		Key:      key,
		Secret:   secret,
		Endpoint: endpoint,
	}
}

type ApiParams struct {
	Method string `url:"method"`
	Nonce  string `url:"nonce"`
}

func newPrivateApiParams(method string) ApiParams {
	return ApiParams{
		Method: method,
		Nonce:  zaif.GetNonce(),
	}
}

type ApiError struct {
	Message string
}

func (err ApiError) Error() string {
	return err.Message
}

type ApiResponse struct {
	Success int    `json:"success"`
	Error   string `json:"error"`
}

func (api *ApiClient) Request(method string, values url.Values, out interface{}) error {
	apiParams := newPrivateApiParams(method)
	values.Add("method", apiParams.Method)
	values.Add("nonce", apiParams.Nonce)

	encodeParams := values.Encode()
	req, err := http.NewRequest("POST", api.Endpoint, strings.NewReader(encodeParams))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Key", api.Key)
	req.Header.Add("Sign", zaif.Sign(encodeParams, api.Secret))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	decoder := zaif.DecodeBody(resp, out)
	return decoder
}

func SetParam(param interface{}) url.Values {
	var ret url.Values
	ret = url.Values{}
	v := reflect.Indirect(reflect.ValueOf(param))

	num := reflect.ValueOf(param).Type().NumField()

	for i := 0; i < num; i++ {
		f := reflect.TypeOf(param).Field(i)

		tag := f.Tag.Get("url")
		v := v.Field(i).Interface()
		var value string
		switch t := v.(type) {
		case string:
			if t == "" {
				continue
			}
			value = t
		case bool:
			if t == false {
				continue
			}
			value = fmt.Sprintf("%v", t)
		case float64, float32:
			if t == 0.0 {
				continue
			}
			value = fmt.Sprintf("%v", t)
		case int64, uint64, int32, uint32, int, uint:
			if t == 0 {
				continue
			}
			value = fmt.Sprintf("%d", t)
		default:
			continue
		}
		ret[tag] = []string{value}
	}
	return ret
}
