package public

import (
	"encoding/json"
	"net/http"
)

type ApiClient struct {
	Endpoint string
}

func NewApiClient(endpoint string) *ApiClient {
	return &ApiClient{endpoint}
}

func (api *ApiClient) GetRequest(method string, param string, ret interface{}) error {
	res, err := http.Get(api.Endpoint + "/" + method + "/" + param)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	decode := json.NewDecoder(res.Body)
	return decode.Decode(ret)
}
