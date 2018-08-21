package public

import (
	"encoding/json"
	"net/http"
)

const publicApiEndpoint = "https://api.zaif.jp/api/1"

type ApiClient struct{}

func NewApiClient() *ApiClient {
	return &ApiClient{}
}

func (api *ApiClient) Request(method string, param string, ret interface{}) error {
	res, err := http.Get(publicApiEndpoint + "/" + method + "/" + param)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	decode := json.NewDecoder(res.Body)
	return decode.Decode(ret)
}
