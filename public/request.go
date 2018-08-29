package public

import (
	"github.com/naoki-maeda/zaif-go"
	"net/http"
)

type ApiClient struct {
	Endpoint string
}

func NewApiClient(endpoint string) *ApiClient {
	return &ApiClient{endpoint}
}

func (api *ApiClient) GetRequest(method string, param string, out interface{}) error {
	resp, err := http.Get(api.Endpoint + "/" + method + "/" + param)
	if err != nil {
		return err
	}
	decoder := zaif.DecodeBody(resp, out)
	return decoder
}
