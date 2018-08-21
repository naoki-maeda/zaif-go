package private

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/naoki-maeda/zaif-go"
	"net/http"
	"strings"
)

const tradeApiEndpoint = "https://api.zaif.jp/tapi"

type ApiClient struct {
	Key    string
	Secret string
}

func NewApiClient(key string, secret string) *ApiClient {
	return &ApiClient{
		Key:    key,
		Secret: secret,
	}
}

type ApiParams struct {
	Method string `url:"method"`
	Nonce  string `url:"nonce"`
}

func newTradeApiParams(method string) ApiParams {
	return ApiParams{
		Method: method,
		Nonce:  zaif_go.GetNonce(),
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

func (api *ApiClient) Request(method string, param interface{}, ret interface{}) error {
	values, err := query.Values(newTradeApiParams(method))
	if err != nil {
		return err
	}
	encodedParams := values.Encode()
	if param != nil {
		params, err := query.Values(param)
		if err != nil {
			return err
		}
		encodedParams += "&" + params.Encode()
	}

	req, err := http.NewRequest("POST", tradeApiEndpoint, strings.NewReader(encodedParams))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Key", api.Key)
	req.Header.Add("Sign", zaif_go.Sign(encodedParams, api.Secret))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	decode := json.NewDecoder(res.Body)
	return decode.Decode(ret)
}
