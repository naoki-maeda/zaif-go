package private

import (
	"github.com/naoki-maeda/zaif-go"
	"net/http"
	"net/url"
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

func (api *ApiClient) Request(method string, param interface{}, out interface{}) error {
	params := newPrivateApiParams(method)
	values := url.Values{}
	values.Add("method", params.Method)
	values.Add("nonce", params.Nonce)

	req, err := http.NewRequest("POST", api.Endpoint, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Key", api.Key)
	req.Header.Add("Sign", zaif.Sign(values.Encode(), api.Secret))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	decoder := zaif.DecodeBody(resp, out)
	return decoder
}
