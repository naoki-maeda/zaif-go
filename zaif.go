package zaif_go

import (
	"net/http"
	"io/ioutil"
	"strings"
	"net/url"
)

const (
	publicEndpoint = "https://api.zaif.jp/api/1"
	privateEndpoint = "https://api.zaif.jp/tapi"
)

type Key struct {
	Key string
	Secret string
}


func PublicRequest(method string, param string) string {
	req, _ := http.NewRequest("GET", publicEndpoint + "/" + method + "/" + param, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "http request error!!"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}


func (k Key) PrivateRequest(method string, params map[string] string) (string) {
	values := url.Values{}
	values.Add("method", method)
	values.Add("nonce", GetNonce())
	for k, v := range params {
		values.Add(k, v)
	}
	encodedParams := values.Encode()
	sign := Sign(encodedParams, k.Secret)
	req, _ := http.NewRequest("POST", privateEndpoint, strings.NewReader(encodedParams))
	req.Header.Set("Key", k.Key)
	req.Header.Set("Sign", sign)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "http request error!!"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
